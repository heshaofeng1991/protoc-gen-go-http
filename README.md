# protoc-gen-go-http

## install

```shell
$ go get github.com/yusank/protoc-gen-go-http
```

## how to use

```shell
$ make proto
# or
$ protoc -I/usr/local/include -I$(GOPATH)/src/github.com/googleapis/googleapis\
 	--proto_path=$(GOPATH)/src:. --go_out=. --go-http_out=. --go-grpc_out=.\
 	 path/to/file.proto
```

基于protoc工具生成golang语言http代码
