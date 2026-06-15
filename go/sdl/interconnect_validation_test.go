package sdl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// CS-5: parser-level cross-field validations for interconnect.
//
// Each test builds a full SDL via Read() — the same entry point the
// CLI/provider use — so we exercise the real parse + validate pipeline.

const sdlInterconnectFixtureHeader = `---
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
	return sdlInterconnectFixtureHeader + `
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
    fabric:
      attributes:
` + placementAttrs + `
      pricing:
        inference-head:   { denom: uact, amount: 1000000 }
        inference-worker: { denom: uact, amount: 1000000 }
deployment:
  inference-head:
    fabric: { profile: inference-head,   count: 1 }
  inference-worker:
    fabric: { profile: inference-worker, count: 1 }
`
}

func TestSDL_Interconnect_Rule1_InterconnectRequiresPlacementCapability(t *testing.T) {
	// Profiles declare gpu.attributes.interconnect: true but placement does NOT
	// require capabilities/gpu-interconnect=true — rule 1 must reject.
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true`,
		`        capabilities/gpu: nvidia`,
	)

	_, err := Read([]byte(body))
	require.Error(t, err)
	require.True(t,
		strings.Contains(err.Error(), "does not require capabilities/gpu-interconnect=true"),
		"unexpected error: %v", err)
}

func TestSDL_Interconnect_Rule1_PassesWhenPlacementRequiresInterconnect(t *testing.T) {
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.NoError(t, err)
}

func TestSDL_Interconnect_Rule2_InterconnectGroupRequiresInterconnectOnSameProfile(t *testing.T) {
	body := sdlBody(
		// head: interconnect: true + interconnect_group set (OK)
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true
            interconnect_group: pair1`,
		// worker: interconnect_group set but interconnect is NOT true — rule 2 must reject.
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect_group: pair1`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.Error(t, err)
	require.True(t,
		strings.Contains(err.Error(), "interconnect_group") && strings.Contains(err.Error(), "interconnect: true"),
		"unexpected error: %v", err)
}

func TestSDL_Interconnect_Rule3_NoMixingExplicitAndImplicitGroup(t *testing.T) {
	body := sdlBody(
		// head: interconnect: true with no interconnect_group (implicit __default__).
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true`,
		// worker: interconnect: true with interconnect_group explicitly set.
		// Rule 3: cannot mix; reject.
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true
            interconnect_group: pair1`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.Error(t, err)
	require.True(t,
		strings.Contains(err.Error(), "mixes explicit and implicit interconnect_group"),
		"unexpected error: %v", err)
}

func TestSDL_Interconnect_Rule3_AllImplicitOK(t *testing.T) {
	// Simple head/worker with no interconnect_group on either profile — the implicit
	// default group is fine.
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.NoError(t, err)
}

func TestSDL_Interconnect_Rule3_AllExplicitOK(t *testing.T) {
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true
            interconnect_group: pair1`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: true
            interconnect_group: pair1`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.NoError(t, err)
}
