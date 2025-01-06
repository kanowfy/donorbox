// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/kanowfy/donorbox/internal/dto"
	goth "github.com/markbates/goth"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kanowfy/donorbox/internal/model"
)

// Auth is an autogenerated mock type for the Auth type
type Auth struct {
	mock.Mock
}

type Auth_Expecter struct {
	mock *mock.Mock
}

func (_m *Auth) EXPECT() *Auth_Expecter {
	return &Auth_Expecter{mock: &_m.Mock}
}

// ActivateAccount provides a mock function with given fields: ctx, activationToken
func (_m *Auth) ActivateAccount(ctx context.Context, activationToken string) error {
	ret := _m.Called(ctx, activationToken)

	if len(ret) == 0 {
		panic("no return value specified for ActivateAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, activationToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Auth_ActivateAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ActivateAccount'
type Auth_ActivateAccount_Call struct {
	*mock.Call
}

// ActivateAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - activationToken string
func (_e *Auth_Expecter) ActivateAccount(ctx interface{}, activationToken interface{}) *Auth_ActivateAccount_Call {
	return &Auth_ActivateAccount_Call{Call: _e.mock.On("ActivateAccount", ctx, activationToken)}
}

func (_c *Auth_ActivateAccount_Call) Run(run func(ctx context.Context, activationToken string)) *Auth_ActivateAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Auth_ActivateAccount_Call) Return(_a0 error) *Auth_ActivateAccount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Auth_ActivateAccount_Call) RunAndReturn(run func(context.Context, string) error) *Auth_ActivateAccount_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: ctx, request
func (_m *Auth) Login(ctx context.Context, request dto.LoginRequest) (string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) (string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) string); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.LoginRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Auth_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type Auth_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - ctx context.Context
//   - request dto.LoginRequest
func (_e *Auth_Expecter) Login(ctx interface{}, request interface{}) *Auth_Login_Call {
	return &Auth_Login_Call{Call: _e.mock.On("Login", ctx, request)}
}

func (_c *Auth_Login_Call) Run(run func(ctx context.Context, request dto.LoginRequest)) *Auth_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dto.LoginRequest))
	})
	return _c
}

func (_c *Auth_Login_Call) Return(_a0 string, _a1 error) *Auth_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Auth_Login_Call) RunAndReturn(run func(context.Context, dto.LoginRequest) (string, error)) *Auth_Login_Call {
	_c.Call.Return(run)
	return _c
}

// LoginOAuth provides a mock function with given fields: ctx, oauthUser
func (_m *Auth) LoginOAuth(ctx context.Context, oauthUser goth.User) (string, error) {
	ret := _m.Called(ctx, oauthUser)

	if len(ret) == 0 {
		panic("no return value specified for LoginOAuth")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, goth.User) (string, error)); ok {
		return rf(ctx, oauthUser)
	}
	if rf, ok := ret.Get(0).(func(context.Context, goth.User) string); ok {
		r0 = rf(ctx, oauthUser)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, goth.User) error); ok {
		r1 = rf(ctx, oauthUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Auth_LoginOAuth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoginOAuth'
type Auth_LoginOAuth_Call struct {
	*mock.Call
}

// LoginOAuth is a helper method to define mock.On call
//   - ctx context.Context
//   - oauthUser goth.User
func (_e *Auth_Expecter) LoginOAuth(ctx interface{}, oauthUser interface{}) *Auth_LoginOAuth_Call {
	return &Auth_LoginOAuth_Call{Call: _e.mock.On("LoginOAuth", ctx, oauthUser)}
}

func (_c *Auth_LoginOAuth_Call) Run(run func(ctx context.Context, oauthUser goth.User)) *Auth_LoginOAuth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(goth.User))
	})
	return _c
}

func (_c *Auth_LoginOAuth_Call) Return(_a0 string, _a1 error) *Auth_LoginOAuth_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Auth_LoginOAuth_Call) RunAndReturn(run func(context.Context, goth.User) (string, error)) *Auth_LoginOAuth_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: ctx, request, hostPath
func (_m *Auth) Register(ctx context.Context, request dto.UserRegisterRequest, hostPath string) (*model.User, error) {
	ret := _m.Called(ctx, request, hostPath)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserRegisterRequest, string) (*model.User, error)); ok {
		return rf(ctx, request, hostPath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserRegisterRequest, string) *model.User); ok {
		r0 = rf(ctx, request, hostPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.UserRegisterRequest, string) error); ok {
		r1 = rf(ctx, request, hostPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Auth_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type Auth_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - ctx context.Context
//   - request dto.UserRegisterRequest
//   - hostPath string
func (_e *Auth_Expecter) Register(ctx interface{}, request interface{}, hostPath interface{}) *Auth_Register_Call {
	return &Auth_Register_Call{Call: _e.mock.On("Register", ctx, request, hostPath)}
}

func (_c *Auth_Register_Call) Run(run func(ctx context.Context, request dto.UserRegisterRequest, hostPath string)) *Auth_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dto.UserRegisterRequest), args[2].(string))
	})
	return _c
}

