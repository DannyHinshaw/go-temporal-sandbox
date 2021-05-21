package main

import (
	"go-temporal-example/app/pkg/activities"
	"go-temporal-example/app/pkg/common"
	"go-temporal-example/app/pkg/workflows"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	// Create the client object just once per process
	hp := common.GetHostPortEnv()
	c, err := client.NewClient(client.Options{HostPort: hp, Namespace: common.Namespaces.WorkerB})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Worker and Activity functions
	w := worker.New(c, common.TaskQueue, worker.Options{})
	w.RegisterWorkflow(workflows.TriggerTestActivity)
	w.RegisterActivity(activities.ReturnSomeJSON)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
