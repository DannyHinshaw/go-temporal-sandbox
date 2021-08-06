package temporal

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"

	"go.temporal.io/sdk/client"
)

// Backoff holds data used for exponential backoff.
type Backoff struct {
	// MaxDelay maximum amount of time to sleep between retries.
	MaxDelay time.Duration
	// Coefficient multiplied against attempt on each run to get exponential sleep duration.
	Coefficient int
	// MaxRetries is the maximum number of times to trying the call before giving up.
	MaxRetries int
	// attempt is used internally to track how many times the function has been called recursively.
	attempt int
}

// ExecuteOptions holds configuration options for handling the trigger of workflow executions.
type ExecuteOptions struct {
	// WorkflowOptions are the Temporal client StartWorkflowOptions required by ExecuteWorkflow.
	WorkflowOptions client.StartWorkflowOptions
	// Workflow is either the workflow function or a string with the name of the workflow function.
	Workflow interface{}
	// Backoff contains the exponential backoff configuration for the ExecuteWorkflowWithRetries handler func.
	Backoff Backoff
}

// ExecuteWorkflowWithRetries handles firing off workflow executions to Temporal with retry/failure handling.
func ExecuteWorkflowWithRetries(ctx context.Context, tClient client.Client, opts *ExecuteOptions, args ...interface{}) (client.WorkflowRun, error) {
	if opts.Backoff.attempt == opts.Backoff.MaxRetries {
		return nil, fmt.Errorf(`error executing workflow "%v", max retries met`, opts.Workflow)
	}

	we, err := tClient.ExecuteWorkflow(ctx, opts.WorkflowOptions, opts.Workflow, args...)
	if err != nil {
		log.Errorf(`error executing workflow "%v": %s`, err)
		opts.Backoff.attempt++
		delay := time.Duration(opts.Backoff.attempt*opts.Backoff.Coefficient) * time.Second
		if delay >= opts.Backoff.MaxDelay {
			delay = opts.Backoff.MaxDelay
		}

		log.Warnf("failed to create Temporal connection for worker (attempt %d), retrying in %s seconds...", opts.Backoff.attempt, delay)
		time.Sleep(delay)

		return ExecuteWorkflowWithRetries(ctx, tClient, opts, args...)
	}

	return we, nil
}
