package handlers

import (
	"github.com/labstack/echo/v4"
)

// HandlerMiddleware is middleware to assign handler values centrally.
func (h *Handler) HandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.Context = c.Request().Context()
		h.EchoContext = c
		return next(c)
	}
}
