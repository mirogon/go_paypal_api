package paypal_api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	mock_http_ "github.com/mirogon/go_http/mocks"
	paypal_api "github.com/mirogon/go_paypal_api"
)

func TestValidateWebhook(t *testing.T) {
	ctrl := gomock.NewController(t)

	requestSenderMock := mock_http_.NewMockHttpRequestSender(ctrl)

	mockResponse := paypal_api.WebhookValidationResponse{VerificationStatus: "SUCCESS"}
	mockResponseJson, _ := json.Marshal(mockResponse)
	mockHttpResponse := http.Response{
		Body:          io.NopCloser(bytes.NewBuffer(mockResponseJson)),
		ContentLength: int64(len(mockResponseJson)),
	}

	validationRequest := paypal_api.WebhookValidationRequest{
		AuthAlgo:         "SHA256",
		TransmissionId:   "transmissionId",
		TransmissionSig:  "transmissionSig",
		TransmissionTime: "transmissionTime",
		CertUrl:          "certUrl",
		WebhookId:        "webhookId",
	}

	expectedRequestBodyJson, _ := json.Marshal(validationRequest)
	expectedRequestBodyJsonReader := strings.NewReader(string(expectedRequestBodyJson))
	expectedHttpRequest, _ := http.NewRequest("POST", "https://example.com", expectedRequestBodyJsonReader)
	expectedHttpRequest.Header.Add("Content-Type", "application/json")
	expectedHttpRequest.Header.Add("Authorization", "Bearer "+"paypalToken")

	requestSenderMock.EXPECT().SendRequest(gomock.Any()).Return(&mockHttpResponse, nil)
	validator := paypal_api.PaypalWebhookValidatorImpl{}
	result, err := validator.ValidateWebHook(validationRequest, "webhookId", requestSenderMock, "https://example.com", "paypalToken")
	if result == false || err != nil {
		t.Error()
	}
}

func TestValidateWebhook_SendRequestFailure(t *testing.T) {
	ctrl := gomock.NewController(t)

	requestSenderMock := mock_http_.NewMockHttpRequestSender(ctrl)

	mockResponse := paypal_api.WebhookValidationResponse{VerificationStatus: "SUCCESS"}
	mockResponseJson, _ := json.Marshal(mockResponse)
	mockHttpResponse := http.Response{
		Body:          io.NopCloser(bytes.NewBuffer(mockResponseJson)),
		ContentLength: int64(len(mockResponseJson)),
	}

	validationRequest := paypal_api.WebhookValidationRequest{
		AuthAlgo:         "SHA256",
		TransmissionId:   "transmissionId",
		TransmissionSig:  "transmissionSig",
		TransmissionTime: "transmissionTime",
		CertUrl:          "certUrl",
		WebhookId:        "webhookId",
	}

	requestSenderMock.EXPECT().SendRequest(gomock.Any()).Return(&mockHttpResponse, errors.New(""))
	validator := paypal_api.PaypalWebhookValidatorImpl{}
	result, _ := validator.ValidateWebHook(validationRequest, "webhookId", requestSenderMock, "https://example.com", "paypalToken")
	if result == true {
		t.Error()
	}
}

func TestValidateWebhook_NoSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)

	requestSenderMock := mock_http_.NewMockHttpRequestSender(ctrl)

	mockResponse := paypal_api.WebhookValidationResponse{VerificationStatus: "FAILURE"}
	mockResponseJson, _ := json.Marshal(mockResponse)
	mockHttpResponse := http.Response{
		Body:          io.NopCloser(bytes.NewBuffer(mockResponseJson)),
		ContentLength: int64(len(mockResponseJson)),
	}

	validationRequest := paypal_api.WebhookValidationRequest{
		AuthAlgo:         "SHA256",
		TransmissionId:   "transmissionId",
		TransmissionSig:  "transmissionSig",
		TransmissionTime: "transmissionTime",
		CertUrl:          "certUrl",
		WebhookId:        "webhookId",
	}

	expectedRequestBodyJson, _ := json.Marshal(validationRequest)
	expectedRequestBodyJsonReader := strings.NewReader(string(expectedRequestBodyJson))
	expectedHttpRequest, _ := http.NewRequest("POST", "https://example.com", expectedRequestBodyJsonReader)
	expectedHttpRequest.Header.Add("Content-Type", "application/json")
	expectedHttpRequest.Header.Add("Authorization", "Bearer "+"paypalToken")

	requestSenderMock.EXPECT().SendRequest(gomock.Any()).Return(&mockHttpResponse, nil)
	validator := paypal_api.PaypalWebhookValidatorImpl{}
	result, _ := validator.ValidateWebHook(validationRequest, "webhookId", requestSenderMock, "https://example.com", "paypalToken")
	if result == true {
		t.Error()
	}
}
