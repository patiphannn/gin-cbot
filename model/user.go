package model

import (
	"github.com/Kamva/mgm/v2"
)

// User is User database model
type User struct {
	mgm.DefaultModel `bson:",inline"`

	Name       string `json:"name" bson:"name"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"-"`
	IsVerified bool   `json:"is_verified" bson:"is_verified"`
}
