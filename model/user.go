package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {

	UserId	primitive.ObjectID `json:"id" bson:"id,omitempty"`

	Username string `json:"username"`

	Email    string `json:"email"`

	Password string `json:"password"`
	
}