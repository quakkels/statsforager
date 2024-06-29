
build:
	@go build  -o bin/statsforagerapi -ldflags "-X main.BuildDate=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.Hash=`git rev-parse HEAD`" .
	
image: build
	@docker build --rm -t statsforagerapi:alpha .

run: image
	@docker run --rm -p 8000:8000 -t statsforagerapi:alpha
	
runnative: build
	@./bin/statsforagerapi

test:
	@go test -v ./...
