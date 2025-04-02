package api

import (
	"ecommerce-order/constants"
	"ecommerce-order/external"
	"ecommerce-order/helpers"
	"ecommerce-order/internal/interfaces"
	"ecommerce-order/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderAPI struct {
	OrderService interfaces.IOrderService
}

func (api *OrderAPI) CreateOrder(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	req := models.Order{}
	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	profileCtx := e.Get("profile")
	profile, ok := profileCtx.(external.Profile)
	if !ok {
		log.Error("failed to get profile from context")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)
	}

	resp, err := api.OrderService.CreateOrder(e.Request().Context(), profile ,&req)
	if err != nil {
		log.Error("failed to create order: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}
