package celeritas

import "fmt"

const Version = "1.0.0"

type Celeritas struct {
	AppName string
	Debug   bool
	Version string
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"handlers", "migrations", "views", "models", "public", "tmp", "logs", "middleware",
		},
	}

	err := c.Init(pathConfig)

	if err != nil {
		return err
	}

	err = c.checkDotEnv(rootPath)
	if err != nil {
		return err 
	}

	// read .env

	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath

	for _, path := range p.folderNames {
		// Create folder if not exists
		err := c.CreateDirIfNotExists(fmt.Sprintf("%s/%s", root, path))

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Celeritas) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err 
	}
	return nil
}
