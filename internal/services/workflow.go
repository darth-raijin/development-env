package services

import "go.temporal.io/sdk/workflow"

type PaymentWorkFlowResult struct {
	state string
}

func PaymentWorkFlow(ctx workflow.Context) (PaymentWorkFlowResult, error) {
	var resultHolder PaymentWorkFlowResult
	workflow.GetLogger(ctx).Info("Starting workflow :wow:")

	return resultHolder, nil
}
