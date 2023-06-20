// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bxcodec/go-clean-arch/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: ctx, us
func (_m *UserRepository) ChangePassword(ctx context.Context, us *domain.UserData) error {
	ret := _m.Called(ctx, us)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserData) error); ok {
		r0 = rf(ctx, us)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UserRepository) DeleteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserProfil provides a mock function with given fields: ctx, id
func (_m *UserRepository) DeleteUserProfil(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUser provides a mock function with given fields: ctx, us
func (_m *UserRepository) FindUser(ctx context.Context, us *domain.UserData) (*domain.User, error) {
	ret := _m.Called(ctx, us)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserData) (*domain.User, error)); ok {
		return rf(ctx, us)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserData) *domain.User); ok {
		r0 = rf(ctx, us)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.UserData) error); ok {
		r1 = rf(ctx, us)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) FindUserByEmail(ctx context.Context, email string) (*domain.UserData, error) {
	ret := _m.Called(ctx, email)

	var r0 *domain.UserData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.UserData, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.UserData); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserById provides a mock function with given fields: ctx, us
func (_m *UserRepository) FindUserById(ctx context.Context, us *domain.UserData) (*domain.User, error) {
	ret := _m.Called(ctx, us)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserData) (*domain.User, error)); ok {
		return rf(ctx, us)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserData) *domain.User); ok {
		r0 = rf(ctx, us)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.UserData) error); ok {
		r1 = rf(ctx, us)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJob provides a mock function with given fields: ctx
func (_m *UserRepository) GetJob(ctx context.Context) ([]*domain.Job, error) {
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

// GetListUser provides a mock function with given fields: ctx
func (_m *UserRepository) GetListUser(ctx context.Context) ([]*domain.User, error) {
	ret := _m.Called(ctx)

	var r0 []*domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
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
func (_m *UserRepository) GetUnit(ctx context.Context) ([]*domain.UnitDTO, error) {
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

// Register provides a mock function with given fields: ctx, us
func (_m *UserRepository) Register(ctx context.Context, us *domain.UserData) error {
	ret := _m.Called(ctx, us)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserData) error); ok {
		r0 = rf(ctx, us)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreProfile provides a mock function with given fields: ctx, us
func (_m *UserRepository) StoreProfile(ctx context.Context, us *domain.UserData) error {
	ret := _m.Called(ctx, us)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserData) error); ok {
		r0 = rf(ctx, us)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
