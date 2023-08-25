package handlers

import (
	"log"

	"github.com/cambo9p/SQLCraft/database"
	"github.com/cambo9p/SQLCraft/internal/dao"
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

  query := "select * from users"
  tableData, err := apiDb.GetQueryData(query)
  if err != nil {
    println(err.Error())
    log.Fatal("oh shit ")
  }
  println(tableData.TableName)

  jsonData, err := dao.MarshalTableData(tableData)
  if err != nil {
    return err
  }
  
  c.Type("application/json")
  return c.Send(jsonData)
}
