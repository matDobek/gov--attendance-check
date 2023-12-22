package vote_store

import (
	"crypto/rand"
	"math"
	"math/big"
	"time"

	"github.com/matDobek/gov--attendance-check/pkg/manager"
)

//----------------------------------
// New
//----------------------------------

func NewMemStore() *MemStore {
	return &MemStore{
    collection: make(map[int]manager.Vote),
  }
}

//----------------------------------
// VoteStore
//----------------------------------

type MemStore struct {
  collection map[int]manager.Vote
}

var (
	_ manager.CreateVoteStore = (*MemStore)(nil)
	_ manager.AllVotesStore   = (*MemStore)(nil)
)

//
//
//

func (s *MemStore) All() ([]manager.Vote, error) {
  collection := []manager.Vote{}

  for _, c := range s.collection {
    collection = append(collection, c)
  }

  return collection, nil
}

//
//
//

func (s *MemStore) Insert(params manager.VoteParams) (manager.Vote, error) {
  id := generateID(s.collection)
  currentTime := time.Now()
  response, _ := params.Response.Unwrap()
  politicianID, _ := params.PoliticianID.Unwrap()
  statueID, _ := params.StatueID.Unwrap()

  result := manager.Vote{
    ID: id,
    CreatedAt: currentTime,
    UpdatedAt: currentTime,
    Response: response,
    PoliticianID: politicianID,
    StatueID: statueID,
  }

  s.collection[id] = result

  return result, nil
}

func generateID[T any](collection map[int]T) int {
  var id int

  for {
    n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
    if err != nil {
      panic(err)
    }

    id = int(n.Int64())

    if _, exists := collection[id]; exists {
      continue
    }

    return id
  }
}
