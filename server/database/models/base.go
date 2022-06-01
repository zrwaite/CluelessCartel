package models

import "go.mongodb.org/mongo-driver/bson"

type Base struct {
	Resources   Resources
	Drugs       Drugs
	Weapons     Weapons
	HexagonRows []HexagonRow `bson:"hexagon_rows"`
}

var startingBase = Base{
	Resources:   StartingResources,
	Drugs:       StartingDrugs,
	Weapons:     StartingWeapons,
	HexagonRows: StartingHexagonRows[:],
}

var baseSchema = bson.M{
	"bsonType": "object",
	"required": []string{"resources", "drugs", "weapons", "hexagon_rows"},
	"properties": bson.M{
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
