package users

import (
	"github.com/Anatol-e/bookstore_users_api/domain/users"
	"github.com/Anatol-e/bookstore_users_api/services"
	"github.com/Anatol-e/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func SearchUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
