package ziface

/*
	抽象IRouter
	路由里的数据都是 IRequest
	处理业务之前的方法
	处理业务的方法
	处理业务之后的方法

	再定义一个具体实现了IRouter接口的 BaseRouter
	（用户想要自定义一个router，继承BaseRouter就好）

 */

type IRouter interface {
	//处理conn业务之前
	PreHandle(r IRequest)
	Handle(r IRequest)
	//处理conn业务之后
	PostHandle(r IRequest)
}