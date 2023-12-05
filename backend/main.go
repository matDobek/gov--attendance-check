package main

import (
	"github.com/matDobek/gov--attendance-check/internal/db"
	"github.com/matDobek/gov--attendance-check/internal/discovery"
	"github.com/matDobek/gov--attendance-check/internal/server"
)

func main() {
	// statues, err := discovery.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	dStatues := []discovery.Statue{
		{1, 1, 1, "1st voting", []discovery.Vote{
			{"Jan Kowalski", "KO", db.VoteResponseYes},
			{"Adam Nowak", "PiS", db.VoteResponseNo},
			{"Marek Zbirek", "Trzecia Droga", db.VoteResponseMaybe},
		}},
		{1, 1, 1, "1st voting", []discovery.Vote{
			{"Jan Kowalski", "KO", db.VoteResponseYes},
			{"Adam Nowak", "PiS", db.VoteResponseNo},
			{"Marek Zbirek", "Trzecia Droga", db.VoteResponseMaybe},
		}},
	}
	statues, votes, politicians := discovery.MapToDBFormat(dStatues)

	db := db.NewGovStore(statues, politicians, votes)
	server := server.NewGovServer(db)
	server.Start()
}
