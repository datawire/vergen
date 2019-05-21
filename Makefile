SHELL=/usr/bin/env bash

IMG_NAME := datawire/vergen

all: clean build

# build produces a binary on the local filesystem. the build itself is performed inside of a docker container
build:
	mkdir -p bin; \
	docker build --target builder -t $(IMG_NAME) .; \
	cid=$$(docker create $(IMG_NAME)); \
	docker cp $$cid:/build/bin/vergen - > bin/vergen.tar; \
	docker rm -v $$cid; \
	tar -xvf bin/vergen.tar -C bin; \
	rm bin/vergen.tar;

build.image: build
	docker build -t $(IMG_NAME):$$(bin/vergen preview --authority=$$USER) .;

clean:
	rm -rf bin
