package manager

import (
	"fmt"
	"strings"
	"time"

	"github.com/matDobek/gov--attendance-check/internal/types"
)

//=======================================================
// Types
//=======================================================

type PoliticianStore interface {
	Insert(PoliticianParams) (Politician, error)
	All() ([]Politician, error)
}

type Politician struct {
	ID            int
  UpdatedAt     time.Time
  CreatedAt     time.Time

	Party string
	Name  string
}

type PoliticianParams struct {
	Party types.Maybe[string]
	Name  types.Maybe[string]
}

func NewPoliticianParams() *PoliticianParams{
  return &PoliticianParams{}
}

func (s *PoliticianParams) WithParty(v string) *PoliticianParams {
  s.Party = types.Some(v)
  return s
}

func (s *PoliticianParams) WithName(v string) *PoliticianParams {
  s.Name = types.Some(v)
  return s
}

func (s *PoliticianParams) IsValid() (bool, error) {
	var err PoliticianErrors

  str, b := s.Party.Unwrap()
  if !b {
		err.Party = append(err.Party, ErrValueRequired)
  }
  if strings.Trim(str, " \t\n") == "" {
		err.Party = append(err.Party, ErrPositiveValue)
	}

  str, b = s.Name.Unwrap()
  if !b {
		err.Name = append(err.Name, ErrValueRequired)
  }
  if strings.Trim(str, " \t\n") == "" {
		err.Name = append(err.Name, ErrPositiveValue)
	}

	if len(err.Party) > 0 || len(err.Name) > 0 {
		return false, err
	}

	return true, nil
}

//=======================================================
// Errors
//=======================================================

type PoliticianErrors struct {
	Party  []error
	Name   []error
}

func (e PoliticianErrors) Error() string {
	var msg string

	msg = fmt.Sprintf("PoliticianErrors: ")

	if len(e.Party) > 0 {
		msg += fmt.Sprintf("party: ")

		for _, err := range e.Party {
			msg += err.Error() + ", "
		}
	}

	if len(e.Name) > 0 {
		msg += fmt.Sprintf("name: ")

		for _, err := range e.Name {
			msg += err.Error() + ", "
		}
	}

	return strings.Trim(msg, ", ")
}

func (e PoliticianErrors) Is(target error) bool {
	_, ok := target.(PoliticianErrors)
	return ok
}

//=======================================================
// Public Functions and
//    Function Specific Types/Interfaces
//=======================================================

//
//
//

func CreatePolitician(store PoliticianStore, params PoliticianParams) (Politician, error) {
  ok, err := params.IsValid()
  if !ok {
    return Politician{}, err
  }

  return store.Insert(params)
}

//
//
//

func AllPoliticians(store PoliticianStore) ([]Politician, error) {
	return store.All()
}
