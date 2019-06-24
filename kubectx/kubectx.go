package kubectx

import (
	"os"
	"os/exec"
)

func Install() (bool, error) {
	kubectx := exec.Command(
		"/bin/bash",
		"-c",
		"brew install kubectx",
	)

	kubectx.Stdout = os.Stdout
	kubectx.Stderr = os.Stderr
	err := kubectx.Run()

	if err != nil {
		return false, err
	}

	return true, nil
}

func Fzf() (bool, error) {
	fzf := exec.Command(
		"/bin/bash",
		"-c",
		"brew install fzf",
	)

	fzf.Stdout = os.Stdout
	fzf.Stderr = os.Stderr
	err := fzf.Run()

	if err != nil {
		return false, err
	}

	return true, nil
}
