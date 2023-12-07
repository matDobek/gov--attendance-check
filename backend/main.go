package main

import (
	"github.com/matDobek/gov--attendance-check/internal/db"
	"github.com/matDobek/gov--attendance-check/internal/discovery"
	"github.com/matDobek/gov--attendance-check/internal/server"
)

func main() {
	// dStatues, err := discovery.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	dStatues := []discovery.Statue{
		{VotingNo: 1, SessionNo: 1, TermNo: 1, Title: "1st voting", Votes: []discovery.Vote{
			{Name: "Jan Kowalski", Party: "KO", Response: db.VoteResponseYes},
			{Name: "Adam Nowak", Party: "PiS", Response: db.VoteResponseNo},
			{Name: "Marek Zbirek", Party: "Trzecia Droga", Response: db.VoteResponseMaybe},
		}},
		{VotingNo: 1, SessionNo: 1, TermNo: 1, Title: "1st voting", Votes: []discovery.Vote{
			{Name: "Jan Kowalski", Party: "KO", Response: db.VoteResponseYes},
			{Name: "Adam Nowak", Party: "PiS", Response: db.VoteResponseNo},
			{Name: "Marek Zbirek", Party: "Trzecia Droga", Response: db.VoteResponseMaybe},
		}},
	}
	statues, votes, politicians := discovery.MapToDBFormat(dStatues)

	db := db.NewGovStore(statues, politicians, votes)
	server := server.NewGovServer(db)
	server.Start()
}
