package db

import "encoding/json"

type DBRow struct {
	itemArray     []interface{}
	itemArrayPtrs []interface{}
	Table         *DBTable
}

func (dr *DBRow) GetItemByIndex(index int) (interface{}, error) {
	// TODO: Add validation

	return dr.itemArray[index], nil
}

func (dr *DBRow) SetItemByIndex(index int, value interface{}) error {
	// TODO: Add validation

	dr.itemArray[index] = value
	return nil
}

func (dr *DBRow) GetItemByColumn(column DBColumn) (interface{}, error) {
	return dr.GetItemByIndex(column.ordinal)
}

func (dr *DBRow) SetItemByColumn(column DBColumn, value interface{}) error {
	return dr.SetItemByIndex(column.ordinal, value)
}

func (dr *DBRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(dr.itemArray)
}

// TODO: UnmarshalJSON methodunu ekle.
