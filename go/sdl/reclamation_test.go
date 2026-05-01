package sdl

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV2_Reclamation(t *testing.T) {
	sdl, err := ReadFile("./_testdata/simple-reclamation.yaml")
	require.NoError(t, err)

	reclamation, err := sdl.Reclamation()
	require.NoError(t, err)
	require.NotNil(t, reclamation)
	assert.Equal(t, 24*time.Hour, reclamation.MinWindow)
}

func TestV2_NoReclamation(t *testing.T) {
	sdl, err := ReadFile("./_testdata/simple.yaml")
	require.NoError(t, err)

	reclamation, err := sdl.Reclamation()
	require.NoError(t, err)
	assert.Nil(t, reclamation)
}

func TestV2_1_Reclamation(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-simple-reclamation.yaml")
	require.NoError(t, err)

	reclamation, err := sdl.Reclamation()
	require.NoError(t, err)
	require.NotNil(t, reclamation)
	assert.Equal(t, 24*time.Hour, reclamation.MinWindow)
}

func TestV2_1_NoReclamation(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-simple.yaml")
	require.NoError(t, err)

	reclamation, err := sdl.Reclamation()
	require.NoError(t, err)
	assert.Nil(t, reclamation)
}

func TestReclamation_InvalidDuration(t *testing.T) {
	r := &v2Reclamation{MinWindow: "abc"}
	_, err := r.toDeploymentReclamation()
	require.Error(t, err)
	assert.ErrorIs(t, err, errSDLInvalid)
}

func TestReclamation_ZeroDuration(t *testing.T) {
	r := &v2Reclamation{MinWindow: "0s"}
	_, err := r.toDeploymentReclamation()
	require.Error(t, err)
	assert.ErrorIs(t, err, errSDLInvalid)
}

func TestReclamation_NegativeDuration(t *testing.T) {
	r := &v2Reclamation{MinWindow: "-1h"}
	_, err := r.toDeploymentReclamation()
	require.Error(t, err)
	assert.ErrorIs(t, err, errSDLInvalid)
}

func TestReclamation_NilReceiver(t *testing.T) {
	var r *v2Reclamation
	result, err := r.toDeploymentReclamation()
	require.NoError(t, err)
	assert.Nil(t, result)
}

func TestV2_ReadFile_InvalidReclamation(t *testing.T) {
	_, err := ReadFile("./_testdata/simple-reclamation-invalid.yaml")
	require.Error(t, err)
	assert.ErrorIs(t, err, errSDLInvalid)
}

func TestV2_ReadFile_ZeroReclamation(t *testing.T) {
	_, err := ReadFile("./_testdata/simple-reclamation-zero.yaml")
	require.Error(t, err)
	assert.ErrorIs(t, err, errSDLInvalid)
}

func TestV2_1_ReadFile_InvalidReclamation(t *testing.T) {
	_, err := ReadFile("./_testdata/v2.1-reclamation-invalid.yaml")
	require.Error(t, err)
	assert.ErrorIs(t, err, errSDLInvalid)
}

func TestV2_1_ReadFile_ZeroReclamation(t *testing.T) {
	_, err := ReadFile("./_testdata/v2.1-reclamation-zero.yaml")
	require.Error(t, err)
	assert.ErrorIs(t, err, errSDLInvalid)
}

func TestReclamation_EmptyString(t *testing.T) {
	r := &v2Reclamation{MinWindow: ""}
	_, err := r.toDeploymentReclamation()
	require.Error(t, err)
	assert.ErrorIs(t, err, errSDLInvalid)
}

func TestReclamation_ValidDurations(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
	}{
		{"1h", 1 * time.Hour},
		{"720h", 720 * time.Hour},
		{"30m", 30 * time.Minute},
		{"1h30m", 1*time.Hour + 30*time.Minute},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			r := &v2Reclamation{MinWindow: tc.input}
			result, err := r.toDeploymentReclamation()
			require.NoError(t, err)
			require.NotNil(t, result)
			assert.Equal(t, tc.expected, result.MinWindow)
		})
	}
}
