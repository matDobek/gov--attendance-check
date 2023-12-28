package utils

import (
	"errors"
	"os"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

//=======================================================
// Errors
//=======================================================

var (
	ErrRootNotFound = errors.New("Project root not found")
)

//=======================================================
// Functions
//=======================================================

//
//
//

func LookupMod() ([]string, error) {
	_, file, _, _ := runtime.Caller(0)

	dirs, err := lookup(strings.Split(file, "/"), "go.mod")
	if err != nil {
		return []string{}, err
	}
	return dirs, nil
}

func lookup(dirs []string, name string) ([]string, error) {
	if len(dirs) == 0 {
		return dirs, ErrRootNotFound
	}

	path := strings.Join(dirs, "/")
	children, err := os.ReadDir(path)
	if err != nil {
		children = []os.DirEntry{}
	}

	for _, file := range children {
		if file.IsDir() || file.Name() != name {
			continue
		}

		return dirs, nil
	}

	return lookup(dirs[:len(dirs)-1], name)
}

//
//
//

func PrimaryDatabaseURL() string {
  path, err := LookupMod()
  if err != nil {
    panic(err)
  }

  path = append(path, ".env")
  spath := strings.Join(path, "/")

  err = godotenv.Load(spath)
  if err != nil {
    panic(err)
  }

  url := os.Getenv("DB__MAIN__URL")
  if strings.Trim(url, " \t\n") == "" {
    panic("primary database url is not set")
  }

  return url
}

//
//
//

func TestPrimaryDatabaseURL() string {
  path, err := LookupMod()
  if err != nil {
    panic(err)
  }

  path = append(path, ".env")
  spath := strings.Join(path, "/")

  err = godotenv.Load(spath)
  if err != nil {
    panic(err)
  }

  url := os.Getenv("DB__MAIN__URL_TEST")
  if strings.Trim(url, " \t\n") == "" {
    panic("testing database url is not set")
  }

  return url
}
