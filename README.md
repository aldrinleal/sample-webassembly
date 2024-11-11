# Sample Go / WASI Contaienr

## Dependencies

Tested on Fedora Workstation, which needs crun-wasmedge and podman set up among others

## Usage

```
make server
make run-local # curl http://localhost:8000 to test
make run-podman # same, but needs a podman ps and podman-kill later
make push # oops you can't do that unless you change :)
```

# TODO

Figure out how to run in Kubernetes
