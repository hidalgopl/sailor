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
