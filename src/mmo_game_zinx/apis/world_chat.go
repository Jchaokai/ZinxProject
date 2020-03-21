package apis

import (
	"ZinxProject/src/mmo_game_zinx/core"
	proto2 "ZinxProject/src/mmo_game_zinx/proto"
	"ZinxProject/src/zinx/ziface"
	"ZinxProject/src/zinx/znet"
	"fmt"
	"github.com/golang/protobuf/proto"
)

/*
	世界聊天业务
*/

type WorldChatRouter struct {
	znet.BaseRouter
}

func (wc *WorldChatRouter) Handle(r ziface.IRequest) {
	//知道发送消息的是哪个玩家
	pid, e := r.GetConn().GetProperty("pid")
	if e != nil {
		fmt.Println("未知玩家发送消息 ", e)
		return
	}
	player := core.WorldObj.GetPlayerByID(pid.(int32))
	//解析聊天proto信息变成以便世界广播
	protoMsg := &proto2.Talk{}
	if e := proto.Unmarshal(r.GetData(), protoMsg); e != nil {
		fmt.Println("world talk proto bytes 解析 error ", e)
		return
	}
	//将聊天广播给所有人
	player.Talk(protoMsg.Content)
}
