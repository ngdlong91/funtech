// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import dto "github.com/ngdlong91/funtech/dto"
import mock "github.com/stretchr/testify/mock"

// ProductUseCase is an autogenerated mock type for the ProductUseCase type
type ProductUseCase struct {
	mock.Mock
}

// Purchase provides a mock function with given fields: payload
func (_m *ProductUseCase) Purchase(payload dto.RequestPurchase) ([]dto.PurchaseResult, error) {
	ret := _m.Called(payload)

	var r0 []dto.PurchaseResult
	if rf, ok := ret.Get(0).(func(dto.RequestPurchase) []dto.PurchaseResult); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.PurchaseResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.RequestPurchase) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
