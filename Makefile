SRC_DIR := cmd/osquerygo
BINARY := osquerygo
BUILD_DIR := build

.PHONY: clean format test

build: 
	mkdir -p ${BUILD_DIR}
	CGO_ENABLED=0 go build -o ${BUILD_DIR}/${BINARY} ${SRC_DIR}/main.go

run: 
	go run ${SRC_DIR}/main.go

clean: 
	go clean 
	rm -rf build

format: 
	golines -m 80 -t 4 -w .
	gofmt -w .

test: 
	go test ./...

deploy:
	git tag ${VERSION}
	git push origin ${VERSION}
