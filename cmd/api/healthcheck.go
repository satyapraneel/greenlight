package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := envelope{"status": "available", "env": app.config.env, "version": version}
	err := app.writeJson(w, http.StatusOK, data, r.Header)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
