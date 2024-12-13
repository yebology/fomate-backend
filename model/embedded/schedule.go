package embedded

import "go.mongodb.org/mongo-driver/bson/primitive"

type Schedule struct {

	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	UserId primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`

	AppName string `json:"appName"`

	StartTime uint64 `json:"startTime"`
	
	EndTime uint64 `json:"endTime"`

}