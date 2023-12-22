package vote_store

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
// VoteStore
//----------------------------------

type SQLStore struct {
	db *sql.DB
}

var (
	_ manager.CreateVoteStore = (*SQLStore)(nil)
	_ manager.AllVotesStore   = (*SQLStore)(nil)
)

//
//
//

func (s *SQLStore) All() ([]manager.Vote, error) {
	var result []manager.Vote

	q := `
		select
      id,
      response,
      statue_id,
      politician_id,
      updated_at,
      created_at
    from votes
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
		Vote := manager.Vote{}

		err := rows.Scan(
			&Vote.ID,
			&Vote.Response,
			&Vote.StatueID,
			&Vote.PoliticianID,
      &Vote.UpdatedAt,
      &Vote.CreatedAt)

		if err != nil {
			return nil, err
		}

		result = append(result, Vote)
	}

	return result, nil
}

//
//
//

func (s *SQLStore) Insert(params manager.VoteParams) (manager.Vote, error) {
  Vote := manager.Vote{}
	q := `
    insert into votes (
        response,
        statue_id,
        politician_id
      )
      values ($1, $2, $3)
      returning
        id,
        response,
        statue_id,
        politician_id,
        updated_at,
        created_at
  `

	stmt, err := s.db.Prepare(q)
	if err != nil {
		return manager.Vote{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		params.Response,
		params.StatueID,
		params.PoliticianID).Scan(
      &Vote.ID,
			&Vote.Response,
			&Vote.StatueID,
			&Vote.PoliticianID,
      &Vote.UpdatedAt,
      &Vote.CreatedAt)
	if err != nil {
		return manager.Vote{}, err
	}

	return Vote, nil
}
