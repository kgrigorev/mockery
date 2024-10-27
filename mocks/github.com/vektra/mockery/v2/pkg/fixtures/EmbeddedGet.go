// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	constraints "github.com/vektra/mockery/v2/pkg/fixtures/constraints"
)

// EmbeddedGet is an autogenerated mock type for the EmbeddedGet type
type EmbeddedGet[T constraints.Signed] struct {
	mock.Mock
}

type EmbeddedGet_Expecter[T constraints.Signed] struct {
	mock *mock.Mock
}

func (_m *EmbeddedGet[T]) EXPECT() *EmbeddedGet_Expecter[T] {
	return &EmbeddedGet_Expecter[T]{mock: &_m.Mock}
}

// Get provides a mock function with given fields:
func (_m *EmbeddedGet[T]) Get() T {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 T
	if rf, ok := ret.Get(0).(func() T); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(T)
		}
	}

	return r0
}

// EmbeddedGet_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type EmbeddedGet_Get_Call[T constraints.Signed] struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *EmbeddedGet_Expecter[T]) Get() *EmbeddedGet_Get_Call[T] {
	return &EmbeddedGet_Get_Call[T]{Call: _e.mock.On("Get")}
}

func (_c *EmbeddedGet_Get_Call[T]) Run(run func()) *EmbeddedGet_Get_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EmbeddedGet_Get_Call[T]) Return(_a0 T) *EmbeddedGet_Get_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EmbeddedGet_Get_Call[T]) RunAndReturn(run func() T) *EmbeddedGet_Get_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewEmbeddedGet creates a new instance of EmbeddedGet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmbeddedGet[T constraints.Signed](t interface {
	mock.TestingT
	Cleanup(func())
}) *EmbeddedGet[T] {
	mock := &EmbeddedGet[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
