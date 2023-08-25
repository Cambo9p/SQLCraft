package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cambo9p/SQLCraft/internal/handlers"
	"github.com/gofiber/fiber/v2"
  _ "github.com/lib/pq"
)

func main() {
  // con to db 
  connStr := "user=postgres dbname=postgres password=newpassword sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
  rows, err := db.Query("SELECT * FROM users")
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  // Iterate over the results
  for rows.Next() {
    var id int
    var name string
    var dob string
    err := rows.Scan(&id, &name, &dob)
    if err != nil {
        panic(err)
    }
    fmt.Println(id, name)
  }
  // set up fiber app
  app := fiber.New()

  handlers.InitAPIRoutes(app)

  app.Static("/", "./static")
  app.Static("/home", "./static/home.html")

  // Start the server
  port := "8080"
  if err := app.Listen(":" + port); err != nil {
    panic(err)
  }
}
