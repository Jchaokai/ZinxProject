package main

import (
	"ZinxProject/src/mmo_game_zinx/apis"
	"ZinxProject/src/mmo_game_zinx/core"
	"ZinxProject/src/zinx/ziface"
	"ZinxProject/src/zinx/znet"
	"fmt"
)

func OnConn(conn ziface.IConn) {
	//创建一个player对象
	player := core.NewPlayer(conn)
	fmt.Println("=======> PlayerID pid = ", player.Pid, " 上线了 <========")
	//给当前conn绑定玩家信息
	conn.SetProperty("pid", player.Pid)
	//给客户端发送 msgID:1 消息
	player.SyncID()
	//给客户端发送 msgID:200 消息
	player.BroadCastStartPos()
	//当玩家刚刚上线时，应该加入世界管理模块
	core.WorldObj.AddPlayer(player)
}
func main() {
	server := znet.NewServer()
	//连接创建销毁的 hook函数
	server.SetOnConnStart(OnConn)
	server.AddRouter(2, &apis.WorldChatRouter{})
	server.Server()
}
