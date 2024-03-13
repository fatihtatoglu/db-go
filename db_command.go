package db

import (
	"context"
	"database/sql"
	"errors"
)

const (
	NIL_CONNECTION = "connection is nil"
)

type DBCommandInterface interface {
	Execute(query string, params ...interface{}) (*sql.Result, error)
	Query(query string, params ...interface{}) (*DBTable, error)
	QueryFirst(query string, params ...interface{}) (*DBRow, error)
}

type dbCommand struct {
	connection DBConnectionInterface
	ctx        context.Context
}

func CreateNewDBCommand(con DBConnectionInterface) (DBCommandInterface, error) {
	if con == nil {
		return nil, errors.New(NIL_CONNECTION)
	}

	return &dbCommand{
		connection: con,
		ctx:        context.Background(),
	}, nil
}

func (cmd *dbCommand) Execute(query string, params ...interface{}) (*sql.Result, error) {
	err := cmd.connection.Open()
	if err != nil {
		return nil, err
	}
	defer cmd.connection.Close()

	res, err := cmd.connection.getConnection().ExecContext(cmd.ctx, query, params...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (cmd *dbCommand) Query(query string, params ...interface{}) (*DBTable, error) {
	resultSet, err := cmd.query(false, query, params...)
	if err != nil {
		return nil, err
	}

	return resultSet, nil
}

func (cmd *dbCommand) QueryFirst(query string, params ...interface{}) (*DBRow, error) {
	resultSet, err := cmd.query(true, query, params...)
	if err != nil {
		return nil, err
	}

	return &resultSet.rows[0], nil
}

// Ref: https://stackoverflow.com/a/17885636
func (cmd *dbCommand) query(returnSingleRow bool, query string, params ...interface{}) (*DBTable, error) {
	err := cmd.connection.Open()
	if err != nil {
		return nil, err
	}
	defer cmd.connection.Close()

	rows, err := cmd.connection.getConnection().QueryContext(cmd.ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columnTypes, _ := rows.ColumnTypes()

	dt := CreateNewDBTable()
	for i, col := range columnTypes {
		c := CreateNewDBColumnFromColumnType(*col, i)
		dt.AddDBColumn(c)
	}

	for rows.Next() {
		r := dt.CreateNewDBRow()

		rows.Scan(r.itemArrayPtrs...)

		for i := range columnTypes {
			rowValue := r.itemArray[i]
			b, ok := rowValue.([]byte)
			var v interface{}
			if ok {
				v = string(b)
			} else {
				v = rowValue
			}

			r.itemArray[i] = v
		}

		dt.AddDBRow(r)

		if returnSingleRow {
			break
		}
	}

	return &dt, nil
}
