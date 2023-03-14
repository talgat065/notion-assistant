package telegram

import "net/http"

const apiUrl = "api.telegram"

type api struct {
	http  http.Client
	token string
}

func (a api) auth() {
	//
}
