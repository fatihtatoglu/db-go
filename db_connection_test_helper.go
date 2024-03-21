package db

import (
	"context"
	"database/sql/driver"
	"errors"
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
