// Code generated by mockery v2.52.2. DO NOT EDIT.

package v1beta3

import (
	authz "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	certv1 "pkg.akt.dev/go/node/cert/v1"

	client "github.com/cosmos/cosmos-sdk/client"

	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	evidencetypes "cosmossdk.io/x/evidence/types"

	feegrant "cosmossdk.io/x/feegrant"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	mock "github.com/stretchr/testify/mock"

	proposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	providerv1beta4 "pkg.akt.dev/go/node/provider/v1beta4"

	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	types "github.com/cosmos/cosmos-sdk/x/auth/types"

	typesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	upgradetypes "cosmossdk.io/x/upgrade/types"

	v1 "pkg.akt.dev/go/node/audit/v1"

	v1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	v1beta4 "pkg.akt.dev/go/node/deployment/v1beta4"

	v1beta5 "pkg.akt.dev/go/node/market/v1beta5"
)

// QueryClient is an autogenerated mock type for the QueryClient type
type QueryClient struct {
	mock.Mock
}

type QueryClient_Expecter struct {
	mock *mock.Mock
}

func (_m *QueryClient) EXPECT() *QueryClient_Expecter {
	return &QueryClient_Expecter{mock: &_m.Mock}
}

// Audit provides a mock function with no fields
func (_m *QueryClient) Audit() v1.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Audit")
	}

	var r0 v1.QueryClient
	if rf, ok := ret.Get(0).(func() v1.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.QueryClient)
		}
	}

	return r0
}

// QueryClient_Audit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Audit'
type QueryClient_Audit_Call struct {
	*mock.Call
}

// Audit is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Audit() *QueryClient_Audit_Call {
	return &QueryClient_Audit_Call{Call: _e.mock.On("Audit")}
}

func (_c *QueryClient_Audit_Call) Run(run func()) *QueryClient_Audit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Audit_Call) Return(_a0 v1.QueryClient) *QueryClient_Audit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Audit_Call) RunAndReturn(run func() v1.QueryClient) *QueryClient_Audit_Call {
	_c.Call.Return(run)
	return _c
}

// Auth provides a mock function with no fields
func (_m *QueryClient) Auth() types.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Auth")
	}

	var r0 types.QueryClient
	if rf, ok := ret.Get(0).(func() types.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.QueryClient)
		}
	}

	return r0
}

// QueryClient_Auth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Auth'
type QueryClient_Auth_Call struct {
	*mock.Call
}

// Auth is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Auth() *QueryClient_Auth_Call {
	return &QueryClient_Auth_Call{Call: _e.mock.On("Auth")}
}

func (_c *QueryClient_Auth_Call) Run(run func()) *QueryClient_Auth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Auth_Call) Return(_a0 types.QueryClient) *QueryClient_Auth_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Auth_Call) RunAndReturn(run func() types.QueryClient) *QueryClient_Auth_Call {
	_c.Call.Return(run)
	return _c
}

// Authz provides a mock function with no fields
func (_m *QueryClient) Authz() authz.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Authz")
	}

	var r0 authz.QueryClient
	if rf, ok := ret.Get(0).(func() authz.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authz.QueryClient)
		}
	}

	return r0
}

// QueryClient_Authz_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Authz'
type QueryClient_Authz_Call struct {
	*mock.Call
}

// Authz is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Authz() *QueryClient_Authz_Call {
	return &QueryClient_Authz_Call{Call: _e.mock.On("Authz")}
}

func (_c *QueryClient_Authz_Call) Run(run func()) *QueryClient_Authz_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Authz_Call) Return(_a0 authz.QueryClient) *QueryClient_Authz_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Authz_Call) RunAndReturn(run func() authz.QueryClient) *QueryClient_Authz_Call {
	_c.Call.Return(run)
	return _c
}

// Bank provides a mock function with no fields
func (_m *QueryClient) Bank() banktypes.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Bank")
	}

	var r0 banktypes.QueryClient
	if rf, ok := ret.Get(0).(func() banktypes.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(banktypes.QueryClient)
		}
	}

	return r0
}

// QueryClient_Bank_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bank'
type QueryClient_Bank_Call struct {
	*mock.Call
}

// Bank is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Bank() *QueryClient_Bank_Call {
	return &QueryClient_Bank_Call{Call: _e.mock.On("Bank")}
}

func (_c *QueryClient_Bank_Call) Run(run func()) *QueryClient_Bank_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Bank_Call) Return(_a0 banktypes.QueryClient) *QueryClient_Bank_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Bank_Call) RunAndReturn(run func() banktypes.QueryClient) *QueryClient_Bank_Call {
	_c.Call.Return(run)
	return _c
}

