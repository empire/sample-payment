// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/sdk/interface.go
//
// Generated by this command:
//
//	mockgen -source=pkg/sdk/interface.go -typed -destination=pkg/sdk/mock_sdk/mock.go
//

// Package mock_sdk is a generated GoMock package.
package mock_sdk

import (
	reflect "reflect"

	sdk "github.com/empire/sample-payment/pkg/sdk"
	gomock "go.uber.org/mock/gomock"
)

// MockPaymentService is a mock of PaymentService interface.
type MockPaymentService struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentServiceMockRecorder
}

// MockPaymentServiceMockRecorder is the mock recorder for MockPaymentService.
type MockPaymentServiceMockRecorder struct {
	mock *MockPaymentService
}

// NewMockPaymentService creates a new mock instance.
func NewMockPaymentService(ctrl *gomock.Controller) *MockPaymentService {
	mock := &MockPaymentService{ctrl: ctrl}
	mock.recorder = &MockPaymentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentService) EXPECT() *MockPaymentServiceMockRecorder {
	return m.recorder
}

// Balance mocks base method.
func (m *MockPaymentService) Balance(accountID sdk.AccountID) (sdk.Amount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Balance", accountID)
	ret0, _ := ret[0].(sdk.Amount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Balance indicates an expected call of Balance.
func (mr *MockPaymentServiceMockRecorder) Balance(accountID any) *MockPaymentServiceBalanceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockPaymentService)(nil).Balance), accountID)
	return &MockPaymentServiceBalanceCall{Call: call}
}

// MockPaymentServiceBalanceCall wrap *gomock.Call
type MockPaymentServiceBalanceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPaymentServiceBalanceCall) Return(arg0 sdk.Amount, arg1 error) *MockPaymentServiceBalanceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPaymentServiceBalanceCall) Do(f func(sdk.AccountID) (sdk.Amount, error)) *MockPaymentServiceBalanceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPaymentServiceBalanceCall) DoAndReturn(f func(sdk.AccountID) (sdk.Amount, error)) *MockPaymentServiceBalanceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Deposit mocks base method.
func (m *MockPaymentService) Deposit(accountID sdk.AccountID, amount sdk.Amount) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deposit", accountID, amount)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deposit indicates an expected call of Deposit.
func (mr *MockPaymentServiceMockRecorder) Deposit(accountID, amount any) *MockPaymentServiceDepositCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deposit", reflect.TypeOf((*MockPaymentService)(nil).Deposit), accountID, amount)
	return &MockPaymentServiceDepositCall{Call: call}
}

// MockPaymentServiceDepositCall wrap *gomock.Call
type MockPaymentServiceDepositCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPaymentServiceDepositCall) Return(arg0 string, arg1 error) *MockPaymentServiceDepositCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPaymentServiceDepositCall) Do(f func(sdk.AccountID, sdk.Amount) (string, error)) *MockPaymentServiceDepositCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPaymentServiceDepositCall) DoAndReturn(f func(sdk.AccountID, sdk.Amount) (string, error)) *MockPaymentServiceDepositCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Withdraw mocks base method.
func (m *MockPaymentService) Withdraw(accountID sdk.AccountID, amount sdk.Amount) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", accountID, amount)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockPaymentServiceMockRecorder) Withdraw(accountID, amount any) *MockPaymentServiceWithdrawCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockPaymentService)(nil).Withdraw), accountID, amount)
	return &MockPaymentServiceWithdrawCall{Call: call}
}

// MockPaymentServiceWithdrawCall wrap *gomock.Call
type MockPaymentServiceWithdrawCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPaymentServiceWithdrawCall) Return(arg0 string, arg1 error) *MockPaymentServiceWithdrawCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPaymentServiceWithdrawCall) Do(f func(sdk.AccountID, sdk.Amount) (string, error)) *MockPaymentServiceWithdrawCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPaymentServiceWithdrawCall) DoAndReturn(f func(sdk.AccountID, sdk.Amount) (string, error)) *MockPaymentServiceWithdrawCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
