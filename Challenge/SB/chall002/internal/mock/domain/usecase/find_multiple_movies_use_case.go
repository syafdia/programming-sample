// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/usecase/find_multiple_movies_use_case.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/syafdia/sb/chal002/internal/domain/entity"
)

// MockFindMultipleMoviesUseCase is a mock of FindMultipleMoviesUseCase interface.
type MockFindMultipleMoviesUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockFindMultipleMoviesUseCaseMockRecorder
}

// MockFindMultipleMoviesUseCaseMockRecorder is the mock recorder for MockFindMultipleMoviesUseCase.
type MockFindMultipleMoviesUseCaseMockRecorder struct {
	mock *MockFindMultipleMoviesUseCase
}

// NewMockFindMultipleMoviesUseCase creates a new mock instance.
func NewMockFindMultipleMoviesUseCase(ctrl *gomock.Controller) *MockFindMultipleMoviesUseCase {
	mock := &MockFindMultipleMoviesUseCase{ctrl: ctrl}
	mock.recorder = &MockFindMultipleMoviesUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFindMultipleMoviesUseCase) EXPECT() *MockFindMultipleMoviesUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockFindMultipleMoviesUseCase) Execute(ctx context.Context, searchWord string, pagination int) ([]entity.MovieSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, searchWord, pagination)
	ret0, _ := ret[0].([]entity.MovieSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockFindMultipleMoviesUseCaseMockRecorder) Execute(ctx, searchWord, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockFindMultipleMoviesUseCase)(nil).Execute), ctx, searchWord, pagination)
}
