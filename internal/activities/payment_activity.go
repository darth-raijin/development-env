package activities

import (
	"context"

	"github.com/google/uuid"
)

type PaymentActivity struct {
	transaction_id uuid.UUID
	status         PaymentState
	amount         int
	user           uuid.UUID
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
	param.transaction_id = uuid.New()
	param.status = Created

	return param, nil

}

func (a *PaymentActivity) CheckBalance(ctx context.Context, param *PaymentActivity) (*PaymentActivity, error) {
	// Check in repo if TX is active, otherwise can -> Based on single or combinaion of keys
	// Card details, sub from JWT or something else
	param.status = Ready

	return param, nil

}

func (a *PaymentActivity) ReserveFunds(ctx context.Context, param PaymentActivity) (*PaymentActivity, error) {
	param.status = Reserved
	return &param, nil
}

func (a *PaymentActivity) WithdrawFunds(ctx context.Context, param PaymentActivity) (*PaymentActivity, error) {
	param.status = Withdrawn

	return &param, nil
}
