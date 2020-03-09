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

func (pr *PingRouter) Handle(r ziface.IRequest){
	e := r.GetConn().SendMsg(200, []byte(" client-handle writer PingRouter内容 "))
	if e != nil {
		fmt.Println("client-handle writer error")
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (pr *HelloRouter) Handle(r ziface.IRequest){
	e := r.GetConn().SendMsg(201, []byte(" client-handle writer HelloRouter内容 "))
	if e != nil {
		fmt.Println("client-handle writer error")
	}
}

func main() {
	//NewServer
	server := znet.NewServer()
	//多路由router
	server.AddRouter(0,&PingRouter{})
	server.AddRouter(1,&HelloRouter{})
	//启动
	server.Server()
}
