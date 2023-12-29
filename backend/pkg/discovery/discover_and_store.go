package discovery

import (
	"github.com/matDobek/gov--attendance-check/pkg/manager"
)

func DiscoverAndStore(statueStore manager.StatueStore, politicianStore manager.PoliticianStore, voteStore manager.VoteStore) error {
	politicians := make(map[[2]string]manager.Politician)
	statues := make(map[[3]int]manager.Statue)

	ps, err := manager.AllPoliticians(politicianStore)
	if err != nil {
		return err
	}

	ss, err := manager.AllStatues(statueStore)
	if err != nil {
		return err
	}

	for _, v := range ps {
		key := [2]string{v.Name, v.Party}

		politicians[key] = v
	}

	for _, v := range ss {
		key := [3]int{v.TermNumber, v.SessionNumber, v.VotingNumber}

		statues[key] = v
	}

	discovered, err := Discover()
	if err != nil {
		return err
	}

	for _, d := range discovered {
		statueKey := [3]int{d.TermNo, d.SessionNo, d.VotingNo}

		_, ok := statues[statueKey]
		if ok {
			continue
		}

		statueParams := manager.
			NewStatueParams().
			WithTermNumber(d.TermNo).
			WithSessionNumber(d.SessionNo).
			WithVotingNumber(d.VotingNo).
			WithTitle(d.Title)

		statue, err := manager.CreateStatue(statueStore, *statueParams)
		if err != nil {
			return err
		}

		for _, v := range d.Votes {
			politicianKey := [2]string{v.Name, v.Party}

			politician, ok := politicians[politicianKey]
			if !ok {
				politicianParams := manager.
					NewPoliticianParams().
					WithName(v.Name).
					WithParty(v.Party)

				politician, err = manager.CreatePolitician(politicianStore, *politicianParams)
				if err != nil {
					return err
				}

				politicians[politicianKey] = politician
			}

			voteParams := manager.
				NewVoteParams().
				WithResponse(v.Response).
				WithStatueID(statue.ID).
				WithPoliticianID(politician.ID)

			_, err = manager.CreateVote(voteStore, *voteParams)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
