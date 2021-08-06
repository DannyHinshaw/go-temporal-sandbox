package temporal

/*
				Simple Convention for Temporal Namespaces & Queues:
	The convention for naming Temporal task queues and namespaces is to use the
	name of the worker service that will be operating on that queue/namespace.
	This will allow us to easily drill down into the workers running workflows
	in the Temporal Web UI via namespaces, simplify task queue naming, and allow
	worker scaling by service (via Temporal global namespaces).
*/

/*
	Namespaces Registry & Types
*/
type Namespace = string

// NamespacesRegistry list of available Temporal worker services currently available in the stack.
type NamespacesRegistry struct {
	WorkerA Namespace
	WorkerB Namespace
}

// Namespaces references the Temporal worker service names (used for namespaces and task queue names).
var Namespaces = NamespacesRegistry{
	WorkerA: "worker-a",
	WorkerB: "worker-b",
}

/*
	WorkflowNames Registry & Types
*/
type WorkflowName = string

// WorkflowNames list of available Temporal workflow names that can be invoked.
type WorkflowNames struct {
	TriggerTestActivity WorkflowName
}

// Workflows central place to store references to Temporal workflow names for invocation.
var Workflows = WorkflowNames{
	TriggerTestActivity: "TriggerTestActivity",
}
