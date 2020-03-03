BINDIR := $(CURDIR)/bin
LDFLAGS := "-extldflags '-static'"

build:
	GOBIN=$(BINDIR) go install -ldflags $(LDFLAGS) ./...
	echo "Build complete. Use ./bin/sailor to run it"
.PHONY: build

clean:
	go clean ./...
	rm -rf $(BINDIR)
	rm -f coverage.*
.PHONY: clean

fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
.PHONY: fmt

test:
	go test -v -race -cover -coverprofile=coverage.out -run . ./...
.PHONY: test

coverage: test
	go tool cover -func=coverage.out
.PHONY: coverage


container: build
	docker build  -t secureapi/sailor:v0.0.2 .
.PHONY: container
