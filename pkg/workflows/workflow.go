package workflows

import (
	"go-temporal-example/app/pkg/activities"
	"go-temporal-example/app/pkg/common"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

// TriggerTestActivity triggers an activity that returns a JSON serializable struct.
func TriggerTestActivity(ctx workflow.Context) (*common.SomeJSON, error) {
	log.Println("THE FUCKING WORKFLOW")
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var res common.SomeJSON
	err := workflow.ExecuteActivity(ctx, activities.ReturnSomeJSON).Get(ctx, &res)

	return &res, err
}
