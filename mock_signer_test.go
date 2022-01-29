// Code generated by mockery v2.10.0. DO NOT EDIT.

package gondlr

import mock "github.com/stretchr/testify/mock"

// MockSigner is an autogenerated mock type for the Signer type
type MockSigner struct {
	mock.Mock
}

// Sign provides a mock function with given fields: data
func (_m *MockSigner) Sign(data []byte) ([]byte, error) {
	ret := _m.Called(data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Verify provides a mock function with given fields: data, sig
func (_m *MockSigner) Verify(data []byte, sig []byte) bool {
	ret := _m.Called(data, sig)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]byte, []byte) bool); ok {
		r0 = rf(data, sig)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}