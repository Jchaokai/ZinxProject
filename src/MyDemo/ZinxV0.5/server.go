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
	e := r.GetConn().SendMsg(1, []byte(" client-handle writer的内容 "))
	if e != nil {
		fmt.Println("client-handle writer error")
	}
}


func main() {
	//NewServer
	server := znet.NewServer()
	//定义router
	server.AddRouter(&PingRouter{})
	//启动
	server.Server()
}
