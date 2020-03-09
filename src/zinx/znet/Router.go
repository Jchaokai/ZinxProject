package znet

import (
	"ZinxProject/src/zinx/ziface"
)

/*
	定义 BaseRouter 是为了以后自定义router
 */
type BaseRouter struct{}

func (br *BaseRouter) PreHandle(r ziface.IRequest) {
	//fmt.Print("BaseRouter PreHandle - ")
}

func (br *BaseRouter) Handle(r ziface.IRequest) {
	//fmt.Print("Handle - ")
}

func (br *BaseRouter) PostHandle(r ziface.IRequest) {
	//fmt.Println("PostHandle ")
}
