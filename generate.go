//go:generate protoc --proto_path=${GOPATH}/src:. --micro_out=api --go_out=api api/api.proto
package main
