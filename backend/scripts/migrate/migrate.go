package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/matDobek/gov--attendance-check/internal/logger"
)

const (
	NewCmd   = "new"   // new create_user_table
	UpCmd    = "up"    // up
	DownCmd  = "down"  // down
	ForceCmd = "force" // force
)

const MigrationDir = "db/migrations/"

//
// We're making heavy use of go migrate
// ref: https://pkg.go.dev/github.com/golang-migrate/migrate/v4#pkg-overview
//

func main() {
	commands := []string{
		NewCmd,
		UpCmd,
		DownCmd,
		ForceCmd,
	}

	if len(os.Args) == 1 {
		panic("command not provided")
	} else if !includes(commands, os.Args[1]) {
		panic("unknown command")
	}

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	var databaseURL string
	if os.Getenv("GBY_ENV") == "test" {
		databaseURL = os.Getenv("DB__MAIN__URL_TEST")
	} else {
		databaseURL = os.Getenv("DB__MAIN__URL")
	}

	switch os.Args[1] {
	case NewCmd:
		if len(os.Args) < 2 {
			if os.Args[2] == "" {
				panic("migration name not provided")
			}
		}

		newMigrationFile(os.Args[2])
	case UpCmd:
		up(databaseURL)
	case DownCmd:
		down(databaseURL)
	case ForceCmd:
		if len(os.Args) < 2 {
			if os.Args[2] == "" {
				panic("migration version not provided")
			}
		}

		force(databaseURL, os.Args[2])
	}
}

func newMigrationFile(name string) {
	timestamp := fmt.Sprintf("%v", time.Now().Unix())
	for _, kind := range []string{"up", "down"} {
		filename := timestamp + "_" + name + "." + kind + ".sql"

		file, err := os.Create(MigrationDir + filename)
		if err != nil {
			panic(err)
		}

		fmt.Println("Created: ", file.Name())
		defer file.Close()
	}
}

func up(databaseURL string) {
	m := initMigration(databaseURL)
	defer m.Close()

	err := m.Up()

	if err != nil && err.Error() == "no change" {
		logger.Info("No new migrations found")
	} else if err != nil {
		panic(err)
	}
}

func down(databaseURL string) {
	m := initMigration(databaseURL)
	defer m.Close()

	err := m.Steps(-1)

	if err != nil && err.Error() == "file does not exist" {
		logger.Info("Nothing to rollback")
	} else if err != nil {
		panic(err)
	}
}

func force(databaseURL string, version string) {
	ver, err := strconv.Atoi(version)
	if err != nil {
		msg := fmt.Sprintf("Unable to parse version: %v to integer", version)
		panic(msg)
	}

	m := initMigration(databaseURL)
	defer m.Close()

	err = m.Force(ver)

	if err != nil {
		panic(err)
	}
}

func initMigration(databaseURL string) *migrate.Migrate {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// FIXME this can be iffy write some tests
	dbPath := strings.Split(databaseURL, "/")
	dbName := strings.Split(dbPath[len(dbPath)-1], "?")[0]
	logger.Info("Connecting to: %v", dbName)

	// "file:///home/cr0xd/main/my_project/backend/db/migrations/"
	path := fmt.Sprintf("file://%v/%v", dir, MigrationDir)
	m, err := migrate.NewWithDatabaseInstance(path, dbName, driver)

	if err != nil {
		panic(err)
	}

	return m
}

func includes(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
