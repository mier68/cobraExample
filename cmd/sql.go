package cmd

import (
	"log"

	"github.com/GoProgramming/internal/sql2struct"
	"github.com/spf13/cobra"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql的转换和处理",
	Long:  "sql的转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql的转换",
	Long:  "sql的转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType : dbType,
			Host : host,
			UserName : username,
			Password: password,
			Charset : charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil{
			log.Fatalf("dbModel.Connect err :%v",err)
		}
		columns,err  := dbModel.GetColumns(dbName,tableName)
		if err != nil{
			log.Fatalf("dbmodel.GetColumns err :%v",err)
		}
		tmpl := sql2struct.NewStructTemplate()
		templateColumns := tmpl.AssemblyColumns(columns)
		err = tmpl.Generate(tableName)
	},
}
