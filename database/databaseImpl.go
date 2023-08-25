package database

import (
	"database/sql"

	"github.com/cambo9p/SQLCraft/internal/dao"
)

type DatabaseImpl struct {
  db *sql.DB
}

func NewDatabase(db *sql.DB) Database {
  return &DatabaseImpl{db: db}
}

// TODO: sanity check ?  -- unused i think
func (d *DatabaseImpl) Query(query string, args ...interface{}) (*sql.Rows, error) {
  return d.db.Query(query, args...)
}

// TODO: refactor into function that marshals rows into DAO -- this should be in the dao pacakage 
// returns the query result in the form of a tabledata dao 
func (d *DatabaseImpl) GetQueryData(query string, args ...interface{}) (dao.TableData, error) {
  var tabledata dao.TableData
  // table name can be anything i suppose since its a query result
  tabledata.TableName = "result"
  
  rows, err := d.db.Query(query, args...)
  if err != nil {
    return tabledata, err
  }

  // get all the column names 
  columns, err := rows.Columns()
  if err != nil {
    return tabledata, err
  }
  tabledata.Columns = columns

  for rows.Next() {
    values := make([]interface{}, len(columns))
    for i := range values {
        var value interface{} // Create a new variable for each element
        values[i] = &value
    }

    if err := rows.Scan(values...); err != nil {
      return tabledata, err
    }

    // converts the interface values into correct types 
    var rowData []interface{}
    for _, value := range values {
      rowData = append(rowData, *value.(*interface{})) // Dereference the interface{} value
    }

    tabledata.Rows = append(tabledata.Rows, rowData)
  }

  return tabledata, nil
}

func (d *DatabaseImpl) InsertTable(query string, args ...interface{}) error {
  
  return nil
}

func (d *DatabaseImpl) InsertData(query string, args ...interface{}) error {

  return nil
}
