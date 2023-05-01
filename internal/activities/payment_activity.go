package activities

import (
	"context"

	"github.com/google/uuid"
)

type PaymentActivity struct {
	transaction_id uuid.UUID
	state          string
	amount         int
	user           uuid.UUID
}

func (a *PaymentActivity) CheckBalance(ctx context.Context)

func (a *PaymentActivity) ReserveFunds(ctx context.Context)

func (a *PaymentActivity) WithdrawFunds(ctx context.Context)
