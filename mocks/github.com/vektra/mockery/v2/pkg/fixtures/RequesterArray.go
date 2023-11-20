// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RequesterArray is an autogenerated mock type for the RequesterArray type
type RequesterArray struct {
	mock.Mock
}

type RequesterArray_Expecter struct {
	mock *mock.Mock
}

func (_m *RequesterArray) EXPECT() *RequesterArray_Expecter {
	return &RequesterArray_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: path
func (_m *RequesterArray) Get(path string) ([2]string, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("Missing Return() function for Get()")
	}

	var r0 [2]string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([2]string, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) [2]string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([2]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequesterArray_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RequesterArray_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - path string
func (_e *RequesterArray_Expecter) Get(path interface{}) *RequesterArray_Get_Call {
	return &RequesterArray_Get_Call{Call: _e.mock.On("Get", path)}
}

func (_c *RequesterArray_Get_Call) Run(run func(path string)) *RequesterArray_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RequesterArray_Get_Call) Return(_a0 [2]string, _a1 error) *RequesterArray_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RequesterArray_Get_Call) RunAndReturn(run func(string) ([2]string, error)) *RequesterArray_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewRequesterArray creates a new instance of RequesterArray. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequesterArray(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequesterArray {
	mock := &RequesterArray{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
