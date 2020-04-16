package service

import (
	"github.com/polnoy/gin-cbot/model"
	"gopkg.in/mgo.v2/bson"
)

// Auth is all user services
type Auth struct{}

// FindEmail is find once with email
func (h *Auth) FindEmail(email string) (model.Auth, error) {
	data := model.Auth{}
	err := DbConnect.Use(userColl).Find(bson.M{"email": email}).One(&data)
	return data, err
}
