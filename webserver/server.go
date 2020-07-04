package main

import (
	"net/http"
)

type Server struct {
	port string
	router *Router
}

func new_server(port string) *Server {
	return &Server{
		port: port,
		router: new_router(),
	}
}

func (s *Server) Listen () error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
