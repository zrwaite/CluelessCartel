package main

import (
	"clueless-cartel-server/api/user"
	"clueless-cartel-server/database"

	"fmt"
	"log"
	"net/http"
)

func main() {
	database.ConnectToMongoDB()
	fileServer := http.FileServer(http.Dir("../client"))
	http.HandleFunc("/api/user", user.UserHandler)
	http.Handle("/", fileServer)
	fmt.Printf("Starting server at port 8004 with hello\n")
	if err := http.ListenAndServe(":8004", nil); err != nil {
		log.Fatal(err)
	}
}
