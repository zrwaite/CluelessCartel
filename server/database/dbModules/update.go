package dbModules

import (
	"clueless-cartel-server/database"
	"clueless-cartel-server/database/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(user *models.GetUserReturn) (*models.GetUserReturn, bool) {
	var newUser *models.GetUserReturn
	update := bson.D{{"$set", user}}
	updateResult, err := database.MongoDatabase.Collection("users").UpdateOne(context.TODO(), CreateUsernameFilter(user.Username), update)
	if err != nil {
		fmt.Println(err.Error())
		return newUser, false
	}
	fmt.Println(updateResult)
	return newUser, true
}
