package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	routes := httprouter.New()
	//method not allowed custom
	routes.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	//not found custom
	routes.NotFound = http.HandlerFunc(app.notFoundResponse)
	routes.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	routes.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovie)
	routes.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	return routes
}
