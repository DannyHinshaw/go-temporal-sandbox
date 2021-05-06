package handlers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-temporal-example/app/pkg/common"
	"go.temporal.io/sdk/client"
	"log"
	"net/http"
)

// TestActivityResponseJSON is a test route for triggering an activity that returns a bad response.
func (h *Handler) TestActivityResponseJSON(c echo.Context) error {
	ctx := h.Context
	taskQueue := common.TaskQueue
	workflow := "TriggerBadActivity"
	log.Printf("starting workflow `%s` for task queue `%s`", workflow, taskQueue)

	options := client.StartWorkflowOptions{TaskQueue: taskQueue}
	we, err := h.TemporalClient.ExecuteWorkflow(ctx, options, workflow)
	if err != nil {
		log.Printf("unable to execute workflow `%s`: %s", workflow, err.Error())
		return fmt.Errorf("error executing workflow in `%s`: %w", taskQueue, err)
	}

	var resp *common.BadJSON
	if err := we.Get(context.Background(), &resp); err != nil {
		log.Println("unable to get workflow result", err)
		log.Printf("unable to get result from workflow `%s`: %s", workflow, err.Error())
		return fmt.Errorf("error executing workflow in `%s`: %w", taskQueue, err)
	}

	return c.JSON(http.StatusOK, &map[string]string{"val": we.GetID()})
}
