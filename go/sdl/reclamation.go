package sdl

import (
	"fmt"
	"time"

	dv1 "pkg.akt.dev/go/node/deployment/v1"
)

type v2Reclamation struct {
	MinWindow string `yaml:"min_window"`
}

func (r *v2Reclamation) toDeploymentReclamation() (*dv1.DeploymentReclamation, error) {
	if r == nil {
		return nil, nil
	}

	d, err := time.ParseDuration(r.MinWindow)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid reclamation min_window %q: %v", errSDLInvalid, r.MinWindow, err)
	}

	if d <= 0 {
		return nil, fmt.Errorf("%w: reclamation min_window must be > 0", errSDLInvalid)
	}

	return &dv1.DeploymentReclamation{MinWindow: d}, nil
}
