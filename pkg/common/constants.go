package common

type TemporalNamespaces struct {
	WorkerA string
	WorkerB string
}

var Namespaces = TemporalNamespaces{
	WorkerA: "worker-a",
	WorkerB: "worker-b",
}
