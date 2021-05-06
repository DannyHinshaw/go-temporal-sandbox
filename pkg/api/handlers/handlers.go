package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.temporal.io/sdk/client"
)

type Handler struct {
	Context        context.Context
	EchoContext    echo.Context
	TemporalClient client.Client
}

// NewHandler constructor func creates a new Handler with DI for services.
func NewHandler(temporalClient client.Client) *Handler {
	return &Handler{
		TemporalClient: temporalClient,
	}
}

// RegisterRouteHandlers registers REST API endpoints handler functions.
func (h *Handler) RegisterRouteHandlers(v1 *echo.Echo) {

	// API Health check
	v1.GET("/health", h.GETHealth)

	// Test Workflow endpoints
	v1.GET("/workflow", h.TestActivityResponseJSON)
}
