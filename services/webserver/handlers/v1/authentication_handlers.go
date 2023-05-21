package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ungame/go-signup/pb/auth"
	"github.com/ungame/go-signup/services/webserver"
	"github.com/ungame/go-signup/services/webserver/jsonext"
	"net/http"
)

type authenticationHandler struct {
	authenticationClient auth.AuthenticationServiceClient
	inputValidator       jsonext.InputValidator
}

func NewAuthenticationHandler(authenticationClient auth.AuthenticationServiceClient, inputValidator jsonext.InputValidator) webserver.Handler {
	return &authenticationHandler{
		authenticationClient: authenticationClient,
		inputValidator:       inputValidator,
	}

}

func (h *authenticationHandler) Register(app *fiber.App) {
	app.Post("/v1/auth", h.CreateAuthentication)
	app.Post("/v1/login", h.Login)
}

func (h *authenticationHandler) CreateAuthentication(ctx *fiber.Ctx) error {

	var (
		authClient     = h.authenticationClient
		inputValidator = h.inputValidator
	)

	input := new(jsonext.CreateAuthenticationInput)

	err := ctx.BodyParser(input)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	err = inputValidator.Validate(input)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	res, err := authClient.CreateAuthentication(ctx.Context(), input.ToProto())
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	output := jsonext.NewAuthenticationUserOutputFromProto(res)

	ctx.Set("location", fmt.Sprintf("/users/%s", res.Id))

	return ctx.
		Status(http.StatusCreated).
		JSON(output)
}

func (h *authenticationHandler) Login(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusNoContent)
}
