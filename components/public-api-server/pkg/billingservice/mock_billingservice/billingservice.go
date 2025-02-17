// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_billingservice is a generated GoMock package.
package mock_billingservice

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// FinalizeInvoice mocks base method.
func (m *MockInterface) FinalizeInvoice(ctx context.Context, invoiceId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeInvoice", ctx, invoiceId)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeInvoice indicates an expected call of FinalizeInvoice.
func (mr *MockInterfaceMockRecorder) FinalizeInvoice(ctx, invoiceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeInvoice", reflect.TypeOf((*MockInterface)(nil).FinalizeInvoice), ctx, invoiceId)
}

// CancelSubscription mocks base method.
func (m *MockInterface) CancelSubscription(ctx context.Context, subscriptionId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelSubscription", ctx, subscriptionId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelSubscription indicates an expected call of CancelSubscription.
func (mr *MockInterfaceMockRecorder) CancelSubscription(ctx, subscriptionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelSubscription", reflect.TypeOf((*MockInterface)(nil).CancelSubscription), ctx, subscriptionId)
}
