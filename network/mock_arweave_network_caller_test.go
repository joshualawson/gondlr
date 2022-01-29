// Code generated by mockery v2.10.0. DO NOT EDIT.

package network

import (
	arweave "github.com/joshualawson/arweave-api"
	mock "github.com/stretchr/testify/mock"
)

// MockArweaveNetworkCaller is an autogenerated mock type for the ArweaveNetworkCaller type
type MockArweaveNetworkCaller struct {
	mock.Mock
}

// Info provides a mock function with given fields:
func (_m *MockArweaveNetworkCaller) Info() (arweave.Info, error) {
	ret := _m.Called()

	var r0 arweave.Info
	if rf, ok := ret.Get(0).(func() arweave.Info); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(arweave.Info)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transaction provides a mock function with given fields: id
func (_m *MockArweaveNetworkCaller) Transaction(id string) (arweave.Transaction, error) {
	ret := _m.Called(id)

	var r0 arweave.Transaction
	if rf, ok := ret.Get(0).(func(string) arweave.Transaction); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(arweave.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionStatus provides a mock function with given fields: id
func (_m *MockArweaveNetworkCaller) TransactionStatus(id string) (arweave.TransactionStatus, error) {
	ret := _m.Called(id)

	var r0 arweave.TransactionStatus
	if rf, ok := ret.Get(0).(func(string) arweave.TransactionStatus); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(arweave.TransactionStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}