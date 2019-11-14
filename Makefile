GO = go
GOBUILD = $(GO) build -mod=readonly

PIXIU = pixiu
GITHASH = $(shell git rev-parse --short=8 HEAD)

.PHONY: build
build:
	@$(GOBUILD) -ldflags "-X main.githash=$(GITHASH)" -o $(PIXIU) ./cmd

.PHONY: clean
clean:
	@$(RM) $(PIXIU)