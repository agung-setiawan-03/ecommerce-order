package cmd

import (
	"ecommerce-order/external"
	"ecommerce-order/helpers"
	"ecommerce-order/internal/api"
	"ecommerce-order/internal/interfaces"
	"ecommerce-order/internal/repository"
	"ecommerce-order/internal/services"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	d := dependencyInject()
	healthcheckAPI := &api.HealthCheckAPI{}

	e := echo.New()
	e.GET("/healthcheck", healthcheckAPI.HealthCheck)

	orderV1 := e.Group("orders/v1")
	orderV1.POST("", d.OrderAPI.CreateOrder, d.MiddlewareValidateAuth)

	e.Start(":" + helpers.GetEnv("PORT", "9001"))
}

type Dependency struct {
	External       interfaces.IExternal
	HealthcheckAPI *api.HealthCheckAPI

	OrderAPI interfaces.IOrderAPI
}

func dependencyInject() Dependency {
	external := &external.External{}

	orderRepo := &repository.OrderRepo{
		DB: helpers.DB,
	}
	orderSvc := &services.OrderService{
		OrderRepo: orderRepo,
		External:  external,
	}
	orderAPI := &api.OrderAPI{
		OrderService: orderSvc,
	}

	return Dependency{
		External:       external,
		HealthcheckAPI: &api.HealthCheckAPI{},

		OrderAPI: orderAPI,
	}
}
