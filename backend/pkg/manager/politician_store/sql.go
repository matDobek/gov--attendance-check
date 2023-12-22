package politician_store

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
// PoliticianStore
//----------------------------------

type SQLStore struct {
	db *sql.DB
}

var (
	_ manager.CreatePoliticianStore = (*SQLStore)(nil)
	_ manager.AllPoliticiansStore   = (*SQLStore)(nil)
)

//
//
//

func (s *SQLStore) All() ([]manager.Politician, error) {
	var result []manager.Politician

	q := `
		select
      id,
      name,
      party,
      updated_at,
      created_at
    from politicians
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
		Politician := manager.Politician{}

		err := rows.Scan(
			&Politician.ID,
			&Politician.Name,
			&Politician.Party,
      &Politician.UpdatedAt,
      &Politician.CreatedAt)

		if err != nil {
			return nil, err
		}

		result = append(result, Politician)
	}

	return result, nil
}

//
//
//

func (s *SQLStore) Insert(params manager.PoliticianParams) (manager.Politician, error) {
  Politician := manager.Politician{}
	q := `
    insert into politicians (
        name,
        party
      )
      values ($1, $2)
      returning
        id,
        name,
        party,
        updated_at,
        created_at
  `

	stmt, err := s.db.Prepare(q)
	if err != nil {
		return manager.Politician{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		params.Name,
		params.Party).Scan(
      &Politician.ID,
			&Politician.Name,
			&Politician.Party,
      &Politician.UpdatedAt,
      &Politician.CreatedAt)
	if err != nil {
		return manager.Politician{}, err
	}

	return Politician, nil
}
