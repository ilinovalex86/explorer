package explorer

import (
	"log"
	"os/exec"
	"strings"
)

var System = "linux"

var Sep = "/"

var User = execCommand("whoami")

var BasePath = getBasePath(User)

func getBasePath(user string) string {
	if user == "root" {
		log.Fatal("don't run with root")
	}
	return "/home/" + user
}

func execCommand(command string) string {
	bash := exec.Command(command)
	buffer, err := bash.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(buffer))
}
