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
			name:          "Column_NotFoundError",
			errorFunc:     ColumnNotFoundError,
			expectedError: errors.New(Column_NotFoundErrorMessage),
		},
		{
			name:          "Column_IndexOutOfRangeError",
			errorFunc:     ColumnIndexOutOfRangeError,
			expectedError: errors.New(Column_IndexOutOfRangeErrorMessage),
		},
		{
			name:          "Connection_InvalidDriverError",
			errorFunc:     ConnectionInvalidDriverError,
			expectedError: errors.New(Connection_InvalidDriverErrorMessage),
		},
		{
			name:          "Connection_EmptyDSNError",
			errorFunc:     ConnectionEmptyDSNError,
			expectedError: errors.New(Connection_EmptyDSNErrorMessage),
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
