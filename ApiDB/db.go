package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strings"
)

func InitDb(pathToInitScript string) {
	db, err := sql.Open("mysql", mysqlServerAddr)
	LogError(err)
	defer db.Close()

	err = db.Ping()
	LogError(err)

	sqlScript, err := ioutil.ReadFile(pathToInitScript)
	LogError(err)

	requests := strings.Split(string(sqlScript), ";")

	for _, req := range requests {
		_, err := db.Exec(req)
		LogError(err)
	}
}

func DropDb(pathToDropScript string) {
	db, err := sql.Open("mysql", mysqlServerAddr)
	LogError(err)
	defer db.Close()

	err = db.Ping()
	LogError(err)

	sqlScript, err := ioutil.ReadFile(pathToDropScript)
	LogError(err)

	requests := strings.Split(string(sqlScript), ";")

	for _, req := range requests {
		_, err := db.Exec(req)
		LogError(err)
	}
}

func DbConn(dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", mysqlServerAddr + dbName)
	if err != nil {
		 return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}