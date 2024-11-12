// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/kanowfy/donorbox/internal/dto"
	filters "github.com/kanowfy/donorbox/internal/filters"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kanowfy/donorbox/internal/model"

	uuid "github.com/google/uuid"
)

// Project is an autogenerated mock type for the Project type
type Project struct {
	mock.Mock
}

type Project_Expecter struct {
	mock *mock.Mock
}

func (_m *Project) EXPECT() *Project_Expecter {
	return &Project_Expecter{mock: &_m.Mock}
}

// CreateProject provides a mock function with given fields: ctx, userID, req
func (_m *Project) CreateProject(ctx context.Context, userID uuid.UUID, req dto.CreateProjectRequest) (*model.Project, error) {
	ret := _m.Called(ctx, userID, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateProject")
	}

	var r0 *model.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, dto.CreateProjectRequest) (*model.Project, error)); ok {
		return rf(ctx, userID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, dto.CreateProjectRequest) *model.Project); ok {
		r0 = rf(ctx, userID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, dto.CreateProjectRequest) error); ok {
		r1 = rf(ctx, userID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Project_CreateProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProject'
type Project_CreateProject_Call struct {
	*mock.Call
}

// CreateProject is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - req dto.CreateProjectRequest
func (_e *Project_Expecter) CreateProject(ctx interface{}, userID interface{}, req interface{}) *Project_CreateProject_Call {
	return &Project_CreateProject_Call{Call: _e.mock.On("CreateProject", ctx, userID, req)}
}

func (_c *Project_CreateProject_Call) Run(run func(ctx context.Context, userID uuid.UUID, req dto.CreateProjectRequest)) *Project_CreateProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(dto.CreateProjectRequest))
	})
	return _c
}

func (_c *Project_CreateProject_Call) Return(_a0 *model.Project, _a1 error) *Project_CreateProject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Project_CreateProject_Call) RunAndReturn(run func(context.Context, uuid.UUID, dto.CreateProjectRequest) (*model.Project, error)) *Project_CreateProject_Call {
	_c.Call.Return(run)
	return _c
}

// CreateProjectUpdate provides a mock function with given fields: ctx, userID, req
func (_m *Project) CreateProjectUpdate(ctx context.Context, userID uuid.UUID, req dto.CreateProjectUpdateRequest) (*model.ProjectUpdate, error) {
	ret := _m.Called(ctx, userID, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateProjectUpdate")
	}

	var r0 *model.ProjectUpdate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, dto.CreateProjectUpdateRequest) (*model.ProjectUpdate, error)); ok {
		return rf(ctx, userID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, dto.CreateProjectUpdateRequest) *model.ProjectUpdate); ok {
		r0 = rf(ctx, userID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProjectUpdate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, dto.CreateProjectUpdateRequest) error); ok {
		r1 = rf(ctx, userID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Project_CreateProjectUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProjectUpdate'
type Project_CreateProjectUpdate_Call struct {
	*mock.Call
}

// CreateProjectUpdate is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - req dto.CreateProjectUpdateRequest
func (_e *Project_Expecter) CreateProjectUpdate(ctx interface{}, userID interface{}, req interface{}) *Project_CreateProjectUpdate_Call {
	return &Project_CreateProjectUpdate_Call{Call: _e.mock.On("CreateProjectUpdate", ctx, userID, req)}
}

func (_c *Project_CreateProjectUpdate_Call) Run(run func(ctx context.Context, userID uuid.UUID, req dto.CreateProjectUpdateRequest)) *Project_CreateProjectUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(dto.CreateProjectUpdateRequest))
	})
	return _c
}

func (_c *Project_CreateProjectUpdate_Call) Return(_a0 *model.ProjectUpdate, _a1 error) *Project_CreateProjectUpdate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Project_CreateProjectUpdate_Call) RunAndReturn(run func(context.Context, uuid.UUID, dto.CreateProjectUpdateRequest) (*model.ProjectUpdate, error)) *Project_CreateProjectUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProject provides a mock function with given fields: ctx, userID, projectID
func (_m *Project) DeleteProject(ctx context.Context, userID uuid.UUID, projectID uuid.UUID) error {
	ret := _m.Called(ctx, userID, projectID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, userID, projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Project_DeleteProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProject'
type Project_DeleteProject_Call struct {
	*mock.Call
}

// DeleteProject is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - projectID uuid.UUID
func (_e *Project_Expecter) DeleteProject(ctx interface{}, userID interface{}, projectID interface{}) *Project_DeleteProject_Call {
	return &Project_DeleteProject_Call{Call: _e.mock.On("DeleteProject", ctx, userID, projectID)}
}

func (_c *Project_DeleteProject_Call) Run(run func(ctx context.Context, userID uuid.UUID, projectID uuid.UUID)) *Project_DeleteProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(uuid.UUID))
	})
	return _c
}

