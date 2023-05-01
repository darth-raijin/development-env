package main

import (
	"fmt"

	"github.com/go-temporal-laboratory/internal"
	"github.com/go-temporal-laboratory/internal/services"
	"github.com/google/uuid"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func main() {
	defer services.DisconnectTemporal()

	// Instantiate singleton config
	internal.GetConfig()

	// Setting up to start a Workflow -> Workflow is a set of tasks (activities)
	temporal := worker.New(
		services.ConnectTemporal(),
		fmt.Sprintf("%v-payment-flow-stripe-api", uuid.New()), // map uuid
		worker.Options{})

	temporalWorkerOptions := workflow.RegisterOptions{
		Name: "Purchase item workflow",
	}

	// Register Workflows -> TODO find optimization
	temporal.RegisterWorkflowWithOptions(services.PaymentWorkFlow, temporalWorkerOptions)

	// Register Activities

}
