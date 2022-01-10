package explorer

import (
	"bufio"
	"os"
)

// FileByString - читает файл и возвращает массив его строк
func FileByString(fileName string) []string {
	var data []string
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
