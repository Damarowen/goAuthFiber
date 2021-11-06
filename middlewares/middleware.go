package middlewares

import (
	"auth-go-fiber/controller"
	fiber "github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"

)

const jwtSecret = controller.SecretKey

func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}
func AuthSuccess(c *fiber.Ctx) error {
	c.Next()
	return nil
}

func AuthorizationRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		// Filter:         nil,
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte(jwtSecret),
		// SigningKeys:   nil,
		SigningMethod: "HS256",
		// ContextKey:    nil,
		// Claims:        nil,
		// TokenLookup:   nil,
		// AuthScheme:    nil,
	})
}
