// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/api (interfaces: OpenStorageNodeServer,OpenStorageClusterServer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
)

// MockOpenStorageNodeServer is a mock of OpenStorageNodeServer interface
type MockOpenStorageNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageNodeServerMockRecorder
}

// MockOpenStorageNodeServerMockRecorder is the mock recorder for MockOpenStorageNodeServer
type MockOpenStorageNodeServerMockRecorder struct {
	mock *MockOpenStorageNodeServer
}

// NewMockOpenStorageNodeServer creates a new mock instance
func NewMockOpenStorageNodeServer(ctrl *gomock.Controller) *MockOpenStorageNodeServer {
	mock := &MockOpenStorageNodeServer{ctrl: ctrl}
	mock.recorder = &MockOpenStorageNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStorageNodeServer) EXPECT() *MockOpenStorageNodeServerMockRecorder {
	return m.recorder
}

// Enumerate mocks base method
func (m *MockOpenStorageNodeServer) Enumerate(arg0 context.Context, arg1 *api.SdkNodeEnumerateRequest) (*api.SdkNodeEnumerateResponse, error) {
	ret := m.ctrl.Call(m, "Enumerate", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkNodeEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enumerate indicates an expected call of Enumerate
func (mr *MockOpenStorageNodeServerMockRecorder) Enumerate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enumerate", reflect.TypeOf((*MockOpenStorageNodeServer)(nil).Enumerate), arg0, arg1)
}

// EnumerateWithFilters mocks base method
func (m *MockOpenStorageNodeServer) EnumerateWithFilters(arg0 context.Context, arg1 *api.SdkNodeEnumerateWithFiltersRequest) (*api.SdkNodeEnumerateWithFiltersResponse, error) {
	ret := m.ctrl.Call(m, "EnumerateWithFilters", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkNodeEnumerateWithFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateWithFilters indicates an expected call of EnumerateWithFilters
func (mr *MockOpenStorageNodeServerMockRecorder) EnumerateWithFilters(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateWithFilters", reflect.TypeOf((*MockOpenStorageNodeServer)(nil).EnumerateWithFilters), arg0, arg1)
}

// Inspect mocks base method
func (m *MockOpenStorageNodeServer) Inspect(arg0 context.Context, arg1 *api.SdkNodeInspectRequest) (*api.SdkNodeInspectResponse, error) {
	ret := m.ctrl.Call(m, "Inspect", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkNodeInspectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect
func (mr *MockOpenStorageNodeServerMockRecorder) Inspect(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockOpenStorageNodeServer)(nil).Inspect), arg0, arg1)
}

// InspectCurrent mocks base method
func (m *MockOpenStorageNodeServer) InspectCurrent(arg0 context.Context, arg1 *api.SdkNodeInspectCurrentRequest) (*api.SdkNodeInspectCurrentResponse, error) {
	ret := m.ctrl.Call(m, "InspectCurrent", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkNodeInspectCurrentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectCurrent indicates an expected call of InspectCurrent
func (mr *MockOpenStorageNodeServerMockRecorder) InspectCurrent(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectCurrent", reflect.TypeOf((*MockOpenStorageNodeServer)(nil).InspectCurrent), arg0, arg1)
}

// MockOpenStorageClusterServer is a mock of OpenStorageClusterServer interface
type MockOpenStorageClusterServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageClusterServerMockRecorder
}

// MockOpenStorageClusterServerMockRecorder is the mock recorder for MockOpenStorageClusterServer
type MockOpenStorageClusterServerMockRecorder struct {
	mock *MockOpenStorageClusterServer
}

// NewMockOpenStorageClusterServer creates a new mock instance
func NewMockOpenStorageClusterServer(ctrl *gomock.Controller) *MockOpenStorageClusterServer {
	mock := &MockOpenStorageClusterServer{ctrl: ctrl}
	mock.recorder = &MockOpenStorageClusterServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStorageClusterServer) EXPECT() *MockOpenStorageClusterServerMockRecorder {
	return m.recorder
}

// InspectCurrent mocks base method
func (m *MockOpenStorageClusterServer) InspectCurrent(arg0 context.Context, arg1 *api.SdkClusterInspectCurrentRequest) (*api.SdkClusterInspectCurrentResponse, error) {
	ret := m.ctrl.Call(m, "InspectCurrent", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkClusterInspectCurrentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectCurrent indicates an expected call of InspectCurrent
func (mr *MockOpenStorageClusterServerMockRecorder) InspectCurrent(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectCurrent", reflect.TypeOf((*MockOpenStorageClusterServer)(nil).InspectCurrent), arg0, arg1)
}
