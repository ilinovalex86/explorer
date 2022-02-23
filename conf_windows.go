package explorer

import (
	"golang.org/x/text/encoding/charmap"
	"log"
	"os/exec"
	"strings"
)

var System = "windows"

var Sep = "/"

var User = execCommand("echo %username%")

var BasePath = execCommand("echo %homedrive%%homepath%")

//execCommand - Выполнение команды в консоли ОС Windows
func execCommand(command string) string {
	cmd := exec.Command("cmd", "/C", command)
	buffer, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	decoder := charmap.CodePage866.NewDecoder()
	res, err := decoder.Bytes(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(res))
}

var translateFoldersWindows = map[string]string{
	"Videos":    "Видео",
	"Documents": "Документы",
	"Downloads": "Загрузки",
	"Favorites": "Избранное",
	"Pictures":  "Изображения",
	"Contacts":  "Контакты",
	"Music":     "Музыка",
	"Desktop":   "Рабочий стол",
}
