package service

import (
	"errors"

	"github.com/Kamva/mgm/v2"
	"github.com/Kamva/mgm/v2/operator"
	"github.com/polnoy/gin-cbot/common"
	"github.com/polnoy/gin-cbot/model"
	"gopkg.in/mgo.v2/bson"
)

// User is all user services
type User struct{}

// collection is mongo collection name
const collection = "users"

// Gets is find all
func (h *User) Gets() ([]model.User, error) {
	result := []model.User{}
	err := mgm.Coll(&model.User{}).SimpleFind(&result, bson.M{})

	return result, err
}

// Get is find once
func (h *User) Get(_id string) (*model.User, error) {
	result := &model.User{}
	err := mgm.Coll(result).FindByID(_id, result)
	return result, err
}

// FindEmail is find once with email
func (h *User) FindEmail(email string) (*model.User, error) {
	result := &model.User{}
	err := mgm.Coll(result).First(bson.M{"email": email}, result)
	return result, err
}

// Create is create data
func (h *User) Create(data model.User) error {
	user, _ := h.FindEmail(data.Email)
	if user.Email != "" {
		return errors.New("Email " + data.Email + " already exists.")
	}

	data.Password = common.GeneratePasswordHash([]byte(data.Password))
	return mgm.Coll(&data).Create(&data)
}

// Update is update data
func (h *User) Update(_id string, data model.User) error {
	dup := &model.User{}
	err := mgm.Coll(dup).First(bson.M{"_id": bson.M{operator.Ne: _id}, "email": data.Email}, dup)

	if dup.Email != "" {
		return errors.New("Email " + data.Email + " already exists.")
	}

	user, err := h.Get(_id)
	if err != nil {
		return errors.New(err.Error())
	}

	if data.Name != "" {
		user.Name = data.Name
	}
	if data.Email != "" {
		user.Email = data.Email
	}
	if data.Password != "" {
		user.Password = common.GeneratePasswordHash([]byte(data.Password))
	}

	return mgm.Coll(user).Update(user)
}

// DeleteByID is delete by id
func (h *User) DeleteByID(_id string) error {
	data, err := h.Get(_id)
	if err != nil {
		return errors.New(err.Error())
	}
	return mgm.Coll(data).Delete(data)
}
