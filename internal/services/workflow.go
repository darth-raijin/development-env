package services

import (
	"fmt"
	"time"

	"github.com/go-temporal-laboratory/internal/activities"
	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
)

type PaymentWorkFlowResult struct {
	state string
}

type PaymentWorkFlowParam struct {
	user   uuid.UUID
	amount float32
}

var (
	user = uuid.New()
)

func PaymentWorkFlow(ctx workflow.Context, param PaymentWorkFlowParam) (PaymentWorkFlowResult, error) {
	var resultHolder PaymentWorkFlowResult
	workflow.GetLogger(ctx).Info("Starting workflow :wow:")

	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}

	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	activityParams := activities.PaymentActivity{
		Transaction_id: uuid.New(),
		Status:         activities.Created,
		Amount:         500,
		User:           user,
	}

	workflow.GetLogger(ctx).Info(fmt.Sprintf("%v: %v",
		activityParams.Transaction_id,
		activityParams.Amount))

	// Activity flow starts
	workflow.ExecuteActivity(ctx, activityParams.CheckBalance)
	workflow.ExecuteActivity(ctx, activityParams.ReserveFunds)
	workflow.ExecuteActivity(ctx, activityParams.WithdrawFunds)

	return resultHolder, nil
}
