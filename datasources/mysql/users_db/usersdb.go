package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var ClientDB *sql.DB

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"1qaz2wsx",
		"localhost",
		"users_db",
	)
	var err error
	ClientDB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = ClientDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured")
}
