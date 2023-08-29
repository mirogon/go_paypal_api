package paypal_api

import (
	"net/http"

	http_ "github.com/mirogon/go_http"
)

type WebhookValidationRequest struct {
	AuthAlgo         string      `json:"auth_algo"`
	TransmissionId   string      `json:"transmission_id"`
	TransmissionSig  string      `json:"transmission_sig"`
	TransmissionTime string      `json:"transmission_time"`
	CertUrl          string      `json:"cert_url"`
	WebhookId        string      `json:"webhook_id"`
	WebhookEvent     interface{} `json:"webhook_event"`
}

type WebhookValidationResponse struct {
	VerificationStatus string `json:"verification_status"`
}

type PaypalWebhookValidator interface {
	ValidateWebHook(webHookValidationReq WebhookValidationRequest, webHookId string, requestSender http_.HttpRequestSender, paypalApiAddress string, token string) (bool, error)
	GetWebhookData(req *http.Request) (WebhookValidationRequest, WebhookNotification, error)
}
