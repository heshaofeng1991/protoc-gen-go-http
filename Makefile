GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc -I=. -I=${GOPATH}/pkg/mod \
	--go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --govalidators_out=. --govalidators_opt=paths=source_relative \
    --go-http_out=. --go-http_opt=paths=source_relative \
   proto/example/example.proto
