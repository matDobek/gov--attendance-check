package statue_store

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
    collection: make(map[int]manager.Statue),
  }
}

//----------------------------------
// Store
//----------------------------------

type MemStore struct {
  collection map[int]manager.Statue
}

var (
	_ manager.StatueStore = (*MemStore)(nil)
)

//
//
//

func (s *MemStore) All() ([]manager.Statue, error) {
  collection := []manager.Statue{}

  for _, c := range s.collection {
    collection = append(collection, c)
  }

  return collection, nil
}

//
//
//

func (s *MemStore) Insert(params manager.StatueParams) (manager.Statue, error) {
  id := generateID(s.collection)
  currentTime := time.Now()
  title, _ := params.Title.Unwrap()
  votingNumber, _ := params.VotingNumber.Unwrap()
  sessionNumber, _ := params.SessionNumber.Unwrap()
  termNumber, _ := params.TermNumber.Unwrap()

  result := manager.Statue{
    ID: id,
    CreatedAt: currentTime,
    UpdatedAt: currentTime,
    Title: title,
    VotingNumber: votingNumber,
    SessionNumber: sessionNumber,
    TermNumber: termNumber,
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
