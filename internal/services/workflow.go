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
	User   uuid.UUID
	Amount float32
}

func PaymentWorkFlow(ctx workflow.Context, param PaymentWorkFlowParam) (PaymentWorkFlowResult, error) {
	var resultHolder PaymentWorkFlowResult
	workflow.GetLogger(ctx).Info("Starting workflow :wow:")

	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}

	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// Setting Param for TX
	activityParams := activities.PaymentActivity{
		Transaction_id: uuid.New(),
		Status:         activities.Created,
	}

	workflow.GetLogger(ctx).Info(fmt.Sprintf("%v: %v",
		activityParams.Transaction_id,
		activityParams.Amount))

	// Activity flow starts
	workflow.GetLogger(ctx).Info("Started transaction")
	workflow.ExecuteActivity(ctx, activityParams.StartTransaction).Get(ctx, &activityParams)

	workflow.GetLogger(ctx).Info("Checking balance")
	workflow.ExecuteActivity(ctx, activityParams.CheckBalance).Get(ctx, &activityParams)

	workflow.GetLogger(ctx).Info("Reserving funds")
	workflow.ExecuteActivity(ctx, activityParams.ReserveFunds).Get(ctx, &activityParams)

	workflow.GetLogger(ctx).Info("Withdrawing funds")
	workflow.ExecuteActivity(ctx, activityParams.WithdrawFunds).Get(ctx, &activityParams)

	return resultHolder, nil
}
