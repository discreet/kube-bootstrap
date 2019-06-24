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

func Install(version string) (bool, error) {
	kubectlURL := fmt.Sprintf(
		"https://storage.googleapis.com/kubernetes-release/release/%s/bin/darwin/amd4/kubectl",
		version,
	)

	resp, err := http.DefaultClient.Get(kubectlURL)

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile("/usr/local/bin/kubectl", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return false, err
	}

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
