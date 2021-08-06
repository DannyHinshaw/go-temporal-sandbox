package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// health endpoint to make sure API is healthy and functional.
func (h *handler) health(c echo.Context) error {
	// Should probably do some sort of Temporal health check here:
	// https://community.temporal.io/t/temporal-client-worker-health-check/205/9
	return c.JSON(http.StatusOK, "ok")
}
