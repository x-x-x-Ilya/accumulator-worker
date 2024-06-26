// Code generated by MockGen. DO NOT EDIT.
// Source: dependencies.go
//
// Generated by this command:
//
//	mockgen -source dependencies.go -destination mock.go -package pipeline
//

// Package pipeline is a generated GoMock package.
package pipeline

import (
	big "math/big"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// Mockaccumulator is a mock of accumulator interface.
type Mockaccumulator struct {
	ctrl     *gomock.Controller
	recorder *MockaccumulatorMockRecorder
}

// MockaccumulatorMockRecorder is the mock recorder for Mockaccumulator.
type MockaccumulatorMockRecorder struct {
	mock *Mockaccumulator
}

// NewMockaccumulator creates a new mock instance.
func NewMockaccumulator(ctrl *gomock.Controller) *Mockaccumulator {
	mock := &Mockaccumulator{ctrl: ctrl}
	mock.recorder = &MockaccumulatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockaccumulator) EXPECT() *MockaccumulatorMockRecorder {
	return m.recorder
}

// Accumulate mocks base method.
func (m *Mockaccumulator) Accumulate(input []int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Accumulate", input)
}

// Accumulate indicates an expected call of Accumulate.
func (mr *MockaccumulatorMockRecorder) Accumulate(input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accumulate", reflect.TypeOf((*Mockaccumulator)(nil).Accumulate), input)
}

// Get mocks base method.
func (m *Mockaccumulator) Get() big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(big.Int)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockaccumulatorMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Mockaccumulator)(nil).Get))
}

// Mockpublisher is a mock of publisher interface.
type Mockpublisher struct {
	ctrl     *gomock.Controller
	recorder *MockpublisherMockRecorder
}

// MockpublisherMockRecorder is the mock recorder for Mockpublisher.
type MockpublisherMockRecorder struct {
	mock *Mockpublisher
}

// NewMockpublisher creates a new mock instance.
func NewMockpublisher(ctrl *gomock.Controller) *Mockpublisher {
	mock := &Mockpublisher{ctrl: ctrl}
	mock.recorder = &MockpublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockpublisher) EXPECT() *MockpublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *Mockpublisher) Publish(data any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Publish", data)
}

// Publish indicates an expected call of Publish.
func (mr *MockpublisherMockRecorder) Publish(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*Mockpublisher)(nil).Publish), data)
}

// Mockworker is a mock of worker interface.
type Mockworker struct {
	ctrl     *gomock.Controller
	recorder *MockworkerMockRecorder
}

// MockworkerMockRecorder is the mock recorder for Mockworker.
type MockworkerMockRecorder struct {
	mock *Mockworker
}

// NewMockworker creates a new mock instance.
func NewMockworker(ctrl *gomock.Controller) *Mockworker {
	mock := &Mockworker{ctrl: ctrl}
	mock.recorder = &MockworkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockworker) EXPECT() *MockworkerMockRecorder {
	return m.recorder
}

// ProcessJob mocks base method.
func (m *Mockworker) ProcessJob(job func() error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessJob", job)
}

// ProcessJob indicates an expected call of ProcessJob.
func (mr *MockworkerMockRecorder) ProcessJob(job any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessJob", reflect.TypeOf((*Mockworker)(nil).ProcessJob), job)
}

// Mockprocessor is a mock of processor interface.
type Mockprocessor struct {
	ctrl     *gomock.Controller
	recorder *MockprocessorMockRecorder
}

// MockprocessorMockRecorder is the mock recorder for Mockprocessor.
type MockprocessorMockRecorder struct {
	mock *Mockprocessor
}

// NewMockprocessor creates a new mock instance.
func NewMockprocessor(ctrl *gomock.Controller) *Mockprocessor {
	mock := &Mockprocessor{ctrl: ctrl}
	mock.recorder = &MockprocessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockprocessor) EXPECT() *MockprocessorMockRecorder {
	return m.recorder
}

// ThreeMaxElements mocks base method.
func (m *Mockprocessor) ThreeMaxElements(input []int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ThreeMaxElements", input)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ThreeMaxElements indicates an expected call of ThreeMaxElements.
func (mr *MockprocessorMockRecorder) ThreeMaxElements(input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ThreeMaxElements", reflect.TypeOf((*Mockprocessor)(nil).ThreeMaxElements), input)
}

// MocksliceGenerator is a mock of sliceGenerator interface.
type MocksliceGenerator struct {
	ctrl     *gomock.Controller
	recorder *MocksliceGeneratorMockRecorder
}

// MocksliceGeneratorMockRecorder is the mock recorder for MocksliceGenerator.
type MocksliceGeneratorMockRecorder struct {
	mock *MocksliceGenerator
}

// NewMocksliceGenerator creates a new mock instance.
func NewMocksliceGenerator(ctrl *gomock.Controller) *MocksliceGenerator {
	mock := &MocksliceGenerator{ctrl: ctrl}
	mock.recorder = &MocksliceGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksliceGenerator) EXPECT() *MocksliceGeneratorMockRecorder {
	return m.recorder
}

// MakeRandomInt mocks base method.
func (m *MocksliceGenerator) MakeRandomInt(length, maxValue, minValue int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRandomInt", length, maxValue, minValue)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeRandomInt indicates an expected call of MakeRandomInt.
func (mr *MocksliceGeneratorMockRecorder) MakeRandomInt(length, maxValue, minValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRandomInt", reflect.TypeOf((*MocksliceGenerator)(nil).MakeRandomInt), length, maxValue, minValue)
}
