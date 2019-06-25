package main

import (
	"fmt"
	"log"

	"github.com/discreet/kube-bootstrap/brew"
	"github.com/discreet/kube-bootstrap/checks"
	"github.com/discreet/kube-bootstrap/helm"
	"github.com/discreet/kube-bootstrap/kubectl"
	"github.com/discreet/kube-bootstrap/kubectx"
)

func main() {
	kubectlVersion := "v1.11.6"
	helmVersion := "v2.10.0"

	if !(checks.Env("HTTP_PROXY") && checks.Env("http_proxy")) {
		log.Fatal("Remember to set your proxy")
	}

	if !checks.App("brew") {
		resp := brew.Prompt()
		switch resp {
		case "yes":
			if _, err := brew.Install(); err != nil {
				log.Fatal(err)
			}
			log.Println("Homebrew has been installed")
		case "no":
			log.Fatal("We cannot proceed without Homebrew. You are on your own.")
		default:
			log.Fatal("Unknown response. Response must be 'yes' or 'no'.")
		}
	}

	if !(checks.App("kubectl")) {
		log.Println("Installing kubectl", kubectlVersion)
		installer := kubectl.NewInstaller()
		installer.URLTemplate = fmt.Sprintf(
			"https://storage.googleapis.com/kubernetes-release/release/%s/bin/darwin/amd4/kubectl",
			kubectlVersion,
		)
		installer.FilePath = "/usr/local/bin"

		if _, err := installer.Install(); err != nil {
			log.Fatal(err)
		}
		log.Println("kubectl", kubectlVersion, "Installed")
	} else {
		currVersion, err := kubectl.Version()
		if err != nil {
			log.Fatal(err)
		}
		if checks.Version(currVersion, kubectlVersion) {
			log.Println("kubectl version", currVersion, "is supported")
		} else {
			log.Println("Installing kubectl", kubectlVersion)
			installer := kubectl.NewInstaller()
			installer.URLTemplate = fmt.Sprintf(
				"https://storage.googleapis.com/kubernetes-release/release/%s/bin/darwin/amd4/kubectl",
				kubectlVersion,
			)
			installer.FilePath = "/usr/local/bin"

			if _, err := installer.Install(); err != nil {
				log.Fatal(err)
			}
			if checks.Version(currVersion, kubectlVersion) {
				log.Println("kubectl is now at a supported version:", kubectlVersion)
			} else {
				log.Println("Failed to bring kubectl to a supported version")
			}
		}
	}

	if !(checks.App("helm")) {
		log.Println("Installing helm", helmVersion)
		installer := helm.NewInstaller()
		installer.URLTemplate = fmt.Sprintf(
			"https://get.helm.sh/helm-%s-darwin-amd64.tar.gz",
			helmVersion,
		)
		installer.FilePath = "/usr/local/bin"

		if _, err := installer.Install(); err != nil {
			log.Fatal(err)
		}
		log.Println("helm", helmVersion, "installed")
	} else {
		currVersion, err := helm.Version()
		if err != nil {
			log.Fatal(err)
		}
		if checks.Version(currVersion, helmVersion) {
			log.Println("helm version", currVersion, "is supported")
		} else {
			installer := helm.NewInstaller()
			installer.URLTemplate = fmt.Sprintf(
				"https://get.helm.sh/helm-%s-darwin-amd64.tar.gz",
				helmVersion,
			)
			installer.FilePath = "/usr/local/bin"

			if _, err := installer.Install(); err != nil {
				log.Fatal(err)
			}
			if checks.Version(currVersion, helmVersion) {
				log.Println("helm is now at a supported version:", helmVersion)
			} else {
				log.Println("Failed to bring helm to a supported version")
			}
		}
	}

	if !checks.App("kubectx") {
		log.Println("Installing kubectx")
		_, err := kubectx.Install()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("kubectx installed")
	}

	if !checks.App("fzf") {
		log.Println("Installing fzf")
		_, err := kubectx.Fzf()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("fzf installed")
	}
}
