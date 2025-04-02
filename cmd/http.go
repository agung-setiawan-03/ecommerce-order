package cmd

import (
	"ecommerce-order/external"
	"ecommerce-order/helpers"
	"ecommerce-order/internal/api"
	"ecommerce-order/internal/interfaces"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	// d := dependencyInject()
	healthcheckAPI := &api.HealthCheckAPI{}

	e := echo.New()
	e.GET("/healthcheck", healthcheckAPI.HealthCheck)

	e.Start(":" + helpers.GetEnv("PORT", "9001"))
}

type Dependency struct {
	External       interfaces.IExternal
	HealthcheckAPI *api.HealthCheckAPI
}

func dependencyInject() Dependency {
	external := &external.External{}

	return Dependency{
		External:       external,
		HealthcheckAPI: &api.HealthCheckAPI{},
	}
}
