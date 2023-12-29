package main

import (
	"github.com/matDobek/gov--attendance-check/internal/server"
	"github.com/matDobek/gov--attendance-check/internal/storage"
	"github.com/matDobek/gov--attendance-check/internal/utils"
	"github.com/matDobek/gov--attendance-check/pkg/manager/politician_store"
	"github.com/matDobek/gov--attendance-check/pkg/manager/statue_store"
	"github.com/matDobek/gov--attendance-check/pkg/manager/vote_store"
)

func main() {
	storage := storage.NewStorage(utils.PrimaryDatabaseURL())

	statueStore := statue_store.NewSQLStore(storage.PrimaryDB)
	politicianStore := politician_store.NewSQLStore(storage.PrimaryDB)
	voteStore := vote_store.NewSQLStore(storage.PrimaryDB)

	server := server.NewGovServer(statueStore, politicianStore, voteStore)
	server.Start()
}
