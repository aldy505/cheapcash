package cheapcash

import (
	"io/ioutil"
	"os"
)

func (c *Cache) Read(key string) ([]byte, error) {
	check, err := c.Exists(c.Path + key)
	if err != nil {
		return []byte{}, err
	}

	if !check {
		return []byte{}, ErrNotExists
	}

	file, err := os.Open(sanitizePath(c.Path + key))
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
