package settings

import (
	"log"

	"github.com/joho/godotenv"
)

const DEV = false

func MatchDev() {
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envDev := envs["DEV"]
	if DEV && envDev != "true" {
		log.Fatal("Backend production environment does not match environment")
	} else if !DEV && envDev != "false" {
		log.Fatal("Backend development environment does not match environment")
	}
}
