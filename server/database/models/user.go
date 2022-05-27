package models

import (
	"clueless-cartel-server/auth"

	"github.com/jinzhu/copier"
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
	Username string
	Password string
}

type PostUser struct {
	Username    string
	Hash        string
	Cash        int
	Resources   Resources
	Drugs       Drugs
	Weapons     Weapons
	HexagonRows []HexagonRow `bson:"hexagon_rows"`
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
