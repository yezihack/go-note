package main

import (
	"log"
	"net"
	"os"

	"context"

	chat "github.com/yezihack/studyGo/gprc/server/proto"
	"google.golang.org/grpc"
)

var (
	port = ":8009"
)

func main() {
	log.SetOutput(os.Stdout)
	//新建一个tcp监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	//起一个服务
	s := grpc.NewServer()
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	chat.RegisterChatServiceServer(s, &Chats{})
	log.Printf("server port %s start...\n", port)
	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

type Chats struct {
}

func (c *Chats) Send(ctx context.Context, in *chat.SendRequest) (*chat.SendReply, error) {
	out := chat.SendReply{
		Msg: "hello " + in.Content,
	}
	return &out, nil
}
