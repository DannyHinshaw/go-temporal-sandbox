package common

const TaskQueue = "TASK_QUEUE"

type TemporalNamespaces struct {
	WorkerA string
	WorkerB string
}

var Namespaces = TemporalNamespaces{
	WorkerA: "worker-a",
	WorkerB: "worker-b",
}
