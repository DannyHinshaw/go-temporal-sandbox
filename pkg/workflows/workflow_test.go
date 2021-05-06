package workflows

import (
	"go-temporal-example/app/pkg/activities"
	"go-temporal-example/app/pkg/common"
	"testing"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	env.OnActivity(activities.ReturnNonSerializableJSON).Return(nil, nil)
	env.ExecuteWorkflow(TriggerBadActivity)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var res common.BadJSON
	require.NoError(t, env.GetWorkflowResult(&res))
	require.NotNil(t, res)
}
