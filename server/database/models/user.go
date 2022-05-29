package models

import (
	"clueless-cartel-server/auth"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
)

var GetUserQuery = struct {
	Username    int
	Cash        int
	Resources   int
	Drugs       int
	Weapons     int
	HexagonRows int `bson:"hexagon_rows"`
}{1, 1, 1, 1, 1, 1}

var GetUserReturn = struct {
	Username    string
	Cash        int
	Resources   Resources
	Drugs       Drugs
	Weapons     Weapons
	HexagonRows []Hexagon `bson:"hexagon_rows"`
}{}

type PostUserParams struct {
	Dev      string
	Username string
	Password string
}

type PostUser struct {
	Username    string `bson:"username"`
	Hash        string
	Cash        int
	Resources   Resources
	Drugs       Drugs
	Weapons     Weapons
	HexagonRows []HexagonRow `bson:"hexagon_rows"`
}

var userSchema = bson.M{
	"bsonType": "object",
	"required": []string{"username", "hash", "cash", "resources", "drugs", "weapons", "hexagon_rows"},
	"properties": bson.M{
		"username":  bson.M{"bsonType": "string"},
		"hash":      bson.M{"bsonType": "string"},
		"cash":      bson.M{"bsonType": "int"},
		"resources": resourcesSchema,
		"drugs":     drugsSchema,
		"weapons":   weaponsStruct,
		"hexagon_rows": bson.M{
			"bsonType":    "array",
			"uniqueItems": false,
			"items":       hexagonRowsSchema,
		},
	},
}

func (user *PostUser) InitData(userParams *PostUserParams) {
	copier.Copy(user, userParams)
	user.Cash = 5000
	user.Resources = StartingResources
	user.Drugs = StartingDrugs
	user.Weapons = StartingWeapons
	user.HexagonRows = StartingHexagonRows[:]
}
func (user *PostUser) CreateHash(password string) error {
	var err error
	user.Hash, err = auth.HashPassword(password)
	return err
}
