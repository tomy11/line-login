package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id           bson.ObjectId `json:"Id" bson:"_id"`
	ProductName  string        `json:"ProductName" bson:"product_name"`
	Image        string        `json:"Image" bson:"image"`
	ProductPoint string        `json:"ProductPoint" bson:"product_point"`
	CreateBy     string        `json:"CreateBy" bson:"create_by"`
	CreatedAt    time.Time     `json:"CreatedAt" bson:"create_at"`
	UpdatedAt    time.Time     `json:"UpdatedAt" bson:"updated_at"`
}
