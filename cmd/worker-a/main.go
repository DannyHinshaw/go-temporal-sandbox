package main

import (
	itemporal "go-temporal-example/pkg/common/temporal"
	"go-temporal-example/pkg/worker-a/activities"
	"go-temporal-example/pkg/worker-a/workflows"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

var namespace = itemporal.Namespaces.WorkerA

func main() {

	// Create the client object just once per process
	hp := itemporal.GetHostPortEnv()
	c, err := client.NewClient(client.Options{HostPort: hp, Namespace: namespace})
	if err != nil {
		log.Fatalln("error creating Temporal client", err)
	}
	defer c.Close()

	// Create namespaced temporal worker client.
	temporalWorker, err := itemporal.NewTemporalWorkerClient(hp, namespace, worker.Options{})
	if err != nil {
		log.Fatalf(`unable to create Temporal client for namespace "%s": %s`, namespace, err)
	}

	// This worker hosts both Workflow and Activity functions
	temporalWorker.RegisterWorkflow(workflows.TriggerTestActivity)
	temporalWorker.RegisterActivity(activities.ReturnSomeJSON)

	// Start listening to the Task Queue
	err = temporalWorker.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf(`error starting temporal worker "%s": %s`, namespace, err)
	}
}
