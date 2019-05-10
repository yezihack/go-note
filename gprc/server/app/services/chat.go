package services

import (
	"context"

	chat "github.com/yezihack/studyGo/gprc/server/proto"
)

//定义接口
type ChatServer interface {
	Send(ctx context.Context, req *chat.SendRequest) (reply *chat.SendReply, err error)
	Receive(ctx context.Context, req *chat.ReceiveRequest) (reply *chat.ReceiveReply, err error)
}

//定义一个接口体,实现接口
type ChatServe struct {
}

//实现接口里的方法: send
func (cs *ChatServe) Send(ctx context.Context, req *chat.SendRequest) (reply *chat.SendReply, err error) {
	reply.Status = 200
	reply.Msg = "发送成功, 你发送的消息, content: " + req.Content
	return
}

//实现接口里的方法: receive
func (cs *ChatServe) Receive(ctx context.Context, req *chat.ReceiveRequest) (reply *chat.ReceiveReply, err error) {

	return
}

//对外提供接口调用的方法
func GetChatServe() ChatServer {
	return &ChatServe{}
}
