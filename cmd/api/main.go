package main

import (
	"go-temporal-example/app/pkg/common"
	"go.temporal.io/sdk/client"
	"log"

	"github.com/labstack/echo/v4"

	"go-temporal-example/app/pkg/api/handlers"
)

func main() {

	// Create the client object just once per process
	hp := common.GetHostPortEnv()
	c, err := client.NewClient(client.Options{HostPort: hp})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// Setup Echo framework
	e := echo.New()
	h := handlers.NewHandler(c)

	// Register middlewares
	e.Use(h.HandlerMiddleware)

	// Register the URL endpoints to Handler
	h.RegisterRouteHandlers(e)

	// Serve
	e.Logger.Fatal(e.Start(":8080"))
}
