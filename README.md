# kube-bootstrap

A little tool to help you get set up for interacting with Kubernetes.

### homebrew

The only dependency for using `kube-bootstrap` is `homebrew`, which is used to
install one of the helpful packages for interacting with a Kubernetes cluster.
If `homebrew` is not already installed, you will be prompted to have `homebrew`
installed for you. You can learn more about `homebrew` at https://brew.sh/.

### kubectl

Part of this set up is installing the correct version of `kubectl` for
interacting with Kuhernetes. If `kubectl` is not already installed, a supported
version will be installed for you. If you already have a version of `kubectl`
installed, then the installed version will be validated against a supported
version. If the installed version is not supported, `kubectl` will be changed to
a supported version.

### helm

Part of this setup is installing the correct version of `helm` for interacting
with Kubernetes. If `helm` is not already installed, a supported version will be
installed for you. If you already have a version of `helm` installed, then the
installed version will be validated against a supported version. If the
installed version is not supported, `helm` will be changed to a supported
version.

### kubectx

Part of this setup is to help make managing contexts simpler. `kubectx` is a
tool that makes is easy to manage multiple Kubernetes contexts with a single
command. If `kubectx` is not alredy installed it will be installed via
`homebrew` along with `fzf` for fuzzy matching on cluster contexts. You can
learn more about `kubectx` from their [GitHub page](https://github.com/ahmetb/kubectx).

### TODO

- [ ] Generate `kube.config`
- [ ] Unit tests
- [ ] Configure `helm
