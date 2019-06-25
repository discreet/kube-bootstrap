package helm

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Installer struct {
	URLTemplate string
	FilePath    string
}

func NewInstaller() *Installer {
	return &Installer{}
}

func (i *Installer) Install() (bool, error) {
	helmURL := i.URLTemplate

	resp, err := http.DefaultClient.Get(helmURL)

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	ret, err := extractHelm(resp.Body, i.FilePath)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func extractHelm(archive io.Reader, path string) (bool, error) {
	archive, err := gzip.NewReader(archive)
	if err != nil {
		return false, err
	}

	r := regexp.MustCompile(`^darwin-amd64\/helm$`)
	tr := tar.NewReader(archive)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return false, err
		}

		if r.MatchString(hdr.Name) {
			continue
		}

		helmPath := fmt.Sprintf("%s/helm", path)

		f, err := os.OpenFile(helmPath, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return false, err
		}

		if _, err := io.Copy(f, tr); err != nil {
			return false, err
		}
		break
	}
	return true, nil
}

func Version() (string, error) {
	currVersion := exec.Command(
		"/bin/bash",
		"-c",
		"helm version --client --short | awk '{print $2}'",
	)

	buf := new(bytes.Buffer)
	currVersion.Stdout = buf
	err := currVersion.Run()
	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(`^v\d{1,2}.\d{1,2}.{1,2}`)
	return r.FindString(strings.TrimSpace(buf.String())), nil
}
