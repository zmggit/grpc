package main

import (
	"context"
	"fmt"
	pb "github.com/zmggit/go-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	Address  = "127.0.0.1:50052"
)

func main()  {
	//TLS连接  记得把server name改成你写的服务器地址
	creds , err := credentials.NewClientTLSFromFile("../../keys/server.crt","a.example.com")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	//初始化客户端
	c := pb.NewHelloClient(conn)

	//调用方法
	req := &pb.HelloRequest{Name:"gRPC"}
	res,err := c.SayHello(context.Background(),req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	fmt.Println(res.Message,"打印错误")

}

