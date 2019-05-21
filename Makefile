SHELL=/usr/bin/env bash

all: clean build

# build produces a binary on the local filesystem. the build itself is performed inside of a docker container
build:
	mkdir -p bin; \
	docker build --target builder -t plombardi89/vergen .; \
	cid=$$(docker create plombardi89/vergen); \
	docker cp $$cid:/build/bin/vergen - > bin/vergen.tar; \
	docker rm -v $$cid; \
	tar -xvf bin/vergen.tar -C bin; \
	rm bin/vergen.tar;

build.image: build
	docker build -t plombardi89/vergen:$$(bin/vergen preview --authority=$$USER) .;

clean:
	rm -rf bin
