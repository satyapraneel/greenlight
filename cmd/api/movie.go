package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/satyapraneel/greenlight/internal/data"
)

func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieId, err := app.readIdFromParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		Id:        movieId,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	// envelope["movie"] = movie -> otherway to add value to map
	err = app.writeJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
