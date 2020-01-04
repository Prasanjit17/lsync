package types

import "os"

type (
	FileWalk chan string
)

func (f FileWalk) Walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		f <- path
	}

	return nil
}
