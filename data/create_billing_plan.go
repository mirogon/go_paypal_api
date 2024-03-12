package paypal_api_data

type CreateBillingPlanRequest struct {
	ProductId          string             `json:"product_id"`          //REQUIRED
	Name               string             `json:"name"`                //REQUIRED
	BillingCycles      []BillingCycle     `json:"billing_cycles"`      //REQUIRED
	PaymentPreferences PaymentPreferences `json:"payment_preferences"` //REQUIRED
	Description        string             `json:"description"`
	Status             string             `json:"status"`
}

type BillingCycle struct {
	TenureType    string           `json:"tenure_type"` //REQUIRED
	Sequence      int              `json:"sequence"`    //REQUIRED
	Frequency     BillingFrequency `json:"frequency"`   //REQUIRED
	TotalCycles   int              `json:"total_cycles"`
	PricingScheme PricingScheme    `json:"pricing_scheme"`
}

type BillingFrequency struct {
	IntervalUnit  string `json:"interval_unit"`
	IntervalCount int    `json:"interval_count"`
}

type PricingScheme struct {
	FixedPrice FixedPrice `json:"fixed_price"`
}

type FixedPrice struct {
	Value        string `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

type PaymentPreferences struct {
	AutoBillOutstanding      bool       `json:"auto_bill_outstanding"`
	SetupFee                 FixedPrice `json:"setup_fee"`
	SetupFeeFailureAction    string     `json:"setup_fee_failure_action"`
	PaymentFailureThreshhold int        `json:"payment_failure_threshold"`
}

type CreateBillingPlanResponse struct {
	Id                 string             `json:"id"`
	ProductId          string             `json:"product_id"`
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	Status             string             `json:"status"`
	BillingCycles      []BillingCycle     `json:"billing_cycles"`
	PaymentPreferences PaymentPreferences `json:"payment_preferences"`
	Taxes              Taxes              `json:"taxes"`
	CreateTime         string             `json:"create_time"`
	UpdateTime         string             `json:"update_time"`
	Links              []Link             `json:"links"`
}

type Taxes struct {
	Percentage string `json:"percentage"`
	Inclusive  bool   `json:"inclusive"`
}
