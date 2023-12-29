### Setup

```
psql -U postgres
create database gov;
create database gov_test;

go run ./scripts/migrate/migrate.go up
GBY_ENV=test go run ./scripts/migrate/migrate.go up
```

### Comments

```go
//=======================================================
// This is Header 1
//=======================================================

//----------------------------------
// This is Header 2
//----------------------------------

//
// This is Header 3
//

//

// fib
//=
//==
//===
//=====
//========
//=============
//=====================
//==================================
//=======================================================
```

### Directory Structure

```
cmd/
    main/
        main.go
scripts/
    migrate/
        migrate.go
api/
    v1/
        server.go
        routes.go
        votes.go
pkg/
    session/
    votes/
        storage/
            sql.go
            redis.go
            memory.go
        votes.go
        create.go
        delete.go
        get_all.go
        utils.go
internal/
    db/
    constrains/
    enumerable/
    predicates/
    logger/
    ...
```

```go
// pkg/votes/votes.go
package votes

// defines response objects
type Vote struct {
    ID          uint
    Name        string
    Description string
    ...
}

// define shared entry objects
type VoteParams struct {
    Name        string
    Description string
    ...
}

// and specific ones ( base on function name )
type CreateVoteParams struct {
    Name        string
    Description string
    ...
}

var (
    // errors
)

// all functions

func CreateVote(params CreateVoteParams) (Vote, error) {
    createVote(params)
}

// ...
```

```go
// pkg/votes/create_vote.go
package votes

type CreateVoteStore interface {
    insert(CreateVoteParams) (Vote, error)
}

func createVote(store CreateVoteStore, params CreateVoteParams) (Vote, error) {
    // ...
}

// ...
```

```go
// basic storage interface
storage.GetBy(params)
storage.All(params)
storage.Create(params)
storage.Update(params)
storage.Delete(params)
```
