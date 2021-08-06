package handlers

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.temporal.io/sdk/client"
)

type handler struct {

	// Easy access for common request contexts (handle by mw).
	context     context.Context
	echoContext echo.Context

	// Example to show housing multiple clients.
	temporalClientA client.Client
	temporalClientB client.Client
}

// NewHandler constructor func creates a new handler with DI for services.
func NewHandler(temporalClientA client.Client, temporalClientB client.Client) *handler {
	return &handler{
		temporalClientA: temporalClientA,
		temporalClientB: temporalClientB,
	}
}

// RegisterRouteHandlers registers REST API endpoints handler functions.
func (h *handler) RegisterRouteHandlers(v1 *echo.Echo) {

	// API Health check
	v1.GET("/health", h.health)

	// Test Workflow endpoints
	v1.GET("/workflow-a", h.testWorkerA)
	v1.GET("/workflow-b", h.testWorkerB)
}
