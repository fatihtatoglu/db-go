package db_errors

import "errors"

const (
	ColumnNotFoundErrorMessage          = "column: the column specified by columnName cannot be found"
	ColumnIndexOutOfRangeErrorMessage   = "column: the columnIndex argument is out of range"
	ConnectionInvalidDriverErrorMessage = "connection: the driver is invalid"
)

func ColumnNotFoundError() error {
	return errors.New(ColumnNotFoundErrorMessage)
}

func ColumnIndexOutOfRangeError() error {
	return errors.New(ColumnIndexOutOfRangeErrorMessage)
}

func ConnectionInvalidDriverError() error {
	return errors.New(ConnectionInvalidDriverErrorMessage)
}
