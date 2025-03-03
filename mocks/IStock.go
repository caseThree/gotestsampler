// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	stocks "github.com/caseThree/sampleserver/stocks"
	mock "github.com/stretchr/testify/mock"
)

// IStock is an autogenerated mock type for the IStock type
type IStock struct {
	mock.Mock
}

// GetTopTwoStocks provides a mock function with given fields:
func (_m *IStock) GetTopTwoStocks() []stocks.Stock {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTopTwoStocks")
	}

	var r0 []stocks.Stock
	if rf, ok := ret.Get(0).(func() []stocks.Stock); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stocks.Stock)
		}
	}

	return r0
}

// NewIStock creates a new instance of IStock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIStock(t interface {
	mock.TestingT
	Cleanup(func())
}) *IStock {
	mock := &IStock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
