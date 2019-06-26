package helm_test

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/discreet/kube-bootstrap/helm"
)

func TestInstall(t *testing.T) {
	s := "echo 0"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hdr := tar.Header{
			Name: "helm-test",
			Mode: 0644,
			Size: int64(len(s)),
		}

		gw := gzip.NewWriter(w)
		defer gw.Close()

		tw := tar.NewWriter(gw)
		defer tw.Close()

		if err := tw.WriteHeader(&hdr); err != nil {
			t.Error(err)
			return
		}

		_, err := fmt.Fprint(tw, s)
		if err != nil {
			t.Errorf("unexpected error copying zip file in server: %v", err)
		}
	}))

	installer := helm.NewInstaller()
	installer.DownloadURL = ts.URL
	installer.DownloadPath = "/tmp/helm-test"
	installer.HelmRegex = `^echo\w0$`

	_, err := os.Stat(installer.DownloadPath)
	if !os.IsNotExist(err) {
		os.Remove(installer.DownloadPath)
	}

	if _, err := installer.Install(); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(installer.DownloadPath)

	info, err := os.Stat(installer.DownloadPath)

	if err != nil {
		t.Fatal(err)
	}

	if info.Mode() != 0755 {
		t.Errorf("Want permissions 0755, got %v", info.Mode())
	}

	b, err := ioutil.ReadFile(installer.DownloadPath)
	if err != nil {
		t.Fatal(err)
	}

	content := string(b)

	if strings.Contains(content, s) != true {
		t.Error("Content mismatch for file")
	}
}
