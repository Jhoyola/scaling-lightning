// Code generated by mockery v2.29.0. DO NOT EDIT.

package mocks

import (
	btcjson "github.com/btcsuite/btcd/btcjson"
	btcutil "github.com/btcsuite/btcd/btcutil"

	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"

	mock "github.com/stretchr/testify/mock"

	rpcclient "github.com/btcsuite/btcd/rpcclient"
)

// RpcClient is an autogenerated mock type for the rpcClient type
type RpcClient struct {
	mock.Mock
}

// CreateWallet provides a mock function with given fields: name, opts
func (_m *RpcClient) CreateWallet(name string, opts ...rpcclient.CreateWalletOpt) (*btcjson.CreateWalletResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *btcjson.CreateWalletResult
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...rpcclient.CreateWalletOpt) (*btcjson.CreateWalletResult, error)); ok {
		return rf(name, opts...)
	}
	if rf, ok := ret.Get(0).(func(string, ...rpcclient.CreateWalletOpt) *btcjson.CreateWalletResult); ok {
		r0 = rf(name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*btcjson.CreateWalletResult)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...rpcclient.CreateWalletOpt) error); ok {
		r1 = rf(name, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateToAddress provides a mock function with given fields: numBlocks, address, maxTries
func (_m *RpcClient) GenerateToAddress(numBlocks int64, address btcutil.Address, maxTries *int64) ([]*chainhash.Hash, error) {
	ret := _m.Called(numBlocks, address, maxTries)

	var r0 []*chainhash.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, btcutil.Address, *int64) ([]*chainhash.Hash, error)); ok {
		return rf(numBlocks, address, maxTries)
	}
	if rf, ok := ret.Get(0).(func(int64, btcutil.Address, *int64) []*chainhash.Hash); ok {
		r0 = rf(numBlocks, address, maxTries)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*chainhash.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, btcutil.Address, *int64) error); ok {
		r1 = rf(numBlocks, address, maxTries)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNewAddress provides a mock function with given fields: account
func (_m *RpcClient) GetNewAddress(account string) (btcutil.Address, error) {
	ret := _m.Called(account)

	var r0 btcutil.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (btcutil.Address, error)); ok {
		return rf(account)
	}
	if rf, ok := ret.Get(0).(func(string) btcutil.Address); ok {
		r0 = rf(account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(btcutil.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWalletInfo provides a mock function with given fields:
func (_m *RpcClient) GetWalletInfo() (*btcjson.GetWalletInfoResult, error) {
	ret := _m.Called()

	var r0 *btcjson.GetWalletInfoResult
	var r1 error
	if rf, ok := ret.Get(0).(func() (*btcjson.GetWalletInfoResult, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *btcjson.GetWalletInfoResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*btcjson.GetWalletInfoResult)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRpcClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewRpcClient creates a new instance of RpcClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRpcClient(t mockConstructorTestingTNewRpcClient) *RpcClient {
	mock := &RpcClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}