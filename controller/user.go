package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polnoy/gin-cbot/model"
	"github.com/polnoy/gin-cbot/service"
)

// User is user service
type User struct {
}

// Gets is find all
func (h *User) Gets(c *gin.Context) {
	service := new(service.User)
	data, err := service.Gets()
	if err != nil {
		log.Println("error user Gets: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": &data,
	})
}

// Get is find once
func (h *User) Get(c *gin.Context) {
	service := new(service.User)
	_id := c.Param("_id")
	data, err := service.Get(_id)
	if err != nil {
		log.Println("error user Get", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": &data,
	})
}

// Create is create data
func (h *User) Create(c *gin.Context) {
	service := new(service.User)

	// JSON
	data := model.User{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("error user Create", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Create
	err = service.Create(data)
	if err != nil {
		log.Println("error user Create", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
	})
}

// Update is update data
func (h *User) Update(c *gin.Context) {
	service := new(service.User)
	_id := c.Param("_id")
	data := model.User{}
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println("error user Update", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := service.Update(_id, data); err != nil {
		log.Println("error user Update", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

// DeleteByID is delete by id
func (h *User) DeleteByID(c *gin.Context) {
	service := new(service.User)
	_id := c.Param("_id")
	err := service.DeleteByID(_id)

	if err != nil {
		log.Println("error user Delete", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{
		"status": true,
	})
}
