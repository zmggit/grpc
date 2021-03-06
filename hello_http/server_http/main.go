package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"context"
	"google.golang.org/grpc"
	"net/http"
	"github.com/zmggit/grpc"
)

func main()  {
	ctx := context.Background()
	ctx , cancel := context.WithCancel(ctx)
	defer cancel()

	// grpc服务地址a
	endpoint := "127.0.0.1:50052"
    mux :=  runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}

	// HTTP转grpc
	err := gw.RegisterHelloHTTPHandlerFromEndpoint(ctx,mux,endpoint,opts)
	if err != nil {
		fmt.Println("Register handler err:%v\n", err)
	}
	fmt.Println("HTTP Listen on 8080")
	http.ListenAndServe(":8080",mux)

}