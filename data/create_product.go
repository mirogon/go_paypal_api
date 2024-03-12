package paypal_api_data

type CreateProductRequest struct {
	Name        string `json:"name"` //REQUIRED
	Type        string `json:"type"` //REQUIRED
	Description string `json:"description"`
	Category    string `json:"category"`
	//ImageUrl    string `json:"image_url"`
	//HomeUrl     string `json:"home_url"`
}

type CreateProductResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	ImageUrl    string `json:"image_url"`
	HomeUrl     string `json:"home_url"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	Links       []Link `json:"links"`
}
