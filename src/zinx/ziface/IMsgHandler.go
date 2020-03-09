package ziface

/*
	为了实现多路由功能
	消息管理抽象成，根据不同的msgID使用不同的router处理
 */

type IMsgHandler interface {
	//不同消息 使用不同路由
	DoMsgHandle(r IRequest)
	//添加路由
	AddRouter(msgID uint64,r IRouter)
	StartWorkerPool()
	//将 request 交给 taskQueue,由worker处理
	SendMsgToTaskQueue(req IRequest)
}