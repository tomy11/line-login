package models

import (
	"time"
	
	"gopkg.in/mgo.v2/bson"
)


type SlipUser struct {
	Id        bson.ObjectId  `json:"Id" bson:"_id,mitempt"`
	UserId    string       `json:"UserId" bson:"userid"`
	LineId    string        `json:"LineId" bson:"lineid"`
	Images    string        `json:"Images" bson:"images"`
	CreatedAt time.Time     `json:"CreatedAt" bson:"creat_at"`
	UpdatedAt time.Time     `json:"UpdatedAt" bson:"updtd_at"`
	Users []User `json:"Users" bson:"users"`
}
