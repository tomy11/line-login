package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type SlipUser struct {
	Id        bson.ObjectId `json:"Id,omitempty" bson:"_id,omitempty"`
	LineId    string        `json:"LineId" bson:"lineid"`
	Images    string        `json:"Images" bson:"images"`
	CreatedAt time.Time     `json:"CreatedAt" bson:"creat_at"`
	UpdatedAt time.Time     `json:"UpdatedAt" bson:"updtd_at"`
	UserId    Auth          `json:"UserId" bson:"userId"`
}

type InputSlipUser struct {
	LineId string `json:"LineId" bson:"lineid" valid:"length(3|300)"`
	Images string `json:"Images" bson:"images" valid:"length(3|300)"`
}
