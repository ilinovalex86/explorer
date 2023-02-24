//Package explorer пакет для работы с файлами, папками, путями.
//Работает в ОС Linux, ОС Windows
// var System string тип ОС
// var User string имя пользователя
// var BasePath string домашняя папка
// var Sep string тип слеша ОС
package explorer

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Explorer - Принимает путь и возвращет содержание каталога
func Explorer(path string) (map[string][]LinkAndName, error) {
	data := make(map[string][]LinkAndName)
	if path == "C:" {
		path = "C:\\"
	}
	_, err := os.Stat(path)
	if err != nil {
		return data, err
	}
	paths, _ := ioutil.ReadDir(path)
	for _, file := range paths {
		filePath := filepath.Join(path, file.Name())
		fileName := file.Name()
		if file.IsDir() {
			if path == BasePath && translateFoldersWindows[fileName] == "" {
				continue
			}
			if path == BasePath && translateFoldersWindows[fileName] != "" {
				fileName = translateFoldersWindows[fileName]
			}
			data["dirs"] = append(data["dirs"], LinkAndName{Link: filePath, Name: fileName})
		} else {
			if path == BasePath || fileName[0] == '~' {
				continue
			}
			data["files"] = append(data["files"], LinkAndName{Link: filePath, Name: fileName})
		}
	}
	return data, nil
}
