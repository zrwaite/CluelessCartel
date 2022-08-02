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
	SusLevel     int
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
	Username string
	Hash     string
	Cash     int
	Bases    []Base
	SusLevel int `bson:"susLevel"`
}

var userSchema = bson.M{
	"bsonType": "object",
	"required": []string{"username", "hash", "cash", "bases", "susLevel"},
	"properties": bson.M{
		"username": bsonString,
		"hash":     bsonString,
		"cash":     bsonInt,
		"bases":    bsonArray(baseSchema),
		"susLevel": bsonInt,
	},
}

const STARTING_USER_CASH = 15000
const STARTING_USER_SUS_LEVEL = 10

func (user *PostUser) InitData(userParams *PostUserParams) {
	copier.Copy(user, userParams)
	user.Cash = STARTING_USER_CASH
	user.SusLevel = STARTING_USER_SUS_LEVEL
	user.Bases = []Base{}
}
func (user *PostUser) CreateHash(password string) error {
	var err error
	user.Hash, err = auth.HashPassword(password)
	return err
}
