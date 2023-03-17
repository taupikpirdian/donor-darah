// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bxcodec/go-clean-arch/domain"
	mock "github.com/stretchr/testify/mock"
)

// DonorRepository is an autogenerated mock type for the DonorRepository type
type DonorRepository struct {
	mock.Mock
}

// DonorRegister provides a mock function with given fields: ctx, donor
func (_m *DonorRepository) DonorRegister(ctx context.Context, donor *domain.DonorRegister) (int64, error) {
	ret := _m.Called(ctx, donor)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.DonorRegister) (int64, error)); ok {
		return rf(ctx, donor)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.DonorRegister) int64); ok {
		r0 = rf(ctx, donor)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.DonorRegister) error); ok {
		r1 = rf(ctx, donor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DonorRegisterQuestioner provides a mock function with given fields: ctx, donor, donorRegisterId
func (_m *DonorRepository) DonorRegisterQuestioner(ctx context.Context, donor *domain.DonorRegisterQuestioner, donorRegisterId int64) error {
	ret := _m.Called(ctx, donor, donorRegisterId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.DonorRegisterQuestioner, int64) error); ok {
		r0 = rf(ctx, donor, donorRegisterId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListAgenda provides a mock function with given fields: ctx, userId
func (_m *DonorRepository) ListAgenda(ctx context.Context, userId int64) ([]*domain.DonorRegisterDTO, error) {
	ret := _m.Called(ctx, userId)

	var r0 []*domain.DonorRegisterDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*domain.DonorRegisterDTO, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*domain.DonorRegisterDTO); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.DonorRegisterDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SingleAgenda provides a mock function with given fields: ctx, id
func (_m *DonorRepository) SingleAgenda(ctx context.Context, id int64) (*domain.DonorRegisterDTO, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.DonorRegisterDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*domain.DonorRegisterDTO, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.DonorRegisterDTO); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.DonorRegisterDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDonorRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewDonorRepository creates a new instance of DonorRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDonorRepository(t mockConstructorTestingTNewDonorRepository) *DonorRepository {
	mock := &DonorRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
