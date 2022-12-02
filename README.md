# protoc-gen-go-http
```text
基于protoc工具生成golang语言gin web框架的 http代码，非常感谢kratos这个框架。

目前是制定了gin的方式，其实是可以大部分的框架都是可以的，可以自行修改生成依赖的。

1. 基于pb生成grpc validate现在有很多脚手架，成熟框架了。
```

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

## 后续可以在拓展的方向
```text
protoc-gen-errors
protoc-gen-swagger
```
