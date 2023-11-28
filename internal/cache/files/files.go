package files

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/matDobek/gov--attendance-check/internal/cache"
)

//===================
// Defaults
//===================

var (
	defaultDirectory = []string{"tmp", "cache"}
)

//===================
// Errors
//===================

var (
	ErrRootNotFound = errors.New("Project root not found")
)

//===================
// Types
//===================

type FileCache struct {
	dir string
}

// verify interface implementation
var _ cache.Cache = (*FileCache)(nil)

//===================
// Public Functions
//===================

//
//
//

func ConfigurableNew(path string) (FileCache, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return FileCache{}, err
	}

	return FileCache{dir: path}, nil
}

//
//
//

func New() (FileCache, error) {
	projectRoot, err := lookupMod()
	if err != nil {
		return FileCache{}, err
	}

	path := strings.Join(append(projectRoot, defaultDirectory...), "/")
	return ConfigurableNew(path)
}

func lookupMod() ([]string, error) {
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

//===================
// Interface Functions
//===================

//
//
//

func (fc *FileCache) Put(id string, content string) error {
	path := pathTo(*fc, id)

	err := os.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return cache.ErrCouldNotWrite
	}

	return nil
}

func pathTo(fc FileCache, id string) string {
	h := sha1.New()
	h.Write([]byte(id))
	fid := fmt.Sprintf("%x", h.Sum(nil))

	return fc.dir + "/" + fid
}

//
//
//

func (fc *FileCache) Get(id string) (string, error) {
	path := pathTo(*fc, id)

	val, err := os.ReadFile(path)
	if err != nil {
		return "", cache.ErrCouldNotRead
	}

	return string(val), nil
}

//
//
//

func (fc *FileCache) Clear() error {
	// As we don't want to remove some files by accident
	// Let's make sure our path contains:
	//   * tmp dir
	//   * cache dir
	contains_dir_tmp := false

	for _, dir := range strings.Split(fc.dir, "/") {
		switch dir {
		case "tmp":
			contains_dir_tmp = true
		}
	}

	if contains_dir_tmp {
		files, err := os.ReadDir(fc.dir)
		if err != nil {
			return err
		}

		for _, file := range files {
			pathToRemove := strings.Join([]string{fc.dir, file.Name()}, "/")
			os.Remove(pathToRemove)
		}
	}

	return nil
}
