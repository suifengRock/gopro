export GOPATH := ${PWD}
export PATH := ${PATH}:${GOPATH}/bin
export GOBIN := ${GOPATH}/bin
main:
	go run main.go

pro:
	go run gopro.go test

build:
	go install gopro.go
