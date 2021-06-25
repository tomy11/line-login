package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Point struct {
	Id        bson.ObjectId `json:"Id" bson:"_id"`
	UserId    string        `json:"UserId" bson:"userid"`
	SlipId    string        `json:"SlipId" bson:"slipid"`
	Point     string        `json:"Point" bson:"point"`
	CreatedAt time.Time     `json:"CreatedAt" bson:"create_at"`
	UpdatedAt time.Time     `json:"UpdatedAt" bson:"updated_at"`
}
