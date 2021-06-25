package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type PointToProduct struct {
	Id        bson.ObjectId `json:"Id" bson:"_id"`
	ProductId string        `json:"ProductId" bson:"product"`
	UserId    string        `json:"UserId" bson:"user_id"`
	Point     string        `json:"Point" bson:"point"`
	CreatedAt time.Time     `json:"CreatedAt" bson:"create_at"`
	UpdatedAt time.Time     `json:"UpdatedAt" bson:"updated_at"`
}