func (_c *Project_DeleteProject_Call) Return(_a0 error) *Project_DeleteProject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Project_DeleteProject_Call) RunAndReturn(run func(context.Context, uuid.UUID, uuid.UUID) error) *Project_DeleteProject_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllCategories provides a mock function with given fields: ctx
func (_m *Project) GetAllCategories(ctx context.Context) ([]model.Category, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllCategories")
	}

	var r0 []model.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Category, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Category); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Project_GetAllCategories_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllCategories'
type Project_GetAllCategories_Call struct {
	*mock.Call
}

// GetAllCategories is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Project_Expecter) GetAllCategories(ctx interface{}) *Project_GetAllCategories_Call {
	return &Project_GetAllCategories_Call{Call: _e.mock.On("GetAllCategories", ctx)}
}

func (_c *Project_GetAllCategories_Call) Run(run func(ctx context.Context)) *Project_GetAllCategories_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Project_GetAllCategories_Call) Return(_a0 []model.Category, _a1 error) *Project_GetAllCategories_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Project_GetAllCategories_Call) RunAndReturn(run func(context.Context) ([]model.Category, error)) *Project_GetAllCategories_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllProjects provides a mock function with given fields: ctx, pageNum, pageSize, categoryIndex
func (_m *Project) GetAllProjects(ctx context.Context, pageNum int, pageSize int, categoryIndex int) ([]model.Project, filters.Metadata, error) {
	ret := _m.Called(ctx, pageNum, pageSize, categoryIndex)

	if len(ret) == 0 {
		panic("no return value specified for GetAllProjects")
	}

	var r0 []model.Project
	var r1 filters.Metadata
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) ([]model.Project, filters.Metadata, error)); ok {
		return rf(ctx, pageNum, pageSize, categoryIndex)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) []model.Project); ok {
		r0 = rf(ctx, pageNum, pageSize, categoryIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, int) filters.Metadata); ok {
		r1 = rf(ctx, pageNum, pageSize, categoryIndex)
	} else {
		r1 = ret.Get(1).(filters.Metadata)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int, int, int) error); ok {
		r2 = rf(ctx, pageNum, pageSize, categoryIndex)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Project_GetAllProjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllProjects'
type Project_GetAllProjects_Call struct {
	*mock.Call
}

// GetAllProjects is a helper method to define mock.On call
//   - ctx context.Context
//   - pageNum int
//   - pageSize int
//   - categoryIndex int
func (_e *Project_Expecter) GetAllProjects(ctx interface{}, pageNum interface{}, pageSize interface{}, categoryIndex interface{}) *Project_GetAllProjects_Call {
	return &Project_GetAllProjects_Call{Call: _e.mock.On("GetAllProjects", ctx, pageNum, pageSize, categoryIndex)}
}

func (_c *Project_GetAllProjects_Call) Run(run func(ctx context.Context, pageNum int, pageSize int, categoryIndex int)) *Project_GetAllProjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *Project_GetAllProjects_Call) Return(_a0 []model.Project, _a1 filters.Metadata, _a2 error) *Project_GetAllProjects_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *Project_GetAllProjects_Call) RunAndReturn(run func(context.Context, int, int, int) ([]model.Project, filters.Metadata, error)) *Project_GetAllProjects_Call {
	_c.Call.Return(run)
	return _c
}

// GetEndedProjects provides a mock function with given fields: ctx
func (_m *Project) GetEndedProjects(ctx context.Context) ([]model.Project, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetEndedProjects")
	}

	var r0 []model.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Project, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Project); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Project_GetEndedProjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEndedProjects'
type Project_GetEndedProjects_Call struct {
	*mock.Call
}

