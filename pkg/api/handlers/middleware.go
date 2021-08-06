package handlers

import (
	"github.com/labstack/echo/v4"
)

// HandlerMiddleware is middleware to assign handler values centrally.
func (h *handler) HandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.context = c.Request().Context()
		h.echoContext = c
		return next(c)
	}
}
