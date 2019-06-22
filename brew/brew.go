package brew

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Homebrew is required. Would you like to install Homebrew? [yes/no] ")
	resp, _ := reader.ReadString('\n')
	resp = strings.Replace(resp, "\n", "", -1)
	return resp
}

func Install() {
	log.Println("Installing Homebrew")
	install := exec.Command(
		"/bin/sh",
		"-c",
		"/usr/bin/ruby -e \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)\"",
	)
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	err := install.Run()
	if err != nil {
		log.Fatalf("Homebrew install failed with:\n%s\n", err)
	}
}
