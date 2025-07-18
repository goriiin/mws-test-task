// Code generated by MockGen. DO NOT EDIT.
// Source: root.go
//
// Generated by this command:
//
//	mockgen -source=root.go -destination=../cmd/mock/mock_repo.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	domain "mws/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProfileRepo is a mock of ProfileRepo interface.
type MockProfileRepo struct {
	ctrl     *gomock.Controller
	recorder *MockProfileRepoMockRecorder
	isgomock struct{}
}

// MockProfileRepoMockRecorder is the mock recorder for MockProfileRepo.
type MockProfileRepoMockRecorder struct {
	mock *MockProfileRepo
}

// NewMockProfileRepo creates a new mock instance.
func NewMockProfileRepo(ctrl *gomock.Controller) *MockProfileRepo {
	mock := &MockProfileRepo{ctrl: ctrl}
	mock.recorder = &MockProfileRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileRepo) EXPECT() *MockProfileRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProfileRepo) Create(profile domain.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", profile)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockProfileRepoMockRecorder) Create(profile any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProfileRepo)(nil).Create), profile)
}

// Delete mocks base method.
func (m *MockProfileRepo) Delete(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProfileRepoMockRecorder) Delete(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProfileRepo)(nil).Delete), name)
}

// Get mocks base method.
func (m *MockProfileRepo) Get(name string) (domain.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(domain.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProfileRepoMockRecorder) Get(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProfileRepo)(nil).Get), name)
}

// List mocks base method.
func (m *MockProfileRepo) List() ([]domain.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]domain.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProfileRepoMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProfileRepo)(nil).List))
}
