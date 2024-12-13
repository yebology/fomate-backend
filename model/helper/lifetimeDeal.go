package helper

import "go.mongodb.org/mongo-driver/bson/primitive"

type LifetimeDeal struct {

	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	UserId primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	
}