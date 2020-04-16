package service

import (
	"errors"
	"time"

	"github.com/polnoy/gin-cbot/common"
	"github.com/polnoy/gin-cbot/model"
	"gopkg.in/mgo.v2/bson"
)

// User is all user services
type User struct{}

// userColl is mongo collection name
const userColl = "users"

// Gets is find all
func (h *User) Gets() ([]model.User, error) {
	data := []model.User{}
	err := DbConnect.Use(userColl).Find(bson.M{}).All(&data)
	return data, err
}

// Get is find once
func (h *User) Get(_id string) (model.User, error) {
	data := model.User{}
	objectID := bson.ObjectIdHex(_id)
	err := DbConnect.Use(userColl).FindId(objectID).One(&data)
	return data, err
}

// FindEmail is find once with email
func (h *User) FindEmail(email string) (model.User, error) {
	data := model.User{}
	err := DbConnect.Use(userColl).Find(bson.M{"email": email}).One(&data)
	return data, err
}

// Create is create data
func (h *User) Create(data model.User) error {
	user := model.User{}
	err := DbConnect.Use(userColl).Find(bson.M{"email": data.Email}).One(&user)
	if err != nil && err.Error() != "not found" {
		return err
	}
	if user.Email != "" {
		return errors.New("Email " + data.Email + " already exists.")
	}

	data.Password = common.GeneratePasswordHash([]byte(data.Password))
	data.CreatedTime = time.Now()
	data.UpdatedTime = data.CreatedTime
	return DbConnect.Use(userColl).Insert(data)
}

// Update is update data
func (h *User) Update(_id string, data model.User) error {
	objectID := bson.ObjectIdHex(_id)

	user := model.User{}
	err := DbConnect.Use(userColl).Find(bson.M{"_id": bson.M{"$ne": objectID}, "email": data.Email}).One(&user)
	if err != nil && err.Error() != "not found" {
		return err
	}
	if user.Email != "" {
		return errors.New("Email " + data.Email + " already exists.")
	}

	newData := bson.M{
		"$set": bson.M{
			"name":         data.Name,
			"email":        data.Email,
			"password":     common.GeneratePasswordHash([]byte(data.Password)),
			"is_verified":  data.IsVerified,
			"updated_time": time.Now(),
		},
	}
	return DbConnect.Use(userColl).UpdateId(objectID, newData)
}

// DeleteByID is delete by id
func (h *User) DeleteByID(_id string) error {
	objectID := bson.ObjectIdHex(_id)
	return DbConnect.Use(userColl).RemoveId(objectID)
}
