package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDoc struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

type MessageDoc struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"user_id"`
	Type      string             `bson:"type"`
	Text      string             `bson:"text"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	CreatedBy primitive.ObjectID `bson:"created_by"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
	UpdatedBy primitive.ObjectID `bson:"updated_by"`
}
