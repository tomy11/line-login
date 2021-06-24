package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"Id" bson:"_id"`
	Email     string        `json:"Email" bson:"email"`
	Password  string        `json:"Password" bson:"password"`
	CreatedAt time.Time     `json:"CreatedAt" bson:"create_at"`
	UpdatedAt time.Time     `json:"UpdatedAt" bson:"updated_at"`
}
