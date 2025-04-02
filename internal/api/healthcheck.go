package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckAPI struct {
}

func (api *HealthCheckAPI) HealthCheck(e echo.Context) error {
	return e.JSON(http.StatusOK, nil)
}