// Certs provides a mock function with no fields
func (_m *QueryClient) Certs() certv1.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Certs")
	}

	var r0 certv1.QueryClient
	if rf, ok := ret.Get(0).(func() certv1.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(certv1.QueryClient)
		}
	}

	return r0
}

// QueryClient_Certs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Certs'
type QueryClient_Certs_Call struct {
	*mock.Call
}

// Certs is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Certs() *QueryClient_Certs_Call {
	return &QueryClient_Certs_Call{Call: _e.mock.On("Certs")}
}

func (_c *QueryClient_Certs_Call) Run(run func()) *QueryClient_Certs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Certs_Call) Return(_a0 certv1.QueryClient) *QueryClient_Certs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Certs_Call) RunAndReturn(run func() certv1.QueryClient) *QueryClient_Certs_Call {
	_c.Call.Return(run)
	return _c
}

// ClientContext provides a mock function with no fields
func (_m *QueryClient) ClientContext() client.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ClientContext")
	}

	var r0 client.Context
	if rf, ok := ret.Get(0).(func() client.Context); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(client.Context)
	}

	return r0
}

// QueryClient_ClientContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientContext'
type QueryClient_ClientContext_Call struct {
	*mock.Call
}

// ClientContext is a helper method to define mock.On call
func (_e *QueryClient_Expecter) ClientContext() *QueryClient_ClientContext_Call {
	return &QueryClient_ClientContext_Call{Call: _e.mock.On("ClientContext")}
}

func (_c *QueryClient_ClientContext_Call) Run(run func()) *QueryClient_ClientContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_ClientContext_Call) Return(_a0 client.Context) *QueryClient_ClientContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_ClientContext_Call) RunAndReturn(run func() client.Context) *QueryClient_ClientContext_Call {
	_c.Call.Return(run)
	return _c
}

// Deployment provides a mock function with no fields
func (_m *QueryClient) Deployment() v1beta4.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Deployment")
	}

	var r0 v1beta4.QueryClient
	if rf, ok := ret.Get(0).(func() v1beta4.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta4.QueryClient)
		}
	}

	return r0
}

// QueryClient_Deployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Deployment'
type QueryClient_Deployment_Call struct {
	*mock.Call
}

// Deployment is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Deployment() *QueryClient_Deployment_Call {
	return &QueryClient_Deployment_Call{Call: _e.mock.On("Deployment")}
}

func (_c *QueryClient_Deployment_Call) Run(run func()) *QueryClient_Deployment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Deployment_Call) Return(_a0 v1beta4.QueryClient) *QueryClient_Deployment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Deployment_Call) RunAndReturn(run func() v1beta4.QueryClient) *QueryClient_Deployment_Call {
	_c.Call.Return(run)
	return _c
}

// Distribution provides a mock function with no fields
func (_m *QueryClient) Distribution() distributiontypes.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Distribution")
	}

	var r0 distributiontypes.QueryClient
	if rf, ok := ret.Get(0).(func() distributiontypes.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(distributiontypes.QueryClient)
		}
	}

	return r0
}

// QueryClient_Distribution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Distribution'
type QueryClient_Distribution_Call struct {
	*mock.Call
}

// Distribution is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Distribution() *QueryClient_Distribution_Call {
	return &QueryClient_Distribution_Call{Call: _e.mock.On("Distribution")}
}

func (_c *QueryClient_Distribution_Call) Run(run func()) *QueryClient_Distribution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Distribution_Call) Return(_a0 distributiontypes.QueryClient) *QueryClient_Distribution_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Distribution_Call) RunAndReturn(run func() distributiontypes.QueryClient) *QueryClient_Distribution_Call {
	_c.Call.Return(run)
	return _c
}

// Evidence provides a mock function with no fields
func (_m *QueryClient) Evidence() evidencetypes.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Evidence")
	}

	var r0 evidencetypes.QueryClient
	if rf, ok := ret.Get(0).(func() evidencetypes.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(evidencetypes.QueryClient)
		}
	}

	return r0
}

// QueryClient_Evidence_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Evidence'
type QueryClient_Evidence_Call struct {
	*mock.Call
}

// Evidence is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Evidence() *QueryClient_Evidence_Call {
	return &QueryClient_Evidence_Call{Call: _e.mock.On("Evidence")}
}

func (_c *QueryClient_Evidence_Call) Run(run func()) *QueryClient_Evidence_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Evidence_Call) Return(_a0 evidencetypes.QueryClient) *QueryClient_Evidence_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Evidence_Call) RunAndReturn(run func() evidencetypes.QueryClient) *QueryClient_Evidence_Call {
	_c.Call.Return(run)
	return _c
}

