package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ungame/go-signup/logext"
	"github.com/ungame/go-signup/services/authentication"
	"github.com/ungame/go-signup/services/webserver"
	v1 "github.com/ungame/go-signup/services/webserver/handlers/v1"
	"github.com/ungame/go-signup/services/webserver/jsonext"
	"github.com/ungame/go-signup/utils"
	"log"
	"net/http"
)

func init() {
	webserver.LoadJSONConfigFromFlags(flag.CommandLine)
	flag.Parse()
}

func main() {
	webserverLogger, err := logext.New("webserver")
	if err != nil {
		log.Fatalln(err)
	}
	defer webserverLogger.Close()

	var (
		configs          = webserver.GetConfigs()
		closeHandler     = utils.NewCloseHandler(true)
		middlewareLogger = webserver.NewLogger(webserverLogger)
	)

	authClient, err := authentication.NewServiceClient(&configs.AuthService)
	if err != nil {
		webserverLogger.Fatal("error on create authentication service client: %s", err.Error())
	}
	defer closeHandler.Close(authClient)

	app := fiber.New()

	middlewareLogger.Register(app)

	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).SendString("OK")
	})

	authHandler := v1.NewAuthenticationHandler(authClient.Client(), jsonext.NewInputValidator())
	authHandler.Register(app)

	if err := app.Listen(configs.Port.Address()); err != nil {
		webserverLogger.Fatal("http serve error: %s", err.Error())
	}

}
