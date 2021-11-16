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

func Default() *Cache {
	return &Cache{
		Path: "/tmp/cheapcash/",
	}
}

func New(path string) *Cache {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	return &Cache{
		Path: path,
	}
}
