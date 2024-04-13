package paypal_api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/logpacker/paypal-go-sdk"
	es "github.com/mirogon/go_error_system"
	http_ "github.com/mirogon/go_http"
	paypal_api_data "github.com/mirogon/go_paypal_api/data"
)

type PaypalClientImpl struct {
	PaypalClient *paypal.Client
	ApiBase      string
	isSandbox    bool
}

func CreatePaypalClient(clientId string, clientSecret string, apiBase string, isSandbox bool) (PaypalClientImpl, es.Error) {
	paypalClient, err := paypal.NewClient(clientId, clientSecret, apiBase)
	if err != nil {
		return PaypalClientImpl{}, es.NewError("lKtWIm", "CreatePaypalClient_NewClient_"+err.Error(), nil)
	}
	_, err = paypalClient.GetAccessToken()
	if err != nil {
		return PaypalClientImpl{}, es.NewError("XyR3Wd", "CreatePaypalClient_GetAccessToken_"+err.Error(), nil)
	}
	client := PaypalClientImpl{PaypalClient: paypalClient, ApiBase: apiBase, isSandbox: isSandbox}
	go UpdateTokenEveryHour(&client)
	return client, nil
}

func (client PaypalClientImpl) AccessToken() string {
	return client.PaypalClient.Token.Token
}

func UpdateTokenEveryHour(paypalClient *PaypalClientImpl) {
	dur, _ := time.ParseDuration("30s")
	for {
		time.Sleep(dur)
		paypalClient.UpdateToken()
	}
}

func (paypalClient *PaypalClientImpl) UpdateToken() es.Error {
	_, err := paypalClient.PaypalClient.GetAccessToken()
	fmt.Println("New Token: " + paypalClient.AccessToken())
	if err != nil {
		return es.NewError("ABHh56", "UpdateToken_"+err.Error(), nil)
	}
	return nil
}

func (paypalClient PaypalClientImpl) IsSandbox() bool {
	return paypalClient.isSandbox
}

func (paypalClient PaypalClientImpl) GetOrder(orderId string) (*paypal.Order, es.Error) {
	order, err := paypalClient.PaypalClient.GetOrder(orderId)
	if err != nil {
		return order, es.NewError("as7oXs", "GetOrder_"+err.Error(), nil)
	}

	return order, nil
}

func (paypalClient PaypalClientImpl) CreateOrder(referenceId string, price string, buyerFirstName string, buyerLastName string, buyerEmail string, intent string, brandName string, returnUrl string, cancelUrl string) (*paypal.Order, es.Error) {
	purchaseUnitAmount := paypal.PurchaseUnitAmount{Currency: "USD", Value: price}
	purchaseUnits := []paypal.PurchaseUnitRequest{{ReferenceID: referenceId, Amount: &purchaseUnitAmount}}
	payerName := paypal.CreateOrderPayerName{GivenName: buyerFirstName, Surname: buyerLastName}
	payer := paypal.CreateOrderPayer{Name: &payerName, EmailAddress: buyerEmail}
	appContext := paypal.ApplicationContext{UserAction: "PAY_NOW", BrandName: brandName, ReturnURL: returnUrl, CancelURL: cancelUrl}
	order, err := paypalClient.PaypalClient.CreateOrder(intent, purchaseUnits, &payer, &appContext)
	return order, es.NewError("Q2Z6ju", "CreateOrder_"+err.Error(), nil)
}

func (paypalClient PaypalClientImpl) CaptureOrder(orderId string) (*paypal.CaptureOrderResponse, es.Error) {
	log.Printf("Capture order with id: %s", orderId)
	resp, err := paypalClient.PaypalClient.CaptureOrder(orderId, paypal.CaptureOrderRequest{})
	if err != nil {
		return resp, es.NewError("MHRHfj", "CaptureOrder_"+err.Error(), nil)
	}
	return resp, nil
}

func (paypalClient PaypalClientImpl) GetAccessToken() string {
	return paypalClient.PaypalClient.Token.Token
}

func (paypalClient PaypalClientImpl) CreateProduct(productName string, productType string) (paypal_api_data.CreateProductResponse, es.Error) {
	createProductReq := paypal_api_data.CreateProductRequest{
		Name:        productName,
		Type:        productType,
		Description: "None",
		Category:    "SOFTWARE",
	}

	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/catalogs/products", createProductReq, paypalClient.AccessToken())
	if err != nil {
		return paypal_api_data.CreateProductResponse{}, es.NewError("3XUvQb", "CreateProduct_", err)
	}

	responseBody, err := getResponseBody[paypal_api_data.CreateProductResponse](response)
	if err != nil {
		return responseBody, es.NewError("LnEbsw", "CreateProduct_", err)
	}

	return responseBody, nil
}

