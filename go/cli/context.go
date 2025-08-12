package cli

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.akt.dev/go/node/client/v1beta3"
)

const (
	ClientContextKey = sdk.ContextKey("client.context")
	ServerContextKey = sdk.ContextKey("server.context")
)

type ContextType string

const (
	ContextTypeClient         = ContextType("context-client")
	ContextTypeQueryClient    = ContextType("context-query-client")
	ContextTypeAddressCodec   = ContextType("address-codec")
	ContextTypeValidatorCodec = ContextType("validator-codec")
	ContextTypeRPCURI         = ContextType("rpc-uri")
	ContextTypeRPCClient      = ContextType("rpc-client")
)

func ClientFromContext(ctx context.Context) (v1beta3.Client, error) {
	val := ctx.Value(ContextTypeClient)
	if val == nil {
		return nil, errors.New("context does not have client set")
	}

	res, valid := val.(v1beta3.Client)
	if !valid {
		return nil, fmt.Errorf("invalid context value, expected \"v1beta3.Client\", actual \"%s\"", reflect.TypeOf(val))
	}

	return res, nil
}

func MustClientFromContext(ctx context.Context) v1beta3.Client {
	cl, err := ClientFromContext(ctx)
	if err != nil {
		panic(err.Error())
	}

	return cl
}

func LightClientFromContext(ctx context.Context) (v1beta3.LightClient, error) {
	val := ctx.Value(ContextTypeQueryClient)
	if val == nil {
		val = ctx.Value(ContextTypeClient)
		if val == nil {
			return nil, errors.New("context does not have client set")
		}
	}

	switch cl := val.(type) {
	case v1beta3.LightClient:
		return cl, nil
	case v1beta3.Client:
		return cl, nil
	default:
		return nil, fmt.Errorf("invalid context value. expected \"v1beta3.Client|v1beta3.LightClient\" actual %s", reflect.TypeOf(val).String())
	}
}

func MustLightClientFromContext(ctx context.Context) v1beta3.LightClient {
	cl, err := LightClientFromContext(ctx)
	if err != nil {
		panic(err.Error())
	}

	return cl
}

func MustAddressCodecFromContext(ctx context.Context) address.Codec {
	val := ctx.Value(ContextTypeAddressCodec)
	if val == nil {
		panic("context does not have address codec set")
	}

	res, valid := val.(address.Codec)
	if !valid {
		panic("invalid context value")
	}

	return res
}

func MustValidatorCodecFromContext(ctx context.Context) address.Codec {
	val := ctx.Value(ContextTypeValidatorCodec)
	if val == nil {
		panic("context does not have validator codec set")
	}

	res, valid := val.(address.Codec)
	if !valid {
		panic("invalid context value")
	}

	return res
}
