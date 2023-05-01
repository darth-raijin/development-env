package services

import (
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

func PaymentWorkFlow(ctx workflow.Context, param PaymentWorkFlowParam) (PaymentWorkFlowResult, error) {
	var resultHolder PaymentWorkFlowResult
	workflow.GetLogger(ctx).Info("Starting workflow :wow:")

	return resultHolder, nil
}
