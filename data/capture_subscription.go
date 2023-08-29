package paypal_api_data

type CaptureSubscriptionRequest struct {
	Note        string     `json:"note"`
	CaptureType string     `json:"capture_type"`
	Amount      FixedPrice `json:"amount"`
}
