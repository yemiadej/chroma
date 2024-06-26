// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GrpcServer is an autogenerated mock type for the GrpcServer type
type GrpcServer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *GrpcServer) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Port provides a mock function with given fields:
func (_m *GrpcServer) Port() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Port")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// NewGrpcServer creates a new instance of GrpcServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGrpcServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *GrpcServer {
	mock := &GrpcServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
