package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DbConn(dbName string) *sql.DB {
	mysqlString := loginMySql + ":" + passwordMySql + "@/" + dbName
	db, err := sql.Open("mysql", mysqlString)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}