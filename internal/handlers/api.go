package handlers

import "github.com/gofiber/fiber/v2"


func InitAPIRoutes(app *fiber.App) {
  app.Get("/api/message", sendMessage)
}

func sendMessage(c *fiber.Ctx) error {
  return c.SendString("ooga booga it worked")
}

