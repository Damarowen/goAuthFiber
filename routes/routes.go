package routes

import (
	"auth-go-fiber/controller"
	"auth-go-fiber/middlewares"
	fiber "github.com/gofiber/fiber/v2"
)


func Setup(app *fiber.App) {

	/*	app.Use("/api/home", jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(key),
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return fiber.ErrUnauthorized
		},
	}))*/

	app.Get("/", controller.Hello)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Get("/api/home", middlewares.AuthorizationRequired(), controller.Home)
	app.Get("/api/logout", controller.Logout)

}
