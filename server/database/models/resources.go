package models

type Resources struct {
	Metal     int
	Plants    int
	Chemicals int
}

var StartingResources = Resources{0, 10, 10}
