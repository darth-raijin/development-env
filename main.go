package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-temporal-laboratory/internal"
	"github.com/go-temporal-laboratory/internal/activities"
	"github.com/go-temporal-laboratory/internal/services"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func main() {
	defer services.DisconnectTemporal()
	temporalClient := services.ConnectTemporal()

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
	activities := &activities.PaymentActivity{}
	temporal.RegisterActivity(activities.StartTransaction)
	temporal.RegisterActivity(activities.CheckBalance)
	temporal.RegisterActivity(activities.ReserveFunds)
	temporal.RegisterActivity(activities.WithdrawFunds)

	// RUN IT
	err := temporal.Run(worker.InterruptCh())

	if err != nil {
		log.Fatalln("Unable to start Worker", err)
	}

	workflowOptions := client.StartWorkflowOptions{}
	workflowParam := services.PaymentWorkFlowParam{
		User:   uuid.New(),
		Amount: 500,
	}

	temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, services.PaymentWorkFlow, &workflowParam)

}
