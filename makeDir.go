package explorer

import "os"

// MakeDir - Создает папку
func MakeDir(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, 0755)
			check(err)
		} else {
			check(err)
		}
	}
}
