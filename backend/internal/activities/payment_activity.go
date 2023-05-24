package activities

import (
	"context"

	"github.com/google/uuid"
)

type PaymentActivity struct {
	Transaction_id uuid.UUID
	Status         PaymentState
	Amount         int
	User           uuid.UUID
}

type PaymentState string

const (
	Unknown   PaymentState = ""
	Created   PaymentState = "created"
	Ready     PaymentState = "ready"
	Reserved  PaymentState = "reserved"
	Withdrawn PaymentState = "withdrawn"
)

func (a *PaymentActivity) StartTransaction(ctx context.Context, param *PaymentActivity) (*PaymentActivity, error) {
	if param.Status == Created {
		return param, ctx.Err() // Find ud af hvordan med fejl
	}

	param.Status = Created

	return param, nil

}

func (a *PaymentActivity) CheckBalance(ctx context.Context, param *PaymentActivity) (*PaymentActivity, error) {
	// Check in repo if TX is active, otherwise can -> Based on single or combinaion of keys
	// Card details, sub from JWT or something else
	param.Status = Ready

	return param, nil

}

func (a *PaymentActivity) ReserveFunds(ctx context.Context, param PaymentActivity) (*PaymentActivity, error) {
	param.Status = Reserved
	return &param, nil
}

func (a *PaymentActivity) WithdrawFunds(ctx context.Context, param PaymentActivity) (*PaymentActivity, error) {
	param.Status = Withdrawn

	return &param, nil
}
