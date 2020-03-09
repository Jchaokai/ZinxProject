package ziface

// server 接口
type IServer interface {
	//start
	Start()
	//stop
	Stop()
	//run
	Server()
	//给当前的服务注册一个 路由方法，供client 连接使用
	AddRouter(msgID uint64, r IRouter)
	GetConnManager() IConnManager
}
