package paypal_api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/logpacker/paypal-go-sdk"
	es "github.com/mirogon/go_error_system"
	http_ "github.com/mirogon/go_http"
	paypal_api_data "github.com/mirogon/go_paypal_api/data"
)

type PaypalClientImpl struct {
	paypalClient *paypal.Client
	AccessToken  string
	ApiBase      string
	isSandbox    bool
}

func CreatePaypalClient(clientId string, clientSecret string, apiBase string, isSandbox bool) (PaypalClientImpl, error) {
	paypalClient, err := paypal.NewClient(clientId, clientSecret, apiBase)
	if err != nil {
		return PaypalClientImpl{}, err
	}
	token, _ := paypalClient.GetAccessToken()
	return PaypalClientImpl{paypalClient: paypalClient, AccessToken: token.Token, ApiBase: apiBase, isSandbox: isSandbox}, nil
}

func (paypalClient PaypalClientImpl) IsSandbox() bool {
	return paypalClient.isSandbox
}

func (paypalClient PaypalClientImpl) GetOrder(orderId string) (*paypal.Order, error) {
	return paypalClient.paypalClient.GetOrder(orderId)
}

func (paypalClient PaypalClientImpl) CreateOrder(referenceId string, price string, buyerFirstName string, buyerLastName string, buyerEmail string, intent string, brandName string, returnUrl string, cancelUrl string) (*paypal.Order, error) {
	purchaseUnitAmount := paypal.PurchaseUnitAmount{Currency: "USD", Value: price}
	purchaseUnits := []paypal.PurchaseUnitRequest{{ReferenceID: referenceId, Amount: &purchaseUnitAmount}}
	payerName := paypal.CreateOrderPayerName{GivenName: buyerFirstName, Surname: buyerLastName}
	payer := paypal.CreateOrderPayer{Name: &payerName, EmailAddress: buyerEmail}
	appContext := paypal.ApplicationContext{UserAction: "PAY_NOW", BrandName: brandName, ReturnURL: returnUrl, CancelURL: cancelUrl}
	order, err := paypalClient.paypalClient.CreateOrder(intent, purchaseUnits, &payer, &appContext)
	return order, err
}

func (paypalClient PaypalClientImpl) CaptureOrder(orderId string) (*paypal.CaptureOrderResponse, error) {
	log.Printf("Capture order with id: %s", orderId)
	return paypalClient.paypalClient.CaptureOrder(orderId, paypal.CaptureOrderRequest{})
}

func (paypalClient PaypalClientImpl) GetAccessToken() string {
	return paypalClient.AccessToken
}

func (paypalClient PaypalClientImpl) CreateProduct(productName string, productType string) (paypal_api_data.CreateProductResponse, error) {
	createProductReq := paypal_api_data.CreateProductRequest{
		Name:        productName,
		Type:        productType,
		Description: "None",
		Category:    "SOFTWARE",
	}

	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/catalogs/products", createProductReq, paypalClient.AccessToken)
	if err != nil {
		return paypal_api_data.CreateProductResponse{}, err
	}

	responseBody, err := getResponseBody[paypal_api_data.CreateProductResponse](response)
	if err != nil {
		return responseBody, err
	}

	return responseBody, nil
}

func (paypalClient PaypalClientImpl) CreateBillingPlan(productId string, pricePerMonth string, name string, description string) (paypal_api_data.CreateBillingPlanResponse, error) {
	frequency := paypal_api_data.BillingFrequency{IntervalUnit: "MONTH", IntervalCount: 1}
	fixedPrice := paypal_api_data.FixedPrice{Value: pricePerMonth, CurrencyCode: "USD"}
	pricingScheme := paypal_api_data.PricingScheme{FixedPrice: fixedPrice}
	billingCycle := paypal_api_data.BillingCycle{TenureType: "REGULAR", Sequence: 1, Frequency: frequency, TotalCycles: 0, PricingScheme: pricingScheme}
	createPlanRequest := paypal_api_data.CreateBillingPlanRequest{
		ProductId:     productId,
		Name:          name,
		BillingCycles: []paypal_api_data.BillingCycle{billingCycle},
		PaymentPreferences: paypal_api_data.PaymentPreferences{
			AutoBillOutstanding:      true,
			SetupFee:                 paypal_api_data.FixedPrice{Value: "0", CurrencyCode: "USD"},
			SetupFeeFailureAction:    "CANCEL",
			PaymentFailureThreshhold: 2,
		},
		Description: description,
		Status:      "ACTIVE",
	}

	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/billing/plans", createPlanRequest, paypalClient.AccessToken)
	if err != nil {
		return paypal_api_data.CreateBillingPlanResponse{}, err
	}

	responseBody, err := getResponseBody[paypal_api_data.CreateBillingPlanResponse](response)
	if err != nil {
		return responseBody, err
	}
	return responseBody, nil
}

