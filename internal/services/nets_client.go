package services

import "github.com/go-temporal-laboratory/internal/models"

func NetsCreatePayment() models.NetsPayment {
	payment := models.NetsPayment{
		Order: models.Order{
			Items: []models.Item{
				{
					Reference:        "Ref1234",
					Name:             "Product Name",
					Quantity:         2,
					Unit:             "pcs",
					UnitPrice:        100,
					GrossTotalAmount: 200,
					NetTotalAmount:   180,
				},
				{
					Reference:        "Ref5678",
					Name:             "Another Product",
					Quantity:         1,
					Unit:             "pcs",
					UnitPrice:        50,
					GrossTotalAmount: 50,
					NetTotalAmount:   45,
				},
			},
			Amount:   250,
			Currency: "USD",
		},
		Checkout: models.Checkout{
			URL:      "https://checkout.example.com",
			TermsURL: "https://checkout.example.com/terms",
		},
	}

	return payment
}
