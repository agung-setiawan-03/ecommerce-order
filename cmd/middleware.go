package cmd

import (
	"ecommerce-order/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d *Dependency) MiddlewareValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		auth := e.Request().Header.Get("Authorization")
		if auth == "" {
			helpers.Logger.Errorf("Authorization is empty")
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "Unauthorize", nil)
		}
		profile, err := d.External.GetProfile(e.Request().Context(), auth)
		if err != nil {
			helpers.Logger.Error("Failed to get profile: ", err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "Unauthorize", nil)
		}

		e.Set("profile", profile)

		return next(e)
	}
}
