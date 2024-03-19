package db_errors

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		name           string
		errorFunc      func() error
		expectedError  error
		expectedErrMsg string
	}{
		{
			name:          "ColumnNotFoundError",
			errorFunc:     ColumnNotFoundError,
			expectedError: errors.New(ColumnNotFoundErrorMessage),
		},
		{
			name:          "ColumnIndexOutOfRangeError",
			errorFunc:     ColumnIndexOutOfRangeError,
			expectedError: errors.New(ColumnIndexOutOfRangeErrorMessage),
		},
		{
			name:          "ConnectionInvalidDriverError",
			errorFunc:     ConnectionInvalidDriverError,
			expectedError: errors.New(ConnectionInvalidDriverErrorMessage),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualError := test.errorFunc()

			if actualError.Error() != test.expectedError.Error() {
				t.Errorf("Expected error: %v, but got: %v", test.expectedError, actualError)
			}
		})
	}
}