// GetEndedProjects is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Project_Expecter) GetEndedProjects(ctx interface{}) *Project_GetEndedProjects_Call {
	return &Project_GetEndedProjects_Call{Call: _e.mock.On("GetEndedProjects", ctx)}
}

func (_c *Project_GetEndedProjects_Call) Run(run func(ctx context.Context)) *Project_GetEndedProjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Project_GetEndedProjects_Call) Return(_a0 []model.Project, _a1 error) *Project_GetEndedProjects_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Project_GetEndedProjects_Call) RunAndReturn(run func(context.Context) ([]model.Project, error)) *Project_GetEndedProjects_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectDetails provides a mock function with given fields: ctx, projectID
func (_m *Project) GetProjectDetails(ctx context.Context, projectID uuid.UUID) (*model.Project, []model.Backing, []model.ProjectUpdate, *model.User, error) {
	ret := _m.Called(ctx, projectID)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectDetails")
	}

	var r0 *model.Project
	var r1 []model.Backing
	var r2 []model.ProjectUpdate
	var r3 *model.User
	var r4 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*model.Project, []model.Backing, []model.ProjectUpdate, *model.User, error)); ok {
		return rf(ctx, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.Project); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) []model.Backing); ok {
		r1 = rf(ctx, projectID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]model.Backing)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, uuid.UUID) []model.ProjectUpdate); ok {
		r2 = rf(ctx, projectID)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]model.ProjectUpdate)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, uuid.UUID) *model.User); ok {
		r3 = rf(ctx, projectID)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(*model.User)
		}
	}

	if rf, ok := ret.Get(4).(func(context.Context, uuid.UUID) error); ok {
		r4 = rf(ctx, projectID)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// Project_GetProjectDetails_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectDetails'
type Project_GetProjectDetails_Call struct {
	*mock.Call
}

// GetProjectDetails is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID uuid.UUID
func (_e *Project_Expecter) GetProjectDetails(ctx interface{}, projectID interface{}) *Project_GetProjectDetails_Call {
	return &Project_GetProjectDetails_Call{Call: _e.mock.On("GetProjectDetails", ctx, projectID)}
}

func (_c *Project_GetProjectDetails_Call) Run(run func(ctx context.Context, projectID uuid.UUID)) *Project_GetProjectDetails_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Project_GetProjectDetails_Call) Return(_a0 *model.Project, _a1 []model.Backing, _a2 []model.ProjectUpdate, _a3 *model.User, _a4 error) *Project_GetProjectDetails_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3, _a4)
	return _c
}

func (_c *Project_GetProjectDetails_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*model.Project, []model.Backing, []model.ProjectUpdate, *model.User, error)) *Project_GetProjectDetails_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectUpdates provides a mock function with given fields: ctx, projectID
func (_m *Project) GetProjectUpdates(ctx context.Context, projectID uuid.UUID) ([]model.ProjectUpdate, error) {
	ret := _m.Called(ctx, projectID)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectUpdates")
	}

	var r0 []model.ProjectUpdate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]model.ProjectUpdate, error)); ok {
		return rf(ctx, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []model.ProjectUpdate); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ProjectUpdate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Project_GetProjectUpdates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectUpdates'
type Project_GetProjectUpdates_Call struct {
	*mock.Call
}

// GetProjectUpdates is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID uuid.UUID
func (_e *Project_Expecter) GetProjectUpdates(ctx interface{}, projectID interface{}) *Project_GetProjectUpdates_Call {
	return &Project_GetProjectUpdates_Call{Call: _e.mock.On("GetProjectUpdates", ctx, projectID)}
}

func (_c *Project_GetProjectUpdates_Call) Run(run func(ctx context.Context, projectID uuid.UUID)) *Project_GetProjectUpdates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Project_GetProjectUpdates_Call) Return(_a0 []model.ProjectUpdate, _a1 error) *Project_GetProjectUpdates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Project_GetProjectUpdates_Call) RunAndReturn(run func(context.Context, uuid.UUID) ([]model.ProjectUpdate, error)) *Project_GetProjectUpdates_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectsForUser provides a mock function with given fields: ctx, userID
func (_m *Project) GetProjectsForUser(ctx context.Context, userID uuid.UUID) ([]model.Project, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectsForUser")
	}

	var r0 []model.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]model.Project, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []model.Project); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Project_GetProjectsForUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectsForUser'
