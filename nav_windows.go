package explorer

import "strings"

// NavFunc - Разбивает путь ОС Windows
func NavFunc(path string) []LinkAndName {
	if path[len(path)-1] == '\\' {
		path = path[:len(path)-1]
	}
	var links []LinkAndName
	tempPath := ""
	if path == BasePath {
		return []LinkAndName{{Link: BasePath, Name: "Компьютер"}}
	}
	if len(path) > len(BasePath) && path[:len(BasePath)] == BasePath {
		links = append(links, LinkAndName{BasePath, "Компьютер"})
		path = path[len(BasePath)+1:]
		tempPath = BasePath
	}
	paths := strings.Split(path, "\\")
	for i := range paths {
		if tempPath == "" {
			links = append(links, LinkAndName{paths[i], paths[i]})
			tempPath = paths[i]
		} else if tempPath == BasePath {
			if translateFoldersWindows[paths[i]] == "" {
				links = append(links, LinkAndName{tempPath + "\\" + paths[i], paths[i]})
			} else {
				links = append(links, LinkAndName{tempPath + "\\" + paths[i], translateFoldersWindows[paths[i]]})
			}
			tempPath = tempPath + "\\" + paths[i]
		} else {
			links = append(links, LinkAndName{tempPath + "\\" + paths[i], paths[i]})
			tempPath = tempPath + "\\" + paths[i]
		}
	}
	return links
}
