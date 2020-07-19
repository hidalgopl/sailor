BINDIR := $(CURDIR)/bin
LDFLAGS := "-extldflags '-static'"

set-envs:
	envsubst < secrets/secrets.yaml.template >> secrets/staging.yaml

build-staging: set-envs build

build:
	GOBIN=$(BINDIR) CGO_ENABLED=0 go install -ldflags $(LDFLAGS) ./...
	echo "Build complete. Use ./bin/sailor to run it"
.PHONY: build

set-locals:
	 NATS_URL=nats://locahost:4222 FRONT_URL=http://localhost:3000 API_URL=http://localhost:8000 envsubst < secrets/secrets.yaml.template >> secrets/staging.yaml
.PHONY: set-locals

build-local: set-locals build
.PHONY: build-local

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
