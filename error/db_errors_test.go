package db_errors

import (
	"errors"
	"testing"
)

func TestColumnNotFoundError(t *testing.T) {
	expectedError := errors.New(ColumnNotFoundErrorMessage)
	actualError := ColumnNotFoundError()

	if actualError.Error() != expectedError.Error() {
		t.Errorf("Expected error message: %s, but got: %s", expectedError.Error(), actualError.Error())
	}
}

func TestColumnIndexOutOfRangeError(t *testing.T) {
	expectedError := errors.New(ColumnIndexOutOfRangeErrorMessage)
	actualError := ColumnIndexOutOfRangeError()

	if actualError.Error() != expectedError.Error() {
		t.Errorf("Expected error message: %s, but got: %s", expectedError.Error(), actualError.Error())
	}
}
