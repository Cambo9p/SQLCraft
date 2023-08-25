package database


import "database/sql"

type DatabaseImpl struct {
  db *sql.DB
}

func NewDatabase(db *sql.DB) Database {
  return &DatabaseImpl{db: db}
}

// TODO: sanity check ? 
func (d *DatabaseImpl) Query(query string, args ...interface{}) (*sql.Rows, error) {
  return d.db.Query(query, args...)
}


