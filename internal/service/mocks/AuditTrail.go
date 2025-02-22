// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/kanowfy/donorbox/internal/model"
	mock "github.com/stretchr/testify/mock"

	service "github.com/kanowfy/donorbox/internal/service"
)

// AuditTrail is an autogenerated mock type for the AuditTrail type
type AuditTrail struct {
	mock.Mock
}

type AuditTrail_Expecter struct {
	mock *mock.Mock
}

func (_m *AuditTrail) EXPECT() *AuditTrail_Expecter {
	return &AuditTrail_Expecter{mock: &_m.Mock}
}

// GetAuditHistory provides a mock function with given fields: ctx
func (_m *AuditTrail) GetAuditHistory(ctx context.Context) ([]model.AuditTrail, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAuditHistory")
	}

	var r0 []model.AuditTrail
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.AuditTrail, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.AuditTrail); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AuditTrail)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuditTrail_GetAuditHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuditHistory'
type AuditTrail_GetAuditHistory_Call struct {
	*mock.Call
}

// GetAuditHistory is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AuditTrail_Expecter) GetAuditHistory(ctx interface{}) *AuditTrail_GetAuditHistory_Call {
	return &AuditTrail_GetAuditHistory_Call{Call: _e.mock.On("GetAuditHistory", ctx)}
}

func (_c *AuditTrail_GetAuditHistory_Call) Run(run func(ctx context.Context)) *AuditTrail_GetAuditHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AuditTrail_GetAuditHistory_Call) Return(_a0 []model.AuditTrail, _a1 error) *AuditTrail_GetAuditHistory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuditTrail_GetAuditHistory_Call) RunAndReturn(run func(context.Context) ([]model.AuditTrail, error)) *AuditTrail_GetAuditHistory_Call {
	_c.Call.Return(run)
	return _c
}

// LogAction provides a mock function with given fields: ctx, params
func (_m *AuditTrail) LogAction(ctx context.Context, params service.LogActionParams) error {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for LogAction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, service.LogActionParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuditTrail_LogAction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogAction'
type AuditTrail_LogAction_Call struct {
	*mock.Call
}

// LogAction is a helper method to define mock.On call
//   - ctx context.Context
//   - params service.LogActionParams
func (_e *AuditTrail_Expecter) LogAction(ctx interface{}, params interface{}) *AuditTrail_LogAction_Call {
	return &AuditTrail_LogAction_Call{Call: _e.mock.On("LogAction", ctx, params)}
}

func (_c *AuditTrail_LogAction_Call) Run(run func(ctx context.Context, params service.LogActionParams)) *AuditTrail_LogAction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(service.LogActionParams))
	})
	return _c
}

func (_c *AuditTrail_LogAction_Call) Return(_a0 error) *AuditTrail_LogAction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuditTrail_LogAction_Call) RunAndReturn(run func(context.Context, service.LogActionParams) error) *AuditTrail_LogAction_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuditTrail creates a new instance of AuditTrail. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuditTrail(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuditTrail {
	mock := &AuditTrail{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
