package main

import (
	"ZinxProject/src/zinx/ziface"
	"ZinxProject/src/zinx/znet"
	"fmt"
)

/*
	自定义一个router 继承baseRouter
*/

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) Handle(r ziface.IRequest) {
	e := r.GetConn().SendMsg(200, []byte(" client-handle writer PingRouter内容 "))
	if e != nil {
		fmt.Println("client-handle writer error")
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (pr *HelloRouter) Handle(r ziface.IRequest) {
	e := r.GetConn().SendMsg(201, []byte(" client-handle writer HelloRouter内容 "))
	if e != nil {
		fmt.Println("client-handle writer error")
	}
}

//SetOnConnStart hookFunc
func ConnStartHook(conn ziface.IConn) {
	if err := conn.SendMsg(111, []byte("处理conn start的hook函数注册成功")); err != nil {
		fmt.Println("SetOnConnStart hookFunc server sendMsg error :", err)
		return
	}
}

//SetOnConnStop hookFunc
func ConnStopHook(conn ziface.IConn) {
	if err := conn.SendMsg(111, []byte("处理 conn:"+conn.GetRemoteAddr().String()+" 连接断开的hook函数注册成功")); err != nil {
		fmt.Println("SetOnConnStop hookFunc server sendMsg error :", err)
		return
	}
}
func main() {
	//NewServer
	server := znet.NewServer()
	//注册处理conn的hook函数
	server.SetOnConnStart(ConnStartHook)
	server.SetOnConnStop(ConnStopHook)
	//多路由router
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})
	//启动
	server.Server()
}
