package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"os"
	"testing"

	db_config "github.com/fatihtatoglu/db-go/config"
	db_errors "github.com/fatihtatoglu/db-go/error"
)

type MockDriver struct {
	connection            *MockConn
	closeErrorMessage     string
	connectorErrorMessage string
	pingerErrorMessage    string
}

type MockConn struct {
	driver *MockDriver
}

type MockConnector struct {
	driver     *MockDriver
	connection *MockConn
}

func (d *MockDriver) CleanErrorMessage() {
	d.closeErrorMessage = ""
	d.connectorErrorMessage = ""
	d.pingerErrorMessage = ""
}

func (d *MockDriver) SetCloseErrorMessage(message string) {
	d.closeErrorMessage = message
}

func (d *MockDriver) SetConnectorErrorMessage(message string) {
	d.connectorErrorMessage = message
}

func (d *MockDriver) SetPingerErrorMessage(message string) {
	d.pingerErrorMessage = message
}

// Fake-it
func (m *MockConn) Begin() (driver.Tx, error)                 { return nil, nil }
func (m *MockConn) Close() error                              { return nil }
func (m *MockConn) Prepare(query string) (driver.Stmt, error) { return nil, nil }

// Fake-it
func (m *MockConn) Ping(ctx context.Context) error {
	if m.driver.pingerErrorMessage != "" {
		return errors.New(m.driver.pingerErrorMessage)
	}

	return nil
}

// Fake-it
func (d *MockDriver) Open(name string) (driver.Conn, error) {
	if d.closeErrorMessage != "" {
		return nil, errors.New(d.closeErrorMessage)
	}

	d.connection = &MockConn{
		driver: d,
	}

	return d.connection, nil
}

// Fake-it
func (d *MockDriver) OpenConnector(name string) (driver.Connector, error) {
	if d.connectorErrorMessage != "" {
		return nil, errors.New(d.connectorErrorMessage)
	}

	// to initialize the conn struct
	d.Open(name)

	return &MockConnector{
		driver:     d,
		connection: d.connection,
	}, nil
}

// Fake-it
func (m *MockConnector) Connect(context.Context) (driver.Conn, error) {
	if m.driver.connectorErrorMessage != "" {
		return nil, errors.New(m.driver.connectorErrorMessage)
	}

	return m.connection, nil
}

// Fake-it
func (m *MockConnector) Driver() driver.Driver {
	return m.driver
}

// Fake-it
func (m *MockConnector) Close() error {
	if m.driver.closeErrorMessage != "" {
		return errors.New(m.driver.closeErrorMessage)
	}

	return nil
}

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
	config := getDBConfig()
	// Act
	connection, err := CreateNewDBConnection(driverName, *config)

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
		{"Empty driver", "", db_errors.ConnectionInvalidDriverErrorMessage},
		{"Invalid driver", "invalid-driver", db_errors.ConnectionInvalidDriverErrorMessage},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			// Arrange
			config := getDBConfig()

			// Act
			_, err := CreateNewDBConnection(tc.driver, *config)

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
			config := getDBConfig()

			connection, _ := CreateNewDBConnection(driverName, *config)

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
	config := getDBConfig()

	connection, _ := CreateNewDBConnection(driverName, *config)

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
	config := getDBConfig()

	connection, _ := CreateNewDBConnection(driverName, *config)

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
	config := getDBConfig()

	connection, _ := CreateNewDBConnection(driverName, *config)

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
	config := getDBConfig()

	connection, _ := CreateNewDBConnection(driverName, *config)

	// Act
	db := connection.getConnection()

	// Assert
	if db != nil {
		t.Error("Expected nil connection, but got non-nil")
	}
}

func TestClose_Without_Opening(t *testing.T) {
	// Arrange
	config := getDBConfig()

	connection, _ := CreateNewDBConnection(driverName, *config)

	// Act
	err := connection.Close()

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func getDBConfig() *db_config.DBConfig {
	config := db_config.CreateNewDBConfig("testuser", "testpassword", "localhost", 3306, "test")
	return config
}
