package routers

import (
	"url-shortener/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/api/shorten", controllers.CreateShortUrl).Methods("POST")
	router.HandleFunc("/api/{shortUrl}", controllers.GetShortUrl).Methods("GET")

	return router
}
