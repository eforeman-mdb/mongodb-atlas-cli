// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: DataFederationLister,DataFederationDescriber,DataFederationCreator,DataFederationUpdater,DataFederationDeleter)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/admin"
)

// MockDataFederationLister is a mock of DataFederationLister interface.
type MockDataFederationLister struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationListerMockRecorder
}

// MockDataFederationListerMockRecorder is the mock recorder for MockDataFederationLister.
type MockDataFederationListerMockRecorder struct {
	mock *MockDataFederationLister
}

// NewMockDataFederationLister creates a new mock instance.
func NewMockDataFederationLister(ctrl *gomock.Controller) *MockDataFederationLister {
	mock := &MockDataFederationLister{ctrl: ctrl}
	mock.recorder = &MockDataFederationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationLister) EXPECT() *MockDataFederationListerMockRecorder {
	return m.recorder
}

// DataFederationList mocks base method.
func (m *MockDataFederationLister) DataFederationList(arg0, arg1 string) ([]admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationList", arg0, arg1)
	ret0, _ := ret[0].([]admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationList indicates an expected call of DataFederationList.
func (mr *MockDataFederationListerMockRecorder) DataFederationList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationList", reflect.TypeOf((*MockDataFederationLister)(nil).DataFederationList), arg0, arg1)
}

// MockDataFederationDescriber is a mock of DataFederationDescriber interface.
type MockDataFederationDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationDescriberMockRecorder
}

// MockDataFederationDescriberMockRecorder is the mock recorder for MockDataFederationDescriber.
type MockDataFederationDescriberMockRecorder struct {
	mock *MockDataFederationDescriber
}

// NewMockDataFederationDescriber creates a new mock instance.
func NewMockDataFederationDescriber(ctrl *gomock.Controller) *MockDataFederationDescriber {
	mock := &MockDataFederationDescriber{ctrl: ctrl}
	mock.recorder = &MockDataFederationDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationDescriber) EXPECT() *MockDataFederationDescriberMockRecorder {
	return m.recorder
}

// DataFederation mocks base method.
func (m *MockDataFederationDescriber) DataFederation(arg0, arg1 string) (*admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederation", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederation indicates an expected call of DataFederation.
func (mr *MockDataFederationDescriberMockRecorder) DataFederation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederation", reflect.TypeOf((*MockDataFederationDescriber)(nil).DataFederation), arg0, arg1)
}

// MockDataFederationCreator is a mock of DataFederationCreator interface.
type MockDataFederationCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationCreatorMockRecorder
}

// MockDataFederationCreatorMockRecorder is the mock recorder for MockDataFederationCreator.
type MockDataFederationCreatorMockRecorder struct {
	mock *MockDataFederationCreator
}

// NewMockDataFederationCreator creates a new mock instance.
func NewMockDataFederationCreator(ctrl *gomock.Controller) *MockDataFederationCreator {
	mock := &MockDataFederationCreator{ctrl: ctrl}
	mock.recorder = &MockDataFederationCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationCreator) EXPECT() *MockDataFederationCreatorMockRecorder {
	return m.recorder
}

// CreateDataFederation mocks base method.
func (m *MockDataFederationCreator) CreateDataFederation(arg0 string, arg1 *admin.DataLakeTenant) (*admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDataFederation", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDataFederation indicates an expected call of CreateDataFederation.
func (mr *MockDataFederationCreatorMockRecorder) CreateDataFederation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDataFederation", reflect.TypeOf((*MockDataFederationCreator)(nil).CreateDataFederation), arg0, arg1)
}

// MockDataFederationUpdater is a mock of DataFederationUpdater interface.
type MockDataFederationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationUpdaterMockRecorder
}

// MockDataFederationUpdaterMockRecorder is the mock recorder for MockDataFederationUpdater.
type MockDataFederationUpdaterMockRecorder struct {
	mock *MockDataFederationUpdater
}

// NewMockDataFederationUpdater creates a new mock instance.
func NewMockDataFederationUpdater(ctrl *gomock.Controller) *MockDataFederationUpdater {
	mock := &MockDataFederationUpdater{ctrl: ctrl}
	mock.recorder = &MockDataFederationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationUpdater) EXPECT() *MockDataFederationUpdaterMockRecorder {
	return m.recorder
}

// UpdateDataFederation mocks base method.
func (m *MockDataFederationUpdater) UpdateDataFederation(arg0, arg1 string, arg2 *admin.DataLakeTenant) (*admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDataFederation", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDataFederation indicates an expected call of UpdateDataFederation.
func (mr *MockDataFederationUpdaterMockRecorder) UpdateDataFederation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDataFederation", reflect.TypeOf((*MockDataFederationUpdater)(nil).UpdateDataFederation), arg0, arg1, arg2)
}

// MockDataFederationDeleter is a mock of DataFederationDeleter interface.
type MockDataFederationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationDeleterMockRecorder
}

// MockDataFederationDeleterMockRecorder is the mock recorder for MockDataFederationDeleter.
type MockDataFederationDeleterMockRecorder struct {
	mock *MockDataFederationDeleter
}

// NewMockDataFederationDeleter creates a new mock instance.
func NewMockDataFederationDeleter(ctrl *gomock.Controller) *MockDataFederationDeleter {
	mock := &MockDataFederationDeleter{ctrl: ctrl}
	mock.recorder = &MockDataFederationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationDeleter) EXPECT() *MockDataFederationDeleterMockRecorder {
	return m.recorder
}

// DeleteDataFederation mocks base method.
func (m *MockDataFederationDeleter) DeleteDataFederation(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataFederation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDataFederation indicates an expected call of DeleteDataFederation.
func (mr *MockDataFederationDeleterMockRecorder) DeleteDataFederation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataFederation", reflect.TypeOf((*MockDataFederationDeleter)(nil).DeleteDataFederation), arg0, arg1)
}
