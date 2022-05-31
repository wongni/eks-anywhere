// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/networking/cilium/templater.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHelm is a mock of Helm interface.
type MockHelm struct {
	ctrl     *gomock.Controller
	recorder *MockHelmMockRecorder
}

// MockHelmMockRecorder is the mock recorder for MockHelm.
type MockHelmMockRecorder struct {
	mock *MockHelm
}

// NewMockHelm creates a new mock instance.
func NewMockHelm(ctrl *gomock.Controller) *MockHelm {
	mock := &MockHelm{ctrl: ctrl}
	mock.recorder = &MockHelmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelm) EXPECT() *MockHelmMockRecorder {
	return m.recorder
}

// Template mocks base method.
func (m *MockHelm) Template(ctx context.Context, ociURI, version, namespace string, values interface{}, kubeVersion string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template", ctx, ociURI, version, namespace, values, kubeVersion)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template.
func (mr *MockHelmMockRecorder) Template(ctx, ociURI, version, namespace, values, kubeVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockHelm)(nil).Template), ctx, ociURI, version, namespace, values, kubeVersion)
}
