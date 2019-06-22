package kubectx

import (
	"log"
	"os"
	"os/exec"
)

func Install() {
	kubectx := exec.Command(
		"/bin/bash",
		"-c",
		"brew install kubectx fzf",
	)

	kubectx.Stdout = os.Stdout
	kubectx.Stderr = os.Stderr
	err := kubectx.Run()

	if err != nil {
		log.Fatalf("Brew install of kubectx failed with:\n%s\n", err)
	}
}
