package models

type Item struct {
	Reference        string `json:"reference"`
	Name             string `json:"name"`
	Quantity         int    `json:"quantity"`
	Unit             string `json:"unit"`
	UnitPrice        int    `json:"unitPrice"`
	GrossTotalAmount int    `json:"grossTotalAmount"`
	NetTotalAmount   int    `json:"netTotalAmount"`
}

type Order struct {
	Items    []Item `json:"items"`
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type Checkout struct {
	URL      string `json:"url"`
	TermsURL string `json:"termsUrl"`
}

type NetsPayment struct {
	Order    Order    `json:"order"`
	Checkout Checkout `json:"checkout"`
}
