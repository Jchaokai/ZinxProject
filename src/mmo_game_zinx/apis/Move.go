package apis

import (
	"ZinxProject/src/mmo_game_zinx/core"
	proto2 "ZinxProject/src/mmo_game_zinx/proto"
	"ZinxProject/src/zinx/ziface"
	"ZinxProject/src/zinx/znet"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type MoveRouter struct {
	znet.BaseRouter
}

func (m *MoveRouter) Handle(r ziface.IRequest) {
	//解析客户端传递的proto
	position := &proto2.Position{}
	if e := proto.Unmarshal(r.GetData(), position); e != nil {
		fmt.Println("Move proto bytes Unmarshal error ", e)
		return
	}
	//得到发送位置信息的是哪个玩家
	pid, e := r.GetConn().GetProperty("pid")
	if e != nil {
		fmt.Println("未知玩家发送位置信息")
		return
	}
	//fmt.Printf("player pid:%d ,position x:%f y:%f z:%f v:%f \n",pid,position.X,position.Y,position.Z,position.V)
	player := core.WorldObj.GetPlayerByID(pid.(int32))
	//其他玩家向自己的客户端发送当前玩家位置信息广播
	player.UpdatePos(position.X, position.Y, position.Z, position.V)

}
