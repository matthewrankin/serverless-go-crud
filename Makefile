.PHONY: build clean deploy
GOBUILD=env GOOS=linux go build -ldflags="-s -w" -o

help:
	@echo "You can perform the following:"
	@echo ""
	@echo "  check  Format, lint, vet, and test Go code"
	@echo "  build  Builds the Go code"
	@echo "  deploy Build and then deploys to AWS using serverless"

build: cleanbin
	export GO111MODULE=on
	$(GOBUILD) bin/addtodo src/handlers/addtodo/main.go
	$(GOBUILD) bin/listtodos src/handlers/listtodos/main.go
	$(GOBUILD) bin/completetodo src/handlers/completetodo/main.go
	$(GOBUILD) bin/deletetodo src/handlers/deletetodo/main.go

check:
	@echo 'Formatting, linting, vetting, and testing Go code.'
	go fmt ./...
	golint ./...
	go vet ./...
	go test ./...

cleanbin:
	rm -rf ./bin

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
