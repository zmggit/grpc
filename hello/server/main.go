package main

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

//定义HelloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
//func (h helloService) SayHello(ctx context.Context,in *) () {
//
//}