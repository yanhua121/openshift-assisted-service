// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/restapi (interfaces: ManifestsAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	middleware "github.com/go-openapi/runtime/middleware"
	gomock "github.com/golang/mock/gomock"
	manifests "github.com/openshift/assisted-service/restapi/operations/manifests"
)

// MockManifestsAPI is a mock of ManifestsAPI interface.
type MockManifestsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockManifestsAPIMockRecorder
}

// MockManifestsAPIMockRecorder is the mock recorder for MockManifestsAPI.
type MockManifestsAPIMockRecorder struct {
	mock *MockManifestsAPI
}

// NewMockManifestsAPI creates a new mock instance.
func NewMockManifestsAPI(ctrl *gomock.Controller) *MockManifestsAPI {
	mock := &MockManifestsAPI{ctrl: ctrl}
	mock.recorder = &MockManifestsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManifestsAPI) EXPECT() *MockManifestsAPIMockRecorder {
	return m.recorder
}

// V2CreateClusterManifest mocks base method.
func (m *MockManifestsAPI) V2CreateClusterManifest(arg0 context.Context, arg1 manifests.V2CreateClusterManifestParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2CreateClusterManifest", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2CreateClusterManifest indicates an expected call of V2CreateClusterManifest.
func (mr *MockManifestsAPIMockRecorder) V2CreateClusterManifest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2CreateClusterManifest", reflect.TypeOf((*MockManifestsAPI)(nil).V2CreateClusterManifest), arg0, arg1)
}

// V2DeleteClusterManifest mocks base method.
func (m *MockManifestsAPI) V2DeleteClusterManifest(arg0 context.Context, arg1 manifests.V2DeleteClusterManifestParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DeleteClusterManifest", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2DeleteClusterManifest indicates an expected call of V2DeleteClusterManifest.
func (mr *MockManifestsAPIMockRecorder) V2DeleteClusterManifest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DeleteClusterManifest", reflect.TypeOf((*MockManifestsAPI)(nil).V2DeleteClusterManifest), arg0, arg1)
}

// V2DownloadClusterManifest mocks base method.
func (m *MockManifestsAPI) V2DownloadClusterManifest(arg0 context.Context, arg1 manifests.V2DownloadClusterManifestParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DownloadClusterManifest", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2DownloadClusterManifest indicates an expected call of V2DownloadClusterManifest.
func (mr *MockManifestsAPIMockRecorder) V2DownloadClusterManifest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DownloadClusterManifest", reflect.TypeOf((*MockManifestsAPI)(nil).V2DownloadClusterManifest), arg0, arg1)
}

// V2ListClusterManifests mocks base method.
func (m *MockManifestsAPI) V2ListClusterManifests(arg0 context.Context, arg1 manifests.V2ListClusterManifestsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2ListClusterManifests", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2ListClusterManifests indicates an expected call of V2ListClusterManifests.
func (mr *MockManifestsAPIMockRecorder) V2ListClusterManifests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2ListClusterManifests", reflect.TypeOf((*MockManifestsAPI)(nil).V2ListClusterManifests), arg0, arg1)
}

// V2UpdateClusterManifest mocks base method.
func (m *MockManifestsAPI) V2UpdateClusterManifest(arg0 context.Context, arg1 manifests.V2UpdateClusterManifestParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateClusterManifest", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UpdateClusterManifest indicates an expected call of V2UpdateClusterManifest.
func (mr *MockManifestsAPIMockRecorder) V2UpdateClusterManifest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateClusterManifest", reflect.TypeOf((*MockManifestsAPI)(nil).V2UpdateClusterManifest), arg0, arg1)
}
