// Code generated by MockGen. DO NOT EDIT.
// Source: accountService.go

// Package service is a generated GoMock package.
package service

import (
	domain "banking/domain"
	errs "banking/errs"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccountService is a mock of AccountService interface.
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService.
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance.
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccountService) CreateAccount(customerId int64, t domain.AccountType) *errs.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", customerId, t)
	ret0, _ := ret[0].(*errs.AppError)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountServiceMockRecorder) CreateAccount(customerId, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountService)(nil).CreateAccount), customerId, t)
}

// CreateTransaction mocks base method.
func (m *MockAccountService) CreateTransaction(accountId int64, amount float64, t domain.TransactionType) *errs.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", accountId, amount, t)
	ret0, _ := ret[0].(*errs.AppError)
	return ret0
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockAccountServiceMockRecorder) CreateTransaction(accountId, amount, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockAccountService)(nil).CreateTransaction), accountId, amount, t)
}
