package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Content struct {

	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	ContentLink string `json:"contentLink"`

	ContentTitle string `json:"contentTitle"`

	ContentDuration uint64 `json:"contentDuration"`

	ContentDescription string `json:"contentDescription"`

	ContentPrice uint64 `json:"contentPrice"`

	ContentInsights []string `json:"contentInsights"`
	
}