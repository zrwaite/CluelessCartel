package models

import "go.mongodb.org/mongo-driver/bson"

var UsersValidator = bson.M{
	"$jsonSchema": userSchema,
}

var bsonInt = bson.M{"bsonType": "int"}
var bsonBoolean = bson.M{"bsonType": "bool"}
var bsonString = bson.M{"bsonType": "string"}

func bsonArray(items bson.M) bson.M {
	return bson.M{
		"bsonType":    "array",
		"uniqueItems": false,
		"items":       items,
	}
}
