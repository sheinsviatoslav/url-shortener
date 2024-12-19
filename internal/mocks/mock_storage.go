// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sheinsviatoslav/shortener/internal/storage (interfaces: Storage)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/sheinsviatoslav/shortener/internal/storage"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddManyUrls mocks base method.
func (m *MockStorage) AddManyUrls(arg0 storage.InputManyUrls, arg1 string) (storage.OutputManyUrls, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddManyUrls", arg0, arg1)
	ret0, _ := ret[0].(storage.OutputManyUrls)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddManyUrls indicates an expected call of AddManyUrls.
func (mr *MockStorageMockRecorder) AddManyUrls(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddManyUrls", reflect.TypeOf((*MockStorage)(nil).AddManyUrls), arg0, arg1)
}

// AddNewURL mocks base method.
func (m *MockStorage) AddNewURL(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewURL indicates an expected call of AddNewURL.
func (mr *MockStorageMockRecorder) AddNewURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewURL", reflect.TypeOf((*MockStorage)(nil).AddNewURL), arg0, arg1, arg2)
}

// GetOriginalURLByShortURL mocks base method.
func (m *MockStorage) GetOriginalURLByShortURL(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOriginalURLByShortURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOriginalURLByShortURL indicates an expected call of GetOriginalURLByShortURL.
func (mr *MockStorageMockRecorder) GetOriginalURLByShortURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOriginalURLByShortURL", reflect.TypeOf((*MockStorage)(nil).GetOriginalURLByShortURL), arg0)
}

// GetShortURLByOriginalURL mocks base method.
func (m *MockStorage) GetShortURLByOriginalURL(arg0 string) (string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShortURLByOriginalURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetShortURLByOriginalURL indicates an expected call of GetShortURLByOriginalURL.
func (mr *MockStorageMockRecorder) GetShortURLByOriginalURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShortURLByOriginalURL", reflect.TypeOf((*MockStorage)(nil).GetShortURLByOriginalURL), arg0)
}

// GetUserUrls mocks base method.
func (m *MockStorage) GetUserUrls(arg0 string) (storage.UserUrls, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserUrls", arg0)
	ret0, _ := ret[0].(storage.UserUrls)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserUrls indicates an expected call of GetUserUrls.
func (mr *MockStorageMockRecorder) GetUserUrls(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserUrls", reflect.TypeOf((*MockStorage)(nil).GetUserUrls), arg0)
}
