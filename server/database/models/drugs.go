package models

import "go.mongodb.org/mongo-driver/bson"

type Drug struct {
	Cost        []ResourceCost
	BatchAmount int
	UnitPrice   int
}

type DrugAmount struct {
	DrugName string
	Amount   int
}

var drugAmountSchema = bson.M{
	"bsonType": "object",
	"required": []string{"drug_name", "amount"},
	"properties": bson.M{
		"drug_name": bson.M{"bsonType": "string"},
		"amount":    bson.M{"bsonType": "int"},
	},
}
