

WORKDIR ?= cmd/releaser

GOARCH ?= amd64
GOOS ?= darwin

up:
	source ci.sh && cd $(WORKDIR) && go run . upload . cmd/

build:
	cd $(WORKDIR) && CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o gitlab-releaser .
	cd $(WORKDIR) && CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o gitlab-server .

up.server:
	cd cmd/server && go run .

