package workflows

import (
	"go-temporal-example/app/pkg/activities"
	"go-temporal-example/app/pkg/common"
	"go.temporal.io/sdk/workflow"
	"time"
)

// TriggerBadActivity triggers an activity that returns a non JSON serializable struct.
func TriggerBadActivity(ctx workflow.Context) (*common.BadJSON, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var res common.BadJSON
	err := workflow.ExecuteActivity(ctx, activities.ReturnNonSerializableJSON).Get(ctx, &res)

	return &res, err
}
