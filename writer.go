package cheapcash

import "os"

// Write a key with a value.
// If the key already exists in the first place, it will
// delete the existing key and replace it with the new
// value.
//
//      c := cheapcash.Default()
//      err := c.Write("users", []byte("Someone\n"))
//      if err != nil {
//        // handle your error
//      }
//
func (c *Cache) Write(key string, value []byte) error {
	err := checkDir(sanitizePath(c.Path))
	if err != nil {
		return err
	}

	check, err := c.Exists(c.Path + key)
	if err != nil {
		return err
	}

	if check {
		err = c.Delete(key)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(sanitizePath(c.Path + key))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(value)
	if err != nil {
		return err
	}

	return nil
}
