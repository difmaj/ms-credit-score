// Code generated by MockGen. DO NOT EDIT.
// Source: /home/alisson-arus/projects/ms-credit-score/internal/interfaces/middleware.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockIMiddleware is a mock of IMiddleware interface.
type MockIMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockIMiddlewareMockRecorder
}

// MockIMiddlewareMockRecorder is the mock recorder for MockIMiddleware.
type MockIMiddlewareMockRecorder struct {
	mock *MockIMiddleware
}

// NewMockIMiddleware creates a new mock instance.
func NewMockIMiddleware(ctrl *gomock.Controller) *MockIMiddleware {
	mock := &MockIMiddleware{ctrl: ctrl}
	mock.recorder = &MockIMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMiddleware) EXPECT() *MockIMiddlewareMockRecorder {
	return m.recorder
}

// BasicAuth mocks base method.
func (m *MockIMiddleware) BasicAuth() func(*gin.Context) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BasicAuth")
	ret0, _ := ret[0].(func(*gin.Context))
	return ret0
}

// BasicAuth indicates an expected call of BasicAuth.
func (mr *MockIMiddlewareMockRecorder) BasicAuth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BasicAuth", reflect.TypeOf((*MockIMiddleware)(nil).BasicAuth))
}

// ErrorHandler mocks base method.
func (m *MockIMiddleware) ErrorHandler() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrorHandler")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// ErrorHandler indicates an expected call of ErrorHandler.
func (mr *MockIMiddlewareMockRecorder) ErrorHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorHandler", reflect.TypeOf((*MockIMiddleware)(nil).ErrorHandler))
}

// PermissionAuth mocks base method.
func (m *MockIMiddleware) PermissionAuth(context, action string) func(*gin.Context) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PermissionAuth", context, action)
	ret0, _ := ret[0].(func(*gin.Context))
	return ret0
}

// PermissionAuth indicates an expected call of PermissionAuth.
func (mr *MockIMiddlewareMockRecorder) PermissionAuth(context, action interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PermissionAuth", reflect.TypeOf((*MockIMiddleware)(nil).PermissionAuth), context, action)
}