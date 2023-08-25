package dao

import "encoding/json"


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
