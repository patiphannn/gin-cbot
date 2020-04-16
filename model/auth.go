package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Auth defines auth object structure
type Auth struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name" binding:"required"`
	Email       string        `json:"email" bson:"email" binding:"required,email"`
	Password    string        `json:"password" bson:"password" binding:"required,min=5,max=20"`
	IsVerified  bool          `json:"is_verified" bson:"is_verified" binding:"required"`
	CreatedTime time.Time     `json:"created_time" bson:"created_time"`
	UpdatedTime time.Time     `json:"updated_time" bson:"updated_time"`
}

// Signin defines user signin
type Signin struct {
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"-" binding:"required,min=5,max=20"`
}
