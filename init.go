package explorer

import (
	cl "github.com/ilinovalex86/consol"
	"log"
	"runtime"
)

type LinkAndName struct {
	Link string
	Name string
}

//System - Тип ОС
var System string

//Sep - Тип слеша ОС
var Sep string

//User - Имя пользователя
var User string

//BasePath - Домашний каталог
var BasePath string

//NavFunc - Функция разбиения пути
var NavFunc func(path string) []LinkAndName

//Только для ОС Windows
//Перевод папок в домашней папке пользователя на русский язык
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

//if err != nil -> log.Fatal(err)
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	System = runtime.GOOS
	switch System {
	case "windows":
		Sep = "\\"
		User = cl.ConsolWindows("echo %username%")
		BasePath = cl.ConsolWindows("echo %homedrive%%homepath%")
		NavFunc = NavWindows
	case "linux":
		Sep = "/"
		User = cl.ConsolLinux("whoami")
		if User == "root" {
			BasePath = "/root"
		} else {
			BasePath = "/home/" + User
		}
		NavFunc = NavLinux
	}
	if !ExistDir(BasePath) {
		log.Fatal("basePath error")
	}
}
