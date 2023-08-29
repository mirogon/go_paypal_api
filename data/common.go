package paypal_api_data

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type Name struct {
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
}

type FullName struct {
	FullName string `json:"full_name"`
}
