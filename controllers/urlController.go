package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/database"
	"url-shortener/models"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: HomeHandler")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Welcome to the URL Shortener!")
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {

	var url models.URL

	err := json.NewDecoder(r.Body).Decode(&url)

	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)

		return
	}

	query := "INSERT INTO Urls (LongUrl, ShortUrl, HitCount) VALUES (@p1, @p2, @p3)"

	_, err = database.DB.Exec(query, url.LongURL, url.ShortURL, url.HitCount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(url)
}

func GetShortUrl(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ShortURL := params["shortUrl"]

	if ShortURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	var url models.URL

	query := "SELECT * FROM Urls WHERE ShortUrl = @p1"

	err := database.DB.QueryRow(query, ShortURL).Scan(&url.ID, &url.LongURL, &url.ShortURL, &url.HitCount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}
