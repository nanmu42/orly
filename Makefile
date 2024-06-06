VERSION := $(shell git describe --tags --dirty --always)
BUILD := $(shell date +%FT%T%z)

.PHONY: config clean dir all frontend rly fonts covers

all: clean config frontend fonts covers rly

dir:
	mkdir -p bin && \
	mkdir -p bin/web && \
	mkdir -p bin/fonts && \
	mkdir -p bin/cover-images

assets:
	mkdir -p assets && \
	cd assets && \
	wget -nc https://github.com/nanmu42/orly/releases/download/1.5.0-beta/cover-images.tar.xz && \
	wget -nc https://github.com/nanmu42/orly/releases/download/1.1.0-beta/fonts.tar.xz

clean:
	rm -rf bin

config: dir
	cd cmd/rly && \
	go run genconfig.go config.go && \
	cp config_example.toml $(PWD)/bin

frontend: dir
	cd frontend && \
	pnpm install && pnpm build && \
	cp -r dist/* $(PWD)/bin/web

rly: rly.bin

fonts: dir assets
	cd assets && \
	tar -xf fonts.tar.xz --skip-old-files -C $(PWD)/bin/fonts

covers: dir assets
	cd assets && \
	tar -xf cover-images.tar.xz --skip-old-files -C $(PWD)/bin/cover-images

%.bin: dir
	cd cmd/$* && \
	CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildDate=$(BUILD)" && \
	cp $* $(PWD)/bin/$*
