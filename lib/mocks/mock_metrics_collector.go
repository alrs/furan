// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dollarshaveclub/furan/lib/metrics (interfaces: MetricsCollector)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// MockMetricsCollector is a mock of MetricsCollector interface
type MockMetricsCollector struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsCollectorMockRecorder
}

// MockMetricsCollectorMockRecorder is the mock recorder for MockMetricsCollector
type MockMetricsCollectorMockRecorder struct {
	mock *MockMetricsCollector
}

// NewMockMetricsCollector creates a new mock instance
func NewMockMetricsCollector(ctrl *gomock.Controller) *MockMetricsCollector {
	mock := &MockMetricsCollector{ctrl: ctrl}
	mock.recorder = &MockMetricsCollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockMetricsCollector) EXPECT() *MockMetricsCollectorMockRecorder {
	return _m.recorder
}

// BuildFailed mocks base method
func (_m *MockMetricsCollector) BuildFailed(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "BuildFailed", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuildFailed indicates an expected call of BuildFailed
func (_mr *MockMetricsCollectorMockRecorder) BuildFailed(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BuildFailed", arg0, arg1)
}

// BuildStarted mocks base method
func (_m *MockMetricsCollector) BuildStarted(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "BuildStarted", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuildStarted indicates an expected call of BuildStarted
func (_mr *MockMetricsCollectorMockRecorder) BuildStarted(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BuildStarted", arg0, arg1)
}

// BuildSucceeded mocks base method
func (_m *MockMetricsCollector) BuildSucceeded(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "BuildSucceeded", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuildSucceeded indicates an expected call of BuildSucceeded
func (_mr *MockMetricsCollectorMockRecorder) BuildSucceeded(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BuildSucceeded", arg0, arg1)
}

// DiskFree mocks base method
func (_m *MockMetricsCollector) DiskFree(_param0 uint64) error {
	ret := _m.ctrl.Call(_m, "DiskFree", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DiskFree indicates an expected call of DiskFree
func (_mr *MockMetricsCollectorMockRecorder) DiskFree(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DiskFree", arg0)
}

// Duration mocks base method
func (_m *MockMetricsCollector) Duration(_param0 string, _param1 string, _param2 string, _param3 []string, _param4 float64) error {
	ret := _m.ctrl.Call(_m, "Duration", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Duration indicates an expected call of Duration
func (_mr *MockMetricsCollectorMockRecorder) Duration(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Duration", arg0, arg1, arg2, arg3, arg4)
}

// FileNodesFree mocks base method
func (_m *MockMetricsCollector) FileNodesFree(_param0 uint64) error {
	ret := _m.ctrl.Call(_m, "FileNodesFree", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FileNodesFree indicates an expected call of FileNodesFree
func (_mr *MockMetricsCollectorMockRecorder) FileNodesFree(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FileNodesFree", arg0)
}

// Float mocks base method
func (_m *MockMetricsCollector) Float(_param0 string, _param1 string, _param2 string, _param3 []string, _param4 float64) error {
	ret := _m.ctrl.Call(_m, "Float", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Float indicates an expected call of Float
func (_mr *MockMetricsCollectorMockRecorder) Float(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Float", arg0, arg1, arg2, arg3, arg4)
}

// GCBytesReclaimed mocks base method
func (_m *MockMetricsCollector) GCBytesReclaimed(_param0 uint64) error {
	ret := _m.ctrl.Call(_m, "GCBytesReclaimed", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GCBytesReclaimed indicates an expected call of GCBytesReclaimed
func (_mr *MockMetricsCollectorMockRecorder) GCBytesReclaimed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GCBytesReclaimed", arg0)
}

// GCFailure mocks base method
func (_m *MockMetricsCollector) GCFailure() error {
	ret := _m.ctrl.Call(_m, "GCFailure")
	ret0, _ := ret[0].(error)
	return ret0
}

// GCFailure indicates an expected call of GCFailure
func (_mr *MockMetricsCollectorMockRecorder) GCFailure() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GCFailure")
}

// GCUntaggedImageRemoved mocks base method
func (_m *MockMetricsCollector) GCUntaggedImageRemoved() error {
	ret := _m.ctrl.Call(_m, "GCUntaggedImageRemoved")
	ret0, _ := ret[0].(error)
	return ret0
}

// GCUntaggedImageRemoved indicates an expected call of GCUntaggedImageRemoved
func (_mr *MockMetricsCollectorMockRecorder) GCUntaggedImageRemoved() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GCUntaggedImageRemoved")
}

// ImageSize mocks base method
func (_m *MockMetricsCollector) ImageSize(_param0 int64, _param1 int64, _param2 string, _param3 string) error {
	ret := _m.ctrl.Call(_m, "ImageSize", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImageSize indicates an expected call of ImageSize
func (_mr *MockMetricsCollectorMockRecorder) ImageSize(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageSize", arg0, arg1, arg2, arg3)
}

// KafkaConsumerFailure mocks base method
func (_m *MockMetricsCollector) KafkaConsumerFailure() error {
	ret := _m.ctrl.Call(_m, "KafkaConsumerFailure")
	ret0, _ := ret[0].(error)
	return ret0
}

// KafkaConsumerFailure indicates an expected call of KafkaConsumerFailure
func (_mr *MockMetricsCollectorMockRecorder) KafkaConsumerFailure() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KafkaConsumerFailure")
}

// KafkaProducerFailure mocks base method
func (_m *MockMetricsCollector) KafkaProducerFailure() error {
	ret := _m.ctrl.Call(_m, "KafkaProducerFailure")
	ret0, _ := ret[0].(error)
	return ret0
}

// KafkaProducerFailure indicates an expected call of KafkaProducerFailure
func (_mr *MockMetricsCollectorMockRecorder) KafkaProducerFailure() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KafkaProducerFailure")
}

// Size mocks base method
func (_m *MockMetricsCollector) Size(_param0 string, _param1 string, _param2 string, _param3 []string, _param4 int64) error {
	ret := _m.ctrl.Call(_m, "Size", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Size indicates an expected call of Size
func (_mr *MockMetricsCollectorMockRecorder) Size(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Size", arg0, arg1, arg2, arg3, arg4)
}
