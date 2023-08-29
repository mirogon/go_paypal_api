package paypal_api

import (
	"github.com/logpacker/paypal-go-sdk"
	paypal_api_data "github.com/mirogon/go_paypal_api/data"
)

type PaypalClient interface {
	CreateOrder(referenceId string, price string, buyerFirstName string, buyerLastName string, buyerEmail string) (*paypal.Order, error)
	GetOrder(orderId string) (*paypal.Order, error)
	CaptureOrder(orderId string) (*paypal.CaptureOrderResponse, error)
	GetAccessToken() string
	CreateProduct(productName string, productType string) (paypal_api_data.CreateProductResponse, error)
	CreateBillingPlan(productId string, pricePerMonth string, name string, description string) (paypal_api_data.CreateBillingPlanResponse, error)
	CreateSubscription(planId string) (paypal_api_data.CreateSubscriptionResponse, error)
	CancelSubscription(subscriptionId string) error
	ShowSubscriptionDetails(subscriptionId string) (paypal_api_data.ShowSubscriptionDetailsResponse, error)
	CaptureSubscription(subscriptionId string, amount string) error
	GetSubscriptionTransactions(subscriptionId string, startTime string, endTime string) (paypal_api_data.GetSubscriptionTransactionsResponse, error)
}
