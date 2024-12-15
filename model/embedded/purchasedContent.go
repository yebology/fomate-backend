package embedded

import "go.mongodb.org/mongo-driver/bson/primitive"

type PurchasedContent struct {

	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	UserId primitive.ObjectID `json:"userId" bson:"userId,omitempty"`

	ContentId primitive.ObjectID `json:"contentId" bson:"contentId,omitempty"`

}