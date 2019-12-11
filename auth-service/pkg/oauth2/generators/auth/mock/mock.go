// Code generated by MockGen. DO NOT EDIT.
// Source: gopkg.in/oauth2.v3 (interfaces: AuthorizeGenerate)

package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	oauth2_v3 "gopkg.in/oauth2.v3"
)

// MockAuthorizeGenerate is a mock of AuthorizeGenerate interface
type MockAuthorizeGenerate struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizeGenerateMockRecorder
}

// MockAuthorizeGenerateMockRecorder is the mock recorder for MockAuthorizeGenerate
type MockAuthorizeGenerateMockRecorder struct {
	mock *MockAuthorizeGenerate
}

// NewMockAuthorizeGenerate creates a new mock instance
func NewMockAuthorizeGenerate(ctrl *gomock.Controller) *MockAuthorizeGenerate {
	mock := &MockAuthorizeGenerate{ctrl: ctrl}
	mock.recorder = &MockAuthorizeGenerateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorizeGenerate) EXPECT() *MockAuthorizeGenerateMockRecorder {
	return m.recorder
}

// Token mocks base method
func (m *MockAuthorizeGenerate) Token(arg0 *oauth2_v3.GenerateBasic) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Token indicates an expected call of Token
func (mr *MockAuthorizeGenerateMockRecorder) Token(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockAuthorizeGenerate)(nil).Token), arg0)
}