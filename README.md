# protoc-gen-go-http

## install

```shell
$ go install github.com/heshaofeng1991/protoc-gen-go-http@latest 
```

## how to use

```shell
$ make proto
# or
$ protoc -I=. -I=${GOPATH}/pkg/mod \
	--go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --govalidators_out=. --govalidators_opt=paths=source_relative \
    --go-http_out=. --go-http_opt=paths=source_relative \
   proto/example/example.proto
```

基于protoc工具生成golang语言http代码
