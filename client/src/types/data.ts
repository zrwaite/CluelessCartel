import { GameData } from './game'
import { User } from './user'

export const defaultUserData: User = {
	"Username": "Example",
	"Cash": 15000,
	"Bases": [
		{
			"Location": {
				"Name": "Nevada",
				"LandMaterial": {
					"Name": "Sand"
				}
			},
			"Index": 0,
			"Resources": [
				{
					"ResourceName": "Metal",
					"Amount": 10
				},
				{
					"ResourceName": "",
					"Amount": 10
				}
			],
			"Drugs": [
				{
					"ResourceName": "Weed",
					"Amount": 10
				},
				{
					"ResourceName": "Fentanyl",
					"Amount": 10
				},
				{
					"ResourceName": "Meth",
					"Amount": 10
				}
			],
			"Weapons": {
				"Guns": 1,
				"Explosives": 1,
				"Ammunition": 10
			},
			"HexagonRows": [
				{
					"Hexagons": [
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 5,
							"X": -2,
							"Owned": false,
							"Buyable": false
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 3,
							"X": -1,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 5,
							"X": 0,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 5,
							"X": 1,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 1,
							"X": 2,
							"Owned": false,
							"Buyable": false
						}
					],
					"Y": -2
				},
				{
					"Hexagons": [
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 0,
							"X": -2,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 1,
							"X": -1,
							"Owned": true,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 2,
							"X": 0,
							"Owned": true,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 4,
							"X": 1,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 0,
							"X": 2,
							"Owned": false,
							"Buyable": false
						}
					],
					"Y": -1
				},
				{
					"Hexagons": [
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 2,
							"X": -2,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 1,
							"X": -1,
							"Owned": true,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 0,
							"X": 0,
							"Owned": true,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 5,
							"X": 1,
							"Owned": true,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 4,
							"X": 2,
							"Owned": false,
							"Buyable": true
						}
					],
					"Y": 0
				},
				{
					"Hexagons": [
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 2,
							"X": -2,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 3,
							"X": -1,
							"Owned": true,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 5,
							"X": 0,
							"Owned": true,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 5,
							"X": 1,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 2,
							"X": 2,
							"Owned": false,
							"Buyable": false
						}
					],
					"Y": 1
				},
				{
					"Hexagons": [
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 5,
							"X": -2,
							"Owned": false,
							"Buyable": false
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 0,
							"X": -1,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 2,
							"X": 0,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 4,
							"X": 1,
							"Owned": false,
							"Buyable": true
						},
						{
							"LandMaterial": {
								"Name": "Sand"
							},
							"Structure": {
								"Moveable": true,
								"Enemy": false,
								"Natural": false,
								"Name": "Empty",
								"LandMaterials": [],
								"ResourceCapacity": []
							},
							"Rotation": 3,
							"X": 2,
							"Owned": false,
							"Buyable": false
						}
					],
					"Y": 2
				}
			]
		}
	],
	"NewBasePrice": 0
}

export const defaultGameData:GameData = {
	AllDrugs:     [],
	AllLocations: [],
	AllResources: [],
	AllStructures: [],
	AllLandMaterials: [],
}