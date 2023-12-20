package manager

import (
	"errors"
	"fmt"
	"strings"
)

//=======================================================
// Types
//=======================================================

type Statue struct {
	ID            int    `json:"id"`
	VotingNumber  int    `json:"votingNumber"`
	SessionNumber int    `json:"sessionNumber"`
	TermNumber    int    `json:"termNumber"`
	Title         string `json:"title"`
}

type StatueParams struct {
	VotingNumber  int
	SessionNumber int
	TermNumber    int
	Title         string
}

//=======================================================
// Errors
//=======================================================

var (
	ErrNonZeroValue  = errors.New("expected non zero value")
	ErrPositiveValue = errors.New("expected positive value")
	ErrNegativeValue = errors.New("expected negative zero value")
)

type StatueErrors struct {
	VotingNumber  []error
	SessionNumber []error
	TermNumber    []error
	Title         []error
}

func (e StatueErrors) Error() string {
	var msg string

	msg = fmt.Sprintf("StatueErrors: ")

	if len(e.Title) > 0 {
		msg += fmt.Sprintf("title: ")

		for _, err := range e.Title {
			msg += err.Error() + ", "
		}
	}

	if len(e.VotingNumber) > 0 {
		msg += fmt.Sprintf("voting number: ")

		for _, err := range e.VotingNumber {
			msg += err.Error() + ", "
		}
	}

	if len(e.SessionNumber) > 0 {
		msg += fmt.Sprintf("session number: ")

		for _, err := range e.SessionNumber {
			msg += err.Error() + ", "
		}
	}

	if len(e.TermNumber) > 0 {
		msg += fmt.Sprintf("term number: ")

		for _, err := range e.TermNumber {
			msg += err.Error() + ", "
		}
	}

	return strings.Trim(msg, ", ")
}

func (e StatueErrors) Is(target error) bool {
	_, ok := target.(StatueErrors)
	return ok
}

//=======================================================
// Public Functions and
//    Function Specific Types/Interfaces
//=======================================================

//
//
//

func ValidateStatue(s StatueParams) (bool, error) {
	var err StatueErrors

	if s.VotingNumber <= 0 {
		err.VotingNumber = append(err.VotingNumber, ErrNonZeroValue)
		err.VotingNumber = append(err.VotingNumber, ErrPositiveValue)
	}

	if s.SessionNumber <= 0 {
		err.SessionNumber = append(err.SessionNumber, ErrNonZeroValue)
		err.SessionNumber = append(err.SessionNumber, ErrPositiveValue)
	}

	if s.TermNumber <= 0 {
		err.TermNumber = append(err.TermNumber, ErrNonZeroValue)
		err.TermNumber = append(err.TermNumber, ErrPositiveValue)
	}

	if strings.Trim(s.Title, " \t\n") == "" {
		err.Title = append(err.Title, ErrNonZeroValue)
	}

	if len(err.Title) > 0 ||
		len(err.VotingNumber) > 0 ||
		len(err.SessionNumber) > 0 ||
		len(err.TermNumber) > 0 {
		return false, err
	}

	return true, nil
}

//
//
//

type CreateStatueStore interface {
	Insert(Statue) (Statue, error)
}

func CreateStatue(store CreateStatueStore, params StatueParams) (Statue, error) {
	statue := buildStatue(params)

	return store.Insert(statue)
}

func buildStatue(p StatueParams) Statue {
	return Statue{
		VotingNumber:  p.VotingNumber,
		SessionNumber: p.SessionNumber,
		TermNumber:    p.TermNumber,
		Title:         p.Title,
	}
}

//
//
//

type AllStatuesStore interface {
	All() ([]Statue, error)
}

func AllStatues(store AllStatuesStore) ([]Statue, error) {
	return store.All()
}

//
//
//

// type Vote struct {
// 	ID           int    `json:"id"`
// 	PoliticianId int    `json:"politicianId"`
// 	StatueId     int    `json:"statueId"`
// 	Response     string `json:"response"`
// }
//
// type Politician struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Party string `json:"party"`
// }
//

// const (
// 	VoteResponseAbsent = "Nieobecny"
// 	VoteResponseYes    = "Za"
// 	VoteResponseNo     = "Przeciw"
// 	VoteResponseMaybe  = "Wstrzymał się"
// )
//

// truthTeller(statues, "Nieobecny", "Najwiekszy obibok [nieobecny]")
// truthTeller(statues, "Wstrzymał się", "Najbardziej wstrzemięźliwy [wstrzymał się]")
// truthTeller(statues, "Za", "yes men [za]")
// truthTeller(statues, "Przeciw", "no men [przeciw]")
// func truthTeller(statues []Statue, response string, msg string) {
// 	logger.Info("========  %v  ========", msg)
// 	votes := make(map[string]int)
// 	for _, s := range statues {
// 		for _, v := range s.Votes {
// 			if v.Response == response {
// 				votes[v.Name] += 1
// 			}
// 		}
// 	}
//
// 	keys := make([]string, 0, len(votes))
// 	for k := range votes {
// 		keys = append(keys, k)
// 	}
//
// 	sort.Slice(keys, func(i, j int) bool {
// 		return votes[keys[i]] > votes[keys[j]]
// 	})
//
// 	l := len(keys)
// 	if l > 5 {
// 		l = 5
// 	}
// 	for _, k := range keys[:l] {
// 		logger.Info("%v/%v - %v", votes[k], len(statues), k)
// 	}
//
// }
