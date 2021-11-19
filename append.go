package cheapcash

import "os"

// Directly append a value into an existing key.
// If a key doesn't exists, it will return an error
// with a type of ErrNotExists.
//
//      c := cheapcash.Default()
//      err := c.Append("users", []byte("Someone\n"))
//      if err != nil {
//        if errors.Is(err, cheapcash.ErrNotExists) {
//          // Handle if file does not exists!
//        }
//        // Handle any other errors
//      }
//
func (c *Cache) Append(key string, value []byte) error {
	check, err := c.Exists(c.Path + key)
	if err != nil {
		return err
	}

	if !check {
		return ErrNotExists
	}

	file, err := os.OpenFile(sanitizePath(c.Path+key), os.O_APPEND|os.O_WRONLY, 0644)
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
