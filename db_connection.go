package db

import (
	"database/sql"
	"fmt"

	db_config "github.com/fatihtatoglu/db-go/config"
	db_errors "github.com/fatihtatoglu/db-go/error"
)

type DBConnectionInterface interface {
	Open() error
	Close() error
	getConnection() *sql.DB
}

type dbConnection struct {
	dsn    string
	driver string
	db     *sql.DB
}

func CreateNewDBConnection(driver string, config db_config.DBConfig) (DBConnectionInterface, error) {
	if driver == "" {
		return nil, db_errors.ConnectionInvalidDriverError()
	}

	var dsn string
	switch driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetDatabaseName())
	default:
		dsn = ""
	}

	if dsn == "" {
		return nil, db_errors.ConnectionInvalidDriverError()
	}

	return &dbConnection{
		dsn:    dsn,
		driver: driver,
	}, nil
}

func (c *dbConnection) Open() error {
	connection, err := sql.Open(c.driver, c.dsn)
	if err != nil {
		return err
	}

	err = connection.Ping()
	if err != nil {
		return err
	}

	c.db = connection
	return nil
}

func (c *dbConnection) Close() error {
	if c.db == nil {
		return nil
	}

	return c.db.Close()
}

func (c *dbConnection) getConnection() *sql.DB {
	return c.db
}
