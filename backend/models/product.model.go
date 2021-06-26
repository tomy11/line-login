package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id           bson.ObjectId `json:"Id,omitempty" bson:"_id,omitempty"`
	ProductName  string        `json:"ProductName" bson:"product_name"`
	Image        string        `json:"Image" bson:"image"`
	ProductPoint int64         `json:"ProductPoint" bson:"product_point"`
	CreateBy     string        `json:"CreateBy" bson:"create_by"`
	UpdateBy     string        `json:"UpdateBy" bson:"update_by"`
	CreatedAt    time.Time     `json:"CreatedAt" bson:"create_at"`
	UpdatedAt    time.Time     `json:"UpdatedAt" bson:"updated_at"`
}

type InputProduct struct {
	ProductName  string `json:"ProductName" bson:"product_name"`
	Image        string `json:"Image" bson:"image"`
	ProductPoint int64  `json:"ProductPoint" bson:"product_point"`
}

type ProductWich struct {
	Id           bson.ObjectId `json:"Id,omitempty" bson:"_id,omitempty"`
	ProductName  string        `json:"ProductName" bson:"product_name"`
	Image        string        `json:"Image" bson:"image"`
	ProductPoint int64         `json:"ProductPoint" bson:"product_point"`
	CreateBy     string        `json:"CreateBy" bson:"create_by"`
}
