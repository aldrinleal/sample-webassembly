# Sample Go / WASI Container

## Dependencies

Tested on Fedora Workstation, which needs crun-wasmedge / buildah and podman set up among others

## Usage

```
make server
make run-local # curl http://localhost:8000 to test
make run-podman # same, but needs a podman ps and podman-kill later
make push # oops you can't do that unless you change :)
```

# TODO

https://github.com/sirupsen/logrus/issues/1394

Figure out how to run in Kubernetes (k3s/containerd)
Automate build with Github Actions

