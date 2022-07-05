package models

import (
	// "gopkg.in/mgo.v2/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)


type Task struct{
	ID primitive.ObjectID 	`json:"_id" bson:"_id,omitempty"`
	Name string 	 		`json:"name"`
	Done bool 		 		`json:"done"`
	InsertedAt time.Time 	`json:"inserted_at" bson:"inserted_at"`
	UpdatedAt time.Time 	`json:"updated_at" bson:"updated_at"`
}

type Tasks []Task