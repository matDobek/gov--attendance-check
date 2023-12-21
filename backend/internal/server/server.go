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

	s.router.Handle("/api/v1/statues/", http.HandlerFunc(s.handleStatues))
	s.router.Handle("/api/v1/politicians/", http.HandlerFunc(s.handlePoliticians))
	s.router.Handle("/api/v1/votes/", http.HandlerFunc(s.handleVotes))

	return s
}

func (s *GovServer) handleStatues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

  err := json.NewEncoder(w).Encode(s.store.GetStatues())
  if err != nil {
    logger.Fatal(err)
  }
}

func (s *GovServer) handleVotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

  err := json.NewEncoder(w).Encode(s.store.GetVotes())
  if err != nil {
    logger.Fatal(err)
  }
}

func (s *GovServer) handlePoliticians(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

  err := json.NewEncoder(w).Encode(s.store.GetPoliticians())
  if err != nil {
    logger.Fatal(err)
  }
}