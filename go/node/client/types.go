package client

import (
	"pkg.akt.dev/go/node/client/v1beta3"
)

type Client interface {
	v1beta3.Client
}

type LightClient interface {
	v1beta3.LightClient
}

type QueryClient interface {
	v1beta3.QueryClient
}
