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

func getUserData(username string, projection interface{}, userReturn interface{}) int {
	opts := options.FindOne().SetProjection(projection)
	cursor := database.MongoDatabase.Collection("users").FindOne(context.TODO(), CreateUsernameFilter(username), opts)
	if err := cursor.Decode(userReturn); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return 404
		} else {
			fmt.Println("Failed to get user " + username + " ; " + err.Error())
			return 400
		}
	} else {
		return 200
	}
}

func GetUserGameData(username string) (user *models.GetUserReturn, status int) {
	user = new(models.GetUserReturn)
	status = getUserData(username, models.GetUserOpts, &user)
	return
}

func GetUserIdData(username string) (user *models.GetIdStruct, status int) {
	user = new(models.GetIdStruct)
	status = getUserData(username, models.GetIdOpts, &user)
	return
}

func GetUserAuthData(username string) (user *models.GetUserAuthReturn, status int) {
	user = new(models.GetUserAuthReturn)
	status = getUserData(username, models.GetUserAuthOpts, &user)
	return
}
