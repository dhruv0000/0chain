// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	encryption "0chain.net/core/encryption"
	mock "github.com/stretchr/testify/mock"
)

// AggregateSignatureScheme is an autogenerated mock type for the AggregateSignatureScheme type
type AggregateSignatureScheme struct {
	mock.Mock
}

// Aggregate provides a mock function with given fields: ss, idx, signature, hash
func (_m *AggregateSignatureScheme) Aggregate(ss encryption.SignatureScheme, idx int, signature string, hash string) error {
	ret := _m.Called(ss, idx, signature, hash)

	var r0 error
	if rf, ok := ret.Get(0).(func(encryption.SignatureScheme, int, string, string) error); ok {
		r0 = rf(ss, idx, signature, hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Verify provides a mock function with given fields:
func (_m *AggregateSignatureScheme) Verify() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}