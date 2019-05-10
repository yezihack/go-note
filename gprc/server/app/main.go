package main

import (
	"log"
	"net"
	"os"

	"github.com/yezihack/studyGo/gprc/server/app/services"
	chat "github.com/yezihack/studyGo/gprc/server/proto"

	"google.golang.org/grpc"
)

var (
	port = ":8009"
)

func main() {
	log.SetOutput(os.Stdout)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	chat.RegisterChatServiceServer(s, services.GetChatServe())
	log.Println("server start...")
	s.Serve(lis)
}
