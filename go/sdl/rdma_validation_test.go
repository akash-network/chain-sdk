package sdl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// CS-5: parser-level cross-field validations for RDMA.
//
// Each test builds a full SDL via Read() — the same entry point the
// CLI/provider use — so we exercise the real parse + validate pipeline.

const sdlRDMAFixtureHeader = `---
version: "2.0"
services:
  inference-head:
    image: nvidia/cuda:latest
    expose:
      - port: 30000
        as: 80
        to: [{ global: true }]
  inference-worker:
    image: nvidia/cuda:latest
    expose:
      - port: 30000
        to: [{ service: inference-head }]
`

func sdlBody(headGPU, workerGPU, placementAttrs string) string {
	return sdlRDMAFixtureHeader + `
profiles:
  compute:
    inference-head:
      resources:
        cpu: { units: 1 }
        memory: { size: 1Gi }
        gpu:
` + headGPU + `
        storage:
          - size: 512Mi
    inference-worker:
      resources:
        cpu: { units: 1 }
        memory: { size: 1Gi }
        gpu:
` + workerGPU + `
        storage:
          - size: 512Mi
  placement:
    rdma:
      attributes:
` + placementAttrs + `
      pricing:
        inference-head:   { denom: uact, amount: 1000000 }
        inference-worker: { denom: uact, amount: 1000000 }
deployment:
  inference-head:
    rdma: { profile: inference-head,   count: 1 }
  inference-worker:
    rdma: { profile: inference-worker, count: 1 }
`
}

func TestSDL_RDMA_Rule1_RDMARequiresPlacementCapability(t *testing.T) {
	// Profiles declare gpu.attributes.rdma: true but placement does NOT
	// require capabilities/rdma=true — rule 1 must reject.
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true`,
		`        capabilities/gpu: nvidia`,
	)

	_, err := Read([]byte(body))
	require.Error(t, err)
	require.True(t,
		strings.Contains(err.Error(), "does not require capabilities/rdma=true"),
		"unexpected error: %v", err)
}

func TestSDL_RDMA_Rule1_PassesWhenPlacementRequiresRDMA(t *testing.T) {
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true`,
		`        capabilities/rdma: "true"`,
	)

	_, err := Read([]byte(body))
	require.NoError(t, err)
}

func TestSDL_RDMA_Rule2_RDMAGroupRequiresRDMAOnSameProfile(t *testing.T) {
	body := sdlBody(
		// head: rdma: true + rdma_group set (OK)
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true
            rdma_group: pair1`,
		// worker: rdma_group set but rdma is NOT true — rule 2 must reject.
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma_group: pair1`,
		`        capabilities/rdma: "true"`,
	)

	_, err := Read([]byte(body))
	require.Error(t, err)
	require.True(t,
		strings.Contains(err.Error(), "rdma_group") && strings.Contains(err.Error(), "rdma: true"),
		"unexpected error: %v", err)
}

func TestSDL_RDMA_Rule3_NoMixingExplicitAndImplicitGroup(t *testing.T) {
	body := sdlBody(
		// head: rdma: true with no rdma_group (implicit __default__).
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true`,
		// worker: rdma: true with rdma_group explicitly set.
		// Rule 3: cannot mix; reject.
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true
            rdma_group: pair1`,
		`        capabilities/rdma: "true"`,
	)

	_, err := Read([]byte(body))
	require.Error(t, err)
	require.True(t,
		strings.Contains(err.Error(), "mixes explicit and implicit rdma_group"),
		"unexpected error: %v", err)
}

func TestSDL_RDMA_Rule3_AllImplicitOK(t *testing.T) {
	// Simple head/worker with no rdma_group on either profile — the implicit
	// default group is fine.
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true`,
		`        capabilities/rdma: "true"`,
	)

	_, err := Read([]byte(body))
	require.NoError(t, err)
}

func TestSDL_RDMA_Rule3_AllExplicitOK(t *testing.T) {
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true
            rdma_group: pair1`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            rdma: true
            rdma_group: pair1`,
		`        capabilities/rdma: "true"`,
	)

	_, err := Read([]byte(body))
	require.NoError(t, err)
}
