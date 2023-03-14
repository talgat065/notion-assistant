package webserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/talgat065/notion-assistant/internal/telegram"
	"io"
	"log"
	"net/http"
)

func MakeGetUpdatesHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		defer r.Body.Close()
		if err != nil {
			msg := fmt.Sprintf("Failed to parse request body")
			http.Error(w, msg, http.StatusBadRequest)
			fmt.Println(msg)
		}
		// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
		// response body. A request body larger than that will now result in
		// Decode() returning a "http: request body too large" error.
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		// Setup the decoder and call the DisallowUnknownFields() method on it.
		// This will cause Decode() to return a "json: unknown field ..." error
		// if it encounters any extra unexpected fields in the JSON. Strictly
		// speaking, it returns an error for "keys which do not match any
		// non-ignored, exported fields in the destination".
		dec := json.NewDecoder(r.Body)

		var u telegram.Update
		err = dec.Decode(&u)
		if err != nil {
			var syntaxError *json.SyntaxError
			var unmarshalTypeError *json.UnmarshalTypeError

			switch {
			// Catch any syntax errors in the JSON and send an error message
			// which interpolates the location of the problem to make it
			// easier for the client to fix.
			case errors.As(err, &syntaxError):
				msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
				http.Error(w, msg, http.StatusBadRequest)
				fmt.Println(msg)

			// In some circumstances Decode() may also return an
			// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
			// is an open issue regarding this at
			// https://github.com/golang/go/issues/25956.
			case errors.Is(err, io.ErrUnexpectedEOF):
				msg := fmt.Sprintf("Request body contains badly-formed JSON")
				http.Error(w, msg, http.StatusBadRequest)
				fmt.Println(msg)
			// Catch any type errors, like trying to assign a string in the
			// JSON request body to a int field in our Person struct. We can
			// interpolate the relevant field name and position into the error
			// message to make it easier for the client to fix.
			case errors.As(err, &unmarshalTypeError):
				msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
				http.Error(w, msg, http.StatusBadRequest)
				fmt.Println(msg)

			// An io.EOF error is returned by Decode() if the request body is
			// empty.
			case errors.Is(err, io.EOF):
				msg := "Request body must not be empty"
				http.Error(w, msg, http.StatusBadRequest)
				fmt.Println(msg)

			// Catch the error caused by the request body being too large. Again
			// there is an open issue regarding turning this into a sentinel
			// error at https://github.com/golang/go/issues/30715.
			case err.Error() == "http: request body too large":
				msg := "Request body must not be larger than 1MB"
				http.Error(w, msg, http.StatusRequestEntityTooLarge)
				fmt.Println(msg)

			// Otherwise default to logging the error and sending a 500 Internal
			// Server Error response.
			default:
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}

		fmt.Println(u.UpdateID)
		return
	}
}
