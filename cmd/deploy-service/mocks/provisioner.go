// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/deploy-service (interfaces: Provisioner)

// Package mocks is a generated GoMock package.
package mocks

import (
	provision "code.cloudfoundry.org/cfdev/provision"
	workspace "code.cloudfoundry.org/cfdev/workspace"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockProvisioner is a mock of Provisioner interface
type MockProvisioner struct {
	ctrl     *gomock.Controller
	recorder *MockProvisionerMockRecorder
}

// MockProvisionerMockRecorder is the mock recorder for MockProvisioner
type MockProvisionerMockRecorder struct {
	mock *MockProvisioner
}

// NewMockProvisioner creates a new mock instance
func NewMockProvisioner(ctrl *gomock.Controller) *MockProvisioner {
	mock := &MockProvisioner{ctrl: ctrl}
	mock.recorder = &MockProvisionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvisioner) EXPECT() *MockProvisionerMockRecorder {
	return m.recorder
}

// DeployServices mocks base method
func (m *MockProvisioner) DeployServices(arg0 provision.UI, arg1 []workspace.Service, arg2 []string) error {
	ret := m.ctrl.Call(m, "DeployServices", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployServices indicates an expected call of DeployServices
func (mr *MockProvisionerMockRecorder) DeployServices(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployServices", reflect.TypeOf((*MockProvisioner)(nil).DeployServices), arg0, arg1, arg2)
}

// GetWhiteListedService mocks base method
func (m *MockProvisioner) GetWhiteListedService(arg0 string, arg1 []workspace.Service) (*workspace.Service, error) {
	ret := m.ctrl.Call(m, "GetWhiteListedService", arg0, arg1)
	ret0, _ := ret[0].(*workspace.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWhiteListedService indicates an expected call of GetWhiteListedService
func (mr *MockProvisionerMockRecorder) GetWhiteListedService(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWhiteListedService", reflect.TypeOf((*MockProvisioner)(nil).GetWhiteListedService), arg0, arg1)
}

// Ping mocks base method
func (m *MockProvisioner) Ping(arg0 time.Duration) error {
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockProvisionerMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockProvisioner)(nil).Ping), arg0)
}
