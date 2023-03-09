package webserver

import (
	"net/http"
)

func MakeDefaultHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//
	}
}
