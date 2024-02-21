package paypal_api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/logpacker/paypal-go-sdk"
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
		ImageUrl:    "https://example.com/streaming.jpg",
		HomeUrl:     "https://example.com/home",
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
	frequency := paypal_api_data.BillingFrequency{IntervalUnit: "Month", IntervalCount: 1}
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
			SetupFeeFailureAction:    "CONTINUE",
			PaymentFailureThreshhold: 3,
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

func (paypalClient PaypalClientImpl) CreateSubscription(planId string) (paypal_api_data.CreateSubscriptionResponse, error) {
	createSubRequest := paypal_api_data.CreateSubscriptionRequestRequiredOnly{PlanId: planId}
	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/billing/subscriptions", createSubRequest, paypalClient.AccessToken)
	if err != nil {
		return paypal_api_data.CreateSubscriptionResponse{}, err
	}

	responseBody, err := getResponseBody[paypal_api_data.CreateSubscriptionResponse](response)
	if err != nil {
		return responseBody, err
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

func (paypalClient PaypalClientImpl) ShowSubscriptionDetails(subscriptionId string) (paypal_api_data.ShowSubscriptionDetailsResponse, error) {
	response, err := sendRequest("GET", paypalClient.ApiBase+"/v1/billing/subscriptions/"+subscriptionId, nil, paypalClient.AccessToken)
	if err != nil {
		return paypal_api_data.ShowSubscriptionDetailsResponse{}, err
	}

	responseBody, err := getResponseBody[paypal_api_data.ShowSubscriptionDetailsResponse](response)
	if err != nil {
		return paypal_api_data.ShowSubscriptionDetailsResponse{}, err
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

func sendRequest(requestMethod string, requestUrl string, requestBody interface{}, paypalAccessToken string) (*http.Response, error) {
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	stringReader := strings.NewReader(string(requestJson))
	requestSender := http_.HttpRequestSenderImpl{}
	request, err := http.NewRequest(requestMethod, requestUrl, stringReader)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+paypalAccessToken)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	response, err := requestSender.SendRequest(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func getResponseBody[responseType any](response *http.Response) (responseType, error) {
	buffer, _ := io.ReadAll(response.Body)

	var responseBody responseType
	err := json.Unmarshal(buffer, &responseBody)
	if err != nil {
		return responseBody, err
	}
	return responseBody, nil
}
