package models

type Drugs struct {
	Opioids int
	Weed    int
	Meth    int
}

var StartingDrugs = Drugs{10, 10, 0}
