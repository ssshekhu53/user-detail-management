// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=mock_interface.go -package=service
//

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	models "github.com/ssshekhu53/user-detail-management/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUser) Create(arg0 *models.UserRequest) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserMockRecorder) Create(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUser)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockUser) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUser)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockUser) Get() []models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].([]models.User)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockUserMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUser)(nil).Get))
}

// GetByID mocks base method.
func (m *MockUser) GetByID(arg0 int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserMockRecorder) GetByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUser)(nil).GetByID), arg0)
}

// Search mocks base method.
func (m *MockUser) Search(filters *models.Filters) []models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", filters)
	ret0, _ := ret[0].([]models.User)
	return ret0
}

// Search indicates an expected call of Search.
func (mr *MockUserMockRecorder) Search(filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUser)(nil).Search), filters)
}

// Update mocks base method.
func (m *MockUser) Update(arg0 *models.UserUpdateRequest) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserMockRecorder) Update(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUser)(nil).Update), arg0)
}
