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
	//给当前conn 自定义属性
	conn.SetProperty("name", "吉安娜")
	conn.SetProperty("role", "法师")
	conn.SetProperty("level", 69)

}

//SetOnConnStop hookFunc
func ConnStopHook(conn ziface.IConn) {
	fmt.Println("# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # \n" +
		"处理 conn:" + conn.GetRemoteAddr().String() + " 连接断开的hook函数注册成功")
	name, _ := conn.GetProperty("name")
	role, _ := conn.GetProperty("role")
	level, _ := conn.GetProperty("level")
	fmt.Printf("# # # # # # # # # # # # # # # # # # # # # # # # # # # \n"+
		"用户 : %s ,职业： %s ,等级： %d   下线了 !\n", name, role, level)
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
