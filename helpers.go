package celeritas

import "os"

func (c *Celeritas) CreateDirIfNotExists(path string) error {
	const mode = 0755

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)

		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Celeritas) CreateFileIfNotExist(path string) error {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)

		if err != nil {
			return err
		}

		defer file.Close() 
	}

	return nil
}
