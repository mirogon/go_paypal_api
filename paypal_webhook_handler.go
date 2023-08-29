package paypal_api

import (
	"net/http"

	http_ "github.com/mirogon/go_http"
)

type PaypalWebhookHandler interface {
	HandlePaypalWebhooks(responseWriter http_.HttpResponseWriter, req *http.Request) (WebhookNotification, error)
}
