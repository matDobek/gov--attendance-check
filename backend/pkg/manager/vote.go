package manager

import (
	"fmt"
  "time"
	"strings"

	"github.com/matDobek/gov--attendance-check/internal/types"
)

//=======================================================
// Types
//=======================================================

const (
	VoteResponseAbsent = "Nieobecny"
	VoteResponseYes    = "Za"
	VoteResponseNo     = "Przeciw"
	VoteResponseMaybe  = "Wstrzymał się"
)

type VoteStore interface {
	Insert(VoteParams) (Vote, error)
	All() ([]Vote, error)
}


type Vote struct {
	ID            int
  UpdatedAt     time.Time
  CreatedAt     time.Time

	PoliticianID int
	StatueID     int

	Response string
}

type VoteParams struct {
	StatueID      types.Maybe[int]
	PoliticianID  types.Maybe[int]
	Response      types.Maybe[string]
}

func NewVoteParams() *VoteParams {
  return &VoteParams{}
}

func (s *VoteParams) WithStatueID(v int) *VoteParams {
  s.StatueID = types.Some(v)
  return s
}

func (s *VoteParams) WithPoliticianID(v int) *VoteParams {
  s.PoliticianID = types.Some(v)
  return s
}

func (s *VoteParams) WithResponse(v string) *VoteParams {
  s.Response = types.Some(v)
  return s
}

func (s *VoteParams) IsValid() (bool, error) {
	var err VoteErrors

  v, b := s.StatueID.Unwrap()
  if !b {
		err.StatueID = append(err.StatueID, ErrValueRequired)
  }
  if v <= 0 {
		err.StatueID = append(err.StatueID, ErrPositiveValue)
	}

  v, b = s.PoliticianID.Unwrap()
  if !b {
		err.PoliticianID = append(err.PoliticianID, ErrValueRequired)
  }
  if v <= 0 {
		err.PoliticianID = append(err.PoliticianID, ErrPositiveValue)
	}

  str, b := s.Response.Unwrap()
  if !b {
		err.Response = append(err.StatueID, ErrValueRequired)
  }
  if strings.Trim(str, " \t\n") == "" {
		err.Response = append(err.Response, ErrNonZeroValue)
	}

	if len(err.Response) > 0 ||
		len(err.StatueID) > 0 ||
		len(err.PoliticianID) > 0 {
		return false, err
	}

	return true, nil
}

//=======================================================
// Errors
//=======================================================

type VoteErrors struct {
	StatueID     []error
	PoliticianID []error
	Response     []error
}

func (e VoteErrors) Error() string {
	var msg string

	msg = fmt.Sprintf("VoteErrors: ")

	if len(e.Response) > 0 {
		msg += fmt.Sprintf("Response: ")

		for _, err := range e.Response {
			msg += err.Error() + ", "
		}
	}

	if len(e.StatueID) > 0 {
		msg += fmt.Sprintf("statue id: ")

		for _, err := range e.StatueID {
			msg += err.Error() + ", "
		}
	}

	if len(e.PoliticianID) > 0 {
		msg += fmt.Sprintf("politician id: ")

		for _, err := range e.PoliticianID {
			msg += err.Error() + ", "
		}
	}

	return strings.Trim(msg, ", ")
}

func (e VoteErrors) Is(target error) bool {
	_, ok := target.(VoteErrors)
	return ok
}

//=======================================================
// Public Functions and
//    Function Specific Types/Interfaces
//=======================================================

//
//
//

func CreateVote(store VoteStore, params VoteParams) (Vote, error) {
  ok, err := params.IsValid()
  if !ok {
    return Vote{}, err
  }

  return store.Insert(params)
}

//
//
//

func AllVotes(store VoteStore) ([]Vote, error) {
	return store.All()
}
