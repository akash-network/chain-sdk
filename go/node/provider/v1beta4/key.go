package v1beta4

const (
	// ModuleName is the module name constant used in many places
	ModuleName = "provider"

	// StoreKey is the store key string for provider
	StoreKey = ModuleName

	// RouterKey is the message route for provider
	RouterKey = ModuleName
)

func ProviderPrefix() []byte {
	return []byte{0x01}
}
