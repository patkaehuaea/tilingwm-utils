GODOC_PORT=:6060
BINARY_NAME=twmu
BINARY_OUT=$(GOPATH)/bin/$(BINARY_NAME)
BINARY_386=$(BINARY_OUT)-linux-386
BINARY_AMD64=$(BINARY_OUT)-linux-amd64

all: fmt install

build:
	GOPATH=$(GOPATH) go build -o $(BINARY_OUT)

clean:
	GOPATH=$(GOPATH) go clean
	rm -f $(BINARY_OUT)
	rm -f $(BINARY_386)
	rm -f $(BINARY_AMD64)

doc:
	GOPATH=$(GOPATH) godoc -v --http=$(GODOC_PORT) --index=true

fmt:
	GOPATH=$(GOPATH) go fmt $(PACKAGES)

install:
	GOPATH=$(GOPATH) go build -o $(BINARY_OUT)

release:
	GOPATH=$(GOPATH) GOOS=linux GOARCH=386 go build -o $(BINARY_386)
	GOPATH=$(GOPATH) GOOS=linux GOARCH=amd64 go build -o $(BINARY_AMD64)

test:
	GOPATH=$(GOPATH) go test $(TEST_PACKAGES)
