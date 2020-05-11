package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"time"

	chat "github.com/yezihack/studyGo/gprc/client/proto"
	"google.golang.org/grpc"
)

const (
	Address = "localhost:8009"
)

func main() {
	log.SetOutput(os.Stdout)
	//建立链接 , grpc.WithInsecure() 必须要传
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	//建立一个客户端
	c := chat.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := chat.SendRequest{}
	req.Content = "world"
	//调用 proto里的方法
	reply, err := c.Send(ctx, &req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(reply.Msg)
}
