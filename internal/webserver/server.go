package webserver

import (
	"net/http"
)

type Server struct {
	Port string
}

func (s *Server) initializeRoutes() {
	http.HandleFunc("/create", MakeDefaultHandler())
}

func (s *Server) Run() error {
	s.initializeRoutes()
	return http.ListenAndServe("127.0.0.1"+s.Port, nil)
}
