package embedded

import "go.mongodb.org/mongo-driver/bson/primitive"

type Schedule struct {

	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	UserId primitive.ObjectID `json:"userId" bson:"userId,omitempty"`

	AppName string `json:"appName"`

	StartTime string `json:"startTime"`
	
	EndTime string `json:"endTime"`

	Timer uint64 `json:"timer"`

}