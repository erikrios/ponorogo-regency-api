// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/erikrios/ponorogo-regency-api/entity"
	mock "github.com/stretchr/testify/mock"
)

// VillageRepository is an autogenerated mock type for the VillageRepository type
type VillageRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: ctx
func (_m *VillageRepository) FindAll(ctx context.Context) ([]entity.Village, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Village
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Village); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Village)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByDistrictID provides a mock function with given fields: ctx, districtID
func (_m *VillageRepository) FindByDistrictID(ctx context.Context, districtID string) ([]entity.Village, error) {
	ret := _m.Called(ctx, districtID)

	var r0 []entity.Village
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Village); ok {
		r0 = rf(ctx, districtID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Village)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, districtID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByDistrictName provides a mock function with given fields: ctx, keyword
func (_m *VillageRepository) FindByDistrictName(ctx context.Context, keyword string) ([]entity.Village, error) {
	ret := _m.Called(ctx, keyword)

	var r0 []entity.Village
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Village); ok {
		r0 = rf(ctx, keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Village)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *VillageRepository) FindByID(ctx context.Context, id string) (entity.Village, error) {
	ret := _m.Called(ctx, id)

	var r0 entity.Village
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Village); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Village)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: ctx, keyword
func (_m *VillageRepository) FindByName(ctx context.Context, keyword string) ([]entity.Village, error) {
	ret := _m.Called(ctx, keyword)

	var r0 []entity.Village
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Village); ok {
		r0 = rf(ctx, keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Village)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
