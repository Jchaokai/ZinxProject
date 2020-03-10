package ziface

import "net"

/*
	连接模块的接口
*/

type IConn interface {
	//启动连接
	Start()
	//停止连接
	Stop()
	//获取当前连接的socket
	GetTcpConn() *net.TCPConn
	//获取当前连接的ID
	GetConnID() uint64
	//获取当前远程客户端 IP Port TCP状态
	GetRemoteAddr() net.Addr
	//发送数据
	SendMsg(uint64, []byte) error
	//设置自定义属性
	SetProperty(key string, value interface{})
	//获取自定义属性
	GetProperty(key string) (interface{}, error)
	//删除自定义属性
	RemoveProperty(key string)
}

//定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
