# protoc-gen-go-http
基于protoc工具生成golang语言http代码

## 安装

```shell
$ go install github.com/heshaofeng1991/protoc-gen-go-http@latest 
```

## 生成代码其它依赖
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
go install github.com/heshaofeng1991/protoc-gen-go-http@latest
```

## 使用说明

```shell
$ make proto
# or
$ protoc -I=. -I=${GOPATH}/pkg/mod \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  --govalidators_out=. --govalidators_opt=paths=source_relative \
  --go-http_out=. --go-http_opt=paths=source_relative \
  --gogofaster_out=plugin=gprc:. --gogofaster_opt=paths=source_relative \
  proto/example/example.proto
```
