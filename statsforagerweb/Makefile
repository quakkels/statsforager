
build:
	@go build  -o bin/statsforagerweb -ldflags "-X main.BuildDate=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.Hash=`git rev-parse HEAD`" .

image: build
	@docker build --rm -t quakkels/statsforagerweb:alpha .

run: image
	@docker run --rm -p 8000:8000 -t quakkels/statsforagerweb:alpha

publish: image
	@docker push quakkels/statsforagerweb

run-native: build
	@./bin/statsforagerweb

test:
	@go test -v ./...

