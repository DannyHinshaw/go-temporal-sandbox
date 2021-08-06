package temporal

import (
	"time"

	log "github.com/sirupsen/logrus"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// waitForTemporalClient recursively waits for a new Temporal connection to be established and returns a namespaced client.
func waitForTemporalClient(temporalAddress string, namespace string, b *Backoff) client.Client {
	temporalWebhooksClient, err := NewNamespacedClient(temporalAddress, namespace)
	if err != nil {
		b.attempt++
		if b.attempt >= b.MaxRetries {
			// TODO: Send error alerts for developers to intercept and fix instead of killing worker.
			log.Fatalf(`failed to create Temporal worker client (attempt #%d), max retries reached with error: %s`, b.MaxRetries, err)
		}

		delay := time.Duration(b.attempt*b.Coefficient) * time.Second
		if delay >= b.MaxDelay {
			delay = b.MaxDelay
		}

		log.Warnf("failed to create Temporal connection for worker, retrying in %s seconds...", delay)
		time.Sleep(delay)

		return waitForTemporalClient(temporalAddress, namespace, b)
	}

	return temporalWebhooksClient
}

// NewTemporalWorkerClient handles creating connection to Temporal service (with retries) and returns a new worker client instance.
func NewTemporalWorkerClient(temporalAddress string, namespace string, options worker.Options) (worker.Worker, error) {
	// 15 max retries with 15sec max exp backoff makes max approx ~5min before giving up.
	b := Backoff{
		MaxDelay:    15 * time.Second,
		MaxRetries:  15,
		Coefficient: 3,
	}
	temporalWebhooksClient := waitForTemporalClient(temporalAddress, namespace, &b)
	return worker.New(temporalWebhooksClient, namespace, options), nil
}
