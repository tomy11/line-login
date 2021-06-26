package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"Id,omitempty" bson:"_id,omitempty"`
	Email     string        `json:"Email" bson:"email"`
	Password  string        `json:"Password" bson:"password"`
	CreatedAt time.Time     `json:"CreatedAt" bson:"create_at"`
	UpdatedAt time.Time     `json:"UpdatedAt" bson:"updated_at"`
}

type Auth struct {
	UserId string `json:"UserId,omitempty" bson:"_id,omitempty"`
}

type Role struct {
	UserId   string `json:"UserId,omitempty" bson:"_id,omitempty"`
	RoleName string `json:"RoleName" bson:"role_name"`
}
