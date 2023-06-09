// Code generated by MockGen. DO NOT EDIT.
// Source: model.go

// Package writer is a generated GoMock package.
package writer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOsFileWriter is a mock of OsFileWriter interface.
type MockOsFileWriter struct {
	ctrl     *gomock.Controller
	recorder *MockOsFileWriterMockRecorder
}

// MockOsFileWriterMockRecorder is the mock recorder for MockOsFileWriter.
type MockOsFileWriterMockRecorder struct {
	mock *MockOsFileWriter
}

// NewMockOsFileWriter creates a new mock instance.
func NewMockOsFileWriter(ctrl *gomock.Controller) *MockOsFileWriter {
	mock := &MockOsFileWriter{ctrl: ctrl}
	mock.recorder = &MockOsFileWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOsFileWriter) EXPECT() *MockOsFileWriterMockRecorder {
	return m.recorder
}

// WriteAt mocks base method.
func (m *MockOsFileWriter) WriteAt(arg0 []byte, arg1 int64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteAt", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteAt indicates an expected call of WriteAt.
func (mr *MockOsFileWriterMockRecorder) WriteAt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAt", reflect.TypeOf((*MockOsFileWriter)(nil).WriteAt), arg0, arg1)
}

// MockChunkWriter is a mock of ChunkWriter interface.
type MockChunkWriter struct {
	ctrl     *gomock.Controller
	recorder *MockChunkWriterMockRecorder
}

// MockChunkWriterMockRecorder is the mock recorder for MockChunkWriter.
type MockChunkWriterMockRecorder struct {
	mock *MockChunkWriter
}

// NewMockChunkWriter creates a new mock instance.
func NewMockChunkWriter(ctrl *gomock.Controller) *MockChunkWriter {
	mock := &MockChunkWriter{ctrl: ctrl}
	mock.recorder = &MockChunkWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChunkWriter) EXPECT() *MockChunkWriterMockRecorder {
	return m.recorder
}

// WriteChunk mocks base method.
func (m *MockChunkWriter) WriteChunk(arg0 int64, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteChunk", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteChunk indicates an expected call of WriteChunk.
func (mr *MockChunkWriterMockRecorder) WriteChunk(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteChunk", reflect.TypeOf((*MockChunkWriter)(nil).WriteChunk), arg0, arg1)
}
