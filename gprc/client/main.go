package main

import (
	"log"
	"os"
	"google.golang.org/grpc"
	"github.com/yezihack/studyGo/gprc/server/proto"
)

const (
	Address = "localhost:8009"
)
func main() {
 	log.SetOutput(os.Stdout)
 	conn, err := grpc.Dial(Address)
 	if err != nil {
 		log.Fatalln(err)
	}
 	defer conn.Close()
 	c := chat.NewChatServiceClient(conn)


}
