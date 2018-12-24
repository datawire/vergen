GOBUILD := go build -a -tags netgo -ldflags '-s -w "-extldflags=-static"'

all: clean build

build:
	GOOS=linux  GOARCH=amd64 $(GOBUILD) -o bin/vergen-linux-amd64
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/vergen-darwin-amd64
	ln -sf $(PWD)/bin/vergen-$(shell go env GOOS)-$(shell go env GOARCH) bin/vergen

clean:
	rm -rf bin
