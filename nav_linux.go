package explorer

import "strings"

// NavFunc - Разбивает путь ОС Linux
func NavFunc(path string) []LinkAndName {
	var links []LinkAndName
	tempPath := ""
	if path == "/" {
		return []LinkAndName{{Link: "/", Name: "/"}}
	}
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	links = append(links, LinkAndName{Link: "/", Name: "/"})
	paths := strings.Split(path, "/")
	for i := 1; i < len(paths); i++ {
		links = append(links, LinkAndName{Link: tempPath + "/" + paths[i], Name: paths[i]})
		tempPath = tempPath + "/" + paths[i]
	}
	return links
}
