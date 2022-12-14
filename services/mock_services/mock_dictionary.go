// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hbl-ngocnd1/dictionary/services (interfaces: DictionaryService)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/hbl-ngocnd1/dictionary/models"
)

// MockDictionaryService is a mock of DictionaryService interface.
type MockDictionaryService struct {
	ctrl     *gomock.Controller
	recorder *MockDictionaryServiceMockRecorder
}

// MockDictionaryServiceMockRecorder is the mock recorder for MockDictionaryService.
type MockDictionaryServiceMockRecorder struct {
	mock *MockDictionaryService
}

// NewMockDictionaryService creates a new mock instance.
func NewMockDictionaryService(ctrl *gomock.Controller) *MockDictionaryService {
	mock := &MockDictionaryService{ctrl: ctrl}
	mock.recorder = &MockDictionaryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDictionaryService) EXPECT() *MockDictionaryServiceMockRecorder {
	return m.recorder
}

// GetDetail mocks base method.
func (m *MockDictionaryService) GetDetail(arg0 context.Context, arg1 string, arg2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetail", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetail indicates an expected call of GetDetail.
func (mr *MockDictionaryServiceMockRecorder) GetDetail(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetail", reflect.TypeOf((*MockDictionaryService)(nil).GetDetail), arg0, arg1, arg2)
}

// GetDictionary mocks base method.
func (m *MockDictionaryService) GetDictionary(arg0 context.Context, arg1 string) ([]models.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDictionary", arg0, arg1)
	ret0, _ := ret[0].([]models.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDictionary indicates an expected call of GetDictionary.
func (mr *MockDictionaryServiceMockRecorder) GetDictionary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDictionary", reflect.TypeOf((*MockDictionaryService)(nil).GetDictionary), arg0, arg1)
}

// GetITJapanWonderWork mocks base method.
func (m *MockDictionaryService) GetITJapanWonderWork(arg0 context.Context, arg1 string) ([][]models.WonderWord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetITJapanWonderWork", arg0, arg1)
	ret0, _ := ret[0].([][]models.WonderWord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetITJapanWonderWork indicates an expected call of GetITJapanWonderWork.
func (mr *MockDictionaryServiceMockRecorder) GetITJapanWonderWork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetITJapanWonderWork", reflect.TypeOf((*MockDictionaryService)(nil).GetITJapanWonderWork), arg0, arg1)
}
