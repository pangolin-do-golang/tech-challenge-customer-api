// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	customer "github.com/pangolin-do-golang/tech-challenge-customer-api/internal/core/customer"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *IRepository) Create(ctx context.Context, _a1 *customer.Customer) (*customer.Customer, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *customer.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *customer.Customer) (*customer.Customer, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *customer.Customer) *customer.Customer); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customer.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *customer.Customer) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, customerId
func (_m *IRepository) Delete(ctx context.Context, customerId uuid.UUID) error {
	ret := _m.Called(ctx, customerId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, customerId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *IRepository) GetAll(ctx context.Context) ([]*customer.Customer, error) {
	ret := _m.Called(ctx)

	var r0 []*customer.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*customer.Customer, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*customer.Customer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*customer.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByCpf provides a mock function with given fields: ctx, customerCpf
func (_m *IRepository) GetByCpf(ctx context.Context, customerCpf string) (*customer.Customer, error) {
	ret := _m.Called(ctx, customerCpf)

	var r0 *customer.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*customer.Customer, error)); ok {
		return rf(ctx, customerCpf)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *customer.Customer); ok {
		r0 = rf(ctx, customerCpf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customer.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, customerCpf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *IRepository) Update(ctx context.Context, _a1 *customer.Customer) (*customer.Customer, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *customer.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *customer.Customer) (*customer.Customer, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *customer.Customer) *customer.Customer); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customer.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *customer.Customer) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepository(t mockConstructorTestingTNewIRepository) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
