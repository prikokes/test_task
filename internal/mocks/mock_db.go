// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/models/db.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockDB) Begin() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Begin indicates an expected call of Begin.
func (mr *MockDBMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockDB)(nil).Begin))
}

// Commit mocks base method.
func (m *MockDB) Commit() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockDBMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockDB)(nil).Commit))
}

// Count mocks base method.
func (m *MockDB) Count(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockDBMockRecorder) Count(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDB)(nil).Count), value)
}

// Create mocks base method.
func (m *MockDB) Create(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDBMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDB)(nil).Create), value)
}

// Find mocks base method.
func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockDBMockRecorder) Find(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDB)(nil).Find), varargs...)
}

// First mocks base method.
func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockDBMockRecorder) First(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockDB)(nil).First), varargs...)
}

// Joins mocks base method.
func (m *MockDB) Joins(query string, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Joins", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Joins indicates an expected call of Joins.
func (mr *MockDBMockRecorder) Joins(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Joins", reflect.TypeOf((*MockDB)(nil).Joins), varargs...)
}

// Model mocks base method.
func (m *MockDB) Model(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockDBMockRecorder) Model(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockDB)(nil).Model), value)
}

// Rollback mocks base method.
func (m *MockDB) Rollback() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockDBMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockDB)(nil).Rollback))
}

// Scan mocks base method.
func (m *MockDB) Scan(dest interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", dest)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockDBMockRecorder) Scan(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockDB)(nil).Scan), dest)
}

// Select mocks base method.
func (m *MockDB) Select(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *MockDBMockRecorder) Select(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockDB)(nil).Select), varargs...)
}

// Table mocks base method.
func (m *MockDB) Table(name string) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Table", name)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Table indicates an expected call of Table.
func (mr *MockDBMockRecorder) Table(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Table", reflect.TypeOf((*MockDB)(nil).Table), name)
}

// Update mocks base method.
func (m *MockDB) Update(column string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", column, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDBMockRecorder) Update(column, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDB)(nil).Update), column, value)
}

// Where mocks base method.
func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockDBMockRecorder) Where(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockDB)(nil).Where), varargs...)
}
