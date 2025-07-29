package testutil

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"

	rtypes "pkg.akt.dev/go/node/types/resources/v1beta4"
)

type InterceptState func(codec.Codec, string, json.RawMessage) json.RawMessage

type networkConfigOptions struct {
	interceptState InterceptState
}

type ConfigOption func(*networkConfigOptions)

// WithInterceptState set custom name of the log object
func WithInterceptState(val InterceptState) ConfigOption {
	return func(t *networkConfigOptions) {
		t.interceptState = val
	}
}

func ResourceUnits(_ testing.TB) rtypes.Resources {
	return rtypes.Resources{
		ID: 1,
		CPU: &rtypes.CPU{
			Units: rtypes.NewResourceValue(uint64(RandCPUUnits())),
		},
		Memory: &rtypes.Memory{
			Quantity: rtypes.NewResourceValue(RandMemoryQuantity()),
		},
		GPU: &rtypes.GPU{
			Units: rtypes.NewResourceValue(uint64(RandGPUUnits())),
		},
		Storage: rtypes.Volumes{
			rtypes.Storage{
				Quantity: rtypes.NewResourceValue(RandStorageQuantity()),
			},
		},
	}
}
