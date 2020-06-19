package main

import (
	"github.com/sampado/bookstore_users-api/app"
	"github.com/sampado/bookstore_utils-go/logger"
)

func main() {
	logger.Info("about to start the application...")
	app.StartApplication()
}
