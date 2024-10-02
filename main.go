package main

import (
	"fmt"
	"log"
	"net/http"

	"url-shortener/database"
	"url-shortener/routers"
)

func main() {
	database.ConnectDB() // Call without error handling

	router := routers.SetupRouter()

	fmt.Println("Server is running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
