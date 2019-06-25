package checks_test

import (
	"os"
	"testing"

	"github.com/discreet/kube-bootstrap/checks"
)

func TestVersion(t *testing.T) {
	tresp := checks.Version("foo", "foo")
	fresp := checks.Version("foo", "bar")

	if tresp != true {
		t.Error("The values should have matched and returned true")
	}

	if fresp != false {
		t.Error("The values should have not matched and returned false")
	}
}

func TestEnv(t *testing.T) {
	os.Setenv("FOO", "bar")

	resp := checks.Env("FOO")

	if resp != true {
		t.Error("Environment variable not checked properly")
	}
}

func TestApp(t *testing.T) {
	resp := checks.App("ls")

	if resp != true {
		t.Error("App check is incorrect. ls ships with the OS")
	}
}
