package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/verbeux-ai/whatsapp-business/pkg/whatsapp"
)

// Auth authenticates the app on webhook route
func Auth(sdk whatsapp.Business) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		mode := ctx.QueryParam("hub.mode")
		token := ctx.QueryParam("hub.verify_token")
		challenge := ctx.QueryParam("hub.challenge")

		// check the mode and token sent are correct
		if mode == "subscribe" && sdk.Auth(token) != nil {
			return ctx.String(http.StatusOK, challenge)
		}

		return ctx.NoContent(http.StatusForbidden)
	}
}
