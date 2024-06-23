go build -ldflags "-X main.BuildDate=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.Hash=`git rev-parse HEAD`" .
