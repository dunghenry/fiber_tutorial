package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginSuccess struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Token    string             `json:"token"`
}