// Feegrant provides a mock function with no fields
func (_m *QueryClient) Feegrant() feegrant.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Feegrant")
	}

	var r0 feegrant.QueryClient
	if rf, ok := ret.Get(0).(func() feegrant.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(feegrant.QueryClient)
		}
	}

	return r0
}

// QueryClient_Feegrant_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Feegrant'
type QueryClient_Feegrant_Call struct {
	*mock.Call
}

// Feegrant is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Feegrant() *QueryClient_Feegrant_Call {
	return &QueryClient_Feegrant_Call{Call: _e.mock.On("Feegrant")}
}

func (_c *QueryClient_Feegrant_Call) Run(run func()) *QueryClient_Feegrant_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Feegrant_Call) Return(_a0 feegrant.QueryClient) *QueryClient_Feegrant_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Feegrant_Call) RunAndReturn(run func() feegrant.QueryClient) *QueryClient_Feegrant_Call {
	_c.Call.Return(run)
	return _c
}

// Gov provides a mock function with no fields
func (_m *QueryClient) Gov() typesv1.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Gov")
	}

	var r0 typesv1.QueryClient
	if rf, ok := ret.Get(0).(func() typesv1.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(typesv1.QueryClient)
		}
	}

	return r0
}

// QueryClient_Gov_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Gov'
type QueryClient_Gov_Call struct {
	*mock.Call
}

// Gov is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Gov() *QueryClient_Gov_Call {
	return &QueryClient_Gov_Call{Call: _e.mock.On("Gov")}
}

func (_c *QueryClient_Gov_Call) Run(run func()) *QueryClient_Gov_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Gov_Call) Return(_a0 typesv1.QueryClient) *QueryClient_Gov_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Gov_Call) RunAndReturn(run func() typesv1.QueryClient) *QueryClient_Gov_Call {
	_c.Call.Return(run)
	return _c
}

// GovLegacy provides a mock function with no fields
func (_m *QueryClient) GovLegacy() v1beta1.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GovLegacy")
	}

	var r0 v1beta1.QueryClient
	if rf, ok := ret.Get(0).(func() v1beta1.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.QueryClient)
		}
	}

	return r0
}

// QueryClient_GovLegacy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GovLegacy'
type QueryClient_GovLegacy_Call struct {
	*mock.Call
}

// GovLegacy is a helper method to define mock.On call
func (_e *QueryClient_Expecter) GovLegacy() *QueryClient_GovLegacy_Call {
	return &QueryClient_GovLegacy_Call{Call: _e.mock.On("GovLegacy")}
}

func (_c *QueryClient_GovLegacy_Call) Run(run func()) *QueryClient_GovLegacy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_GovLegacy_Call) Return(_a0 v1beta1.QueryClient) *QueryClient_GovLegacy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_GovLegacy_Call) RunAndReturn(run func() v1beta1.QueryClient) *QueryClient_GovLegacy_Call {
	_c.Call.Return(run)
	return _c
}

// Market provides a mock function with no fields
func (_m *QueryClient) Market() v1beta5.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Market")
	}

	var r0 v1beta5.QueryClient
	if rf, ok := ret.Get(0).(func() v1beta5.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta5.QueryClient)
		}
	}

	return r0
}

// QueryClient_Market_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Market'
type QueryClient_Market_Call struct {
	*mock.Call
}

// Market is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Market() *QueryClient_Market_Call {
	return &QueryClient_Market_Call{Call: _e.mock.On("Market")}
}

func (_c *QueryClient_Market_Call) Run(run func()) *QueryClient_Market_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Market_Call) Return(_a0 v1beta5.QueryClient) *QueryClient_Market_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Market_Call) RunAndReturn(run func() v1beta5.QueryClient) *QueryClient_Market_Call {
	_c.Call.Return(run)
	return _c
}

// Mint provides a mock function with no fields
func (_m *QueryClient) Mint() minttypes.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Mint")
	}

	var r0 minttypes.QueryClient
	if rf, ok := ret.Get(0).(func() minttypes.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(minttypes.QueryClient)
		}
	}

	return r0
}

// QueryClient_Mint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mint'
type QueryClient_Mint_Call struct {
	*mock.Call
}

// Mint is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Mint() *QueryClient_Mint_Call {
	return &QueryClient_Mint_Call{Call: _e.mock.On("Mint")}
}

func (_c *QueryClient_Mint_Call) Run(run func()) *QueryClient_Mint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Mint_Call) Return(_a0 minttypes.QueryClient) *QueryClient_Mint_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Mint_Call) RunAndReturn(run func() minttypes.QueryClient) *QueryClient_Mint_Call {
	_c.Call.Return(run)
	return _c
}

// Params provides a mock function with no fields
func (_m *QueryClient) Params() proposal.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Params")
	}

	var r0 proposal.QueryClient
	if rf, ok := ret.Get(0).(func() proposal.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(proposal.QueryClient)
		}
	}

	return r0
}

