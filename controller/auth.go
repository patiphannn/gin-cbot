package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polnoy/gin-cbot/common"
	"github.com/polnoy/gin-cbot/model"
	"github.com/polnoy/gin-cbot/service"
)

// Auth is auth service
type Auth struct {
}

// Signup is create data
func (h *Auth) Signup(c *gin.Context) {
	service := new(service.User)
	data := model.User{}
	err := c.ShouldBindJSON(&data)

	if err != nil {
		log.Println("error user Signup", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = service.Create(data)

	if err != nil {
		log.Println("error user Signup", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
	})
}

// Signin defines signin
func (h *Auth) Signin(c *gin.Context) {
	service := new(service.Auth)

	data := model.Signin{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("Error Signin ShouldBindJSON", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user, err := service.FindEmail(data.Email)
	if err != nil {
		log.Println("Error Signin FindEmail", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	if err := common.PasswordCompare([]byte(data.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := common.CreateToken(user.ID.String())
	if err != nil {
		log.Println("Error Signin CreateToken", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"_id":           user.ID,
		"name":          user.Name,
		"email":         user.Email,
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"created_time":  user.CreatedTime,
		"updated_time":  user.UpdatedTime,
	})
}

// Refresh defines refresh token
func (h *Auth) Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	token, err := common.RefreshToken(refreshToken)
	if err != nil {
		log.Println("Error Refresh", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}
