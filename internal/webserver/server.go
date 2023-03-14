package webserver

import (
	"log"
	"net/http"
)

func NewServer() Server {
	return Server{":9000"}
}

type Server struct {
	Port string
}

func (s Server) Run() {
	s.initializeRoutes()
	log.Fatal(http.ListenAndServe(s.Port, nil))
}

func (s Server) initializeRoutes() {
	http.HandleFunc("/get-updates", MakeGetUpdatesHandler())
}
