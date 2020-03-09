package main

import "ZinxProject/src/zinx/znet"

func main() {
	//NewServer
	server := znet.NewServer()
	//启动
	server.Server()
}
