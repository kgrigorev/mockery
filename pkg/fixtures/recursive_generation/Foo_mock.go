// Code generated by mockery. DO NOT EDIT.

package recursive_generation

import mock "github.com/stretchr/testify/mock"

// MockFoo is an autogenerated mock type for the Foo type
type MockFoo struct {
	mock.Mock
}

type MockFoo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFoo) EXPECT() *MockFoo_Expecter {
	return &MockFoo_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields:
func (_m *MockFoo) Get() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("Missing Return() function for Get()")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockFoo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockFoo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *MockFoo_Expecter) Get() *MockFoo_Get_Call {
	return &MockFoo_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *MockFoo_Get_Call) Run(run func()) *MockFoo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFoo_Get_Call) Return(_a0 string) *MockFoo_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFoo_Get_Call) RunAndReturn(run func() string) *MockFoo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFoo creates a new instance of MockFoo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFoo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFoo {
	mock := &MockFoo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
