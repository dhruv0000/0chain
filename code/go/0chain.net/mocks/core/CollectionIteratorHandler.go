// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	datastore "0chain.net/core/datastore"
	mock "github.com/stretchr/testify/mock"
)

// CollectionIteratorHandler is an autogenerated mock type for the CollectionIteratorHandler type
type CollectionIteratorHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, ce
func (_m *CollectionIteratorHandler) Execute(ctx context.Context, ce datastore.CollectionEntity) bool {
	ret := _m.Called(ctx, ce)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, datastore.CollectionEntity) bool); ok {
		r0 = rf(ctx, ce)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}