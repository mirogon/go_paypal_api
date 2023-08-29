package paypal_api_data

type ShowSubscriptionDetailsResponse struct {
	Id               string      `json:"id"`
	PlanId           string      `json:"plan_id"`
	StartTime        string      `json:"start_time"`
	Quantity         string      `json:"quantity"`
	ShippingAmount   FixedPrice  `json:"shipping_amount"`
	Subscriber       Subscriber  `json:"subscriber"`
	BillingInfo      BillingInfo `json:"billing_info"`
	CreateTime       string      `json:"create_time"`
	UpdateTime       string      `json:"update_time"`
	Links            []Link      `json:"links"`
	Status           string      `json:"status"`
	StatusUpdateTime string      `json:"status_update_time"`
}

type BillingInfo struct {
	OutstandingBalance  FixedPrice       `json:"outstanding_balance"`
	CycleExecutions     []CycleExecution `json:"cycle_executions"`
	LastPayment         LastPayment      `json:"last_payment"`
	NextBillingTime     string           `json:"next_billing_time"`
	FailedPaymentsCount int              `json:"failed_payments_count"`
}

type CycleExecution struct {
	TenureType      string `json:"tenure_type"`
	Sequence        int    `json:"sequence"`
	CyclesCompleted int    `json:"cycles_completed"`
	CyclesRemaining int    `json:"cycles_remaining"`
	TotalCycles     int    `json:"total_cycles"`
}

type LastPayment struct {
	Amount FixedPrice
	Time   string
}
