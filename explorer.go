//Package explorer пакет для работы с файлами, папками, путями.
//Работает в ОС Linux, ОС Windows
//Определяет тип ОС, имя пользователя, домашний папку,
//тип слеша ОС, функцию навигации пути
//и сохраняет в соответствующие переменные
// var System string
// var User string
// var BasePath string
// var Sep string
// NavFunc - функция навигации пути определенной ОС
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
			if System == "windows" && path == BasePath && translateFoldersWindows[fileName] == "" {
				continue
			}
			if System == "windows" && path == BasePath && translateFoldersWindows[fileName] != "" {
				fileName = translateFoldersWindows[fileName]
			}
			data["dirs"] = append(data["dirs"], LinkAndName{Link: filePath, Name: fileName})
		} else {
			if path == BasePath && System == "windows" {
				continue
			}
			data["files"] = append(data["files"], LinkAndName{Link: filePath, Name: fileName})
		}
	}
	return data, nil
}
