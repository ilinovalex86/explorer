package explorer

import (
	"io/ioutil"
)

// ReadFileFull - читает файл целиком, возвращает строку
func ReadFileFull(fileName string) ([]byte, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return b, err
}