func (paypalClient PaypalClientImpl) CreateBillingPlan(productId string, pricePerMonth string, name string, description string) (paypal_api_data.CreateBillingPlanResponse, es.Error) {
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

	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/billing/plans", createPlanRequest, paypalClient.AccessToken())
	if err != nil {
		return paypal_api_data.CreateBillingPlanResponse{}, es.NewError("oZwLmv", "CreateBillingPlan_", err)
	}

	responseBody, err := getResponseBody[paypal_api_data.CreateBillingPlanResponse](response)
	if err != nil {
		return responseBody, es.NewError("IR61kK", "CreateBillingPlan_", err)
	}
	return responseBody, nil
}

func (paypalClient PaypalClientImpl) CreateSubscription(planId string) (paypal_api_data.CreateSubscriptionResponse, es.Error) {
	createSubRequest := paypal_api_data.CreateSubscriptionRequestRequiredOnly{PlanId: planId}
	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/billing/subscriptions", createSubRequest, paypalClient.AccessToken())
	if err != nil {
		return paypal_api_data.CreateSubscriptionResponse{}, es.NewError("dAxHC9", "CreateSubscription_", err)
	}

	responseBody, err := getResponseBody[paypal_api_data.CreateSubscriptionResponse](response)
	if err != nil {
		return responseBody, es.NewError("TCmoq1", "CreateSubscription_", err)
	}
	return responseBody, nil
}

func (paypalClient PaypalClientImpl) CancelSubscription(subscriptionId string) es.Error {
	cancelRequest := paypal_api_data.RequestCancelSubscription{Reason: "Not satisfied with the service"}
	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/billing/subscriptions/"+subscriptionId+"/cancel", cancelRequest, paypalClient.AccessToken())
	if err != nil {
		return es.NewError("GDD0Xg", "CancelSubscription_", err)
	}

	if response.StatusCode != 204 {
		return es.NewError("8D26Nj", "CancelSubscription_InvalidStatusCode", nil)
	}

	return nil
}

func (paypalClient PaypalClientImpl) ShowSubscriptionDetails(subscriptionId string) (paypal_api_data.ShowSubscriptionDetailsResponse, es.Error) {
	response, err := sendRequest("GET", paypalClient.ApiBase+"/v1/billing/subscriptions/"+subscriptionId, nil, paypalClient.AccessToken())
	if err != nil {
		return paypal_api_data.ShowSubscriptionDetailsResponse{}, es.NewError("f93ff0", "ShowSubscriptionDetails_", err)
	}

	responseBody, err := getResponseBody[paypal_api_data.ShowSubscriptionDetailsResponse](response)
	if err != nil {
		return paypal_api_data.ShowSubscriptionDetailsResponse{}, es.NewError("FbP9ia", "ShowSubscriptionDetails_", err)
	}

	return responseBody, nil
}

func (paypalClient PaypalClientImpl) CaptureSubscription(subscriptionId string, amount string) es.Error {
	requestBody := paypal_api_data.CaptureSubscriptionRequest{}
	response, err := sendRequest("POST", paypalClient.ApiBase+"/v1/billing/subscriptions/"+subscriptionId+"/capture", requestBody, paypalClient.AccessToken())
	if err != nil {
		return es.NewError("yX93IT", "CaptureSubscription_", err)
	}
	if response.StatusCode != 202 {
		return es.NewError("GufcqQ", "CaptureSubscription_", err)
	}
	return nil
}

func (paypalClient PaypalClientImpl) GetSubscriptionTransactions(subscriptionId string, startTime string, endTime string) (paypal_api_data.GetSubscriptionTransactionsResponse, es.Error) {
	response, err := sendRequest("GET", paypalClient.ApiBase+"/v1/billing/subscriptions/"+subscriptionId+"/transactions?start_time="+startTime+"&end_time="+endTime, nil, paypalClient.AccessToken())
	if err != nil {
		return paypal_api_data.GetSubscriptionTransactionsResponse{}, es.NewError("LIpk2W", "GetSubscriptionTransactions_", err)
	}

	responseBody, err := getResponseBody[paypal_api_data.GetSubscriptionTransactionsResponse](response)
	if err != nil {
		return paypal_api_data.GetSubscriptionTransactionsResponse{}, es.NewError("diIQfW", "GetSubscriptionTransactions_", err)
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
		respBody, _ := io.ReadAll(response.Body)
		return nil, es.NewError("DiT6En", "SendRequest_ResponseStatusNotOk_"+fmt.Sprintf("%d", response.StatusCode)+"_"+response.Status+"_"+(string)(respBody), nil)
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
