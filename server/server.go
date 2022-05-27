package main

import (
	"clueless-cartel-server/api/hello"
	"clueless-cartel-server/api/user"
	"clueless-cartel-server/database"

	"fmt"
	"log"
	"net/http"
)

func main() {
	database.ConnectToMongoDB()
	fileServer := http.FileServer(http.Dir("../client"))
	http.HandleFunc("/hello", hello.HelloHandler)
	http.HandleFunc("/api/user", user.UserHandler)
	http.Handle("/", fileServer)
	fmt.Printf("Starting server at port 6969 with hello\n")
	if err := http.ListenAndServe(":6969", nil); err != nil {
		log.Fatal(err)
	}
}
