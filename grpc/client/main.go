package main

import (
	pb "go-examples/grpc/client/myclient/protoc/hi" // 引入proto包
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:9999"
)

// ====== 添加的逻辑  START ==============
var IsTls = true

// myCredential 自定义认证
type myCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c myCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "myappid",
		"appkey": "mykey",
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c myCredential) RequireTransportSecurity() bool {
	return IsTls
}

// ====== 添加的逻辑  END  ==============

func main() {
	log.SetFlags(log.Ltime | log.Llongfile)
	// ====== 添加的逻辑  START ==============
	var err error
	var opts []grpc.DialOption

	if IsTls {
		//打开tls 走tls认证
		creds, err := credentials.NewClientTLSFromFile(
			"./grpc/server/keys/server.crt",
			"localhost",
		)
		if err != nil {
			log.Panicf("Failed to create TLS mycredentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	opts = append(opts, grpc.WithPerRPCCredentials(new(myCredential)))

	// TLS连接  记得把xxx改成你写的服务器地址,可以默认写127.0.0.1
	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		grpclog.Fatalln(err)
	}
	// ====== 添加的逻辑  END  ==============
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHiClient(conn)

	// 调用方法
	req := &pb.HiRequest{Name: "gRPC"}
	res, err := c.SayHi(context.Background(), req)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res.Message)
}
