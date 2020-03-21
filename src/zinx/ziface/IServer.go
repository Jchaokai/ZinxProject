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
	AddRouter(msgID uint32, r IRouter)
	GetConnManager() IConnManager
	//注册 OnConnStart hook函数
	SetOnConnStart(func(conn IConn))
	//注册 OnConnStop hook函数
	SetOnConnStop(func(conn IConn))
	//调用 OnConnStart hook函数
	CallOnConnStart(conn IConn)
	//调用 OnConnStart hook函数
	CallOnConnStop(conn IConn)
}
