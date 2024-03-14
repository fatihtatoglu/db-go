package db

import (
	"encoding/json"
)

type DBTable struct {
	columns []DBColumn
	rows    []DBRow
}

func CreateNewDBTable() DBTable {
	return DBTable{
		columns: make([]DBColumn, 0),
		rows:    make([]DBRow, 0),
	}
}

func (dt *DBTable) CreateNewDBRow() DBRow {
	row := DBRow{
		itemArray:     make([]interface{}, len(dt.columns)),
		itemArrayPtrs: make([]interface{}, len(dt.columns)),
		Table:         dt,
	}

	for i := range dt.columns {
		row.itemArrayPtrs[i] = &row.itemArray[i]
	}

	return row
}

func (dt *DBTable) AddDBColumn(column DBColumn) error {
	column.Table = dt

	dt.columns = append(dt.columns, column)
	return nil
}

func (dt *DBTable) AddDBRow(row DBRow) error {
	dt.rows = append(dt.rows, row)
	return nil
}

func (dt *DBTable) MarshalJSON() ([]byte, error) {
	tableMap := map[string]interface{}{
		"columns": dt.columns,
		"rows":    dt.rows,
	}

	return json.Marshal(tableMap)
}

func (dt *DBTable) GetRows() []DBRow {
	return dt.rows
}

// TODO: UnmarshalJSON methodunu ekle.
