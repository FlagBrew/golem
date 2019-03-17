# enable go module support

export GO111MODULE=on

fetch:
	go mod download
	go mod tidy
	go mod vendor

test:
	go test -test.v
