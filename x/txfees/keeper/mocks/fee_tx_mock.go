// Code generated by MockGen. DO NOT EDIT.
// Source: /root/workspace/cosmos-sdk/types/tx_msg.go

// Package mock_types is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)


// MockFeeTx is a mock of FeeTx interface.
type MockFeeTx struct {
	ctrl     *gomock.Controller
	recorder *MockFeeTxMockRecorder
}

// MockFeeTxMockRecorder is the mock recorder for MockFeeTx.
type MockFeeTxMockRecorder struct {
	mock *MockFeeTx
}

// NewMockFeeTx creates a new mock instance.
func NewMockFeeTx(ctrl *gomock.Controller) *MockFeeTx {
	mock := &MockFeeTx{ctrl: ctrl}
	mock.recorder = &MockFeeTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeeTx) EXPECT() *MockFeeTxMockRecorder {
	return m.recorder
}

// FeeGranter mocks base method.
func (m *MockFeeTx) FeeGranter() types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeeGranter")
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}

// FeeGranter indicates an expected call of FeeGranter.
func (mr *MockFeeTxMockRecorder) FeeGranter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeeGranter", reflect.TypeOf((*MockFeeTx)(nil).FeeGranter))
}

// FeePayer mocks base method.
func (m *MockFeeTx) FeePayer() types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeePayer")
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}

// FeePayer indicates an expected call of FeePayer.
func (mr *MockFeeTxMockRecorder) FeePayer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeePayer", reflect.TypeOf((*MockFeeTx)(nil).FeePayer))
}

// GetFee mocks base method.
func (m *MockFeeTx) GetFee() types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFee")
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// GetFee indicates an expected call of GetFee.
func (mr *MockFeeTxMockRecorder) GetFee() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFee", reflect.TypeOf((*MockFeeTx)(nil).GetFee))
}

// GetGas mocks base method.
func (m *MockFeeTx) GetGas() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGas")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetGas indicates an expected call of GetGas.
func (mr *MockFeeTxMockRecorder) GetGas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGas", reflect.TypeOf((*MockFeeTx)(nil).GetGas))
}

// GetMsgs mocks base method.
func (m *MockFeeTx) GetMsgs() []types0.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMsgs")
	ret0, _ := ret[0].([]types0.Msg)
	return ret0
}

// GetMsgs indicates an expected call of GetMsgs.
func (mr *MockFeeTxMockRecorder) GetMsgs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMsgs", reflect.TypeOf((*MockFeeTx)(nil).GetMsgs))
}

// ValidateBasic mocks base method.
func (m *MockFeeTx) ValidateBasic() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateBasic")
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateBasic indicates an expected call of ValidateBasic.
func (mr *MockFeeTxMockRecorder) ValidateBasic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateBasic", reflect.TypeOf((*MockFeeTx)(nil).ValidateBasic))
}