package models

import "go.mongodb.org/mongo-driver/bson"

var UsersValidator = bson.M{
	"$jsonSchema": userSchema,
}
