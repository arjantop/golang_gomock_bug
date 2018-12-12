// Code generated by MockGen. DO NOT EDIT.
// Source: foo.go

// Package mock_foo is a generated GoMock package.
package mock_foo

import (
	x "."
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFoo is a mock of Foo interface
type MockFoo struct {
	ctrl     *gomock.Controller
	recorder *MockFooMockRecorder
}

// MockFooMockRecorder is the mock recorder for MockFoo
type MockFooMockRecorder struct {
	mock *MockFoo
}

// NewMockFoo creates a new mock instance
func NewMockFoo(ctrl *gomock.Controller) *MockFoo {
	mock := &MockFoo{ctrl: ctrl}
	mock.recorder = &MockFooMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFoo) EXPECT() *MockFooMockRecorder {
	return m.recorder
}

// Bar mocks base method
func (m *MockFoo) Bar(a x.Arg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bar", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bar indicates an expected call of Bar
func (mr *MockFooMockRecorder) Bar(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bar", reflect.TypeOf((*MockFoo)(nil).Bar), a)
}