package server

import (
	"encoding/json"
	"net/http"

	"github.com/matDobek/gov--attendance-check/internal/logger"
	"github.com/matDobek/gov--attendance-check/pkg/manager"
)

type GovServer struct {
	router *http.ServeMux

	statueStore     manager.StatueStore
	politicianStore manager.PoliticianStore
	voteStore       manager.VoteStore
}

func (s *GovServer) Start() {
	logger.Info("Starting gov server on port 8080")
	logger.Fatal(http.ListenAndServe(":8080", s.router))
}

func NewGovServer(statueStore manager.StatueStore, politicianStore manager.PoliticianStore, voteStore manager.VoteStore) *GovServer {
	s := &GovServer{
		router: http.NewServeMux(),

		statueStore:     statueStore,
		politicianStore: politicianStore,
		voteStore:       voteStore,
	}

	s.router.Handle("/api/v1/statues/", http.HandlerFunc(s.handleStatues))
	s.router.Handle("/api/v1/politicians/", http.HandlerFunc(s.handlePoliticians))
	s.router.Handle("/api/v1/votes/", http.HandlerFunc(s.handleVotes))

	return s
}

func (s *GovServer) handleStatues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := manager.AllStatues(s.statueStore)
	if err != nil {
		logger.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Fatal(err)
	}
}

func (s *GovServer) handleVotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := manager.AllVotes(s.voteStore)
	if err != nil {
		logger.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Fatal(err)
	}
}

func (s *GovServer) handlePoliticians(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := manager.AllPoliticians(s.politicianStore)
	if err != nil {
		logger.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Fatal(err)
	}
}
