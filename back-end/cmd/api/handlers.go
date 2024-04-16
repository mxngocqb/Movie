package main

import (
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

	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *application) AllMovie(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil{
		app.errorJSON(w, err)
		log.Fatal(err)
	}
	
	_ = app.writeJSON(w, http.StatusOK, movies)
}
