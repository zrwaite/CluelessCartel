package main

import (
	"clueless-cartel-server/api"
	"clueless-cartel-server/database"
	"clueless-cartel-server/settings"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	settings.MatchDev()
	database.ConnectToMongoDB()
	database.InitializeDatabase()
	fileServer := http.FileServer(http.Dir("../client"))
	http.HandleFunc("/api/", api.APIHandler)
	http.Handle("/", fileServer)
	fmt.Printf("Starting server at port 8004 with hello\n")
	if err := http.ListenAndServe(":8004", nil); err != nil {
		log.Fatal(err)
	}
}
