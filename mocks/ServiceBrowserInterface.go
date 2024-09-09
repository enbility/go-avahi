// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	avahi "github.com/enbility/go-avahi"
	dbus "github.com/godbus/dbus/v5"

	mock "github.com/stretchr/testify/mock"
)

// ServiceBrowserInterface is an autogenerated mock type for the ServiceBrowserInterface type
type ServiceBrowserInterface struct {
	mock.Mock
}

type ServiceBrowserInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceBrowserInterface) EXPECT() *ServiceBrowserInterface_Expecter {
	return &ServiceBrowserInterface_Expecter{mock: &_m.Mock}
}

// AddChannel provides a mock function with given fields:
func (_m *ServiceBrowserInterface) AddChannel() chan avahi.Service {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AddChannel")
	}

	var r0 chan avahi.Service
	if rf, ok := ret.Get(0).(func() chan avahi.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan avahi.Service)
		}
	}

	return r0
}

// ServiceBrowserInterface_AddChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddChannel'
type ServiceBrowserInterface_AddChannel_Call struct {
	*mock.Call
}

// AddChannel is a helper method to define mock.On call
func (_e *ServiceBrowserInterface_Expecter) AddChannel() *ServiceBrowserInterface_AddChannel_Call {
	return &ServiceBrowserInterface_AddChannel_Call{Call: _e.mock.On("AddChannel")}
}

func (_c *ServiceBrowserInterface_AddChannel_Call) Run(run func()) *ServiceBrowserInterface_AddChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServiceBrowserInterface_AddChannel_Call) Return(_a0 chan avahi.Service) *ServiceBrowserInterface_AddChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceBrowserInterface_AddChannel_Call) RunAndReturn(run func() chan avahi.Service) *ServiceBrowserInterface_AddChannel_Call {
	_c.Call.Return(run)
	return _c
}

// DispatchSignal provides a mock function with given fields: signal
func (_m *ServiceBrowserInterface) DispatchSignal(signal *dbus.Signal) error {
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

// ServiceBrowserInterface_DispatchSignal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DispatchSignal'
type ServiceBrowserInterface_DispatchSignal_Call struct {
	*mock.Call
}

// DispatchSignal is a helper method to define mock.On call
//   - signal *dbus.Signal
func (_e *ServiceBrowserInterface_Expecter) DispatchSignal(signal interface{}) *ServiceBrowserInterface_DispatchSignal_Call {
	return &ServiceBrowserInterface_DispatchSignal_Call{Call: _e.mock.On("DispatchSignal", signal)}
}

func (_c *ServiceBrowserInterface_DispatchSignal_Call) Run(run func(signal *dbus.Signal)) *ServiceBrowserInterface_DispatchSignal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*dbus.Signal))
	})
	return _c
}

func (_c *ServiceBrowserInterface_DispatchSignal_Call) Return(_a0 error) *ServiceBrowserInterface_DispatchSignal_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceBrowserInterface_DispatchSignal_Call) RunAndReturn(run func(*dbus.Signal) error) *ServiceBrowserInterface_DispatchSignal_Call {
	_c.Call.Return(run)
	return _c
}

// Free provides a mock function with given fields:
func (_m *ServiceBrowserInterface) Free() {
	_m.Called()
}

// ServiceBrowserInterface_Free_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Free'
type ServiceBrowserInterface_Free_Call struct {
	*mock.Call
}

// Free is a helper method to define mock.On call
func (_e *ServiceBrowserInterface_Expecter) Free() *ServiceBrowserInterface_Free_Call {
	return &ServiceBrowserInterface_Free_Call{Call: _e.mock.On("Free")}
}

func (_c *ServiceBrowserInterface_Free_Call) Run(run func()) *ServiceBrowserInterface_Free_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServiceBrowserInterface_Free_Call) Return() *ServiceBrowserInterface_Free_Call {
	_c.Call.Return()
	return _c
}

func (_c *ServiceBrowserInterface_Free_Call) RunAndReturn(run func()) *ServiceBrowserInterface_Free_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectPath provides a mock function with given fields:
func (_m *ServiceBrowserInterface) GetObjectPath() dbus.ObjectPath {
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

// ServiceBrowserInterface_GetObjectPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectPath'
type ServiceBrowserInterface_GetObjectPath_Call struct {
	*mock.Call
}

// GetObjectPath is a helper method to define mock.On call
func (_e *ServiceBrowserInterface_Expecter) GetObjectPath() *ServiceBrowserInterface_GetObjectPath_Call {
	return &ServiceBrowserInterface_GetObjectPath_Call{Call: _e.mock.On("GetObjectPath")}
}

func (_c *ServiceBrowserInterface_GetObjectPath_Call) Run(run func()) *ServiceBrowserInterface_GetObjectPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServiceBrowserInterface_GetObjectPath_Call) Return(_a0 dbus.ObjectPath) *ServiceBrowserInterface_GetObjectPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceBrowserInterface_GetObjectPath_Call) RunAndReturn(run func() dbus.ObjectPath) *ServiceBrowserInterface_GetObjectPath_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveChannel provides a mock function with given fields:
func (_m *ServiceBrowserInterface) RemoveChannel() chan avahi.Service {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoveChannel")
	}

	var r0 chan avahi.Service
	if rf, ok := ret.Get(0).(func() chan avahi.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan avahi.Service)
		}
	}

	return r0
}

// ServiceBrowserInterface_RemoveChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveChannel'
type ServiceBrowserInterface_RemoveChannel_Call struct {
	*mock.Call
}

// RemoveChannel is a helper method to define mock.On call
func (_e *ServiceBrowserInterface_Expecter) RemoveChannel() *ServiceBrowserInterface_RemoveChannel_Call {
	return &ServiceBrowserInterface_RemoveChannel_Call{Call: _e.mock.On("RemoveChannel")}
}

func (_c *ServiceBrowserInterface_RemoveChannel_Call) Run(run func()) *ServiceBrowserInterface_RemoveChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServiceBrowserInterface_RemoveChannel_Call) Return(_a0 chan avahi.Service) *ServiceBrowserInterface_RemoveChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceBrowserInterface_RemoveChannel_Call) RunAndReturn(run func() chan avahi.Service) *ServiceBrowserInterface_RemoveChannel_Call {
	_c.Call.Return(run)
	return _c
}

// NewServiceBrowserInterface creates a new instance of ServiceBrowserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceBrowserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceBrowserInterface {
	mock := &ServiceBrowserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
