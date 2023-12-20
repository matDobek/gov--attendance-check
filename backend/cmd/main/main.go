package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/matDobek/gov--attendance-check/internal/logger"
	"github.com/matDobek/gov--attendance-check/internal/storage"
	"github.com/matDobek/gov--attendance-check/pkg/manager"
	"github.com/matDobek/gov--attendance-check/pkg/manager/statue_store"
)

func main() {
	// dStatues, err := discovery.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// dStatues := []discovery.Statue{
	// 	{VotingNo: 1, SessionNo: 1, TermNo: 1, Title: "1st voting", Votes: []discovery.Vote{
	// 		{Name: "Jan Kowalski", Party: "KO", Response: db.VoteResponseYes},
	// 		{Name: "Adam Nowak", Party: "PiS", Response: db.VoteResponseNo},
	// 		{Name: "Marek Zbirek", Party: "Trzecia Droga", Response: db.VoteResponseMaybe},
	// 	}},
	// 	{VotingNo: 1, SessionNo: 1, TermNo: 1, Title: "1st voting", Votes: []discovery.Vote{
	// 		{Name: "Jan Kowalski", Party: "KO", Response: db.VoteResponseYes},
	// 		{Name: "Adam Nowak", Party: "PiS", Response: db.VoteResponseNo},
	// 		{Name: "Marek Zbirek", Party: "Trzecia Droga", Response: db.VoteResponseMaybe},
	// 	}},
	// }
	//
	err := godotenv.Load()
	if err != nil {
		logger.Fatal(errors.New("error loading .env file"))
		return
	}

	databaseURL := os.Getenv("DB__MAIN__URL")
	if databaseURL == "" {
		logger.Fatal(errors.New("no database url"))
		return
	}

	storage := storage.NewStorage(databaseURL)
	s := statue_store.NewSQLStatusStore(storage)
	// v := manager.StatueParams{
	// 	SessionNumber: 1,
	// 	TermNumber:    1,
	// 	VotingNumber:  1,
	// 	Title:         "1st voting",
	// }
	//
	// _, errs := manager.CreateStatue(s, v)
	// if errs != nil {
	// 	logger.Fatal(errs)
	// }
	//

	statues, err := manager.AllStatues(s)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("%#v", statues)

	// db := db.NewDB(databaseURL)

	// statues, votes, politicians := discovery.MapToDBFormat(dStatues)
	// db := db.NewGovStore(statues, politicians, votes)
	// server := server.NewGovServer(db)
	// server.Start()
}
