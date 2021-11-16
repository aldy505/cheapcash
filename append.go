package cheapcash

import "os"

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
