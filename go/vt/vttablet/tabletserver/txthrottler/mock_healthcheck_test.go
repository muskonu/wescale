/*
Copyright ApeCloud, Inc.
Licensed under the Apache v2(found in the LICENSE file in the root directory).
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: vitess.io/vitess/go/vt/discovery (interfaces: HealthCheck)

// Package txthrottler is a generated GoMock package.
package txthrottler

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	discovery "vitess.io/vitess/go/vt/discovery"
	query "vitess.io/vitess/go/vt/proto/query"
	topodata "vitess.io/vitess/go/vt/proto/topodata"
	queryservice "vitess.io/vitess/go/vt/vttablet/queryservice"
)

// MockHealthCheck is a mock of HealthCheck interface.
type MockHealthCheck struct {
	ctrl     *gomock.Controller
	recorder *MockHealthCheckMockRecorder
}

func (m *MockHealthCheck) GetAllHealthyTabletStats() []*discovery.TabletHealth {
	return nil
}

// MockHealthCheckMockRecorder is the mock recorder for MockHealthCheck.
type MockHealthCheckMockRecorder struct {
	mock *MockHealthCheck
}

// NewMockHealthCheck creates a new mock instance.
func NewMockHealthCheck(ctrl *gomock.Controller) *MockHealthCheck {
	mock := &MockHealthCheck{ctrl: ctrl}
	mock.recorder = &MockHealthCheckMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthCheck) EXPECT() *MockHealthCheckMockRecorder {
	return m.recorder
}

// AddTablet mocks base method.
func (m *MockHealthCheck) AddTablet(arg0 *topodata.Tablet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTablet", arg0)
}

// AddTablet indicates an expected call of AddTablet.
func (mr *MockHealthCheckMockRecorder) AddTablet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTablet", reflect.TypeOf((*MockHealthCheck)(nil).AddTablet), arg0)
}

// CacheStatus mocks base method.
func (m *MockHealthCheck) CacheStatus() discovery.TabletsCacheStatusList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CacheStatus")
	ret0, _ := ret[0].(discovery.TabletsCacheStatusList)
	return ret0
}

// CacheStatus indicates an expected call of CacheStatus.
func (mr *MockHealthCheckMockRecorder) CacheStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheStatus", reflect.TypeOf((*MockHealthCheck)(nil).CacheStatus))
}

// CacheStatusMap mocks base method.
func (m *MockHealthCheck) CacheStatusMap() map[string]*discovery.TabletsCacheStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CacheStatusMap")
	ret0, _ := ret[0].(map[string]*discovery.TabletsCacheStatus)
	return ret0
}

// CacheStatusMap indicates an expected call of CacheStatusMap.
func (mr *MockHealthCheckMockRecorder) CacheStatusMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheStatusMap", reflect.TypeOf((*MockHealthCheck)(nil).CacheStatusMap))
}

// Close mocks base method.
func (m *MockHealthCheck) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockHealthCheckMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHealthCheck)(nil).Close))
}

// GetAllHealthyTabletStats mocks base method.
func (m *MockHealthCheck) GetAllHealthyTabletStats() []*discovery.TabletHealth {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllHealthyTabletStats")
	ret0, _ := ret[0].([]*discovery.TabletHealth)
	return ret0
}

// GetAllHealthyTabletStats indicates an expected call of GetAllHealthyTabletStats.
func (mr *MockHealthCheckMockRecorder) GetAllHealthyTabletStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllHealthyTabletStats", reflect.TypeOf((*MockHealthCheck)(nil).GetAllHealthyTabletStats))
}

// GetHealthyTabletStats mocks base method.
func (m *MockHealthCheck) GetHealthyTabletStats(arg0 *query.Target) []*discovery.TabletHealth {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHealthyTabletStats", arg0)
	ret0, _ := ret[0].([]*discovery.TabletHealth)
	return ret0
}

// GetHealthyTabletStats indicates an expected call of GetHealthyTabletStats.
func (mr *MockHealthCheckMockRecorder) GetHealthyTabletStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHealthyTabletStats", reflect.TypeOf((*MockHealthCheck)(nil).GetHealthyTabletStats), arg0)
}

// GetTabletHealth mocks base method.
func (m *MockHealthCheck) GetTabletHealth(arg0 discovery.KeyspaceShardTabletType, arg1 *topodata.TabletAlias) (*discovery.TabletHealth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTabletHealth", arg0, arg1)
	ret0, _ := ret[0].(*discovery.TabletHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTabletHealth indicates an expected call of GetTabletHealth.
func (mr *MockHealthCheckMockRecorder) GetTabletHealth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTabletHealth", reflect.TypeOf((*MockHealthCheck)(nil).GetTabletHealth), arg0, arg1)
}

// GetTabletHealthByAlias mocks base method.
func (m *MockHealthCheck) GetTabletHealthByAlias(arg0 *topodata.TabletAlias) (*discovery.TabletHealth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTabletHealthByAlias", arg0)
	ret0, _ := ret[0].(*discovery.TabletHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTabletHealthByAlias indicates an expected call of GetTabletHealthByAlias.
func (mr *MockHealthCheckMockRecorder) GetTabletHealthByAlias(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTabletHealthByAlias", reflect.TypeOf((*MockHealthCheck)(nil).GetTabletHealthByAlias), arg0)
}

// RegisterStats mocks base method.
func (m *MockHealthCheck) RegisterStats() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterStats")
}

// RegisterStats indicates an expected call of RegisterStats.
func (mr *MockHealthCheckMockRecorder) RegisterStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterStats", reflect.TypeOf((*MockHealthCheck)(nil).RegisterStats))
}

// RemoveTablet mocks base method.
func (m *MockHealthCheck) RemoveTablet(arg0 *topodata.Tablet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveTablet", arg0)
}

// RemoveTablet indicates an expected call of RemoveTablet.
func (mr *MockHealthCheckMockRecorder) RemoveTablet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTablet", reflect.TypeOf((*MockHealthCheck)(nil).RemoveTablet), arg0)
}

// ReplaceTablet mocks base method.
func (m *MockHealthCheck) ReplaceTablet(arg0, arg1 *topodata.Tablet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReplaceTablet", arg0, arg1)
}

// ReplaceTablet indicates an expected call of ReplaceTablet.
func (mr *MockHealthCheckMockRecorder) ReplaceTablet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceTablet", reflect.TypeOf((*MockHealthCheck)(nil).ReplaceTablet), arg0, arg1)
}

// Subscribe mocks base method.
func (m *MockHealthCheck) Subscribe() chan *discovery.TabletHealth {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe")
	ret0, _ := ret[0].(chan *discovery.TabletHealth)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockHealthCheckMockRecorder) Subscribe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockHealthCheck)(nil).Subscribe))
}

// TabletConnection mocks base method.
func (m *MockHealthCheck) TabletConnection(arg0 *topodata.TabletAlias, arg1 *query.Target) (queryservice.QueryService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TabletConnection", arg0, arg1)
	ret0, _ := ret[0].(queryservice.QueryService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TabletConnection indicates an expected call of TabletConnection.
func (mr *MockHealthCheckMockRecorder) TabletConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TabletConnection", reflect.TypeOf((*MockHealthCheck)(nil).TabletConnection), arg0, arg1)
}

// Unsubscribe mocks base method.
func (m *MockHealthCheck) Unsubscribe(arg0 chan *discovery.TabletHealth) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unsubscribe", arg0)
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockHealthCheckMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockHealthCheck)(nil).Unsubscribe), arg0)
}

// WaitForAllServingTablets mocks base method.
func (m *MockHealthCheck) WaitForAllServingTablets(arg0 context.Context, arg1 []*query.Target) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForAllServingTablets", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForAllServingTablets indicates an expected call of WaitForAllServingTablets.
func (mr *MockHealthCheckMockRecorder) WaitForAllServingTablets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForAllServingTablets", reflect.TypeOf((*MockHealthCheck)(nil).WaitForAllServingTablets), arg0, arg1)
}
