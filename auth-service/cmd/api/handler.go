package main

import (
	"errors"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user agains the database

	user, err := app.Models.User.GetByEmail(requestPayload.Email)

	if err != nil {
		app.errorJSON(w, errors.New("invalid email or password"), http.StatusBadRequest)
		return
	}

	// validate the password
	valid, err := user.PasswordMatches(requestPayload.Password)

	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid email or password"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Authentication successful",
		Data:    user,
	}

	app.writeJSON(w, http.StatusOK, payload)

}
