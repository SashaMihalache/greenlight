package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, envelope{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem", http.StatusInternalServerError)
	}
}
