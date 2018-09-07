package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strings"
)

func InitDb(pathToInitScript string) {
	db, err := sql.Open("mysql", mysqlServerAddr)
	PanicOnErr(err)
	defer db.Close()

	err = db.Ping()
	PanicOnErr(err)

	sqlScript, err := ioutil.ReadFile(pathToInitScript)
	PanicOnErr(err)

	requests := strings.Split(string(sqlScript), ";")

	for _, req := range requests {
		_, err := db.Exec(req)
		PanicOnErr(err)
	}
}

func DbConn(dbName string) *sql.DB {
	db, err := sql.Open("mysql", mysqlServerAddr + dbName)
	PanicOnErr(err)

	err = db.Ping()
	PanicOnErr(err)

	return db
}