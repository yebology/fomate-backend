package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Schedule struct {

	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	UserId primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`

	AppName string `json:"app_name"`

	StartTime uint64 `json:"start_time"`
	
	EndTime uint64 `json:"end_time"`

}