package main

import (
	"go-temporal-example/app/pkg/common"
	"go.temporal.io/sdk/client"
	"log"

	"github.com/labstack/echo/v4"

	"go-temporal-example/app/pkg/api/handlers"
)

const (
	namespaceA = "worker-a"
	namespaceB = "worker-b"
)

func main() {

	// Create the client object just once per process
	hp := common.GetHostPortEnv()
	ca, err := client.NewClient(client.Options{HostPort: hp, Namespace: namespaceA})
	if err != nil {
		log.Fatalf(`unable to create Temporal client for namespace "%s": %s`, namespaceA, err)
	}
	defer ca.Close()

	cb, err := client.NewClient(client.Options{HostPort: hp, Namespace: namespaceB})
	if err != nil {
		log.Fatalf(`unable to create Temporal client for namespace "%s": %s`, namespaceB, err)
	}
	defer cb.Close()

	// Setup Echo framework
	e := echo.New()
	h := handlers.NewHandler(ca, cb)

	// Register middlewares
	e.Use(h.HandlerMiddleware)

	// Register the URL endpoints to Handler
	h.RegisterRouteHandlers(e)

	// Serve
	e.Logger.Fatal(e.Start(":8080"))
}
