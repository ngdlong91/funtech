// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"
import product "github.com/ngdlong91/funtech/cmd/gin/pkg/product"

// ProductClient is an autogenerated mock type for the ProductClient type
type ProductClient struct {
	mock.Mock
}

// DoPurchase provides a mock function with given fields: ctx, in, opts
func (_m *ProductClient) DoPurchase(ctx context.Context, in *product.PurchaseRequest, opts ...grpc.CallOption) (*product.PurchaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *product.PurchaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *product.PurchaseRequest, ...grpc.CallOption) *product.PurchaseResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.PurchaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *product.PurchaseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}