// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// HashScorer is an autogenerated mock type for the HashScorer type
type HashScorer struct {
	mock.Mock
}

// Score provides a mock function with given fields: hash1, hash2
func (_m *HashScorer) Score(hash1 []byte, hash2 []byte) int32 {
	ret := _m.Called(hash1, hash2)

	var r0 int32
	if rf, ok := ret.Get(0).(func([]byte, []byte) int32); ok {
		r0 = rf(hash1, hash2)
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}