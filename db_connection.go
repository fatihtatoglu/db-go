package db

import (
	"database/sql"
	"errors"
	"fmt"

	db_config "github.com/fatihtatoglu/db-go/config"
)

const (
	INVALID_DRIVER = "driver is empty or invalid"
)

type DBConnectionInterface interface {
	Open() error
	Close() error

	getConnection() *sql.DB
}

type dbConnection struct {
	dsn        string
	driver     string
	connection *sql.DB
}

func CreateNewDBConnection(driver string, config db_config.DBConfig) (DBConnectionInterface, error) {
	if driver == "" {
		return nil, errors.New(INVALID_DRIVER)
	}

	var dsn string
	switch driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetDatabaseName())
	default:
		dsn = ""
	}

	return &dbConnection{
		dsn:    dsn,
		driver: driver,
	}, nil
}

func (dbc *dbConnection) Open() error {
	connection, err := sql.Open(dbc.driver, dbc.dsn)
	if err != nil {
		return err
	}

	err = connection.Ping()
	if err != nil {
		return err
	}

	dbc.connection = connection
	return nil
}

func (dbc *dbConnection) Close() error {
	if dbc.connection == nil {
		return nil
	}

	err := dbc.connection.Close()
	return err
}

func (dbc *dbConnection) getConnection() *sql.DB {
	return dbc.connection
}
