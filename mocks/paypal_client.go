// Code generated by MockGen. DO NOT EDIT.
// Source: paypal_client.go

// Package mock_paypal_api is a generated GoMock package.
package mock_paypal_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	paypal "github.com/logpacker/paypal-go-sdk"
	error_system "github.com/mirogon/go_error_system"
	paypal_api_data "github.com/mirogon/go_paypal_api/data"
)

// MockPaypalClient is a mock of PaypalClient interface.
type MockPaypalClient struct {
	ctrl     *gomock.Controller
	recorder *MockPaypalClientMockRecorder
}

// MockPaypalClientMockRecorder is the mock recorder for MockPaypalClient.
type MockPaypalClientMockRecorder struct {
	mock *MockPaypalClient
}

// NewMockPaypalClient creates a new mock instance.
func NewMockPaypalClient(ctrl *gomock.Controller) *MockPaypalClient {
	mock := &MockPaypalClient{ctrl: ctrl}
	mock.recorder = &MockPaypalClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaypalClient) EXPECT() *MockPaypalClientMockRecorder {
	return m.recorder
}

// CancelSubscription mocks base method.
func (m *MockPaypalClient) CancelSubscription(subscriptionId string) error_system.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelSubscription", subscriptionId)
	ret0, _ := ret[0].(error_system.Error)
	return ret0
}

// CancelSubscription indicates an expected call of CancelSubscription.
func (mr *MockPaypalClientMockRecorder) CancelSubscription(subscriptionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelSubscription", reflect.TypeOf((*MockPaypalClient)(nil).CancelSubscription), subscriptionId)
}

// CaptureOrder mocks base method.
func (m *MockPaypalClient) CaptureOrder(orderId string) (*paypal.CaptureOrderResponse, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureOrder", orderId)
	ret0, _ := ret[0].(*paypal.CaptureOrderResponse)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// CaptureOrder indicates an expected call of CaptureOrder.
func (mr *MockPaypalClientMockRecorder) CaptureOrder(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureOrder", reflect.TypeOf((*MockPaypalClient)(nil).CaptureOrder), orderId)
}

// CaptureSubscription mocks base method.
func (m *MockPaypalClient) CaptureSubscription(subscriptionId, amount string) error_system.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureSubscription", subscriptionId, amount)
	ret0, _ := ret[0].(error_system.Error)
	return ret0
}

// CaptureSubscription indicates an expected call of CaptureSubscription.
func (mr *MockPaypalClientMockRecorder) CaptureSubscription(subscriptionId, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureSubscription", reflect.TypeOf((*MockPaypalClient)(nil).CaptureSubscription), subscriptionId, amount)
}

// CreateBillingPlan mocks base method.
func (m *MockPaypalClient) CreateBillingPlan(productId, pricePerMonth, name, description string) (paypal_api_data.CreateBillingPlanResponse, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBillingPlan", productId, pricePerMonth, name, description)
	ret0, _ := ret[0].(paypal_api_data.CreateBillingPlanResponse)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// CreateBillingPlan indicates an expected call of CreateBillingPlan.
func (mr *MockPaypalClientMockRecorder) CreateBillingPlan(productId, pricePerMonth, name, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBillingPlan", reflect.TypeOf((*MockPaypalClient)(nil).CreateBillingPlan), productId, pricePerMonth, name, description)
}

// CreateOrder mocks base method.
func (m *MockPaypalClient) CreateOrder(referenceId, price, buyerFirstName, buyerLastName, buyerEmail, intent, brandName, returnUrl, cancelUrl string) (*paypal.Order, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", referenceId, price, buyerFirstName, buyerLastName, buyerEmail, intent, brandName, returnUrl, cancelUrl)
	ret0, _ := ret[0].(*paypal.Order)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockPaypalClientMockRecorder) CreateOrder(referenceId, price, buyerFirstName, buyerLastName, buyerEmail, intent, brandName, returnUrl, cancelUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockPaypalClient)(nil).CreateOrder), referenceId, price, buyerFirstName, buyerLastName, buyerEmail, intent, brandName, returnUrl, cancelUrl)
}

// CreateProduct mocks base method.
func (m *MockPaypalClient) CreateProduct(productName, productType string) (paypal_api_data.CreateProductResponse, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", productName, productType)
	ret0, _ := ret[0].(paypal_api_data.CreateProductResponse)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockPaypalClientMockRecorder) CreateProduct(productName, productType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockPaypalClient)(nil).CreateProduct), productName, productType)
}

// CreateSubscription mocks base method.
func (m *MockPaypalClient) CreateSubscription(planId string) (paypal_api_data.CreateSubscriptionResponse, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", planId)
	ret0, _ := ret[0].(paypal_api_data.CreateSubscriptionResponse)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// CreateSubscription indicates an expected call of CreateSubscription.
func (mr *MockPaypalClientMockRecorder) CreateSubscription(planId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockPaypalClient)(nil).CreateSubscription), planId)
}

// GetAccessToken mocks base method.
func (m *MockPaypalClient) GetAccessToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAccessToken indicates an expected call of GetAccessToken.
func (mr *MockPaypalClientMockRecorder) GetAccessToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessToken", reflect.TypeOf((*MockPaypalClient)(nil).GetAccessToken))
}

// GetOrder mocks base method.
func (m *MockPaypalClient) GetOrder(orderId string) (*paypal.Order, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", orderId)
	ret0, _ := ret[0].(*paypal.Order)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockPaypalClientMockRecorder) GetOrder(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockPaypalClient)(nil).GetOrder), orderId)
}

// GetSubscriptionTransactions mocks base method.
func (m *MockPaypalClient) GetSubscriptionTransactions(subscriptionId, startTime, endTime string) (paypal_api_data.GetSubscriptionTransactionsResponse, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionTransactions", subscriptionId, startTime, endTime)
	ret0, _ := ret[0].(paypal_api_data.GetSubscriptionTransactionsResponse)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// GetSubscriptionTransactions indicates an expected call of GetSubscriptionTransactions.
func (mr *MockPaypalClientMockRecorder) GetSubscriptionTransactions(subscriptionId, startTime, endTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionTransactions", reflect.TypeOf((*MockPaypalClient)(nil).GetSubscriptionTransactions), subscriptionId, startTime, endTime)
}

// IsSandbox mocks base method.
func (m *MockPaypalClient) IsSandbox() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSandbox")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSandbox indicates an expected call of IsSandbox.
func (mr *MockPaypalClientMockRecorder) IsSandbox() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSandbox", reflect.TypeOf((*MockPaypalClient)(nil).IsSandbox))
}

// ShowSubscriptionDetails mocks base method.
func (m *MockPaypalClient) ShowSubscriptionDetails(subscriptionId string) (paypal_api_data.ShowSubscriptionDetailsResponse, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowSubscriptionDetails", subscriptionId)
	ret0, _ := ret[0].(paypal_api_data.ShowSubscriptionDetailsResponse)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// ShowSubscriptionDetails indicates an expected call of ShowSubscriptionDetails.
func (mr *MockPaypalClientMockRecorder) ShowSubscriptionDetails(subscriptionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowSubscriptionDetails", reflect.TypeOf((*MockPaypalClient)(nil).ShowSubscriptionDetails), subscriptionId)
}

// UpdateToken mocks base method.
func (m *MockPaypalClient) UpdateToken() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateToken")
}

// UpdateToken indicates an expected call of UpdateToken.
func (mr *MockPaypalClientMockRecorder) UpdateToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockPaypalClient)(nil).UpdateToken))
}
