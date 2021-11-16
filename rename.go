package cheapcash

import "os"

func (c *Cache) Rename(old, new string) error {
	err := checkDir(sanitizePath(c.Path))
	if err != nil {
		return err
	}

	checkOld, err := c.Exists(c.Path + old)
	if err != nil {
		return err
	}

	checkNew, err := c.Exists(c.Path + new)
	if err != nil {
		return err
	}

	if !checkOld {
		return ErrNotExists
	}

	if checkNew {
		return ErrExists
	}

	err = os.Rename(c.Path+old, c.Path+new)
	if err != nil {
		return err
	}

	return nil
}
