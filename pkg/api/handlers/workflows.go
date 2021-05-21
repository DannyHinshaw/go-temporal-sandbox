package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go-temporal-example/app/pkg/common"
	"go.temporal.io/sdk/client"
)

// TestWorker is a test route for triggering an activity that returns an arbitrary response.
func (h *Handler) TestWorker(c echo.Context, worker string, tClient client.Client) error {
	ctx := h.Context
	taskQueue := worker
	workflow := "TriggerTestActivity"
	log.Printf("starting workflow `%s` for task queue `%s`", workflow, taskQueue)

	options := client.StartWorkflowOptions{TaskQueue: taskQueue}
	we, err := tClient.ExecuteWorkflow(ctx, options, workflow)
	if err != nil {
		log.Printf("unable to execute workflow `%s`: %s", workflow, err.Error())
		return fmt.Errorf("error executing workflow in `%s`: %w", taskQueue, err)
	}

	log.Printf("waiting on workflow `%s` response for task queue `%s`", workflow, taskQueue)
	var resp *common.SomeJSON
	if err := we.Get(context.Background(), &resp); err != nil {
		log.Printf("unable to get result from workflow `%s`: %s", workflow, err.Error())
		return fmt.Errorf("error executing workflow in `%s`: %w", taskQueue, err)
	}

	return c.JSON(http.StatusOK, &map[string]string{"val": we.GetID()})
}

func (h *Handler) TestWorkerA(c echo.Context) error {
	return h.TestWorker(c, "worker-a", h.TemporalClientA)
}

func (h *Handler) TestWorkerB(c echo.Context) error {
	return h.TestWorker(c, "worker-b", h.TemporalClientB)
}
