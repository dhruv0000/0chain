// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	datastore "0chain.net/core/datastore"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// CollectionEntity is an autogenerated mock type for the CollectionEntity type
type CollectionEntity struct {
	mock.Mock
}

// ComputeProperties provides a mock function with given fields:
func (_m *CollectionEntity) ComputeProperties() {
	_m.Called()
}

// Delete provides a mock function with given fields: ctx
func (_m *CollectionEntity) Delete(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCollectionDuration provides a mock function with given fields:
func (_m *CollectionEntity) GetCollectionDuration() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// GetCollectionName provides a mock function with given fields:
func (_m *CollectionEntity) GetCollectionName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetCollectionScore provides a mock function with given fields:
func (_m *CollectionEntity) GetCollectionScore() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetCollectionSize provides a mock function with given fields:
func (_m *CollectionEntity) GetCollectionSize() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetEntityMetadata provides a mock function with given fields:
func (_m *CollectionEntity) GetEntityMetadata() datastore.EntityMetadata {
	ret := _m.Called()

	var r0 datastore.EntityMetadata
	if rf, ok := ret.Get(0).(func() datastore.EntityMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.EntityMetadata)
		}
	}

	return r0
}

// GetKey provides a mock function with given fields:
func (_m *CollectionEntity) GetKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetScore provides a mock function with given fields:
func (_m *CollectionEntity) GetScore() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// InitCollectionScore provides a mock function with given fields:
func (_m *CollectionEntity) InitCollectionScore() {
	_m.Called()
}

// Read provides a mock function with given fields: ctx, key
func (_m *CollectionEntity) Read(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCollectionScore provides a mock function with given fields: score
func (_m *CollectionEntity) SetCollectionScore(score int64) {
	_m.Called(score)
}

// SetKey provides a mock function with given fields: key
func (_m *CollectionEntity) SetKey(key string) {
	_m.Called(key)
}

// Validate provides a mock function with given fields: ctx
func (_m *CollectionEntity) Validate(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: ctx
func (_m *CollectionEntity) Write(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}