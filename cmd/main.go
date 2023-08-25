package main

import (
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  app.Static("/", "./static")
  app.Static("/home", "./static/home.html")

  // Start the server
  port := "8080"
  if err := app.Listen(":" + port); err != nil {
    panic(err)
  }
}
