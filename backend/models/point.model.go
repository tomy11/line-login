package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Point struct {
	Id        bson.ObjectId `json:"Id,omitempty" bson:"_id,omitempty"`
	UserId    string        `json:"UserId" bson:"userid"`
	SlipId    string        `json:"SlipId" bson:"slipid"`
	Point     int64         `json:"Point" bson:"point"`
	CreatedAt time.Time     `json:"CreatedAt" bson:"create_at"`
	UpdatedAt time.Time     `json:"UpdatedAt" bson:"updated_at"`
}
type InputPoints struct {
	UserId string `json:"UserId" bson:"userid"`
	SlipId string `json:"SlipId" bson:"slipid"`
	Point  int64  `json:"Point" bson:"point"`
}
