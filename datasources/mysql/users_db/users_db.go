package users_db

import (
	"database/sql"
	"fmt"
	"os"

	// _ imports the package no matter if it's being used here

	"github.com/go-sql-driver/mysql"
	"github.com/sampado/bookstore_utils-go/logger"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

// automatically called when the package is imported
func init() {
	connect()
}

func connect() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, host, schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger(logger.GetLogger())
	logger.Info("DB successfully configured")
}
