package models

type NetsPayment struct {
	Order struct {
		Items []struct {
			Reference        string `json:"reference"`
			Name             string `json:"name"`
			Quantity         int    `json:"quantity"`
			Unit             string `json:"unit"`
			UnitPrice        int    `json:"unitPrice"`
			GrossTotalAmount int    `json:"grossTotalAmount"`
			NetTotalAmount   int    `json:"netTotalAmount"`
		} `json:"items"`
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"order"`
	Checkout struct {
		URL      string `json:"url"`
		TermsURL string `json:"termsUrl"`
	} `json:"checkout"`
}
