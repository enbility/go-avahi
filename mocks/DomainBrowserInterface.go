// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	dbus "github.com/godbus/dbus/v5"
	mock "github.com/stretchr/testify/mock"
)

// DomainBrowserInterface is an autogenerated mock type for the DomainBrowserInterface type
type DomainBrowserInterface struct {
	mock.Mock
}

type DomainBrowserInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *DomainBrowserInterface) EXPECT() *DomainBrowserInterface_Expecter {
	return &DomainBrowserInterface_Expecter{mock: &_m.Mock}
}

// DispatchSignal provides a mock function with given fields: signal
func (_m *DomainBrowserInterface) DispatchSignal(signal *dbus.Signal) error {
	ret := _m.Called(signal)

	if len(ret) == 0 {
		panic("no return value specified for DispatchSignal")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbus.Signal) error); ok {
		r0 = rf(signal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DomainBrowserInterface_DispatchSignal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DispatchSignal'
type DomainBrowserInterface_DispatchSignal_Call struct {
	*mock.Call
}

// DispatchSignal is a helper method to define mock.On call
//   - signal *dbus.Signal
func (_e *DomainBrowserInterface_Expecter) DispatchSignal(signal interface{}) *DomainBrowserInterface_DispatchSignal_Call {
	return &DomainBrowserInterface_DispatchSignal_Call{Call: _e.mock.On("DispatchSignal", signal)}
}

func (_c *DomainBrowserInterface_DispatchSignal_Call) Run(run func(signal *dbus.Signal)) *DomainBrowserInterface_DispatchSignal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*dbus.Signal))
	})
	return _c
}

func (_c *DomainBrowserInterface_DispatchSignal_Call) Return(_a0 error) *DomainBrowserInterface_DispatchSignal_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DomainBrowserInterface_DispatchSignal_Call) RunAndReturn(run func(*dbus.Signal) error) *DomainBrowserInterface_DispatchSignal_Call {
	_c.Call.Return(run)
	return _c
}

// Free provides a mock function with given fields:
func (_m *DomainBrowserInterface) Free() {
	_m.Called()
}

// DomainBrowserInterface_Free_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Free'
type DomainBrowserInterface_Free_Call struct {
	*mock.Call
}

// Free is a helper method to define mock.On call
func (_e *DomainBrowserInterface_Expecter) Free() *DomainBrowserInterface_Free_Call {
	return &DomainBrowserInterface_Free_Call{Call: _e.mock.On("Free")}
}

func (_c *DomainBrowserInterface_Free_Call) Run(run func()) *DomainBrowserInterface_Free_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DomainBrowserInterface_Free_Call) Return() *DomainBrowserInterface_Free_Call {
	_c.Call.Return()
	return _c
}

func (_c *DomainBrowserInterface_Free_Call) RunAndReturn(run func()) *DomainBrowserInterface_Free_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectPath provides a mock function with given fields:
func (_m *DomainBrowserInterface) GetObjectPath() dbus.ObjectPath {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetObjectPath")
	}

	var r0 dbus.ObjectPath
	if rf, ok := ret.Get(0).(func() dbus.ObjectPath); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(dbus.ObjectPath)
	}

	return r0
}

// DomainBrowserInterface_GetObjectPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectPath'
type DomainBrowserInterface_GetObjectPath_Call struct {
	*mock.Call
}

// GetObjectPath is a helper method to define mock.On call
func (_e *DomainBrowserInterface_Expecter) GetObjectPath() *DomainBrowserInterface_GetObjectPath_Call {
	return &DomainBrowserInterface_GetObjectPath_Call{Call: _e.mock.On("GetObjectPath")}
}

func (_c *DomainBrowserInterface_GetObjectPath_Call) Run(run func()) *DomainBrowserInterface_GetObjectPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DomainBrowserInterface_GetObjectPath_Call) Return(_a0 dbus.ObjectPath) *DomainBrowserInterface_GetObjectPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DomainBrowserInterface_GetObjectPath_Call) RunAndReturn(run func() dbus.ObjectPath) *DomainBrowserInterface_GetObjectPath_Call {
	_c.Call.Return(run)
	return _c
}

// NewDomainBrowserInterface creates a new instance of DomainBrowserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDomainBrowserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DomainBrowserInterface {
	mock := &DomainBrowserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
