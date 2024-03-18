package db_errors

import "errors"

const (
	ColumnNotFoundErrorMessage        = "the column specified by columnName cannot be found"
	ColumnIndexOutOfRangeErrorMessage = "the columnIndex argument is out of range"
)

func ColumnNotFoundError() error {
	return errors.New(ColumnNotFoundErrorMessage)
}

func ColumnIndexOutOfRangeError() error {
	return errors.New(ColumnIndexOutOfRangeErrorMessage)
}
