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

// returns the query result in the form of a tabledata dao 
// TODO - i think there will be a lot of overlap between this and get table data
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

  // put the rows from the table into the DAO
  if err := dao.GetDatabaseRowsIntoTableData(rows, &tabledata); err != nil {
    return tabledata, err
  }

  return tabledata, nil
}

// gets the view for the specific table  -- may not need this as the query one has all funcrtoinality
// func (d *DatabaseImpl) GetTableData(query string, args ...interface{}) (*sql.Rows, error) {
// }

// creates a new table -- TODO 
func (d *DatabaseImpl) InsertTable(query string, args ...interface{}) error {
  
  return nil
}

// inserts data into a table -- TODO
func (d *DatabaseImpl) InsertData(query string, args ...interface{}) error {

  return nil
}
