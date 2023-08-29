package paypal_api_data

type CreateSubscriptionRequest struct {
	PlanId             string             `json:"plan_id"` //REQUIRED
	StartTime          string             `json:"start_time"`
	Quantity           string             `json:"quantity"`
	ShippingAmount     FixedPrice         `json:"shipping_amount"`
	Subscriber         Subscriber         `json:"subscriber"`
	ApplicationContext ApplicationContext `json:"application_context"`
}

type CreateSubscriptionRequestRequiredOnly struct {
	PlanId string `json:"plan_id"`
}

type Subscriber struct {
	Name            Name            `json:"name"`
	EmailAddress    string          `json:"email_address"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type ShippingAddress struct {
	Name    FullName `json:"name"`
	Address Address  `json:"address"`
}

type Address struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	AdminArea2   string `json:"admin_area_2"`
	AdminArea1   string `json:"admin_area_1"`
	PostalCode   string `json:"postal_code"`
	CountryCode  string `json:"country_code"`
}

type ApplicationContext struct {
	BrandName          string        `json:"brand_name"`
	Locale             string        `json:"locale"`
	ShippingPreference string        `json:"shipping_preference"`
	UserAction         string        `json:"user_action"`
	PaymentMethod      PaymentMethod `json:"payment_method"`
	ReturnUrl          string        `json:"return_url"`
	CancelUrl          string        `json:"cancel_url"`
}

type PaymentMethod struct {
	PayerSelected  string `json:"payer_selected"`
	PayeePreferred string `json:"payee_preferred"`
}

type CreateSubscriptionResponse struct {
	Id               string     `json:"id"`
	Status           string     `json:"status"`
	StatusUpdateTime string     `json:"status_update_time"`
	PlanId           string     `json:"plan_id"`
	PlanOverridden   bool       `json:"plan_overridden"`
	StartTime        string     `json:"start_time"`
	Quantity         string     `json:"quantity"`
	ShippingAmount   FixedPrice `json:"shipping_amount"`
	Subscriber       Subscriber `json:"subscriber"`
	CreateTime       string     `json:"create_time"`
	Links            []Link     `json:"links"`
}
