package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"net"
	"google.golang.org/grpc/grpclog"
	pb "github.com/zmggit/go-grpc"
	"google.golang.org/grpc/metadata"

)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

//定义HelloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context,in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 解析metada中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		appid  string
		appkey string
	)

	if val, ok := md["appid"]; ok {
		appid = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "101010" || appkey != "i am key" {
		return nil, grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}


	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", in.Name, appid, appkey)

	return resp, nil
}

func main()  {
	listen,err := net.Listen("tcp",Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	//TlS 认证
	creds , err := credentials.NewServerTLSFromFile("../../keys/server.crt","../../keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}
	// 实例化grpc Server并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册HelloService
	pb.RegisterHelloServer(s,HelloService)

	fmt.Println("Listen on " + Address + " with TLS")
	s.Serve(listen)
}
