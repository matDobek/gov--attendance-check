package discovery

import (
	"math/rand"

	"github.com/matDobek/gov--attendance-check/internal/db"
)

func MapToDBFormat(input []Statue) ([]db.Statue, []db.Vote, []db.Politician) {
	var statues []db.Statue
	var votes []db.Vote
	var politicians []db.Politician

	mapOfPoliticians := make(map[string]db.Politician)

	for _, statue := range input {
		sId := generateID()
		statues = append(statues, db.Statue{
			Id:        sId,
			VotingNo:  statue.VotingNo,
			SessionNo: statue.SessionNo,
			TermNo:    statue.TermNo,
			Title:     statue.Title,
		})

		for _, vote := range statue.Votes {
			pID := -1
			pol, ok := mapOfPoliticians[vote.Name]

			if ok {
				pID = pol.Id
			} else {
				pID = generateID()
				mapOfPoliticians[vote.Name] = db.Politician{Id: pID, Name: vote.Name, Party: vote.Party}
			}

			votes = append(votes, db.Vote{
				Id:           generateID(),
				PoliticianId: pID,
				StatueId:     sId,
				Response:     vote.Response,
			})
		}
	}

	for _, v := range mapOfPoliticians {
		politicians = append(politicians, v)
	}

	return statues, votes, politicians
}

func generateID() int {
	// TODO make something better
	//	or at least check if we repeated
	return rand.Intn(100000000)
}
