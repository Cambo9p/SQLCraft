package database


import "database/sql"

type Database interface {
  Query(query string, args ...interface{}) (*sql.Rows, error)
}

