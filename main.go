package main

import (
	"github.com/matDobek/gov--attendance-check/internal/db"
	"github.com/matDobek/gov--attendance-check/internal/server"
)

func main() {
	// statues, err := discovery.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	statues := []db.Statue{
		{1, 1, 1, "1st voting", []db.Vote{}},
		{2, 1, 1, "2nd voting", []db.Vote{}},
	}

	db := db.NewGovStore(statues)
	server := server.NewGovServer(db)
	server.Start()
}
