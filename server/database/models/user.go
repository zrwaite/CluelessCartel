package models

import (
	"clueless-cartel-server/auth"

	"github.com/jinzhu/copier"
)

var GetUserQuery = struct {
	Username  int
	FirstName int `bson:"first_name"`
	LastName  int `bson:"last_name"`
}{1, 1, 1}

var GetUserReturn = struct {
	Username  string
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}{}

type PostUserParams struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
}

type PostUser struct {
	Username  string
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Hash      string
}

func (user *PostUser) ParseData(userParams *PostUserParams) {
	copier.Copy(user, userParams)
}
func (user *PostUser) CreateHash(password string) error {
	var err error
	user.Hash, err = auth.HashPassword(password)
	return err
}
