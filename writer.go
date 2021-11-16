package cheapcash

import "os"

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
