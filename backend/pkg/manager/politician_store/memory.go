package politician_store

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
    collection: make(map[int]manager.Politician),
  }
}

//----------------------------------
// Store
//----------------------------------

type MemStore struct {
  collection map[int]manager.Politician
}

var (
	_ manager.PoliticianStore = (*MemStore)(nil)
)

//
//
//

func (s *MemStore) All() ([]manager.Politician, error) {
  collection := []manager.Politician{}

  for _, c := range s.collection {
    collection = append(collection, c)
  }

  return collection, nil
}

//
//
//

func (s *MemStore) Insert(params manager.PoliticianParams) (manager.Politician, error) {
  id := generateID(s.collection)
  currentTime := time.Now()
  name, _ := params.Name.Unwrap()
  party, _ := params.Party.Unwrap()

  result := manager.Politician{
    ID: id,
    CreatedAt: currentTime,
    UpdatedAt: currentTime,
    Name: name,
    Party: party,
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
