package explorer

import (
	"golang.org/x/text/encoding/charmap"
	"log"
	"os/exec"
	"strings"
)

var System = "windows"

var Sep = "\\"

var codePage = findCodePage()

var User = execCommand("echo %username%")

var BasePath = execCommand("echo %homedrive%%homepath%")

func findCodePage() string {
	cmd := exec.Command("cmd", "/C", "chcp")
	buffer, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(buffer[:])
	return strings.TrimSpace(s[strings.Index(s, ":")+1:])
}

//execCommand - Выполнение команды в консоли ОС Windows
func execCommand(command string) string {
	cmd := exec.Command("cmd", "/C", command)
	buffer, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	var res []byte
	switch codePage {
	case "65001":
		res = buffer[:]
	case "866":
		decoder := charmap.CodePage866.NewDecoder()
		res, err = decoder.Bytes(buffer)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Кодировка не распознана.")
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
