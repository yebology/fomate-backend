package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type App struct {
	Id primitive.ObjectID `json:"id" bson:"id,omitempty"`
	AppName string `json:"app_name"`
}