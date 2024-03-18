package db

import (
	"encoding/json"

	db_errors "github.com/fatihtatoglu/db-go/error"
)

type DBRow struct {
	itemArray     []interface{}
	itemArrayPtrs []interface{}
	Table         *DBTable
}

func (dr *DBRow) SetItemByIndex(index int, value interface{}) error {
	err := dr.validateColumnIndex(index)
	if err != nil {
		return err
	}

	dr.itemArray[index] = value
	return nil
}

func (dr *DBRow) GetItemByIndex(index int) (interface{}, error) {
	err := dr.validateColumnIndex(index)
	if err != nil {
		return nil, err
	}

	return dr.itemArray[index], nil
}

func (dr *DBRow) GetItemByName(columnName string) (interface{}, error) {
	column, err := dr.findColumnByName(columnName)
	if err != nil {
		return nil, err
	}

	return dr.GetItemByIndex(column.ordinal)
}

func (dr *DBRow) SetItemByName(columnName string, value interface{}) error {
	column, err := dr.findColumnByName(columnName)
	if err != nil {
		return err
	}

	return dr.SetItemByIndex(column.ordinal, value)
}

func (dr *DBRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(dr.itemArray)
}

// TODO: UnmarshalJSON methodunu ekle.

func (dr *DBRow) findColumnByName(columnName string) (*DBColumn, error) {
	var column *DBColumn
	for _, c := range dr.Table.columns {
		if c.name == columnName {
			column = &c
			break
		}
	}

	if column == nil {
		return nil, db_errors.ColumnNotFoundError()
	}

	return column, nil
}

func (dr *DBRow) validateColumnIndex(index int) error {
	count := len(dr.Table.columns)
	if index < 0 || index >= count {
		return db_errors.ColumnIndexOutOfRangeError()
	}

	return nil
}
