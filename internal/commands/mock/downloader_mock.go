// Code generated by MockGen. DO NOT EDIT.
// Source: download.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDownloader is a mock of Downloader interface
type MockDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockDownloaderMockRecorder
}

// MockDownloaderMockRecorder is the mock recorder for MockDownloader
type MockDownloaderMockRecorder struct {
	mock *MockDownloader
}

// NewMockDownloader creates a new mock instance
func NewMockDownloader(ctrl *gomock.Controller) *MockDownloader {
	mock := &MockDownloader{ctrl: ctrl}
	mock.recorder = &MockDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDownloader) EXPECT() *MockDownloaderMockRecorder {
	return m.recorder
}

// DownloadGoPackage mocks base method
func (m *MockDownloader) DownloadGoPackage(version string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadGoPackage", version)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadGoPackage indicates an expected call of DownloadGoPackage
func (mr *MockDownloaderMockRecorder) DownloadGoPackage(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadGoPackage", reflect.TypeOf((*MockDownloader)(nil).DownloadGoPackage), version)
}
