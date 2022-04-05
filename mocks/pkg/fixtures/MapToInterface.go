// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MapToInterface is an autogenerated mock type for the MapToInterface type
type MapToInterface struct {
	mock.Mock
}

// Foo provides a mock function with given fields: arg1
func (_m *MapToInterface) Foo(arg1 ...map[string]interface{}) {
	_va := make([]interface{}, len(arg1))
	for _i := range arg1 {
		_va[_i] = arg1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// NewMapToInterface creates a new instance of MapToInterface. It also registers a cleanup function to assert the mocks expectations.
func NewMapToInterface(t testing.TB) *MapToInterface {
	mock := &MapToInterface{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
