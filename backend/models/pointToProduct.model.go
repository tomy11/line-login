package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type PointToProduct struct {
	Id        bson.ObjectId `json:"Id,omitempty" bson:"_id,omitempty"`
	ProductId string        `json:"ProductId" bson:"product"`
	UserId    string        `json:"UserId" bson:"user_id"`
	Point     int64         `json:"Point" bson:"point"`
	CreatedAt time.Time     `json:"CreatedAt" bson:"create_at"`
	UpdatedAt time.Time     `json:"UpdatedAt" bson:"updated_at"`
	Product   ProductWich   `json:"Product" bson:"product"`
}

type InputPointToProduct struct {
	ProductId string      `json:"ProductId" bson:"product"`
	UserId    string      `json:"UserId" bson:"user_id"`
	Point     int64       `json:"Point" bson:"point"`
	Product   ProductWich `json:"Product" bson:"product"`
}
