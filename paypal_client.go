package paypal_api

import (
	"github.com/logpacker/paypal-go-sdk"
	es "github.com/mirogon/go_error_system"
	paypal_api_data "github.com/mirogon/go_paypal_api/data"
)

type PaypalClient interface {
	UpdateToken()
	IsSandbox() bool
	CreateOrder(referenceId string, price string, buyerFirstName string, buyerLastName string, buyerEmail string, intent string, brandName string, returnUrl string, cancelUrl string) (*paypal.Order, es.Error)
	GetOrder(orderId string) (*paypal.Order, es.Error)
	CaptureOrder(orderId string) (*paypal.CaptureOrderResponse, es.Error)
	GetAccessToken() string
	CreateProduct(productName string, productType string) (paypal_api_data.CreateProductResponse, es.Error)
	CreateBillingPlan(productId string, pricePerMonth string, name string, description string) (paypal_api_data.CreateBillingPlanResponse, es.Error)
	CreateSubscription(planId string) (paypal_api_data.CreateSubscriptionResponse, es.Error)
	CancelSubscription(subscriptionId string) es.Error
	ShowSubscriptionDetails(subscriptionId string) (paypal_api_data.ShowSubscriptionDetailsResponse, es.Error)
	CaptureSubscription(subscriptionId string, amount string) es.Error
	GetSubscriptionTransactions(subscriptionId string, startTime string, endTime string) (paypal_api_data.GetSubscriptionTransactionsResponse, es.Error)
}
