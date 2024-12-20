// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	mongo "go.mongodb.org/mongo-driver/v2/mongo"

	options "go.mongodb.org/mongo-driver/v2/mongo/options"
)

// ICollection is an autogenerated mock type for the ICollection type
type ICollection struct {
	mock.Mock
}

// DeleteOne provides a mock function with given fields: ctx, filter, opts
func (_m *ICollection) DeleteOne(ctx context.Context, filter interface{}, opts ...options.Lister[options.DeleteOptions]) (*mongo.DeleteResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.DeleteResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...options.Lister[options.DeleteOptions]) (*mongo.DeleteResult, error)); ok {
		return rf(ctx, filter, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...options.Lister[options.DeleteOptions]) *mongo.DeleteResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.DeleteResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...options.Lister[options.DeleteOptions]) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter, opts
func (_m *ICollection) Find(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.Cursor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...options.Lister[options.FindOptions]) (*mongo.Cursor, error)); ok {
		return rf(ctx, filter, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...options.Lister[options.FindOptions]) *mongo.Cursor); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Cursor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...options.Lister[options.FindOptions]) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, filter, opts
func (_m *ICollection) FindOne(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...options.Lister[options.FindOneOptions]) *mongo.SingleResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.SingleResult)
		}
	}

	return r0
}

// FindOneAndUpdate provides a mock function with given fields: ctx, filter, update, opts
func (_m *ICollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult); ok {
		r0 = rf(ctx, filter, update, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.SingleResult)
		}
	}

	return r0
}

// InsertOne provides a mock function with given fields: ctx, document, opts
func (_m *ICollection) InsertOne(ctx context.Context, document interface{}, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, document)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.InsertOneResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error)); ok {
		return rf(ctx, document, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...options.Lister[options.InsertOneOptions]) *mongo.InsertOneResult); ok {
		r0 = rf(ctx, document, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.InsertOneResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...options.Lister[options.InsertOneOptions]) error); ok {
		r1 = rf(ctx, document, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewICollection interface {
	mock.TestingT
	Cleanup(func())
}

// NewICollection creates a new instance of ICollection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICollection(t mockConstructorTestingTNewICollection) *ICollection {
	mock := &ICollection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
