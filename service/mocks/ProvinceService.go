// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/erikrios/ponorogo-regency-api/model"
	mock "github.com/stretchr/testify/mock"
)

// ProvinceService is an autogenerated mock type for the ProvinceService type
type ProvinceService struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx, keyword
func (_m *ProvinceService) GetAll(ctx context.Context, keyword string) ([]model.Province, error) {
	ret := _m.Called(ctx, keyword)

	var r0 []model.Province
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Province); ok {
		r0 = rf(ctx, keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Province)
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

// GetByID provides a mock function with given fields: ctx, id
func (_m *ProvinceService) GetByID(ctx context.Context, id string) (model.Province, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Province
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Province); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Province)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
