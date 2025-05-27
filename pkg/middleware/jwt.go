// nolint:revive
package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tubfuzzy/banraicore-api/config"
)

func Authorization(allows []string, config config.Configuration) func(*fiber.Ctx) error {
	jwtSecret := config.Authorization.JWTSecret

	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			rolesInterface := claims["roles"].([]interface{})
			roles := make([]string, len(rolesInterface))
			for i, v := range rolesInterface {
				roleMap := v.(map[string]interface{})
				roles[i] = roleMap["role"].(string)
			}
			if len(findIntersection(allows, roles)) > 0 {
				return ctx.Next()
			}

			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{
					"code":    "40100",
					"message": "Unauthorized",
				})
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.
					Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"code":    "40000",
						"message": "Bad Request",
					})
			} else {
				return c.
					Status(fiber.StatusUnauthorized).
					JSON(fiber.Map{
						"code":    "40100",
						"message": "Unauthorized",
					})
			}
		},
	})
}

func findIntersection(arr1, arr2 []string) []string {
	intersection := []string{}

	arr1Map := make(map[string]bool)
	for _, val := range arr1 {
		arr1Map[val] = true
	}

	for _, val := range arr2 {
		if arr1Map[val] {
			intersection = append(intersection, val)
		}
	}

	return intersection
}
