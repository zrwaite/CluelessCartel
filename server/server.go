package main

import (
	"clueless-cartel-server/api/base"
	"clueless-cartel-server/api/user"
	"clueless-cartel-server/database"
	"clueless-cartel-server/settings"

	"fmt"
	"log"
	"net/http"
)

func main() {
	database.ConnectToMongoDB()
	// database.InitializeDatabase()
	settings.MatchDev()
	fileServer := http.FileServer(http.Dir("../client"))
	http.HandleFunc("/api/user", user.UserHandler)
	http.HandleFunc("/api/base", base.BaseHandler)
	http.Handle("/", fileServer)
	fmt.Printf("Starting server at port 8004 with hello\n")
	if err := http.ListenAndServe(":8004", nil); err != nil {
		log.Fatal(err)
	}
}
