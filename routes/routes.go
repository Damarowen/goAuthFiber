package routes

import (
	"auth-go-fiber/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App){
	app.Get("/", controller.Hello)
}