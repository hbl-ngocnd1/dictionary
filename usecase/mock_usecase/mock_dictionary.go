// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hbl-ngocnd1/dictionary/usecase (interfaces: DictUseCase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/hbl-ngocnd1/dictionary/models"
)

// MockDictUseCase is a mock of DictUseCase interface.
type MockDictUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockDictUseCaseMockRecorder
}

// MockDictUseCaseMockRecorder is the mock recorder for MockDictUseCase.
type MockDictUseCaseMockRecorder struct {
	mock *MockDictUseCase
}

// NewMockDictUseCase creates a new mock instance.
func NewMockDictUseCase(ctrl *gomock.Controller) *MockDictUseCase {
	mock := &MockDictUseCase{ctrl: ctrl}
	mock.recorder = &MockDictUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDictUseCase) EXPECT() *MockDictUseCaseMockRecorder {
	return m.recorder
}

// GetDetail mocks base method.
func (m *MockDictUseCase) GetDetail(arg0 context.Context, arg1 string, arg2 int) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetail", arg0, arg1, arg2)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetail indicates an expected call of GetDetail.
func (mr *MockDictUseCaseMockRecorder) GetDetail(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetail", reflect.TypeOf((*MockDictUseCase)(nil).GetDetail), arg0, arg1, arg2)
}

// GetDict mocks base method.
func (m *MockDictUseCase) GetDict(arg0 context.Context, arg1, arg2 int, arg3, arg4, arg5 string, arg6 models.MakeData) ([]models.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDict", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].([]models.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDict indicates an expected call of GetDict.
func (mr *MockDictUseCaseMockRecorder) GetDict(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDict", reflect.TypeOf((*MockDictUseCase)(nil).GetDict), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// GetITJapanWonderWork mocks base method.
func (m *MockDictUseCase) GetITJapanWonderWork(arg0 context.Context) ([][]models.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetITJapanWonderWork", arg0)
	ret0, _ := ret[0].([][]models.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetITJapanWonderWork indicates an expected call of GetITJapanWonderWork.
func (mr *MockDictUseCaseMockRecorder) GetITJapanWonderWork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetITJapanWonderWork", reflect.TypeOf((*MockDictUseCase)(nil).GetITJapanWonderWork), arg0)
}
