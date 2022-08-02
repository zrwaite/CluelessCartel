package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type TaskType struct {
	Name string
}

type Task struct {
	EndTime      int    `bson:"endTime"`
	TaskTypeName string `bson:"taskTypeName"`
	SimName      string `bson:"simName"`
}

var taskSchema = bson.M{
	"bsonType": "object",
	"required": []string{"endTime", "taskTypeName", "simName"},
	"properties": bson.M{
		"endTime":      bsonInt,
		"taskTypeName": bsonString,
		"simName":      bsonString,
	},
}
