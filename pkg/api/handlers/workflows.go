package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"go-temporal-example/pkg/common/models"
	itemporal "go-temporal-example/pkg/common/temporal"

	"go.temporal.io/sdk/client"
)

// testWorker is a test route for triggering a simple activity that returns an arbitrary response.
func (h *handler) testWorker(c echo.Context, worker string, tClient client.Client) error {
	ctx := h.context
	taskQueue := worker
	workflow := itemporal.Workflows.TriggerTestActivity
	log.Printf("starting workflow `%s` for task queue `%s`", workflow, taskQueue)

	exOpts := itemporal.ExecuteOptions{
		Workflow: workflow,
		Backoff: itemporal.Backoff{
			Coefficient: 3,
			MaxRetries:  5,
			MaxDelay:    25 * time.Second,
		},
		WorkflowOptions: client.StartWorkflowOptions{
			TaskQueue: taskQueue,
		},
	}
	we, err := itemporal.ExecuteWorkflowWithRetries(ctx, tClient, &exOpts, workflow)
	if err != nil {
		log.Printf("unable to execute workflow `%s`: %s", workflow, err.Error())
		return fmt.Errorf("error executing workflow in `%s`: %w", taskQueue, err)
	}
	log.Printf("waiting on workflow `%s` response for task queue `%s`", workflow, taskQueue)

	var resp *models.SomeJSON
	if err := we.Get(context.Background(), &resp); err != nil {
		log.Printf("unable to get result from workflow `%s`: %s", workflow, err)
		return fmt.Errorf("error executing workflow in `%s`: %w", taskQueue, err)
	}

	return c.JSON(http.StatusOK, &map[string]string{"val": resp.SomeProp})
}

// testWorkerA calls testWorker with "worker-a" namespace.
func (h *handler) testWorkerA(c echo.Context) error {
	return h.testWorker(c, itemporal.Namespaces.WorkerA, h.temporalClientA)
}

// testWorkerB calls testWorker with "worker-b" namespace.
func (h *handler) testWorkerB(c echo.Context) error {
	return h.testWorker(c, itemporal.Namespaces.WorkerB, h.temporalClientB)
}
