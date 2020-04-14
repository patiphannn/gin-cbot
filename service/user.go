package service

import (
	"time"

	"github.com/polnoy/gin-cbot/common"
	"github.com/polnoy/gin-cbot/model"
	"gopkg.in/mgo.v2/bson"
)

// User is all user services
type User struct{}

// DbConnect is connection database
var DbConnect = common.ConnectDB()

// collection is mongo collection name
const collection = "users"

// Gets is find all
func (h *User) Gets() ([]model.User, error) {
	data := []model.User{}
	err := DbConnect.Use(collection).Find(bson.M{}).All(&data)
	return data, err
}

// Get is find once
func (h *User) Get(_id string) (model.User, error) {
	data := model.User{}
	objectID := bson.ObjectIdHex(_id)
	err := DbConnect.Use(collection).FindId(objectID).One(&data)
	return data, err
}

// FindEmail is find once with email
func (h *User) FindEmail(email string) (model.User, error) {
	data := model.User{}
	err := DbConnect.Use(collection).Find(bson.M{"email": email}).One(&data)
	return data, err
}

// Create is create data
func (h *User) Create(data model.User) error {
	data.Password = common.GeneratePasswordHash([]byte(data.Password))
	data.CreatedTime = time.Now()
	data.UpdatedTime = data.CreatedTime
	return DbConnect.Use(collection).Insert(data)
}

// Update is update data
func (h *User) Update(_id string, data model.User) error {
	objectID := bson.ObjectIdHex(_id)
	newData := bson.M{
		"$set": bson.M{
			"name":         data.Name,
			"email":        data.Email,
			"password":     common.GeneratePasswordHash([]byte(data.Password)),
			"is_verified":  data.IsVerified,
			"updated_time": time.Now(),
		},
	}
	return DbConnect.Use(collection).UpdateId(objectID, newData)
}

// DeleteByID is delete by id
func (h *User) DeleteByID(_id string) error {
	objectID := bson.ObjectIdHex(_id)
	return DbConnect.Use(collection).RemoveId(objectID)
}
