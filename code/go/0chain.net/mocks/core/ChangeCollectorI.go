// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	io "io"

	util "0chain.net/core/util"
	mock "github.com/stretchr/testify/mock"
)

// ChangeCollectorI is an autogenerated mock type for the ChangeCollectorI type
type ChangeCollectorI struct {
	mock.Mock
}

// AddChange provides a mock function with given fields: oldNode, newNode
func (_m *ChangeCollectorI) AddChange(oldNode util.Node, newNode util.Node) {
	_m.Called(oldNode, newNode)
}

// Clone provides a mock function with given fields:
func (_m *ChangeCollectorI) Clone() util.ChangeCollectorI {
	ret := _m.Called()

	var r0 util.ChangeCollectorI
	if rf, ok := ret.Get(0).(func() util.ChangeCollectorI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(util.ChangeCollectorI)
		}
	}

	return r0
}

// DeleteChange provides a mock function with given fields: oldNode
func (_m *ChangeCollectorI) DeleteChange(oldNode util.Node) {
	_m.Called(oldNode)
}

// GetChanges provides a mock function with given fields:
func (_m *ChangeCollectorI) GetChanges() []*util.NodeChange {
	ret := _m.Called()

	var r0 []*util.NodeChange
	if rf, ok := ret.Get(0).(func() []*util.NodeChange); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*util.NodeChange)
		}
	}

	return r0
}

// GetDeletes provides a mock function with given fields:
func (_m *ChangeCollectorI) GetDeletes() []util.Node {
	ret := _m.Called()

	var r0 []util.Node
	if rf, ok := ret.Get(0).(func() []util.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]util.Node)
		}
	}

	return r0
}

// PrintChanges provides a mock function with given fields: w
func (_m *ChangeCollectorI) PrintChanges(w io.Writer) {
	_m.Called(w)
}

// UpdateChanges provides a mock function with given fields: ndb, origin, includeDeletes
func (_m *ChangeCollectorI) UpdateChanges(ndb util.NodeDB, origin util.Sequence, includeDeletes bool) error {
	ret := _m.Called(ndb, origin, includeDeletes)

	var r0 error
	if rf, ok := ret.Get(0).(func(util.NodeDB, util.Sequence, bool) error); ok {
		r0 = rf(ndb, origin, includeDeletes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Validate provides a mock function with given fields:
func (_m *ChangeCollectorI) Validate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}