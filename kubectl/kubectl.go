package kubectl

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Installer struct {
	DownloadURL string
	InstallPath string
}

func NewInstaller() *Installer {
	return &Installer{}
}

func (i *Installer) Install() (bool, error) {
	kubectlURL := i.DownloadURL

	resp, err := http.DefaultClient.Get(kubectlURL)

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	kubectlPath := fmt.Sprintf("%s/kubectl", i.InstallPath)

	f, err := os.OpenFile(kubectlPath, os.O_WRONLY|os.O_CREATE, 0755)

	if err != nil {
		return false, err
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		return false, err
	}

	return true, nil
}

func Version() (string, error) {
	currVersion := exec.Command(
		"/bin/bash",
		"-c",
		"kubectl version --client=true --short | awk '{prbool $3}'",
	)
	buf := new(bytes.Buffer)
	currVersion.Stdout = buf
	err := currVersion.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}
