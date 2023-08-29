package paypal_api_data

type GetSubscriptionTransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
	Links        []Link        `json:"links"`
}

type Transaction struct {
	Id                  string              `json:"id"`
	Status              string              `json:"status"`
	PayerEmail          string              `json:"payer_email"`
	PayerName           Name                `json:"payer_name"`
	AmountWithBreakdown AmountWithBreakdown `json:"amount_with_breakdown"`
	Time                string              `json:"time"`
}

type AmountWithBreakdown struct {
	GrossAmount FixedPrice `json:"gross_amount"`
	FeeAmount   FixedPrice `json:"fee_amount"`
	NetAmount   FixedPrice `json:"net_amount"`
}
