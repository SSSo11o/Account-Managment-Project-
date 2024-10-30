package controllers

import (
	"account_managment/coinkeeper/internal/models"
	"account_managment/coinkeeper/internal/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err)
		return
	}
	err := service.CreateUser(user)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, newDefaultResponse("user created successfully"))
}


func SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err)
		return
	}
	accessToken, err := service.SignIn(user.Username, user.Password)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, accessTokenResponse{accessToken})

}
