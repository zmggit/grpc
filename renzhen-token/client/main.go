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
	// Address gRPC服务地址
	Address  = "127.0.0.1:50052"

	// OpenTLS 是否开启TLS认证
	OpenTLS = true
)

//customCredential 自定义认证
type customCredential struct {}

// GetRequestMetadata 实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string,error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	},nil
}
// RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}

func main()  {
	var err error
	var opts []grpc.DialOption

	if OpenTLS {
		//TLS连接  记得把server name改成你写的服务器地址
		creds , err := credentials.NewClientTLSFromFile("../../keys/server.crt","a.example.com")
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = append(opts,grpc.WithTransportCredentials(creds))
	}else {
		opts = append(opts,grpc.WithInsecure())
	}

	//使用自定义认证

	opts = append(opts,grpc.WithPerRPCCredentials(new(customCredential)))

	conn, err := grpc.Dial(Address,opts...)
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

