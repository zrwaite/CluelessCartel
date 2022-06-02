package dbModules

import (
	"clueless-cartel-server/database"
	"clueless-cartel-server/database/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUsernameFilter(username string) bson.D {
	return bson.D{{
		Key: "username",
		Value: bson.D{{
			Key:   "$eq",
			Value: username,
		}},
	}}
}

func GetUserData(username string, projection interface{}) (*models.GetUserReturn, int) {
	opts := options.FindOne().SetProjection(projection)
	cursor := database.MongoDatabase.Collection("users").FindOne(context.TODO(), CreateUsernameFilter(username), opts)
	userReturn := new(models.GetUserReturn)
	if err := cursor.Decode(userReturn); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return userReturn, 404
		} else {
			fmt.Println("Failed to get user " + username + " ; " + err.Error())
			return userReturn, 400
		}
	} else {
		return userReturn, 200
	}
}
