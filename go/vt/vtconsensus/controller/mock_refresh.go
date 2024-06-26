/*
Copyright ApeCloud, Inc.
Licensed under the Apache v2(found in the LICENSE file in the root directory).
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: refresh.go

// Package mock_controller is a generated GoMock package.
package controller

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"

	topodata "vitess.io/vitess/go/vt/proto/topodata"
	topo "vitess.io/vitess/go/vt/topo"
)

// MockConsensusTopo is a mock of ConsensusTopo interface.
type MockConsensusTopo struct {
	ctrl     *gomock.Controller
	recorder *MockConsensusTopoMockRecorder
}

// MockConsensusTopoMockRecorder is the mock recorder for MockConsensusTopo.
type MockConsensusTopoMockRecorder struct {
	mock *MockConsensusTopo
}

// NewMockConsensusTopo creates a new mock instance.
func NewMockConsensusTopo(ctrl *gomock.Controller) *MockConsensusTopo {
	mock := &MockConsensusTopo{ctrl: ctrl}
	mock.recorder = &MockConsensusTopoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsensusTopo) EXPECT() *MockConsensusTopoMockRecorder {
	return m.recorder
}

// GetShard mocks base method.
func (m *MockConsensusTopo) GetShard(ctx context.Context, keyspace, shard string) (*topo.ShardInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShard", ctx, keyspace, shard)
	ret0, _ := ret[0].(*topo.ShardInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShard indicates an expected call of GetShard.
func (mr *MockConsensusTopoMockRecorder) GetShard(ctx, keyspace, shard interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShard", reflect.TypeOf((*MockConsensusTopo)(nil).GetShard), ctx, keyspace, shard)
}

// GetShardNames mocks base method.
func (m *MockConsensusTopo) GetShardNames(ctx context.Context, keyspace string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShardNames", ctx, keyspace)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShardNames indicates an expected call of GetShardNames.
func (mr *MockConsensusTopoMockRecorder) GetShardNames(ctx, keyspace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardNames", reflect.TypeOf((*MockConsensusTopo)(nil).GetShardNames), ctx, keyspace)
}

// GetTabletMapForShardByCell mocks base method.
func (m *MockConsensusTopo) GetTabletMapForShardByCell(ctx context.Context, keyspace, shard string, cells []string) (map[string]*topo.TabletInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTabletMapForShardByCell", ctx, keyspace, shard, cells)
	ret0, _ := ret[0].(map[string]*topo.TabletInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTabletMapForShardByCell indicates an expected call of GetTabletMapForShardByCell.
func (mr *MockConsensusTopoMockRecorder) GetTabletMapForShardByCell(ctx, keyspace, shard, cells interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTabletMapForShardByCell", reflect.TypeOf((*MockConsensusTopo)(nil).GetTabletMapForShardByCell), ctx, keyspace, shard, cells)
}

// LockShard mocks base method.
func (m *MockConsensusTopo) LockShard(ctx context.Context, keyspace, shard, action string) (context.Context, func(*error), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockShard", ctx, keyspace, shard, action)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(func(*error))
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LockShard indicates an expected call of LockShard.
func (mr *MockConsensusTopoMockRecorder) LockShard(ctx, keyspace, shard, action interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockShard", reflect.TypeOf((*MockConsensusTopo)(nil).LockShard), ctx, keyspace, shard, action)
}

// MockConsensusTmcClient is a mock of ConsensusTmcClient interface.
type MockConsensusTmcClient struct {
	ctrl     *gomock.Controller
	recorder *MockConsensusTmcClientMockRecorder
}

// MockConsensusTmcClientMockRecorder is the mock recorder for MockConsensusTmcClient.
type MockConsensusTmcClientMockRecorder struct {
	mock *MockConsensusTmcClient
}

// NewMockConsensusTmcClient creates a new mock instance.
func NewMockConsensusTmcClient(ctrl *gomock.Controller) *MockConsensusTmcClient {
	mock := &MockConsensusTmcClient{ctrl: ctrl}
	mock.recorder = &MockConsensusTmcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsensusTmcClient) EXPECT() *MockConsensusTmcClientMockRecorder {
	return m.recorder
}

// ChangeType mocks base method.
func (m *MockConsensusTmcClient) ChangeType(ctx context.Context, tablet *topodata.Tablet, dbType topodata.TabletType, semiSync bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeType", ctx, tablet, dbType, semiSync)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeType indicates an expected call of ChangeType.
func (mr *MockConsensusTmcClientMockRecorder) ChangeType(ctx, tablet, dbType, semiSync interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeType", reflect.TypeOf((*MockConsensusTmcClient)(nil).ChangeType), ctx, tablet, dbType, semiSync)
}

// Ping mocks base method.
func (m *MockConsensusTmcClient) Ping(ctx context.Context, tablet *topodata.Tablet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx, tablet)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockConsensusTmcClientMockRecorder) Ping(ctx, tablet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockConsensusTmcClient)(nil).Ping), ctx, tablet)
}
