package main

import (
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
			brew.Install()
		case "no":
			log.Fatal("We cannot proceed without Homebrew. You are on your own.")
		default:
			log.Fatal("Unknown response. Response must be 'yes' or 'no'.")
		}
	}

	if !(checks.App("kubectl")) {
		log.Println("Installing kubectl", kubectlVersion)
		kubectl.Install(kubectlVersion)
	} else {
		currVersion := kubectl.Version()
		if checks.Version(currVersion, kubectlVersion) {
			log.Println("kubectl version", currVersion, "is supported")
		} else {
			kubectl.Install(kubectlVersion)
			if checks.Version(currVersion, kubectlVersion) {
				log.Println("kubectl is now at a supported version:", kubectlVersion)
			} else {
				log.Println("Failed to bring kubectl to a supported version")
			}
		}
	}

	if !(checks.App("helm")) {
		log.Println("Installing helm", helmVersion)
		helm.Install(helmVersion)
	} else {
		currVersion := helm.Version()
		if checks.Version(currVersion, helmVersion) {
			log.Println("helm version", currVersion, "is supported")
		} else {
			helm.Install(helmVersion)
			if checks.Version(currVersion, helmVersion) {
				log.Println("helm is now at a supported version:", helmVersion)
			} else {
				log.Println("Failed to bring helm to a supported version")
			}
		}
	}

	if !(checks.App("kubectx") || checks.App("fzf")) {
		log.Println("Installing kubectx and fzf")
		kubectx.Install()
	}
}
