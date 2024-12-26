package param

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateHealth struct {

	UserId primitive.ObjectID `json:"userId" bson:"userId,omitempty"`

	NewHealth int64 `json:"newHealth"`

} 