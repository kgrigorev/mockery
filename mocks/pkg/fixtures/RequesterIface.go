// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// RequesterIface is an autogenerated mock type for the RequesterIface type
type RequesterIface struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *RequesterIface) Get() io.Reader {
	ret := _m.Called()

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func() io.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	return r0
}

// NewRequesterIface creates a new instance of RequesterIface. It also registers a cleanup function to assert the mocks expectations.
func NewRequesterIface(t testing.TB) *RequesterIface {
	mock := &RequesterIface{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
