

WORKDIR ?= cmd/releaser

up:
	source ci.sh && cd $(WORKDIR) && go run . upload . cmd/