// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bxcodec/go-clean-arch/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// GetJob provides a mock function with given fields: ctx
func (_m *UserUsecase) GetJob(ctx context.Context) ([]*domain.Job, error) {
	ret := _m.Called(ctx)

	var r0 []*domain.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.Job, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Job); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnit provides a mock function with given fields: ctx
func (_m *UserUsecase) GetUnit(ctx context.Context) ([]*domain.UnitDTO, error) {
	ret := _m.Called(ctx)

	var r0 []*domain.UnitDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.UnitDTO, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.UnitDTO); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.UnitDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, us
func (_m *UserUsecase) Login(ctx context.Context, us *domain.DtoRequestLogin) (*domain.Auth, error) {
	ret := _m.Called(ctx, us)

	var r0 *domain.Auth
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.DtoRequestLogin) (*domain.Auth, error)); ok {
		return rf(ctx, us)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.DtoRequestLogin) *domain.Auth); ok {
		r0 = rf(ctx, us)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Auth)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.DtoRequestLogin) error); ok {
		r1 = rf(ctx, us)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, us
func (_m *UserUsecase) Register(ctx context.Context, us *domain.User) error {
	ret := _m.Called(ctx, us)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, us)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
