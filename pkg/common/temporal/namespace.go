package temporal

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"go.temporal.io/api/serviceerror"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
)

var workflowExecutionRetentionPeriod = 24 * time.Hour

// waitForNamespaceReady recursively waits for namespace to be ready in Temporal.
func waitForNamespaceReady(namespaceClient client.NamespaceClient, namespace string, b *Backoff) error {
	b.attempt++
	if b.attempt == b.MaxRetries {
		return fmt.Errorf(`max attempts reached waiting for Temporal namespace "%s" to be ready`, namespace)
	}

	resp, err := namespaceClient.Describe(context.Background(), namespace)
	if err == nil {
		log.Infof(`successfully registered new namespace "%s"`, resp.NamespaceInfo.Name)
		return nil
	}

	log.Errorf(`error from attempt %d to describe namespace "%s": %s`, b.attempt, namespace, err)
	time.Sleep(b.MaxDelay) // wait for namespace cache refresh on temporal-server

	return waitForNamespaceReady(namespaceClient, namespace, b)
}

// registerTemporalNamespace handles registering the namespace for the current worker if necessary.
func registerTemporalNamespace(temporalAddress string, namespace string) error {
	namespaceClient, err := client.NewNamespaceClient(client.Options{HostPort: temporalAddress})
	if err != nil {
		return fmt.Errorf("error creating Temporal namespaceClient: %w", err)
	}
	defer namespaceClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = namespaceClient.Register(ctx, &workflowservice.RegisterNamespaceRequest{
		Namespace:                        namespace,
		WorkflowExecutionRetentionPeriod: &workflowExecutionRetentionPeriod,
	})
	if _, ok := err.(*serviceerror.NamespaceAlreadyExists); ok {
		log.Infof(`namespace "%s" already exists, stopping registration (no need)`, namespace)
		return nil
	}

	if err != nil {
		return fmt.Errorf(`error registering Temporal namespace "%s": %w`, namespace, err)
	}

	// Constant "backoff" (ie no "backoff"), we simply want to refresh Temporal namespace cache
	//	every 20 seconds if it is not registering the namespace.
	b := Backoff{
		MaxDelay:    20 * time.Second,
		MaxRetries:  20,
		Coefficient: 1,
	}
	return waitForNamespaceReady(namespaceClient, namespace, &b)
}

// NewNamespacedClient synchronously calls and waits on required methods for namespace registration and on success returns a new client.
func NewNamespacedClient(temporalAddress string, namespace Namespace) (client.Client, error) {
	err := registerTemporalNamespace(temporalAddress, namespace)
	if err != nil {
		return nil, fmt.Errorf(`unable to create Temporal client for namespace "%s": %w`, namespace, err)
	}

	temporalClient, err := client.NewClient(client.Options{HostPort: temporalAddress, Namespace: namespace})
	if err != nil {
		return nil, fmt.Errorf(`unable to create Temporal client for namespace "%s": %w`, namespace, err)
	}

	return temporalClient, nil
}