// QueryClient_Params_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Params'
type QueryClient_Params_Call struct {
	*mock.Call
}

// Params is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Params() *QueryClient_Params_Call {
	return &QueryClient_Params_Call{Call: _e.mock.On("Params")}
}

func (_c *QueryClient_Params_Call) Run(run func()) *QueryClient_Params_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Params_Call) Return(_a0 proposal.QueryClient) *QueryClient_Params_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Params_Call) RunAndReturn(run func() proposal.QueryClient) *QueryClient_Params_Call {
	_c.Call.Return(run)
	return _c
}

// Provider provides a mock function with no fields
func (_m *QueryClient) Provider() providerv1beta4.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Provider")
	}

	var r0 providerv1beta4.QueryClient
	if rf, ok := ret.Get(0).(func() providerv1beta4.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(providerv1beta4.QueryClient)
		}
	}

	return r0
}

// QueryClient_Provider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Provider'
type QueryClient_Provider_Call struct {
	*mock.Call
}

// Provider is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Provider() *QueryClient_Provider_Call {
	return &QueryClient_Provider_Call{Call: _e.mock.On("Provider")}
}

func (_c *QueryClient_Provider_Call) Run(run func()) *QueryClient_Provider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Provider_Call) Return(_a0 providerv1beta4.QueryClient) *QueryClient_Provider_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Provider_Call) RunAndReturn(run func() providerv1beta4.QueryClient) *QueryClient_Provider_Call {
	_c.Call.Return(run)
	return _c
}

// Slashing provides a mock function with no fields
func (_m *QueryClient) Slashing() slashingtypes.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Slashing")
	}

	var r0 slashingtypes.QueryClient
	if rf, ok := ret.Get(0).(func() slashingtypes.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(slashingtypes.QueryClient)
		}
	}

	return r0
}

// QueryClient_Slashing_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Slashing'
type QueryClient_Slashing_Call struct {
	*mock.Call
}

// Slashing is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Slashing() *QueryClient_Slashing_Call {
	return &QueryClient_Slashing_Call{Call: _e.mock.On("Slashing")}
}

func (_c *QueryClient_Slashing_Call) Run(run func()) *QueryClient_Slashing_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Slashing_Call) Return(_a0 slashingtypes.QueryClient) *QueryClient_Slashing_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Slashing_Call) RunAndReturn(run func() slashingtypes.QueryClient) *QueryClient_Slashing_Call {
	_c.Call.Return(run)
	return _c
}

// Staking provides a mock function with no fields
func (_m *QueryClient) Staking() stakingtypes.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Staking")
	}

	var r0 stakingtypes.QueryClient
	if rf, ok := ret.Get(0).(func() stakingtypes.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stakingtypes.QueryClient)
		}
	}

	return r0
}

// QueryClient_Staking_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Staking'
type QueryClient_Staking_Call struct {
	*mock.Call
}

// Staking is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Staking() *QueryClient_Staking_Call {
	return &QueryClient_Staking_Call{Call: _e.mock.On("Staking")}
}

func (_c *QueryClient_Staking_Call) Run(run func()) *QueryClient_Staking_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Staking_Call) Return(_a0 stakingtypes.QueryClient) *QueryClient_Staking_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Staking_Call) RunAndReturn(run func() stakingtypes.QueryClient) *QueryClient_Staking_Call {
	_c.Call.Return(run)
	return _c
}

// Upgrade provides a mock function with no fields
func (_m *QueryClient) Upgrade() upgradetypes.QueryClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Upgrade")
	}

	var r0 upgradetypes.QueryClient
	if rf, ok := ret.Get(0).(func() upgradetypes.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(upgradetypes.QueryClient)
		}
	}

	return r0
}

// QueryClient_Upgrade_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upgrade'
type QueryClient_Upgrade_Call struct {
	*mock.Call
}

// Upgrade is a helper method to define mock.On call
func (_e *QueryClient_Expecter) Upgrade() *QueryClient_Upgrade_Call {
	return &QueryClient_Upgrade_Call{Call: _e.mock.On("Upgrade")}
}

func (_c *QueryClient_Upgrade_Call) Run(run func()) *QueryClient_Upgrade_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryClient_Upgrade_Call) Return(_a0 upgradetypes.QueryClient) *QueryClient_Upgrade_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryClient_Upgrade_Call) RunAndReturn(run func() upgradetypes.QueryClient) *QueryClient_Upgrade_Call {
	_c.Call.Return(run)
	return _c
}

// NewQueryClient creates a new instance of QueryClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueryClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *QueryClient {
	mock := &QueryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
