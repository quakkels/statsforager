build:
	@go build  -o bin/statsforagerapi -ldflags "-X main.BuildDate=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.Hash=`git rev-parse HEAD`" .

run: build
	@./bin/statsforagerapi

test:
	@go test -v ./...
