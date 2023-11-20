// Code generated by mockery. DO NOT EDIT.

package same_name_arg_and_type

import mock "github.com/stretchr/testify/mock"

// interfaceB0Mock is an autogenerated mock type for the interfaceB0 type
type interfaceB0Mock struct {
	mock.Mock
}

type interfaceB0Mock_Expecter struct {
	mock *mock.Mock
}

func (_m *interfaceB0Mock) EXPECT() *interfaceB0Mock_Expecter {
	return &interfaceB0Mock_Expecter{mock: &_m.Mock}
}

// DoB0 provides a mock function with given fields: interfaceB00
func (_m *interfaceB0Mock) DoB0(interfaceB00 interfaceB0) interfaceB0 {
	ret := _m.Called(interfaceB00)

	if len(ret) == 0 {
		panic("Missing Return() function for DoB0()")
	}

	var r0 interfaceB0
	if rf, ok := ret.Get(0).(func(interfaceB0) interfaceB0); ok {
		r0 = rf(interfaceB00)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaceB0)
		}
	}

	return r0
}

// interfaceB0Mock_DoB0_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoB0'
type interfaceB0Mock_DoB0_Call struct {
	*mock.Call
}

// DoB0 is a helper method to define mock.On call
//   - interfaceB00 interfaceB0
func (_e *interfaceB0Mock_Expecter) DoB0(interfaceB00 interface{}) *interfaceB0Mock_DoB0_Call {
	return &interfaceB0Mock_DoB0_Call{Call: _e.mock.On("DoB0", interfaceB00)}
}

func (_c *interfaceB0Mock_DoB0_Call) Run(run func(interfaceB00 interfaceB0)) *interfaceB0Mock_DoB0_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interfaceB0))
	})
	return _c
}

func (_c *interfaceB0Mock_DoB0_Call) Return(_a0 interfaceB0) *interfaceB0Mock_DoB0_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *interfaceB0Mock_DoB0_Call) RunAndReturn(run func(interfaceB0) interfaceB0) *interfaceB0Mock_DoB0_Call {
	_c.Call.Return(run)
	return _c
}

// newInterfaceB0Mock creates a new instance of interfaceB0Mock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newInterfaceB0Mock(t interface {
	mock.TestingT
	Cleanup(func())
}) *interfaceB0Mock {
	mock := &interfaceB0Mock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
