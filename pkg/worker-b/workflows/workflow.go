package workflows

import (
	"log"
	"time"

	"go-temporal-example/pkg/common/models"
	"go-temporal-example/pkg/worker-b/activities"

	"go.temporal.io/sdk/workflow"
)

// TriggerTestActivity triggers an activity that returns a JSON serializable struct.
func TriggerTestActivity(ctx workflow.Context) (*models.SomeJSON, error) {
	log.Println("WORKFLOW-B:: TriggerTestActivity")
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var res models.SomeJSON
	err := workflow.ExecuteActivity(ctx, activities.ReturnSomeJSON).Get(ctx, &res)
	log.Printf("WORKFLOW-B::res:: %+v", res)

	return &res, err
}
