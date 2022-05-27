package models

type Structure struct {
	Moveable         bool
	Name             string
	Image            string
	Materials        []int
	ResourceCapacity Resources
}

var EmptyStructure = Structure{
	Moveable:         true,
	Name:             "Empty",
	Image:            "",
	Materials:        []int{},
	ResourceCapacity: Resources{},
}
var RV = Structure{
	Moveable:  true,
	Name:      "RV",
	Image:     "CC.svg",
	Materials: []int{},
	ResourceCapacity: Resources{
		Metal: 7,
	},
}
var BuriedStorage = Structure{
	Moveable:  false,
	Name:      "Buried Storage",
	Image:     "CC.svg",
	Materials: []int{LandMaterial.Dirt, LandMaterial.Grass, LandMaterial.Sand},
}
var StorageUnit = Structure{
	Moveable:  false,
	Name:      "Storage Unit",
	Image:     "CC.svg",
	Materials: []int{LandMaterial.Pavement},
}
var Shed = Structure{
	Moveable:  true,
	Name:      "Shed",
	Image:     "CC.svg",
	Materials: []int{},
}
var Armory = Structure{
	Moveable:  false,
	Name:      "Armory",
	Image:     "CC.svg",
	Materials: []int{LandMaterial.Pavement},
}
