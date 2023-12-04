package server

import (
	"encoding/json"
	"net/http"

	"github.com/matDobek/gov--attendance-check/internal/db"
	"github.com/matDobek/gov--attendance-check/internal/logger"
)

type GovServer struct {
	store  *db.GovStore
	router *http.ServeMux
}

func (s *GovServer) Start() {
	logger.Fatal(http.ListenAndServe(":8080", s.router))
}

func NewGovServer(store *db.GovStore) *GovServer {
	s := &GovServer{
		store:  store,
		router: http.NewServeMux(),
	}

	s.router.Handle("/statues/", http.HandlerFunc(s.handleStatues))

	return s
}

func (s *GovServer) handleStatues(w http.ResponseWriter, r *http.Request) {
	statues := s.store.GetStatues()

	json.NewEncoder(w).Encode(statues)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