type Project_GetProjectsForUser_Call struct {
	*mock.Call
}

// GetProjectsForUser is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
func (_e *Project_Expecter) GetProjectsForUser(ctx interface{}, userID interface{}) *Project_GetProjectsForUser_Call {
	return &Project_GetProjectsForUser_Call{Call: _e.mock.On("GetProjectsForUser", ctx, userID)}
}

func (_c *Project_GetProjectsForUser_Call) Run(run func(ctx context.Context, userID uuid.UUID)) *Project_GetProjectsForUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Project_GetProjectsForUser_Call) Return(_a0 []model.Project, _a1 error) *Project_GetProjectsForUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Project_GetProjectsForUser_Call) RunAndReturn(run func(context.Context, uuid.UUID) ([]model.Project, error)) *Project_GetProjectsForUser_Call {
	_c.Call.Return(run)
	return _c
}

// SearchProjects provides a mock function with given fields: ctx, query, pageNum, pageSize
func (_m *Project) SearchProjects(ctx context.Context, query string, pageNum int, pageSize int) ([]model.Project, filters.Metadata, error) {
	ret := _m.Called(ctx, query, pageNum, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for SearchProjects")
	}

	var r0 []model.Project
	var r1 filters.Metadata
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) ([]model.Project, filters.Metadata, error)); ok {
		return rf(ctx, query, pageNum, pageSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) []model.Project); ok {
		r0 = rf(ctx, query, pageNum, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) filters.Metadata); ok {
		r1 = rf(ctx, query, pageNum, pageSize)
	} else {
		r1 = ret.Get(1).(filters.Metadata)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int, int) error); ok {
		r2 = rf(ctx, query, pageNum, pageSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Project_SearchProjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchProjects'
type Project_SearchProjects_Call struct {
	*mock.Call
}

// SearchProjects is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - pageNum int
//   - pageSize int
func (_e *Project_Expecter) SearchProjects(ctx interface{}, query interface{}, pageNum interface{}, pageSize interface{}) *Project_SearchProjects_Call {
	return &Project_SearchProjects_Call{Call: _e.mock.On("SearchProjects", ctx, query, pageNum, pageSize)}
}

func (_c *Project_SearchProjects_Call) Run(run func(ctx context.Context, query string, pageNum int, pageSize int)) *Project_SearchProjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *Project_SearchProjects_Call) Return(_a0 []model.Project, _a1 filters.Metadata, _a2 error) *Project_SearchProjects_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *Project_SearchProjects_Call) RunAndReturn(run func(context.Context, string, int, int) ([]model.Project, filters.Metadata, error)) *Project_SearchProjects_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateProject provides a mock function with given fields: ctx, userID, projectID, req
func (_m *Project) UpdateProject(ctx context.Context, userID uuid.UUID, projectID uuid.UUID, req dto.UpdateProjectRequest) error {
	ret := _m.Called(ctx, userID, projectID, req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID, dto.UpdateProjectRequest) error); ok {
		r0 = rf(ctx, userID, projectID, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Project_UpdateProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateProject'
type Project_UpdateProject_Call struct {
	*mock.Call
}

// UpdateProject is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - projectID uuid.UUID
//   - req dto.UpdateProjectRequest
func (_e *Project_Expecter) UpdateProject(ctx interface{}, userID interface{}, projectID interface{}, req interface{}) *Project_UpdateProject_Call {
	return &Project_UpdateProject_Call{Call: _e.mock.On("UpdateProject", ctx, userID, projectID, req)}
}

func (_c *Project_UpdateProject_Call) Run(run func(ctx context.Context, userID uuid.UUID, projectID uuid.UUID, req dto.UpdateProjectRequest)) *Project_UpdateProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(uuid.UUID), args[3].(dto.UpdateProjectRequest))
	})
	return _c
}

func (_c *Project_UpdateProject_Call) Return(_a0 error) *Project_UpdateProject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Project_UpdateProject_Call) RunAndReturn(run func(context.Context, uuid.UUID, uuid.UUID, dto.UpdateProjectRequest) error) *Project_UpdateProject_Call {
	_c.Call.Return(run)
	return _c
}

// NewProject creates a new instance of Project. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProject(t interface {
	mock.TestingT
	Cleanup(func())
}) *Project {
	mock := &Project{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
