.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/franz main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
