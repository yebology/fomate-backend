package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Content struct {

	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	ContentLink string `json:"content_link"`

	ContentTitle string `json:"content_title"`

	ContentDuration uint64 `json:"content_duration"`

	ContentDescription string `json:"content_description"`

	ContentPrice uint64 `json:"content_price"`

	ContentInsights []string `json:"content_insights"`
	
}