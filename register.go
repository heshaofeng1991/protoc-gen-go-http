package main

//import (
//	"context"
//
//	"github.com/gin-gonic/gin"
//	"google.golang.org/grpc"
//
//	v1 "github.com/heshaofeng1991/protoc-gen-go-http/proto/example"
//)
//
//// Greeter implement v1.HelloServer and v1.HelloHTTPHandler at same time
//type Greeter struct {
//	v1.UnimplementedExampleServer
//}
//
//func (g *Greeter) Add(ctx context.Context, in *v1.AddRequest) (*v1.CommonResponse, error) {
//	panic("implement me")
//}
//
//func (g *Greeter) Get(ctx context.Context, in *v1.GetRequest) (*v1.CommonResponse, error) {
//	panic("implement me")
//}
//
//func register() {
//	// can call method by rpc or http
//	v1.RegisterExampleServer(&grpc.Server{}, &Greeter{})
//	v1.RegisterExampleHTTPHandler(gin.Default().Group("/"), &Greeter{})
//}
