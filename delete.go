package cheapcash

import "os"

// Delete a key from the cache directory.
// And of course would return an error of ErrNotExists
// if the key doesn't exists.
//
//      c := cheapcash.Default()
//      err := c.Write("users", []byte("Someone\n"))
//      // Handle error here
//      err = c.Delete("users")
//      // Handle error here
//
func (c *Cache) Delete(key string) error {
	check, err := c.Exists(c.Path + key)
	if err != nil {
		return err
	}

	if !check {
		return ErrNotExists
	}

	err = os.Remove(sanitizePath(c.Path + key))
	if err != nil {
		return err
	}

	return nil
}
