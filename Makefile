
prepare:
	dep ensure -v

build: prepare
	NOW=$(date +'%Y-%m-%d_%T') && \
	go build -o ./bin/stagectl -v -ldflags "-X github.com/auto-staging/stagectl/cmd.gitSha=`git rev-parse HEAD` -X github.com/auto-staging/stagectl/cmd.buildTime=$NOW -X github.com/auto-staging/stagectl/cmd.version=LOCAL_BUILD"

tests:
	go test ./... -v

run:
	go run main.go
