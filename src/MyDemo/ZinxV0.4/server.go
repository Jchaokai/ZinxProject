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

func (pr *PingRouter) PreHandle(r ziface.IRequest){
	if _, e := r.GetConn().GetTcpConn().Write([]byte("before ping !!  "));e!=nil{
		fmt.Println("preHandle write [before ping !!] error",e.Error())
		return
	}
}
func (pr *PingRouter) Handle(r ziface.IRequest){
	if _, e := r.GetConn().GetTcpConn().Write([]byte("ping !!  "));e!=nil{
		fmt.Println("Handle write [ ping !!] error",e.Error())
		return
	}
}
func (pr *PingRouter) PostHandle(r ziface.IRequest){
	if _, e := r.GetConn().GetTcpConn().Write([]byte("after ping !!  "));e!=nil{
		fmt.Println("preHandle write [after ping !!] error",e.Error())
		return
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
