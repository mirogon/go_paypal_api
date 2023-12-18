package paypal_api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/iancoleman/orderedmap"
	http_ "github.com/mirogon/go_http"
)

type PaypalWebhookValidatorImpl struct {
	PaypalWebhookId string
}

func CreatePaypalWebhookValidator(webhookId string) PaypalWebhookValidatorImpl {
	return PaypalWebhookValidatorImpl{
		PaypalWebhookId: webhookId,
	}
}

func (validator PaypalWebhookValidatorImpl) ValidateWebHook(webHookValidationReq WebhookValidationRequest, webHookId string, requestSender http_.HttpRequestSender, paypalApiAddress string, token string) (bool, error) {
	jsonValidationData, _ := json.Marshal(webHookValidationReq)
	jsonValidationDataReader := strings.NewReader(string(jsonValidationData))
	req, err := http.NewRequest("POST", paypalApiAddress, jsonValidationDataReader)
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	response, err := requestSender.SendRequest(req)
	if err != nil {
		return false, err
	}

	responseLen := response.ContentLength
	if responseLen == -1 {
		responseLen = 4096
	}

	responseBuffer := make([]byte, responseLen)
	actualLen, err := response.Body.Read(responseBuffer)
	responseBuffer = responseBuffer[:actualLen]
	if err != nil {
		//return false, err
	}

	var validationResponse WebhookValidationResponse
	json.Unmarshal(responseBuffer, &validationResponse)
	if validationResponse.VerificationStatus == "SUCCESS" {
		return true, nil
	}
	return false, err
}

func (validator PaypalWebhookValidatorImpl) GetWebhookData(req *http.Request) (WebhookValidationRequest, WebhookNotification, error) {
	paypalTransmissionId := req.Header.Get("Paypal-Transmission-Id")
	paypalTransmissionSig := req.Header.Get("Paypal-Transmission-Sig")
	paypalCertUrl := req.Header.Get("Paypal-Cert-Url")
	paypalAuthAlgo := req.Header.Get("Paypal-Auth-Algo")
	paypalTransmissionTime := req.Header.Get("Paypal-Transmission-Time")
	bodyBuffer := make([]byte, req.ContentLength)
	_, _ = req.Body.Read(bodyBuffer)

	var webHookEventData WebhookNotification
	err := json.Unmarshal(bodyBuffer, &webHookEventData)
	if err != nil {
		return WebhookValidationRequest{}, WebhookNotification{}, err
	}

	log.Printf("Webhook notification buffer string: %s", string(bodyBuffer))

	jsonObject := orderedmap.New()
	err = json.Unmarshal([]byte(bodyBuffer), &jsonObject)
	if err != nil {
		return WebhookValidationRequest{}, WebhookNotification{}, err
	}

	return WebhookValidationRequest{
		AuthAlgo:         paypalAuthAlgo,
		TransmissionId:   paypalTransmissionId,
		TransmissionSig:  paypalTransmissionSig,
		TransmissionTime: paypalTransmissionTime,
		CertUrl:          paypalCertUrl,
		WebhookId:        validator.PaypalWebhookId,
		WebhookEvent:     jsonObject,
	}, webHookEventData, nil
}
