package explorer

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// DelDir - Удаляет папку
func DelDir(path string) error {
	paths, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range paths {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			DelDir(filePath)
		} else {
			err = os.Remove(filePath)
			if err != nil {
				return err
			}
		}
	}
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