func (paypalClient PaypalClientImpl) CreateSubscription(planId string) (paypal_api_data.CreateSubscriptionResponse, es.Error) {
	createSubRequest := paypal_api_data.CreateSubscriptionRequestRequiredOnly{PlanId: planId}
	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/billing/subscriptions", createSubRequest, paypalClient.AccessToken)
	if err != nil {
		return paypal_api_data.CreateSubscriptionResponse{}, es.NewError("dAxHC9", "CreateSubscription_", err)
	}

	responseBody, err := getResponseBody[paypal_api_data.CreateSubscriptionResponse](response)
	if err != nil {
		return responseBody, es.NewError("TCmoq1", "CreateSubscription_", err)
	}
	return responseBody, nil
}

func (paypalClient PaypalClientImpl) CancelSubscription(subscriptionId string) error {
	cancelRequest := paypal_api_data.RequestCancelSubscription{Reason: "Not satisfied with the service"}
	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/billing/subscriptions/"+subscriptionId+"/cancel", cancelRequest, paypalClient.AccessToken)
	if err != nil {
		return err
	}

	if response.StatusCode != 204 {
		return err
	}

	return nil
}

func (paypalClient PaypalClientImpl) ShowSubscriptionDetails(subscriptionId string) (paypal_api_data.ShowSubscriptionDetailsResponse, es.Error) {
	response, err := sendRequest("GET", paypalClient.ApiBase+"/v1/billing/subscriptions/"+subscriptionId, nil, paypalClient.AccessToken)
	if err != nil {
		return paypal_api_data.ShowSubscriptionDetailsResponse{}, es.NewError("f93ff0", "ShowSubscriptionDetails_", err)
	}

	responseBody, err := getResponseBody[paypal_api_data.ShowSubscriptionDetailsResponse](response)
	if err != nil {
		return paypal_api_data.ShowSubscriptionDetailsResponse{}, es.NewError("FbP9ia", "ShowSubscriptionDetails_", err)
	}

	return responseBody, nil
}

func (paypalClient PaypalClientImpl) CaptureSubscription(subscriptionId string, amount string) error {
	requestBody := paypal_api_data.CaptureSubscriptionRequest{}
	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/billing/subscriptions/"+subscriptionId+"/capture", requestBody, paypalClient.AccessToken)
	if err != nil {
		return err
	}
	if response.StatusCode != 202 {
		return err
	}
	return nil
}

func (paypalClient PaypalClientImpl) GetSubscriptionTransactions(subscriptionId string, startTime string, endTime string) (paypal_api_data.GetSubscriptionTransactionsResponse, error) {
	response, err := sendRequest("GET", paypalClient.ApiBase+"/v1/billing/subscriptions/"+subscriptionId+"/transactions?start_time="+startTime+"&end_time="+endTime, nil, paypalClient.AccessToken)
	if err != nil {
		return paypal_api_data.GetSubscriptionTransactionsResponse{}, err
	}

	responseBody, err := getResponseBody[paypal_api_data.GetSubscriptionTransactionsResponse](response)
	if err != nil {
		return paypal_api_data.GetSubscriptionTransactionsResponse{}, err
	}

	return responseBody, nil
}

func sendRequest(requestMethod string, requestUrl string, requestBody interface{}, paypalAccessToken string) (*http.Response, es.Error) {
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, es.NewError("mcOCIo", "SendRequest_Marshal_"+err.Error(), nil)
	}

	stringReader := strings.NewReader(string(requestJson))
	request, err := http.NewRequest(requestMethod, requestUrl, stringReader)
	if err != nil {
		return nil, es.NewError("JlIoqt", "SendRequest_NewRequest_"+err.Error(), nil)
	}

	request.Header.Add("Authorization", "Bearer "+paypalAccessToken)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	requestSender := http_.HttpRequestSenderImpl{}
	response, err := requestSender.SendRequest(request)
	if err != nil {
		return nil, es.NewError("CghUaD", "SendRequest_SendRequest_"+err.Error(), nil)
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, es.NewError("DiT6En", "SendRequest_ResponseStatusNotOk_"+fmt.Sprintf("%d", response.StatusCode), nil)
	}

	return response, nil
}

func getResponseBody[responseType any](response *http.Response) (responseType, es.Error) {
	buffer, err := io.ReadAll(response.Body)
	var responseBody responseType
	if err != nil {
		return responseBody, es.NewError("Oh6D89", "GetResponseBody_ReadAll_ "+err.Error(), nil)
	}

	err = json.Unmarshal(buffer, &responseBody)
	if err != nil {
		return responseBody, es.NewError("JJKf90", "GetResponseBody_Unmarshal_"+err.Error(), nil)
	}
	return responseBody, nil
}
