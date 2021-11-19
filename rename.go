package cheapcash

import "os"

// Rename a key. It's that simple.
// The contents of the cache stays the same, but the key
// is renamed.
//
// It will return 2 different error in case of:
//
// 1. If the old key (first argument parameter) doesn't exists,
// it will return an error of ErrNotExists
//
// 2. If the new key (second argument parameter) already exists,
// it will return an error of ErrExists
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
