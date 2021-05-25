package main

import (
	"context"
	"fmt"
	"go-temporal-example/app/pkg/activities"
	"go-temporal-example/app/pkg/common"
	"go-temporal-example/app/pkg/workflows"
	"go.temporal.io/api/serviceerror"
	"go.temporal.io/api/workflowservice/v1"
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

var namespace = common.Namespaces.WorkerA

const (
	ctxTimeout                    = 15 * time.Second
	namespaceCacheRefreshInterval = 20 * time.Second
	maxWaitForNamespaceAttempts   = 20
)

// waitForNamespaceReady recursively waits for namespace to be ready in Temporal.
func waitForNamespaceReady(namespaceClient client.NamespaceClient, attempt int) error {
	if attempt == maxWaitForNamespaceAttempts {
		return fmt.Errorf(`max attempts reached waiting for namespace "%s" ready`, namespace)
	}

	resp, err := namespaceClient.Describe(context.Background(), namespace)
	if err == nil {
		log.Printf(`successfully registered new namespace "%s"`, resp.NamespaceInfo.Name)
		return nil
	}

	log.Printf(`error from attempt #%d to describe namespace "%s": %s`, attempt, namespace, err)
	time.Sleep(namespaceCacheRefreshInterval) // wait for namespace cache refresh on temporal-server

	return waitForNamespaceReady(namespaceClient, attempt+1)
}

// registerNamespace handles registering the namespace for the current worker if necessary.
func registerNamespace(hp string) {
	namespaceClient, err := client.NewNamespaceClient(client.Options{HostPort: hp})
	if err != nil {
		log.Fatalln("error creating Temporal namespaceClient: ", err)
	}
	defer namespaceClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	retention := 1 * time.Hour * 24
	err = namespaceClient.Register(ctx, &workflowservice.RegisterNamespaceRequest{
		Namespace:                        namespace,
		WorkflowExecutionRetentionPeriod: &retention,
	})
	if _, ok := err.(*serviceerror.NamespaceAlreadyExists); ok {
		log.Printf(`namespace "%s" already exists, cancelling register`, namespace)
		return
	}

	if err != nil {
		log.Fatalf(`error registering Temporal namespace "%s": %s`, namespace, err)
	}

	err = waitForNamespaceReady(namespaceClient, 1)
	if err != nil {
		log.Fatalln("error waiting for Temporal namespace", err)
	}
}

func main() {

	// Create the client object just once per process
	hp := common.GetHostPortEnv()
	registerNamespace(hp)

	c, err := client.NewClient(client.Options{HostPort: hp, Namespace: namespace})
	if err != nil {
		log.Fatalln("error creating Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Worker and Activity functions
	w := worker.New(c, namespace, worker.Options{})
	w.RegisterWorkflow(workflows.TriggerTestActivity)
	w.RegisterActivity(activities.ReturnSomeJSON)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("error starting Worker", err)
	}
}
