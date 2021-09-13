

WORKDIR ?= cmd/releaser

GOARCH ?= amd64
GOOS ?= darwin

up:
	source ci.sh && cd $(WORKDIR) && go run . upload . cmd/

build:
	cd $(WORKDIR) && CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o gitlab-releaser .

docker:
	docker buildx build --push --platform=linux/amd64,linux/arm64 -t hub-dev.rockontrol.com/rk-infrav2/gitlab-releaser:v1.0.7 .	

up.server:
	cd cmd/server && go run .