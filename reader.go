package cheapcash

import (
	"io/ioutil"
	"os"
)

// Read the value of a given key.
// Will return an error of ErrNotExists if the given
// key does not exists.
//
//      c := cheapcash.Default()
//      res, err := c.Read("users")
//      if err != nil {
//        // handle your error here!
//      }
//      log.Println(string(res))
//
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
