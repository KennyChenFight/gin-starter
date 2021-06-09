PATH := ${CURDIR}/bin/cmd:${CURDIR}/bin:$(PATH)
goexe = $(shell go env GOEXE)

.PHONY: build
build: test
	$(shell ls cmd | xargs -I {} go build -o bin/cmd/{} cmd/{}/main.go)

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: codegen
codegen: bin/mockgen$(go_exe)
	go generate ./...

bin/mockgen$(go_exe):
	go build -o $@ github.com/golang/mock/mockgen
