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
	err := workflow.ExecuteActivity(ctx, "StartTransaction", activityParams).Get(ctx, &activityParams)
	if err != nil {
		return resultHolder, err
	}

	workflow.GetLogger(ctx).Info("Checking balance")
	err = workflow.ExecuteActivity(ctx, "CheckBalance", activityParams).Get(ctx, &activityParams)
	if err != nil {
		return resultHolder, err
	}

	workflow.GetLogger(ctx).Info("Reserving funds")
	err = workflow.ExecuteActivity(ctx, "ReserveFunds", activityParams).Get(ctx, &activityParams)
	if err != nil {
		return resultHolder, err
	}

	workflow.GetLogger(ctx).Info("Withdrawing funds")
	err = workflow.ExecuteActivity(ctx, "WithdrawFunds", activityParams).Get(ctx, &activityParams)
	if err != nil {
		return resultHolder, err
	}

	return resultHolder, nil
}
