package cheapcash

import (
	"errors"
	"io/fs"
	"os"
	"strings"
)

// Check whether or not a key exists.
// Returns true if the key exists, false otherwise.
//
// WARNING: You should provide your c.Path value yourself.
//
//      check, err := cache.Exists("something.txt")
//      // will search in ./something.txt
//
//      check, err = cache.Exists(c.Path + "something.txt")
//      // will search relative to c.Path value
func (c *Cache) Exists(key string) (bool, error) {
	file, err := os.Open(sanitizePath(key))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}
	defer file.Close()

	return true, nil
}

// Will validate the existance of a directory.
// If the directory (including its' children) doesn't
// exists, it will create the corresponding directory
// tree from the given directory path.
//
// If the directory already exists, it will return
// a nil value.
func checkDir(path string) error {
	// Remove / from path
	path = strings.TrimSuffix(path, "/")

	dir, err := os.Stat(path)
	if err == nil {
		if dir.IsDir() {
			return nil
		}
	}

	// Create directory with a loop
	separated := strings.Split(path, "/")

	for i := 0; i < len(separated); i++ {
		if separated[i] == "" {
			os.Chdir("/")
			continue
		} else {
			os.Chdir(separated[i-1])
		}

		err = os.Mkdir(separated[i], fs.ModePerm)
		if err != nil {
			if errors.Is(err, os.ErrExist) {
				continue
			}
			return err
		}
	}

	return nil
}