func (_c *Auth_Register_Call) Return(_a0 *model.User, _a1 error) *Auth_Register_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Auth_Register_Call) RunAndReturn(run func(context.Context, dto.UserRegisterRequest, string) (*model.User, error)) *Auth_Register_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterEscrow provides a mock function with given fields: ctx, req
func (_m *Auth) RegisterEscrow(ctx context.Context, req dto.EscrowRegisterRequest) (*model.EscrowUser, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for RegisterEscrow")
	}

	var r0 *model.EscrowUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.EscrowRegisterRequest) (*model.EscrowUser, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.EscrowRegisterRequest) *model.EscrowUser); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.EscrowUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.EscrowRegisterRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Auth_RegisterEscrow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterEscrow'
type Auth_RegisterEscrow_Call struct {
	*mock.Call
}

// RegisterEscrow is a helper method to define mock.On call
//   - ctx context.Context
//   - req dto.EscrowRegisterRequest
func (_e *Auth_Expecter) RegisterEscrow(ctx interface{}, req interface{}) *Auth_RegisterEscrow_Call {
	return &Auth_RegisterEscrow_Call{Call: _e.mock.On("RegisterEscrow", ctx, req)}
}

func (_c *Auth_RegisterEscrow_Call) Run(run func(ctx context.Context, req dto.EscrowRegisterRequest)) *Auth_RegisterEscrow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dto.EscrowRegisterRequest))
	})
	return _c
}

func (_c *Auth_RegisterEscrow_Call) Return(_a0 *model.EscrowUser, _a1 error) *Auth_RegisterEscrow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Auth_RegisterEscrow_Call) RunAndReturn(run func(context.Context, dto.EscrowRegisterRequest) (*model.EscrowUser, error)) *Auth_RegisterEscrow_Call {
	_c.Call.Return(run)
	return _c
}

// ResetPassword provides a mock function with given fields: ctx, request
func (_m *Auth) ResetPassword(ctx context.Context, request dto.ResetPasswordRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ResetPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.ResetPasswordRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Auth_ResetPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResetPassword'
type Auth_ResetPassword_Call struct {
	*mock.Call
}

// ResetPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - request dto.ResetPasswordRequest
func (_e *Auth_Expecter) ResetPassword(ctx interface{}, request interface{}) *Auth_ResetPassword_Call {
	return &Auth_ResetPassword_Call{Call: _e.mock.On("ResetPassword", ctx, request)}
}

func (_c *Auth_ResetPassword_Call) Run(run func(ctx context.Context, request dto.ResetPasswordRequest)) *Auth_ResetPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dto.ResetPasswordRequest))
	})
	return _c
}

func (_c *Auth_ResetPassword_Call) Return(_a0 error) *Auth_ResetPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Auth_ResetPassword_Call) RunAndReturn(run func(context.Context, dto.ResetPasswordRequest) error) *Auth_ResetPassword_Call {
	_c.Call.Return(run)
	return _c
}

// SendResetPasswordToken provides a mock function with given fields: ctx, email, hostPath
func (_m *Auth) SendResetPasswordToken(ctx context.Context, email string, hostPath string) error {
	ret := _m.Called(ctx, email, hostPath)

	if len(ret) == 0 {
		panic("no return value specified for SendResetPasswordToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, email, hostPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Auth_SendResetPasswordToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendResetPasswordToken'
type Auth_SendResetPasswordToken_Call struct {
	*mock.Call
}

// SendResetPasswordToken is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
//   - hostPath string
func (_e *Auth_Expecter) SendResetPasswordToken(ctx interface{}, email interface{}, hostPath interface{}) *Auth_SendResetPasswordToken_Call {
	return &Auth_SendResetPasswordToken_Call{Call: _e.mock.On("SendResetPasswordToken", ctx, email, hostPath)}
}

func (_c *Auth_SendResetPasswordToken_Call) Run(run func(ctx context.Context, email string, hostPath string)) *Auth_SendResetPasswordToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Auth_SendResetPasswordToken_Call) Return(_a0 error) *Auth_SendResetPasswordToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Auth_SendResetPasswordToken_Call) RunAndReturn(run func(context.Context, string, string) error) *Auth_SendResetPasswordToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuth creates a new instance of Auth. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuth(t interface {
	mock.TestingT
	Cleanup(func())
}) *Auth {
	mock := &Auth{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
