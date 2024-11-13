package routes

import "github.com/gofiber/fiber/v2"

// Get API health status
func GetHealth(ctx *fiber.Ctx) error {
	/*
	* !todo:
	* Check database connection or something
	 */
	return ctx.JSON(fiber.Map{"status": "ok"})
}
