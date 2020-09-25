test:
	./scripts/validate-license.sh
	go fmt .
	go mod tidy
	go test .
	golangci-lint run --allow-parallel-runners -v --enable-all --disable testpackage --fix

build:
	docker build . -t paskalmaksim/helm-sluglify:dev