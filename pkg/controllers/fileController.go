package controllers

import (
	"net/http"
)

type Handler interface {
	Load(w http.ResponseWriter, r *http.Request)
}

func (s *Server) LoadHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("Loading file")

	s.ServiceReader.ReadFile("./mock-data/sales.csv")

	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}
