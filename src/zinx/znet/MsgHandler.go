package znet

import (
	"ZinxProject/src/zinx/utils"
	"ZinxProject/src/zinx/ziface"
	"fmt"
	"strconv"
)

type MsgHandler struct {
	//不同msgID 对应的不同路由处理方法
	Apis map[uint32]ziface.IRouter
	//消息队列
	TaskQueue []chan ziface.IRequest
	//工作池
	WorkPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:         map[uint32]ziface.IRouter{},
		WorkPoolSize: utils.GlobalObject.WorkPoolSize,
		TaskQueue:    make([]chan ziface.IRequest, utils.GlobalObject.WorkPoolSize),
	}
}

func (msgHandle *MsgHandler) DoMsgHandle(r ziface.IRequest) {
	router, ok := msgHandle.Apis[r.GetMsgID()]
	if !ok {
		fmt.Println("msgID ", r.GetMsgID(), " --> API router [IS NOT FOUND] , is [Registering ing ...] ")
		return
		//msgHandle.AddRouter(r.GetMsgID(),new(BaseRouter))
		//msgHandle.DoMsgHandle(r)
	}
	go func(req ziface.IRequest) {
		router.PreHandle(req)
		router.Handle(req)
		router.PostHandle(req)
	}(r)
}

func (msgHandle *MsgHandler) AddRouter(msgID uint32, r ziface.IRouter) {
	//1.当前msgID对应的router是否存在
	if _, ok := msgHandle.Apis[msgID]; ok {
		panic("repeat router ,msgID : " + strconv.Itoa(int(msgID)))
	}
	//2.添加
	msgHandle.Apis[msgID] = r
	fmt.Printf("msgID : %d --> API router added success\n", msgID)
}

//启动一个worker pool,zinx框架开启工作池的动作智能发生一次
func (msgHandle *MsgHandler) StartWorkerPool() {
	//根据workerPoolSize 开启workers
	for i := 0; i < int(msgHandle.WorkPoolSize); i++ {
		//1.给当前worker对应的channel 开辟空间,（第0个worker就用地0个channel）
		msgHandle.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskSize)
		//2.启动当前的worker，阻塞等待消息从channel传递出来
		go msgHandle.StartOneWorker(i, msgHandle.TaskQueue[i])
	}
}

//启动一个worker工作流程
func (msgHandle *MsgHandler) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Printf("workerID : %d is starting . . . \n", workerID)
	//不断阻塞等待 对应消息队列的消息
	for {
		select {
		case request := <-taskQueue:
			msgHandle.DoMsgHandle(request)
		}
	}
}

//将 request 交给 taskQueue,由worker处理
func (msgHandle *MsgHandler) SendMsgToTaskQueue(req ziface.IRequest) {
	//将 消息平均分配（根据connID/ requestID）给不同的worker
	workerID := req.GetConn().GetConnID() % msgHandle.WorkPoolSize
	fmt.Printf("conn: %s connID: [%d] request is send to workerID : [%d] is solving\n",
		req.GetConn().GetRemoteAddr().String(), req.GetConn().GetConnID(), workerID)
	//将request发送给对应的worker的taskQueue
	msgHandle.TaskQueue[workerID] <- req
}
