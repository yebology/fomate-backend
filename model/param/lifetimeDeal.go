package param

import "go.mongodb.org/mongo-driver/bson/primitive"

type LifetimeDeal struct {

	UserId primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	
}