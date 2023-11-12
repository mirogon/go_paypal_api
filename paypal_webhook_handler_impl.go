package paypal_api

import (
	"errors"
	"net/http"

	http_ "github.com/mirogon/go_http"
	paypal_api_data "github.com/mirogon/go_paypal_api/data"
)

type WebhookNotification struct {
	Id            string
	Create_Time   string
	Resource_Type string
	Summary       string
	Resource      WebhookNotificationResource
	Links         []paypal_api_data.Link
}

type WebhookNotificationResource struct {
	Id                 string
	BillingAgreementId string `json:"billing_agreement_id"`
	Create_Time        string
	Update_Time        string
	State              string
	Amount             WebhookNotificationResourceAmount
	Parent_Payment     string
	Valid_Until        string
	Links              []paypal_api_data.Link
}

type WebhookNotificationResourceAmount struct {
	Total    string
	Currency string
}

type PaypalWebhookHandlerImpl struct {
	PaypalClient           PaypalClient
	PaypalWebhookValidator PaypalWebhookValidator
	ApiBase                string
}

func CreatePapalWebhookHandler(paypalClient PaypalClient, validator PaypalWebhookValidator, apiBase string) PaypalWebhookHandlerImpl {
	return PaypalWebhookHandlerImpl{PaypalClient: paypalClient, PaypalWebhookValidator: validator, ApiBase: apiBase}
}

func (handler PaypalWebhookHandlerImpl) HandlePaypalWebhooks(responseWriter http_.HttpResponseWriter, req *http.Request) (WebhookNotification, error) {
	validationData, webhookNotification, err := handler.PaypalWebhookValidator.GetWebhookData(req)
	if err != nil {
		return webhookNotification, err
	}
	isValid, err := handler.PaypalWebhookValidator.ValidateWebHook(validationData, webhookNotification.Id, http_.HttpRequestSenderImpl{}, handler.ApiBase+"/v1/notifications/verify-webhook-signature", handler.PaypalClient.GetAccessToken())
	if !isValid {
		return webhookNotification, errors.New("Invalid")
	}
	if err != nil {
		return webhookNotification, err
	}
	responseWriter.Send()
	return webhookNotification, nil
}
