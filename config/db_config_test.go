package db_config

import (
	"testing"
)

func TestCreateNewDBConfig(t *testing.T) {
	user := "testuser"
	pass := "testpass"
	host := "localhost"
	port := 3306
	dbName := "testdb"

	config := CreateNewDBConfig(user, pass, host, port, dbName)

	if config.GetUser() != user {
		t.Errorf("Expected user: %s, got: %s", user, config.GetUser())
	}

	if config.GetPassword() != pass {
		t.Errorf("Expected password: %s, got: %s", pass, config.GetPassword())
	}

	if config.GetHost() != host {
		t.Errorf("Expected host: %s, got: %s", host, config.GetHost())
	}

	if config.GetPort() != port {
		t.Errorf("Expected port: %d, got: %d", port, config.GetPort())
	}

	if config.GetDatabaseName() != dbName {
		t.Errorf("Expected database name: %s, got: %s", dbName, config.GetDatabaseName())
	}
}

func TestMySqlDSN(t *testing.T) {
	// Arrange
	expectedDSN := "testuser:testpass@tcp(localhost:3306)/testdb"

	user := "testuser"
	pass := "testpass"
	host := "localhost"
	port := 3306
	dbName := "testdb"

	config := CreateNewDBConfig(user, pass, host, port, dbName)

	// Act
	dsn := config.GetMysqlDSN()

	// Assert
	if expectedDSN != dsn {
		t.Errorf("Expected DSN: %s, got: %s", expectedDSN, dsn)
	}
}
