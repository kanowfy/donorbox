// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/kanowfy/donorbox/internal/dto"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kanowfy/donorbox/internal/model"

	uuid "github.com/google/uuid"
)

// Escrow is an autogenerated mock type for the Escrow type
type Escrow struct {
	mock.Mock
}

type Escrow_Expecter struct {
	mock *mock.Mock
}

func (_m *Escrow) EXPECT() *Escrow_Expecter {
	return &Escrow_Expecter{mock: &_m.Mock}
}

// ApproveOfProject provides a mock function with given fields: ctx, req
func (_m *Escrow) ApproveOfProject(ctx context.Context, req dto.ProjectApprovalRequest) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for ApproveOfProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.ProjectApprovalRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Escrow_ApproveOfProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApproveOfProject'
type Escrow_ApproveOfProject_Call struct {
	*mock.Call
}

// ApproveOfProject is a helper method to define mock.On call
//   - ctx context.Context
//   - req dto.ProjectApprovalRequest
func (_e *Escrow_Expecter) ApproveOfProject(ctx interface{}, req interface{}) *Escrow_ApproveOfProject_Call {
	return &Escrow_ApproveOfProject_Call{Call: _e.mock.On("ApproveOfProject", ctx, req)}
}

func (_c *Escrow_ApproveOfProject_Call) Run(run func(ctx context.Context, req dto.ProjectApprovalRequest)) *Escrow_ApproveOfProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dto.ProjectApprovalRequest))
	})
	return _c
}

func (_c *Escrow_ApproveOfProject_Call) Return(_a0 error) *Escrow_ApproveOfProject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Escrow_ApproveOfProject_Call) RunAndReturn(run func(context.Context, dto.ProjectApprovalRequest) error) *Escrow_ApproveOfProject_Call {
	_c.Call.Return(run)
	return _c
}

// GetEscrowByID provides a mock function with given fields: ctx, id
func (_m *Escrow) GetEscrowByID(ctx context.Context, id uuid.UUID) (*model.EscrowUser, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetEscrowByID")
	}

	var r0 *model.EscrowUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*model.EscrowUser, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.EscrowUser); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.EscrowUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Escrow_GetEscrowByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEscrowByID'
type Escrow_GetEscrowByID_Call struct {
	*mock.Call
}

// GetEscrowByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *Escrow_Expecter) GetEscrowByID(ctx interface{}, id interface{}) *Escrow_GetEscrowByID_Call {
	return &Escrow_GetEscrowByID_Call{Call: _e.mock.On("GetEscrowByID", ctx, id)}
}

func (_c *Escrow_GetEscrowByID_Call) Run(run func(ctx context.Context, id uuid.UUID)) *Escrow_GetEscrowByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Escrow_GetEscrowByID_Call) Return(_a0 *model.EscrowUser, _a1 error) *Escrow_GetEscrowByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Escrow_GetEscrowByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*model.EscrowUser, error)) *Escrow_GetEscrowByID_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: ctx, req
func (_m *Escrow) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) (string, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) string); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.LoginRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Escrow_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type Escrow_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - ctx context.Context
//   - req dto.LoginRequest
func (_e *Escrow_Expecter) Login(ctx interface{}, req interface{}) *Escrow_Login_Call {
	return &Escrow_Login_Call{Call: _e.mock.On("Login", ctx, req)}
}

func (_c *Escrow_Login_Call) Run(run func(ctx context.Context, req dto.LoginRequest)) *Escrow_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dto.LoginRequest))
	})
	return _c
}

func (_c *Escrow_Login_Call) Return(_a0 string, _a1 error) *Escrow_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Escrow_Login_Call) RunAndReturn(run func(context.Context, dto.LoginRequest) (string, error)) *Escrow_Login_Call {
	_c.Call.Return(run)
	return _c
}

// NewEscrow creates a new instance of Escrow. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEscrow(t interface {
	mock.TestingT
	Cleanup(func())
}) *Escrow {
	mock := &Escrow{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
