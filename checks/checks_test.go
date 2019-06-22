package checks_test

import (
	"os"
	"testing"

	"github.com/discreet/kube-bootstrap/checks"
)

func TestVersion(t *testing.T) {
	resp := checks.Version("foo", "foo")

	if resp != true {
		t.Error("The values should have matched and returned true")
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
