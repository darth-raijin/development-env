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
)

func main() {
	defer services.DisconnectTemporal()
	temporalClient := services.ConnectTemporal()

	// Instantiate singleton config
	internal.GetConfig()

	// Setting up to start a Workflow -> Workflow is a set of tasks (activities)
	taskQueue := fmt.Sprintf("%v-payment-flow-stripe-api", uuid.New()) // map uuid to stripe
	temporal := worker.New(
		temporalClient,
		taskQueue,
		worker.Options{},
	)

	// Register Workflows -> TODO find optimization
	temporal.RegisterWorkflow(services.PaymentWorkFlow)

	// Register Activities
	activities := &activities.PaymentActivity{}
	temporal.RegisterActivity(activities.StartTransaction)
	temporal.RegisterActivity(activities.CheckBalance)
	temporal.RegisterActivity(activities.ReserveFunds)
	temporal.RegisterActivity(activities.WithdrawFunds)

	// RUN IT
	err := temporal.Start()
	if err != nil {
		log.Fatalln("Unable to start Worker", err)
	}

	// Start the Workflow Execution
	workflowOptions := client.StartWorkflowOptions{
		ID:        "payment-workflow",
		TaskQueue: taskQueue,
	}
	workflowParam := services.PaymentWorkFlowParam{
		User:   uuid.New(),
		Amount: 500,
	}

	we, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, services.PaymentWorkFlow, workflowParam)
	if err != nil {
		log.Fatalln("Unable to start Workflow Execution", err)
	}

	log.Println("Started Workflow Execution", we.GetID(), we.GetRunID())

}
