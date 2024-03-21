package db

import (
	"database/sql"
	"os"
	"testing"

	db_errors "github.com/fatihtatoglu/db-go/error"
)

var mockDriver *MockDriver
var driverName string

func TestMain(m *testing.M) {
	driverName = "mysql"
	mockDriver = &MockDriver{}
	sql.Register(driverName, mockDriver)

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCreateNewDBConnection(t *testing.T) {
	// Arrange
	driverName := "mysql"
	dsn := "testuser:testpass@tcp(localhost:3306)/testdb"

	// Act
	connection, err := CreateNewDBConnection(driverName, dsn)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if connection == nil {
		t.Errorf("connection is nil")
	}
}

func TestCreateNewDBConnection_Invalid_Driver(t *testing.T) {
	// Test cases
	testCases := []struct {
		testName       string
		driver         string
		expectedErrMsg string
	}{
		{"Empty driver", "", db_errors.Connection_InvalidDriverErrorMessage},
		{"Invalid driver", "invalid-driver", db_errors.Connection_InvalidDriverErrorMessage},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			// Arrange
			dsn := "testuser:testpass@tcp(localhost:3306)/testdb"

			// Act
			_, err := CreateNewDBConnection(tc.driver, dsn)

			// Assert
			if err == nil {
				t.Error("expected error not occurred")
			}

			if err.Error() != tc.expectedErrMsg {
				t.Errorf("expected error message: %s, got: %s", tc.expectedErrMsg, err.Error())
			}
		})
	}
}

func TestCreateNewDBConnection_Empty_DSN(t *testing.T) {
	// Arrange
	dsn := ""

	// Act
	_, err := CreateNewDBConnection("mysql", dsn)

	// Assert
	if err == nil {
		t.Error("expected error not occurred")
	}
}

func TestOpen(t *testing.T) {
	// Arrange
	tests := []struct {
		name                 string
		driver               *MockDriver
		expcetedErrorMessage string
		arrangeFunc          func()
	}{
		{
			name:                 "Success",
			driver:               mockDriver,
			expcetedErrorMessage: "",
			arrangeFunc:          func() {},
		},
		{
			name:                 "Failed sql.Open",
			driver:               mockDriver,
			expcetedErrorMessage: "connection open failed",
			arrangeFunc: func() {
				mockDriver.SetConnectorErrorMessage("connection open failed")
			},
		},
		{
			name:                 "Failed sql.Ping",
			driver:               mockDriver,
			expcetedErrorMessage: "ping failed",
			arrangeFunc: func() {
				mockDriver.SetPingerErrorMessage("ping failed")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dsn := "testuser:testpass@tcp(localhost:3306)/testdb"

			connection, _ := CreateNewDBConnection(driverName, dsn)

			test.arrangeFunc()

			// Act
			err := connection.Open()

			// Assert
			if test.expcetedErrorMessage == "" && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if test.expcetedErrorMessage != "" && err == nil {
				t.Error("Expected error, but got nil")
			}

			if test.expcetedErrorMessage != "" && err != nil && test.expcetedErrorMessage != err.Error() {
				t.Errorf("Expected error: %v, but got: %v", test.expcetedErrorMessage, err.Error())
			}
		})

		mockDriver.CleanErrorMessage()
	}
}

func TestClose(t *testing.T) {
	// Arrange
	dsn := "testuser:testpass@tcp(localhost:3306)/testdb"

	connection, _ := CreateNewDBConnection(driverName, dsn)

	connection.Open()

	// Act
	err := connection.Close()

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestClose_With_Error(t *testing.T) {
	// Arrange
	dsn := "testuser:testpass@tcp(localhost:3306)/testdb"

	connection, _ := CreateNewDBConnection(driverName, dsn)

	mockDriver.SetCloseErrorMessage("close failed")

	connection.Open()

	// Act
	err := connection.Close()

	// Assert
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func Test_getConnection(t *testing.T) {
	// Arrange
	dsn := "testuser:testpass@tcp(localhost:3306)/testdb"

	connection, _ := CreateNewDBConnection(driverName, dsn)

	connection.Open()

	// Act
	db := connection.getConnection()

	// Assert
	if db == nil {
		t.Error("Expected non-nil connection, but got nil")
	}
}

func Test_getConnection_Without_Open(t *testing.T) {
	// Arrange
	dsn := "testuser:testpass@tcp(localhost:3306)/testdb"

	connection, _ := CreateNewDBConnection(driverName, dsn)

	// Act
	db := connection.getConnection()

	// Assert
	if db != nil {
		t.Error("Expected nil connection, but got non-nil")
	}
}

func TestClose_Without_Opening(t *testing.T) {
	// Arrange
	dsn := "testuser:testpass@tcp(localhost:3306)/testdb"

	connection, _ := CreateNewDBConnection(driverName, dsn)

	// Act
	err := connection.Close()

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
