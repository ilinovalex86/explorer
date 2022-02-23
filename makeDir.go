package explorer

import "os"

// MakeDir - Создает папку
func MakeDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, 0755)
			if err != nil {
				return err
			}
		}
		return err
	}
	return nil
}
