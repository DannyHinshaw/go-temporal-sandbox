package workflows

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go-temporal-example/pkg/common/models"
	"go-temporal-example/pkg/worker-b/activities"

	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	env.OnActivity(activities.ReturnSomeJSON).Return(nil, nil)
	env.ExecuteWorkflow(TriggerTestActivity)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var res models.SomeJSON
	require.NoError(t, env.GetWorkflowResult(&res))
	require.NotNil(t, res)
}
