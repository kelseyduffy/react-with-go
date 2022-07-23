package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// using ...int makes this an optional parameter since it takes 0 or more values. not a breaking change
func (app *application) errorJson(w http.ResponseWriter, err error, status ...int) {

	// Assume status code of 400
	statusCode := http.StatusBadRequest

	// overwrite that if a status code was passed
	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(w, statusCode, theError, "error")
}
