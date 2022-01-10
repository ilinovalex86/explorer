package explorer

import (
	"io/ioutil"
)

// ReadFileFull - читает файл целиком, возвращает строку
func ReadFileFull(fileName string) []byte {
	b, err := ioutil.ReadFile(fileName)
	check(err)
	return b
}
