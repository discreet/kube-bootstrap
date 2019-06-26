package kubectl_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/discreet/kube-bootstrap/kubectl"
)

func TestInstall(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "echo 0")
	}))
	defer ts.Close()

	installer := kubectl.NewInstaller()
	installer.DownloadURL = ts.URL
	installer.DownloadPath = "/tmp/kubectl-test"

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
		t.Errorf("want permissions 0755, got %v", info.Mode())
	}

	b, err := ioutil.ReadFile(installer.DownloadPath)
	if err != nil {
		t.Fatal(err)
	}

	content := string(b)

	if strings.Contains(content, "echo 0") != true {
		t.Error("Content mismatch for file")
	}
}
