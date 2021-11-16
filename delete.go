package cheapcash

import "os"

func (c *Cache) Delete(key string) error {
	check, err := c.Exists(c.Path + "/" + key)
	if err != nil {
		return err
	}

	if !check {
		return ErrNotExists
	}

	err = os.Remove(sanitizePath(c.Path + "/" + key))
	if err != nil {
		return err
	}

	return nil
}
