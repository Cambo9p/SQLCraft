package dao

import (
	"database/sql"
	"encoding/json"
)


type TableData struct {
  TableName string          `json:"table"`
  Columns   []string        `json:"columns"`
  Rows      [][]interface{} `json:"rows"` 
}

// table data dao to json
func  MarshalTableData(data TableData) ([]byte, error) {
  return json.Marshal(data)
}

// json to table data dao
func  UnMrshalTableData(data []byte) (TableData, error) {
  var tableData TableData
  err := json.Unmarshal(data, &tableData)
  return tableData, err
}

// takes rows from the db and puts them inside the tabledata DAO 
func GetDatabaseRowsIntoTableData(rows *sql.Rows, tabledata *TableData) error {
  columns, err := rows.Columns()
  if err != nil { 
  return err
  }
  
  for rows.Next() {
    values := make([]interface{}, len(columns))
    for i := range values {
        var value interface{} // Create a new variable for each element
        values[i] = &value
    }

    if err := rows.Scan(values...); err != nil {
      return err
    }

    // converts the interface values into correct types 
    var rowData []interface{}
    for _, value := range values {
      rowData = append(rowData, *value.(*interface{})) // Dereference the interface{} value
    }

    tabledata.Rows = append(tabledata.Rows, rowData)
  }
  return nil
}
