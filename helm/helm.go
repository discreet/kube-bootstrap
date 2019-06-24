package helm

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/cavaliercoder/grab"
)

func Install(version string) {
	helmURL := fmt.Sprintf(
		"https://get.helm.sh/helm-%s-darwin-amd64.tar.gz",
		version,
	)

	resp, err := grab.Get("/tmp", helmURL)

	if err != nil {
		log.Fatalf("helm download failed with:\n%s\n", err)
	}
}

func moveHelm() {
	currLocation := "/tmp/darwin-amd64/helm"
	newLocation := "/usr/local/bin/helm"

	err := os.Rename(currLocation, newLocation)

	if err != nil {
		log.Fatal("Failed to move", currLocation, "to", newLocation)
	}
	setPerms()
}

func setPerms() {
	helm := "/usr/local/bin/helm"

	err := os.Chmod(helm, 0755)

	if err != nil {
		log.Fatal("Failed to make", helm, "executable")
	}
	cleanupHelm()
}

func cleanupHelm() {
	helmDir := "/tmp/darwin-amd64"

	err := os.RemoveAll(helmDir)

	if err != nil {
		log.Fatal("Failed to cleanup helm artifacts.")
	}
}

func Version() string {
	currVersion := exec.Command(
		"/bin/bash",
		"-c",
		"helm version --client --short | awk '{print $2}'",
	)

	buf := new(bytes.Buffer)
	currVersion.Stdout = buf
	err := currVersion.Run()

	if err != nil {
		log.Fatalf("Failed to get current helm version:\n%s\n", err)
	}

	r := regexp.MustCompile(`^v\d{1,2}.\d{1,2}.{1,2}`)
	return r.FindString(strings.TrimSpace(buf.String()))
}
