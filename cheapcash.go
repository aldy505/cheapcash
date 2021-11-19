package cheapcash

import (
	"errors"
	"strings"

	"sync"
)

type Cache struct {
	sync.Mutex
	Path string
}

var ErrNotExists = errors.New("key does not exist")
var ErrInvalidPath = errors.New("path supplied is invalid")
var ErrDiskFull = errors.New("there was no space left on the device")
var ErrExists = errors.New("key already exists")

// Creates default Cheapcash instance which defaults
// the corresponding cache path to /tmp/cheapcash.
//
// The caveat of using this one is this will be most likely
// only compatible with UNIX-like filesystem.
// Windows devices will most likely happen to have
// an error of the invalid path.
//
// This returns a Cheapcash instance, which method
// ('Append', 'Exists', 'Write', 'Read') you can do
// by just specifying the key, without supplying the
// full path of the cached file.
func Default() *Cache {
	return &Cache{
		Path: "/tmp/cheapcash/",
	}
}

// Creates a new Cheapcash instance with the given
// path from the argument provided.
//
// If path is empty (or an empty string), it will panic
// with ErrInvalidPath error.
//
// The path provided might have '/' as the ending, so
// these are valid and will return the same path:
//
//       New("/tmp/box")
//       New("/tmp/box/")
func New(path string) *Cache {
	if path == "" {
		panic(ErrInvalidPath)
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	return &Cache{
		Path: path,
	}
}
