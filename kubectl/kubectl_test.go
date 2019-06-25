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
	installer.InstallPath = "/tmp"

	_, err := os.Stat("/tmp/kubectl")
	if !os.IsNotExist(err) {
		os.Remove("/tmp/kubectl")
	}

	if _, err := installer.Install(); err != nil {
		t.Error(err)
	}
	defer os.Remove("/tmp/kubectl")

	info, err := os.Stat("/tmp/kubectl")

	if err != nil {
		t.Fatal(err)
	}

	if info.Mode() != 0755 {
		t.Errorf("want permissions 0755, got %v", info.Mode())
	}

	b, err := ioutil.ReadFile("/tmp/kubectl")
	if err != nil {
		t.Fatal(err)
	}

	content := string(b)

	if strings.Contains(content, "echo 0") != true {
		t.Error("Content mismatch for file")
	}
}
