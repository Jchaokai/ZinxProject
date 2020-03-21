package main

import (
	"ZinxProject/src/mmo_game_zinx/core"
	"ZinxProject/src/zinx/ziface"
	"ZinxProject/src/zinx/znet"
	"fmt"
)

func OnConn(conn ziface.IConn) {
	//创建一个player对象
	player := core.NewPlayer(conn)
	fmt.Println("=======> PlayerID pid = ", player.Pid, " 上线了 <========")
	//给客户端发送 msgID:1 消息
	player.SyncID()
	//给客户端发送 msgID:200 消息
	player.BroadCastStartPos()
}
func main() {
	server := znet.NewServer()
	//连接创建销毁的 hook函数
	server.SetOnConnStart(OnConn)
	server.Server()
}
