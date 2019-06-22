package kubectl

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/hashicorp/go-getter"
)

func Install(version string) {
	kubectlURL := fmt.Sprintf(
		"https://storage.googleapis.com/kubernetes-release/release/%s/bin/darwin/amd4/kubectl",
		version,
	)

	err := getter.GetAny("/tmp", kubectlURL)

	if err != nil {
		log.Fatalf("kubectl download failed with:\n%s\n", err)
	}
	moveKubectl()
}

func moveKubectl() {
	moveCMD := fmt.Sprintf(
		"chmod 755 /tmp/kubectl && mv /tmp/kubectl /usr/local/bin",
	)

	install := exec.Command(
		"/bin/sh",
		"-c",
		moveCMD,
	)

	err := install.Run()

	if err != nil {
		log.Fatal("Failed to move kubectl to /usr/local/bin")
	}
}

func Version() string {
	currVersion := exec.Command(
		"/bin/bash",
		"-c",
		"kubectl version --client=true --short | awk '{print $3}'",
	)
	buf := new(bytes.Buffer)
	currVersion.Stdout = buf
	err := currVersion.Run()
	if err != nil {
		log.Fatalf("Failed to get current kubectl version:\n%s\n", err)
	}
	return strings.TrimSpace(buf.String())
}
