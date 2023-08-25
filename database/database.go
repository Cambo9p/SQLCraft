package database

import (

	"github.com/cambo9p/SQLCraft/internal/dao"
)

type Database interface {
  GetQueryData(query string, args ...interface{}) (dao.TableData, error)
  InsertTable(query string, args ...interface{}) error
  InsertData(query string, args ...interface{}) error
  // Query(query string, args ...interface{}) (*sql.Rows, error)
}

