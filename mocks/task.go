// Code generated by MockGen. DO NOT EDIT.
// Source: task.go
//
// Generated by this command:
//
//	mockgen -source=task.go -destination mocks/task.go
//

// Package mock_taskapi is a generated GoMock package.
package mock_taskapi

import (
	reflect "reflect"

	taskapi "github.com/chuksgpfr/task-api"
	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockTaskService is a mock of TaskService interface.
type MockTaskService struct {
	ctrl     *gomock.Controller
	recorder *MockTaskServiceMockRecorder
}

// MockTaskServiceMockRecorder is the mock recorder for MockTaskService.
type MockTaskServiceMockRecorder struct {
	mock *MockTaskService
}

// NewMockTaskService creates a new mock instance.
func NewMockTaskService(ctrl *gomock.Controller) *MockTaskService {
	mock := &MockTaskService{ctrl: ctrl}
	mock.recorder = &MockTaskServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskService) EXPECT() *MockTaskServiceMockRecorder {
	return m.recorder
}

// CompleteTask mocks base method.
func (m *MockTaskService) CompleteTask(user *taskapi.User, slug string) (*taskapi.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteTask", user, slug)
	ret0, _ := ret[0].(*taskapi.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteTask indicates an expected call of CompleteTask.
func (mr *MockTaskServiceMockRecorder) CompleteTask(user, slug any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTask", reflect.TypeOf((*MockTaskService)(nil).CompleteTask), user, slug)
}

// CreateTask mocks base method.
func (m *MockTaskService) CreateTask(body *taskapi.CreateTask, user *taskapi.User) (*taskapi.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", body, user)
	ret0, _ := ret[0].(*taskapi.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskServiceMockRecorder) CreateTask(body, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskService)(nil).CreateTask), body, user)
}

// GetTask mocks base method.
func (m *MockTaskService) GetTask(slug string, user *taskapi.User) (*taskapi.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", slug, user)
	ret0, _ := ret[0].(*taskapi.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockTaskServiceMockRecorder) GetTask(slug, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockTaskService)(nil).GetTask), slug, user)
}

// GetTasks mocks base method.
func (m *MockTaskService) GetTasks(ctx *gin.Context, user *taskapi.User) ([]*taskapi.Task, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks", ctx, user)
	ret0, _ := ret[0].([]*taskapi.Task)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockTaskServiceMockRecorder) GetTasks(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockTaskService)(nil).GetTasks), ctx, user)
}

// UpdateTask mocks base method.
func (m *MockTaskService) UpdateTask(body *taskapi.CreateTask, user *taskapi.User, slug string) (*taskapi.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", body, user, slug)
	ret0, _ := ret[0].(*taskapi.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockTaskServiceMockRecorder) UpdateTask(body, user, slug any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTaskService)(nil).UpdateTask), body, user, slug)
}
