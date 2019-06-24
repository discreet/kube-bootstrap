package kubectl

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/cavaliercoder/grab"
)

func Install(version string) {
	kubectlURL := fmt.Sprintf(
		"https://storage.googleapis.com/kubernetes-release/release/%s/bin/darwin/amd4/kubectl",
		version,
	)

	resp, err := grab.Get("/tmp", kubectlURL)

	if err != nil {
		log.Fatalf("kubectl download failed with:\n%s\n", err)
	}
	setPerms(resp.Filename)
	moveKubectl(resp.Filename)
}

func setPerms(file string) {
	kubectl := file

	err := os.Chmod(kubectl, 0755)

	if err != nil {
		log.Fatal("Failed to make", kubectl, "executable")
	}
}

func moveKubectl(file string) {
	currLocation := file
	newLocation := "/usr/local/bin/kubectl"

	err := os.Rename(currLocation, newLocation)

	if err != nil {
		log.Fatal("Failed to move", currLocation, "to", newLocation)
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
