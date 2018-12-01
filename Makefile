
prepare:
	dep ensure -v

build: prepare
	go build -o ./bin/stagectl -v

tests:
	go test ./... -v

run:
	go run main.go