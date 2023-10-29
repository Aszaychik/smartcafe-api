package middleware

import (
	"aszaychik/smartcafe-api/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterAuth(config *config.AuthConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the value of x-api-key header from the request
			requestAPIKey := c.Request().Header.Get("x-api-key")

			// Check if the header is present and matches the expected API key
			if requestAPIKey == "" || requestAPIKey != config.XAPIKey {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			}

			// Call the next handler in the chain
			return next(c)
		}
	}
}
