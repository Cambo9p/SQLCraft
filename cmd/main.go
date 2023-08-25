package main

import (
	"database/sql"
	// "fmt"
	"log"

	"github.com/cambo9p/SQLCraft/database"
	"github.com/cambo9p/SQLCraft/internal/handlers"
	"github.com/gofiber/fiber/v2"
	 _ "github.com/lib/pq"
)

func main() {
  // connect to db 
  connStr := "user=postgres dbname=postgres password=newpassword sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  dbInstance := database.NewDatabase(db)
  
  // set up fiber app
  app := fiber.New()

  handlers.InitAPIRoutes(app, dbInstance)

  app.Static("/", "./static")
  app.Static("/home", "./static/home.html")

  // Start the server
  port := "8080"
  if err := app.Listen(":" + port); err != nil {
    panic(err)
  }
}
