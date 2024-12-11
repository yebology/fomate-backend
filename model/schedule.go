package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Schedule struct {
	Id primitive.ObjectID `json:"id" bson:"id,omitempty"`
	AppId primitive.ObjectID `json:"app_id" bson:"app_id,omitempty"`
	UserId primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Duration uint64 `json:"duration"`
}