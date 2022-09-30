package api

import (
	"github.com/labstack/echo/v5"
)

// Init Init
func Init(e *echo.Echo) error {
	v1 := e.Group("/api").Group("/v1")
	v1.GET("/ping", ping)

	return nil
}
