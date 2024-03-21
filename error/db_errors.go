package db_errors

import "errors"

const (
	Column_NotFoundErrorMessage        = "column: the column specified by columnName cannot be found"
	Column_IndexOutOfRangeErrorMessage = "column: the columnIndex argument is out of range"

	Connection_InvalidDriverErrorMessage = "connection: the driver is invalid"
	Connection_EmptyDSNErrorMessage      = "connection: the dsn cannot be empty"
)

func ColumnNotFoundError() error {
	return errors.New(Column_NotFoundErrorMessage)
}

func ColumnIndexOutOfRangeError() error {
	return errors.New(Column_IndexOutOfRangeErrorMessage)
}

func ConnectionInvalidDriverError() error {
	return errors.New(Connection_InvalidDriverErrorMessage)
}

func ConnectionEmptyDSNError() error {
	return errors.New(Connection_EmptyDSNErrorMessage)
}
