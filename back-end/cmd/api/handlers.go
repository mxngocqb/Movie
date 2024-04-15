package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and runnung",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func (app *application) AllMovie(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil{
		log.Fatal(err)
	}
	out, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
