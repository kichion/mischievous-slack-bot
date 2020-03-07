.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/events events/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/callback callback/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	serverless deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

lint:
	golangci-lint run --config linter/.golangci.yml

test:
	gotestsum -- -cover ./...
