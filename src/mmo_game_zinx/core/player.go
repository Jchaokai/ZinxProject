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
	X float32 //平面的x坐标
	Y float32 //高度
	Z float32 //平面的y坐标
	V float32 //旋转的角度0-360
}

//告诉客户端playerID
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
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	p.SendMsg(200, msg)
}

func NewPlayer(conn ziface.IConn) *Player {
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()
	return &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(150 + rand.Intn(30)),
		Y:    0,
		Z:    float32(130 + rand.Intn(30)),
		V:    0,
	}
}

//发送给U3D客户端的消息
//主要是将 protobuf数据序列化后，调用zinx conn 的sendMsg
func (p *Player) SendMsg(msgID uint32, data proto.Message) {
	//将proto message结构体序列化
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
		play.SendMsg(200, protoMsg)
	}
}

func (p *Player) SyncSurrounding() {
	//1.获取当前玩家的周围玩家(九宫格)
	playerIDs := WorldObj.Aoi.GetPlayerIDsByPos(p.X, p.Z)
	players := make([]*Player, 0, len(playerIDs))
	for _, pid := range playerIDs {
		players = append(players, WorldObj.GetPlayerByID(int32(pid)))
	}
	//2.周围玩家通过MsgID:200 向各自客户端发送刚上线玩家的位置信息
	//2.1组建 msgID:200 proto数据
	protoMsg := &proto2.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &proto2.BroadCast_P{
			P: &proto2.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	//2.2 周围的玩家都向 各自的客户端发送数据
	for _, p := range players {
		p.SendMsg(200, protoMsg)
	}
	//3.刚上线玩家向自己的客户端发送MagID:202 周围玩家的位置信息
	//3.1 msgID:202 proto 数据
	players_proto_msg := make([]*proto2.Player, 0, len(players))
	for _, p := range players {
		players_proto_msg = append(players_proto_msg, &proto2.Player{
			Pid: p.Pid,
			P: &proto2.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		})
	}
	Syncplayers_proto := &proto2.SyncPlayers{
		Ps: players_proto_msg[:],
	}
	p.SendMsg(202, Syncplayers_proto)
}

//广播并跟新当前玩家的位置信息
func (p *Player) UpdatePos(X float32, Y float32, Z float32, V float32) {
	//更新当前玩家的坐标
	p.X = X
	p.Y = Y
	p.Z = Z
	p.V = V
	//组件广播协议MsgID:200 tp:4
	protomsg := &proto2.BroadCast{
		Pid: p.Pid,
		Tp:  4, //4 移动之后坐标信息更新
		Data: &proto2.BroadCast_P{
			P: &proto2.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	//获取玩家的周围玩家
	players := p.GetSurroundingPlayers()
	//周围玩家向自己的客户端发送当前玩家位置信息
	for _, player := range players {
		player.SendMsg(200, protomsg)
	}
}

func (p *Player) GetSurroundingPlayers() (players []*Player) {
	playerIDs := WorldObj.Aoi.GetPlayerIDsByPos(p.X, p.Z)
	for _, pid := range playerIDs {
		players = append(players, WorldObj.GetPlayerByID(int32(pid)))
	}
	return
}

func (p *Player) Offline() {
	//得到下线玩家 的周围玩家
	players := p.GetSurroundingPlayers()
	//周围玩家想自己的客户端发送 要下线玩家的信息
	protomsg := &proto2.SyncPid{
		Pid: p.Pid,
	}
	for _, player := range players {
		player.SendMsg(201, protomsg)
	}
	//将下线的玩家从世界管理器中删除
	WorldObj.Aoi.RemovePlayFromGridByPos(int(p.Pid), p.X, p.Z)
	WorldObj.RemovePlayerByPid(p.Pid)

}
