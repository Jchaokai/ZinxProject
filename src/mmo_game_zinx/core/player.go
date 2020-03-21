package core

import (
	proto2 "ZinxProject/src/mmo_game_zinx/proto"
	"ZinxProject/src/zinx/ziface"
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"sync"
)

var (
	PidGen int32 = 1
	IdLock sync.Mutex
)

type Player struct {
	Pid  int32
	Conn ziface.IConn
	//U3D的坐标体系 与服务器定义的有区别
	x float32 //平面的x坐标
	y float32 //高度
	z float32 //平面的y坐标
	v float32 //旋转的角度0-360
}

//告诉客户但playerID,同步已经生成的玩家ID 给客户端
func (p *Player) SyncID() {
	proto_data := &proto2.SyncPid{
		Pid: p.Pid,
	}
	p.SendMsg(1, proto_data)
}

//广播玩家自己的出生地
func (p *Player) BroadCastStartPos() {
	msg := &proto2.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &proto2.BroadCast_P{
			P: &proto2.Position{
				X: p.x,
				Y: p.y,
				Z: p.z,
				V: p.v,
			},
		},
	}
	p.SendMsg(200, msg)
}

func NewPlayer(conn ziface.IConn) *Player {
	//没有数据库，先暂时生成一个玩家id
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()
	return &Player{
		Pid:  id,
		Conn: conn,
		//随机在160左右偏移
		x: float32(160 + rand.Intn(10)),
		y: 0,
		z: float32(140 + rand.Intn(20)),
		v: 0,
	}
}

//发送给U3D客户端的消息
//主要是将 protobuf数据序列化后，调用zinx的sendMsg
func (p *Player) SendMsg(msgID uint32, data proto.Message) {
	//将proto message序列化
	bytes, e := proto.Marshal(data)
	if e != nil {
		fmt.Println("proto marshal error", e)
		return
	}
	//通过zinx的sendMsg 将二进制发送给客户端
	if p.Conn == nil {
		fmt.Println("conn of player is closed,dont send proto_bytes", e)
		return
	}
	if e := p.Conn.SendMsg(msgID, bytes); e != nil {
		fmt.Println("player sendMsg error !!", e)
		return
	}
}

//广播世界聊天消息
func (p *Player) Talk(content string) {
	//组建一个msgID:200 的proto数据
	protoMsg := &proto2.BroadCast{
		Pid:  p.Pid,
		Tp:   1,
		Data: &proto2.BroadCast_Content{Content: content},
	}
	//得到所有在线玩家
	allPlayers := WorldObj.GetAllPlayers()
	//发送
	for _, play := range allPlayers {
		go play.SendMsg(200, protoMsg)
	}
}
