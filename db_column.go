package db

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type DBColumn struct {
	name      string
	ordinal   int
	appType   reflect.Type
	dBType    string
	precision int64
	scale     int64
	length    int64
	nullable  bool
	Table     *DBTable
}

func CreateNewDBColumn(name string, ordinal int) DBColumn {
	col := DBColumn{
		name:    name,
		ordinal: ordinal,
	}

	return col
}

func CreateNewDBColumnFromColumnType(ct sql.ColumnType, ordinal int) DBColumn {
	col := DBColumn{
		name:    ct.Name(),
		ordinal: ordinal,
		appType: ct.ScanType(),
		dBType:  ct.DatabaseTypeName(),
	}

	p, s, ok := ct.DecimalSize()
	if ok {
		col.precision = p
		col.scale = s
	}

	l, ok := ct.Length()
	if ok {
		col.length = l
	}

	n, ok := ct.Nullable()
	if ok {
		col.nullable = n
	}

	return col
}

func (cd *DBColumn) GetName() string {
	return cd.name
}

func (cd *DBColumn) GetOrdinal() int {
	return cd.ordinal
}

func (cd *DBColumn) GetType() reflect.Type {
	return cd.appType
}

func (cd *DBColumn) GetDBType() string {
	return cd.dBType
}

func (cd *DBColumn) GetPrecision() int64 {
	return cd.precision
}

func (cd *DBColumn) GetScale() int64 {
	return cd.scale
}

func (cd *DBColumn) GetLength() int64 {
	return cd.length
}

func (cd *DBColumn) GetNullable() bool {
	return cd.nullable
}

func (dc *DBColumn) MarshalJSON() ([]byte, error) {
	columnMap := map[string]interface{}{
		"name":      dc.name,
		"ordinal":   dc.ordinal,
		"appType":   dc.appType.String(),
		"dBType":    dc.dBType,
		"precision": dc.precision,
		"scale":     dc.scale,
		"length":    dc.length,
		"nullable":  dc.nullable,
	}

	return json.Marshal(columnMap)
}

// TODO: UnmarshalJSON methodunu ekle.
