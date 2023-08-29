package paypal_api_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_http_ "github.com/mirogon/go_http/mocks"
	paypal_api "github.com/mirogon/go_paypal_api"
	mock_paypal_api "github.com/mirogon/go_paypal_api/mocks"
	util "github.com/mirogon/go_util"
)

func TestHandlePaypalWebhooks(t *testing.T) {
	ctrl := gomock.NewController(t)

	paypalClientMock := mock_paypal_api.NewMockPaypalClient(ctrl)
	paypalClientMock.EXPECT().GetAccessToken()
	paypalWebhookValidatorMock := mock_paypal_api.NewMockPaypalWebhookValidator(ctrl)
	paypalWebhookValidatorMock.EXPECT().GetWebhookData(gomock.Any())
	paypalWebhookValidatorMock.EXPECT().ValidateWebHook(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)

	mockResponseWriter := mock_http_.NewMockHttpResponseWriter(ctrl)

	req := util.SetupTestRequest(nil, "123")

	handler := paypal_api.CreatePapalWebhookHandler(paypalClientMock, paypalWebhookValidatorMock)
	_, err := handler.HandlePaypalWebhooks(mockResponseWriter, req)
	if err != nil {
		t.Error()
	}
}

func TestHandlePaypalWebhooks_GetWebhookDataError(t *testing.T) {
	ctrl := gomock.NewController(t)

	paypalClientMock := mock_paypal_api.NewMockPaypalClient(ctrl)
	paypalWebhookValidatorMock := mock_paypal_api.NewMockPaypalWebhookValidator(ctrl)
	paypalWebhookValidatorMock.EXPECT().GetWebhookData(gomock.Any()).Return(paypal_api.WebhookValidationRequest{}, paypal_api.WebhookNotification{}, errors.New(""))

	mockResponseWriter := mock_http_.NewMockHttpResponseWriter(ctrl)

	req := util.SetupTestRequest(nil, "123")

	handler := paypal_api.CreatePapalWebhookHandler(paypalClientMock, paypalWebhookValidatorMock)
	_, err := handler.HandlePaypalWebhooks(mockResponseWriter, req)
	if err == nil {
		t.Error()
	}
}

func TestHandlePaypalWebhooks_ValidateWebhookErr(t *testing.T) {
	ctrl := gomock.NewController(t)

	paypalClientMock := mock_paypal_api.NewMockPaypalClient(ctrl)
	paypalClientMock.EXPECT().GetAccessToken()
	paypalWebhookValidatorMock := mock_paypal_api.NewMockPaypalWebhookValidator(ctrl)
	paypalWebhookValidatorMock.EXPECT().GetWebhookData(gomock.Any())
	paypalWebhookValidatorMock.EXPECT().ValidateWebHook(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(true, errors.New(""))

	mockResponseWriter := mock_http_.NewMockHttpResponseWriter(ctrl)

	req := util.SetupTestRequest(nil, "123")

	handler := paypal_api.CreatePapalWebhookHandler(paypalClientMock, paypalWebhookValidatorMock)
	_, err := handler.HandlePaypalWebhooks(mockResponseWriter, req)
	if err == nil {
		t.Error()
	}
}

func TestHandlePaypalWebhooks_Invalid(t *testing.T) {
	ctrl := gomock.NewController(t)

	paypalClientMock := mock_paypal_api.NewMockPaypalClient(ctrl)
	paypalClientMock.EXPECT().GetAccessToken()
	paypalWebhookValidatorMock := mock_paypal_api.NewMockPaypalWebhookValidator(ctrl)
	paypalWebhookValidatorMock.EXPECT().GetWebhookData(gomock.Any())
	paypalWebhookValidatorMock.EXPECT().ValidateWebHook(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)

	mockResponseWriter := mock_http_.NewMockHttpResponseWriter(ctrl)

	req := util.SetupTestRequest(nil, "123")

	handler := paypal_api.CreatePapalWebhookHandler(paypalClientMock, paypalWebhookValidatorMock)
	_, err := handler.HandlePaypalWebhooks(mockResponseWriter, req)
	if err == nil {
		t.Error()
	}
}
