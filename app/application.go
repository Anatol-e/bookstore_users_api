package app

import (
	"github.com/Anatol-e/bookstore_users_api/logger"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	mapUrl()
	logger.Info("about to start the application")
	router.Run(":8080")
}
