package handlers

import (
	"log"

	"github.com/cambo9p/SQLCraft/database"
	"github.com/gofiber/fiber/v2"
)

var apiDb database.Database

func InitAPIRoutes(app *fiber.App, db database.Database) {
  apiDb = db
  app.Get("/api/message", sendMessage)
  app.Get("api/db/get_users", sendCurrUserTable)
}

func sendMessage(c *fiber.Ctx) error {
  return c.SendString("ooga booga it worked")
}


func sendCurrUserTable(c *fiber.Ctx) error {

  rows, err := apiDb.Query("SELECT * FROM users")
  if err != nil {
    log.Fatal("fucked the query")
  }

  // send as json ? 
  names := ""

  for rows.Next() {
    var id int
    var name string
    var dob string
    err := rows.Scan(&id, &name, &dob)
    if err != nil {
        panic(err)
    }
    names += name
  }

  return c.SendString(names)
}
