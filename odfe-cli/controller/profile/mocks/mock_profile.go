// Code generated by MockGen. DO NOT EDIT.
// Source: es-cli/odfe-cli/controller/profile (interfaces: Controller)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "es-cli/odfe-cli/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockController is a mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// CreateProfile mocks base method
func (m *MockController) CreateProfile(arg0 entity.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProfile indicates an expected call of CreateProfile
func (mr *MockControllerMockRecorder) CreateProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockController)(nil).CreateProfile), arg0)
}

// DeleteProfiles mocks base method
func (m *MockController) DeleteProfiles(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfiles", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfiles indicates an expected call of DeleteProfiles
func (mr *MockControllerMockRecorder) DeleteProfiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfiles", reflect.TypeOf((*MockController)(nil).DeleteProfiles), arg0)
}

// GetProfileForExecution mocks base method
func (m *MockController) GetProfileForExecution(arg0 string) (entity.Profile, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileForExecution", arg0)
	ret0, _ := ret[0].(entity.Profile)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProfileForExecution indicates an expected call of GetProfileForExecution
func (mr *MockControllerMockRecorder) GetProfileForExecution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileForExecution", reflect.TypeOf((*MockController)(nil).GetProfileForExecution), arg0)
}

// GetProfileNames mocks base method
func (m *MockController) GetProfileNames() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileNames indicates an expected call of GetProfileNames
func (mr *MockControllerMockRecorder) GetProfileNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileNames", reflect.TypeOf((*MockController)(nil).GetProfileNames))
}

// GetProfiles mocks base method
func (m *MockController) GetProfiles() ([]entity.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfiles")
	ret0, _ := ret[0].([]entity.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfiles indicates an expected call of GetProfiles
func (mr *MockControllerMockRecorder) GetProfiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfiles", reflect.TypeOf((*MockController)(nil).GetProfiles))
}

// GetProfilesMap mocks base method
func (m *MockController) GetProfilesMap() (map[string]entity.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfilesMap")
	ret0, _ := ret[0].(map[string]entity.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfilesMap indicates an expected call of GetProfilesMap
func (mr *MockControllerMockRecorder) GetProfilesMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfilesMap", reflect.TypeOf((*MockController)(nil).GetProfilesMap))
}
