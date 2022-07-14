package models

import (
	"clueless-cartel-server/auth"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
)

var GetUserOpts = struct {
	Username int
	Cash     int
	Bases    int
}{1, 1, 1}

type GetUserReturn struct {
	Username     string
	Cash         int
	Bases        []Base
	NewBasePrice int
}

var GetUserAuthOpts = struct {
	Username int
	Hash     int
}{1, 1}

type GetUserAuthReturn struct {
	Username string
	Hash     string
}

type PostUserParams struct {
	Username string
	Password string
}

type PostUser struct {
	Username string `bson:"username"`
	Hash     string
	Cash     int
	Bases    []Base
}

var userSchema = bson.M{
	"bsonType": "object",
	"required": []string{"username", "hash", "cash", "bases"},
	"properties": bson.M{
		"username": bson.M{"bsonType": "string"},
		"hash":     bson.M{"bsonType": "string"},
		"cash":     bson.M{"bsonType": "int"},
		"bases": bson.M{
			"bsonType":    "array",
			"uniqueItems": false,
			"items":       baseSchema,
		},
	},
}

const STARTING_USER_CASH = 15000

func (user *PostUser) InitData(userParams *PostUserParams) {
	copier.Copy(user, userParams)
	user.Cash = STARTING_USER_CASH
	user.Bases = []Base{}
}
func (user *PostUser) CreateHash(password string) error {
	var err error
	user.Hash, err = auth.HashPassword(password)
	return err
}
