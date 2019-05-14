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
	//建立链接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := chat.SendRequest{}
	req.Content = "world"
	reply, err := c.Send(ctx, &req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(reply.Msg)
}
