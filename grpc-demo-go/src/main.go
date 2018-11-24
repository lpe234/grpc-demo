package main

import (
	"google.golang.org/grpc"
	pb "proto"
	"context"
	"log"
)

func main() {
	// java spring boot 暴露的grpc服务接口
	const addr = "127.0.0.1:10081"


	// 连接服务
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}

	// 确保连接最终被关闭
	defer conn.Close()

	// 建立远程调用客户端
	client := pb.NewUserProviderClient(conn)
	reply, err := client.GetByUserId(context.Background(), &pb.UserIdRequest{Id: 1})
	if err != nil {
		log.Panic(err)
	}

	// 输出结果
	log.Println("user info:", reply.Id, reply.Username)
}

