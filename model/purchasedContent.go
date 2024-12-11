package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PurchasedContent struct {
	Id primitive.ObjectID `json:"id" bson:"id,omitempty"`
	UserId primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	ContentId primitive.ObjectID `json:"content_id" bson:"content_id,omitempty"`
}