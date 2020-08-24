package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysqlUserUsername = "mysql_user_username"
	mysqlUserPassword = "mysql_user_password"
	mysqlUserHost     = "mysql_user_host"
	mysqlUserSchema   = "mysql_user_schema"
)

var (
	ClientDB *sql.DB
	username = os.Getenv(mysqlUserUsername)
	password = os.Getenv(mysqlUserPassword)
	host     = os.Getenv(mysqlUserHost)
	schema   = os.Getenv(mysqlUserSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
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
