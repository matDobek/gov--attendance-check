package statue_store

import (
	"database/sql"

	"github.com/matDobek/gov--attendance-check/pkg/manager"
)

//----------------------------------
// New
//----------------------------------

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{db: db}
}

//----------------------------------
// Store
//----------------------------------

type SQLStore struct {
	db *sql.DB
}

var (
	_ manager.StatueStore   = (*SQLStore)(nil)
)

//
//
//

func (s *SQLStore) All() ([]manager.Statue, error) {
	var result []manager.Statue

	q := `
		select
      id,
      title,
      term_number,
      session_number,
      voting_number,
      updated_at,
      created_at
    from statues
	`

	stmt, err := s.db.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		statue := manager.Statue{}

		err := rows.Scan(
			&statue.ID,
			&statue.Title,
			&statue.TermNumber,
			&statue.SessionNumber,
			&statue.VotingNumber,
      &statue.UpdatedAt,
      &statue.CreatedAt)

		if err != nil {
			return nil, err
		}

		result = append(result, statue)
	}

	return result, nil
}

//
//
//

func (s *SQLStore) Insert(params manager.StatueParams) (manager.Statue, error) {
  statue := manager.Statue{}
	q := `
    insert into statues (
        title,
        term_number,
        session_number,
        voting_number
      )
      values ($1, $2, $3, $4)
      returning
        id,
        title,
        term_number,
        session_number,
        voting_number,
        updated_at,
        created_at
  `

	stmt, err := s.db.Prepare(q)
	if err != nil {
		return manager.Statue{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		params.Title,
		params.TermNumber,
		params.SessionNumber,
		params.VotingNumber).Scan(
      &statue.ID,
			&statue.Title,
			&statue.TermNumber,
			&statue.SessionNumber,
			&statue.VotingNumber,
      &statue.UpdatedAt,
      &statue.CreatedAt)
	if err != nil {
		return manager.Statue{}, err
	}

	return statue, nil
}
