package settings

import (
	"log"
	"os"
)

const DEV = false

func MatchDev() {
	envDev := os.Getenv("DEV")
	if envDev == "" {
		log.Fatal("Failed to load .env file")
	}
	if DEV && envDev != "true" {
		log.Fatal("Backend production environment does not match environment")
	} else if !DEV && envDev != "false" {
		log.Fatal("Backend development environment does not match environment")
	}
}
