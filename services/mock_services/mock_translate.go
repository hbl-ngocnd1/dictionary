// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hbl-ngocnd1/dictionary/services (interfaces: TranslateService)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/hbl-ngocnd1/dictionary/models"
)

// MockTranslateService is a mock of TranslateService interface.
type MockTranslateService struct {
	ctrl     *gomock.Controller
	recorder *MockTranslateServiceMockRecorder
}

// MockTranslateServiceMockRecorder is the mock recorder for MockTranslateService.
type MockTranslateServiceMockRecorder struct {
	mock *MockTranslateService
}

// NewMockTranslateService creates a new mock instance.
func NewMockTranslateService(ctrl *gomock.Controller) *MockTranslateService {
	mock := &MockTranslateService{ctrl: ctrl}
	mock.recorder = &MockTranslateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslateService) EXPECT() *MockTranslateServiceMockRecorder {
	return m.recorder
}

// TranslateData mocks base method.
func (m *MockTranslateService) TranslateData(arg0 context.Context, arg1 []models.Word) []models.Word {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateData", arg0, arg1)
	ret0, _ := ret[0].([]models.Word)
	return ret0
}

// TranslateData indicates an expected call of TranslateData.
func (mr *MockTranslateServiceMockRecorder) TranslateData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateData", reflect.TypeOf((*MockTranslateService)(nil).TranslateData), arg0, arg1)
}
