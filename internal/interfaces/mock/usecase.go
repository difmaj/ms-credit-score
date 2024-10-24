// Code generated by MockGen. DO NOT EDIT.
// Source: /home/alisson-arus/projects/ms-credit-score/internal/interfaces/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/difmaj/ms-credit-score/internal/domain"
	dto "github.com/difmaj/ms-credit-score/internal/dto"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIUsecase is a mock of IUsecase interface.
type MockIUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIUsecaseMockRecorder
}

// MockIUsecaseMockRecorder is the mock recorder for MockIUsecase.
type MockIUsecaseMockRecorder struct {
	mock *MockIUsecase
}

// NewMockIUsecase creates a new mock instance.
func NewMockIUsecase(ctrl *gomock.Controller) *MockIUsecase {
	mock := &MockIUsecase{ctrl: ctrl}
	mock.recorder = &MockIUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUsecase) EXPECT() *MockIUsecaseMockRecorder {
	return m.recorder
}

// ClaimsJWT mocks base method.
func (m *MockIUsecase) ClaimsJWT(token string) (*domain.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClaimsJWT", token)
	ret0, _ := ret[0].(*domain.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClaimsJWT indicates an expected call of ClaimsJWT.
func (mr *MockIUsecaseMockRecorder) ClaimsJWT(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClaimsJWT", reflect.TypeOf((*MockIUsecase)(nil).ClaimsJWT), token)
}

// CreateAsset mocks base method.
func (m *MockIUsecase) CreateAsset(ctx context.Context, userID uuid.UUID, in *dto.CreateAssetInput) (*dto.AssetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAsset", ctx, userID, in)
	ret0, _ := ret[0].(*dto.AssetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAsset indicates an expected call of CreateAsset.
func (mr *MockIUsecaseMockRecorder) CreateAsset(ctx, userID, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAsset", reflect.TypeOf((*MockIUsecase)(nil).CreateAsset), ctx, userID, in)
}

// CreateDebt mocks base method.
func (m *MockIUsecase) CreateDebt(ctx context.Context, userID uuid.UUID, in *dto.CreateDebtInput) (*dto.DebtOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDebt", ctx, userID, in)
	ret0, _ := ret[0].(*dto.DebtOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDebt indicates an expected call of CreateDebt.
func (mr *MockIUsecaseMockRecorder) CreateDebt(ctx, userID, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDebt", reflect.TypeOf((*MockIUsecase)(nil).CreateDebt), ctx, userID, in)
}

// DeleteAsset mocks base method.
func (m *MockIUsecase) DeleteAsset(ctx context.Context, userID uuid.UUID, in *dto.DeleteAssetInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAsset", ctx, userID, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAsset indicates an expected call of DeleteAsset.
func (mr *MockIUsecaseMockRecorder) DeleteAsset(ctx, userID, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAsset", reflect.TypeOf((*MockIUsecase)(nil).DeleteAsset), ctx, userID, in)
}

// DeleteDebt mocks base method.
func (m *MockIUsecase) DeleteDebt(ctx context.Context, userID uuid.UUID, in *dto.DeleteDebtInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDebt", ctx, userID, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDebt indicates an expected call of DeleteDebt.
func (mr *MockIUsecaseMockRecorder) DeleteDebt(ctx, userID, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDebt", reflect.TypeOf((*MockIUsecase)(nil).DeleteDebt), ctx, userID, in)
}

// GetAssetByID mocks base method.
func (m *MockIUsecase) GetAssetByID(ctx context.Context, userID uuid.UUID, in *dto.GetAssetByIDInput) (*dto.AssetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetByID", ctx, userID, in)
	ret0, _ := ret[0].(*dto.AssetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetByID indicates an expected call of GetAssetByID.
func (mr *MockIUsecaseMockRecorder) GetAssetByID(ctx, userID, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetByID", reflect.TypeOf((*MockIUsecase)(nil).GetAssetByID), ctx, userID, in)
}

// GetAssetsByUserID mocks base method.
func (m *MockIUsecase) GetAssetsByUserID(ctx context.Context, userID uuid.UUID) ([]*dto.AssetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetsByUserID", ctx, userID)
	ret0, _ := ret[0].([]*dto.AssetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetsByUserID indicates an expected call of GetAssetsByUserID.
func (mr *MockIUsecaseMockRecorder) GetAssetsByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetsByUserID", reflect.TypeOf((*MockIUsecase)(nil).GetAssetsByUserID), ctx, userID)
}

// GetDebtByID mocks base method.
func (m *MockIUsecase) GetDebtByID(ctx context.Context, userID uuid.UUID, in *dto.GetDebtByIDInput) (*dto.DebtOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDebtByID", ctx, userID, in)
	ret0, _ := ret[0].(*dto.DebtOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDebtByID indicates an expected call of GetDebtByID.
func (mr *MockIUsecaseMockRecorder) GetDebtByID(ctx, userID, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDebtByID", reflect.TypeOf((*MockIUsecase)(nil).GetDebtByID), ctx, userID, in)
}

// GetDebtsByUserID mocks base method.
func (m *MockIUsecase) GetDebtsByUserID(ctx context.Context, userID uuid.UUID) ([]*dto.DebtOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDebtsByUserID", ctx, userID)
	ret0, _ := ret[0].([]*dto.DebtOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDebtsByUserID indicates an expected call of GetDebtsByUserID.
func (mr *MockIUsecaseMockRecorder) GetDebtsByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDebtsByUserID", reflect.TypeOf((*MockIUsecase)(nil).GetDebtsByUserID), ctx, userID)
}

// Login mocks base method.
func (m *MockIUsecase) Login(arg0 context.Context, arg1 *dto.LoginInput) (*dto.LoginOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*dto.LoginOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockIUsecaseMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockIUsecase)(nil).Login), arg0, arg1)
}

// UpdateAsset mocks base method.
func (m *MockIUsecase) UpdateAsset(ctx context.Context, userID, assertID uuid.UUID, in *dto.UpdateAssetInput) (*dto.AssetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAsset", ctx, userID, assertID, in)
	ret0, _ := ret[0].(*dto.AssetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAsset indicates an expected call of UpdateAsset.
func (mr *MockIUsecaseMockRecorder) UpdateAsset(ctx, userID, assertID, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAsset", reflect.TypeOf((*MockIUsecase)(nil).UpdateAsset), ctx, userID, assertID, in)
}

// UpdateDebt mocks base method.
func (m *MockIUsecase) UpdateDebt(ctx context.Context, userID, assertID uuid.UUID, in *dto.UpdateDebtInput) (*dto.DebtOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDebt", ctx, userID, assertID, in)
	ret0, _ := ret[0].(*dto.DebtOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDebt indicates an expected call of UpdateDebt.
func (mr *MockIUsecaseMockRecorder) UpdateDebt(ctx, userID, assertID, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDebt", reflect.TypeOf((*MockIUsecase)(nil).UpdateDebt), ctx, userID, assertID, in)
}
