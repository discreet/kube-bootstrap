package helm

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/hashicorp/go-getter"
)

func Install(version string) {
	helmURL := fmt.Sprintf(
		"https://get.helm.sh/helm-%s-darwin-amd64.tar.gz",
		version,
	)

	err := getter.GetAny("/tmp", helmURL)

	if err != nil {
		log.Fatalf("helm download failed with:\n%s\n", err)
	}
	moveHelm()
}

func moveHelm() {
	moveCMD := fmt.Sprintf(
		"mv /tmp/darwin-amd64/helm /usr/local/bin",
	)

	install := exec.Command(
		"/bin/sh",
		"-c",
		moveCMD,
	)

	err := install.Run()

	if err != nil {
		log.Fatal("Failed to move helm to /usr/local/bin")
	}
	cleanupHelm()
}

func cleanupHelm() {
	cleanupCMD := fmt.Sprintf(
		"rm -rf /tmp/darwin-amd64",
	)

	cleanup := exec.Command(
		"/bin/bash",
		"-c",
		cleanupCMD,
	)

	err := cleanup.Run()

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
