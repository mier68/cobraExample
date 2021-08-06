package sql2struct

import "database/sql"

type DBModel struct {
	DBEngine *sql.DB
	DBInfo *DBInfo
}

type DBInfo struct{
	DBType string
	Host string
	UserName string
	Password string
	Charset string
}

type TableColumn struct{
	ColumnName string
	DataType string
	IsNullable string
	ColumnKey string
	ColumnType string
	ColumnComment string
}

func NewDBModel(Info *DBInfo) *DBModel{
	return &DBModel{
	}
}