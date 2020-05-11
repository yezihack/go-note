#   Grpc


## 资料
- https://grpc.io/docs/quickstart/go/
- https://studygolang.com/articles/16627

## 新建一个proto文件
```
syntax = "proto3";

package chat;

service ChatService {
    //发送消息
    rpc Send(SendRequest) returns (SendReply){};
}
//发送消息 请求结构体
message SendRequest {
    string content = 1;//发送内容
}
//发送消息 响应结构体
message SendReply {
    string msg = 1;//返回消息
}
```

## GRPC客户端
```
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
	//启动服务
	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

//新建一个结构体,实现proto里定义的方法
type Chats struct {
}
//实现proto方法
func (c *Chats) Send(ctx context.Context, in *chat.SendRequest) (*chat.SendReply, error) {
	out := chat.SendReply{
		Msg: "hello " + in.Content,
	}
	return &out, nil
}

```

## 客户端代码
```
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

```