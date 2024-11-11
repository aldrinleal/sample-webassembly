REGISTRY=235368163414.dkr.ecr.us-west-2.amazonaws.com
IMAGE=$(REGISTRY)/sample-go-wasi:latest

.PHONY: server run-local run-podman push

all: server

server:
	GOOS=wasip1 GOARCH=wasm go build -o main.wasm ./cmd/server/
	buildah build --platform=wasi/wasm -t $(IMAGE) .

run-local: server
	wasmedge ./main.wasm

run-podman:
	podman run --net=host --platform wasi/wasm --rm 235368163414.dkr.ecr.us-west-2.amazonaws.com/sample-go-wasi:latest

push:
	aws ecr get-login-password --region us-west-2 | podman login -u AWS --password-stdin $(REGISTRY)
	podman push $(IMAGE)

