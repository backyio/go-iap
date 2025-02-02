// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/awa/go-iap/appstore (interfaces: IAPClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	appstore "github.com/awa/go-iap/appstore"
	jwt "github.com/golang-jwt/jwt/v4"
	gomock "github.com/golang/mock/gomock"
)

// MockIAPClient is a mock of IAPClient interface.
type MockIAPClient struct {
	ctrl     *gomock.Controller
	recorder *MockIAPClientMockRecorder
}

// MockIAPClientMockRecorder is the mock recorder for MockIAPClient.
type MockIAPClientMockRecorder struct {
	mock *MockIAPClient
}

// NewMockIAPClient creates a new mock instance.
func NewMockIAPClient(ctrl *gomock.Controller) *MockIAPClient {
	mock := &MockIAPClient{ctrl: ctrl}
	mock.recorder = &MockIAPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAPClient) EXPECT() *MockIAPClientMockRecorder {
	return m.recorder
}

// ParseNotificationV2 mocks base method.
func (m *MockIAPClient) ParseNotificationV2(arg0 string, arg1 *jwt.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseNotificationV2", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ParseNotificationV2 indicates an expected call of ParseNotificationV2.
func (mr *MockIAPClientMockRecorder) ParseNotificationV2(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseNotificationV2", reflect.TypeOf((*MockIAPClient)(nil).ParseNotificationV2), arg0, arg1)
}

// Verify mocks base method.
func (m *MockIAPClient) Verify(arg0 context.Context, arg1 appstore.IAPRequest, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockIAPClientMockRecorder) Verify(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockIAPClient)(nil).Verify), arg0, arg1, arg2)
}

// VerifyWithStatus mocks base method.
func (m *MockIAPClient) VerifyWithStatus(arg0 context.Context, arg1 appstore.IAPRequest, arg2 interface{}) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyWithStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyWithStatus indicates an expected call of VerifyWithStatus.
func (mr *MockIAPClientMockRecorder) VerifyWithStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyWithStatus", reflect.TypeOf((*MockIAPClient)(nil).VerifyWithStatus), arg0, arg1, arg2)
}
