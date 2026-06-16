package sdl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// AKT-492: parser-level cross-field validations for interconnect.
//
// Each test builds a full SDL via Read() — the same entry point the
// CLI/provider use — so we exercise the real parse + validate pipeline.
// See docs/sdl-interconnect-spec.md for the rules being enforced.

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

// Rule 1: any interconnect opt-in requires the placement to require
// capabilities/gpu-interconnect=true. Tested here with the implicit
// form; the explicit form goes through the same code path.
func TestSDL_Interconnect_Rule1_InterconnectRequiresPlacementCapability(t *testing.T) {
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: []`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: []`,
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
            interconnect: []`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: []`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.NoError(t, err)
}

// Rule 2: the reserved group name `auto` cannot be written explicitly
// under `{group: ...}`. The parser rejects this at the per-profile
// UnmarshalYAML level — verified here through the full Read() pipeline.
func TestSDL_Interconnect_Rule2_AutoReservedName(t *testing.T) {
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect:
              group: auto`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect:
              group: auto`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.Error(t, err)
	require.True(t,
		strings.Contains(err.Error(), "reserved name"),
		"unexpected error: %v", err)
}

// Rule 3: within one placement, no mixing of implicit (`interconnect: []`)
// and explicit (`interconnect: { group: ... }`) opt-in forms.
func TestSDL_Interconnect_Rule3_NoMixingImplicitAndExplicit(t *testing.T) {
	body := sdlBody(
		// head uses implicit form
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: []`,
		// worker uses explicit form — same placement, must reject
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect:
              group: pair0`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.Error(t, err)
	require.True(t,
		strings.Contains(err.Error(), "mixes implicit") && strings.Contains(err.Error(), "explicit"),
		"unexpected error: %v", err)
}

// Rule 3 (positive): every service uses the implicit form — accepted,
// both profiles end up in the shared `auto` group.
func TestSDL_Interconnect_Rule3_AllImplicitOK(t *testing.T) {
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: []`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect: []`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.NoError(t, err)
}

// Rule 3 (positive): every service uses the explicit form, sharing one
// named group — accepted.
func TestSDL_Interconnect_Rule3_AllExplicitOK(t *testing.T) {
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect:
              group: pair1`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect:
              group: pair1`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.NoError(t, err)
}

// Rule 3 (positive): two explicit groups in one placement is fine —
// the no-mixing rule blocks implicit-vs-explicit, not multiple disjoint
// explicit groups.
func TestSDL_Interconnect_Rule3_TwoDisjointExplicitGroupsOK(t *testing.T) {
	body := sdlBody(
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect:
              group: pair0`,
		`          units: 1
          attributes:
            vendor: { nvidia: [{ model: a100 }] }
            interconnect:
              group: pair1`,
		`        capabilities/gpu-interconnect: "true"`,
	)

	_, err := Read([]byte(body))
	require.NoError(t, err)
}
