package main

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"

	pb "go-examples/grpc/server/myserver/protoc/hi"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:9999"
)

// 定义helloService并实现约定的接口
type HiService struct{}

// HiService Hello服务
var HiSer = HiService{}

// SayHello 实现Hello服务接口
func (h HiService) SayHi(ctx context.Context, in *pb.HiRequest) (*pb.HiResponse, error) {

	// ====== 添加的逻辑  START ==============

	// 解析metada中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "no token ")
	}

	var (
		appId  string
		appKey string
	)

	// md 是一个 map[string][]string 类型的
	if val, ok := md["appid"]; ok {
		appId = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appKey = val[0]
	}

	if appId != "myappid" || appKey != "mykey" {
		return nil, status.Errorf(codes.Unauthenticated, "token invalide: appid=%s, appkey=%s", appId, appKey)
	}

	// ====== 添加的逻辑  END ==============

	resp := new(pb.HiResponse)
	resp.Message = fmt.Sprintf("Hi %s.", in.Name)

	return resp, nil
}

func main() {
	log.SetFlags(log.Ltime | log.Llongfile)
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Panicf("Failed to listen: %v", err)
	}

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile(
		"./grpc/server/keys/server.crt",
		"./grpc/server/keys/server.key",
	)
	if err != nil {
		log.Panicf("Failed to generate credentials %v", err)
	}

	// 实例化grpc Server, 并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册HelloService
	pb.RegisterHiServer(s, HiSer)

	log.Println("Listen on " + Address + " with TLS")

	s.Serve(listen)
}
