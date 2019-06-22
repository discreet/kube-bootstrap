package checks

import (
	"os"
	"os/exec"
)

func Env(envVar string) bool {
	_, set := os.LookupEnv(envVar)
	return set
}

func App(app string) bool {
	_, installed := exec.LookPath(app)
	return installed == nil
}

func Version(currVersion string, newVersion string) bool {
	return currVersion == newVersion
}
