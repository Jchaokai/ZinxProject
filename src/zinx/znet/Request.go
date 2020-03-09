package znet

/*
	绑定 IConn，data
*/
import (
	"ZinxProject/src/zinx/ziface"
)

type Request struct {
	conn ziface.IConn
	msg  ziface.IMessage
}

func (r *Request) GetConn() ziface.IConn {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}
func (r *Request) GetMsgID() uint64 {
	return r.msg.GetMsgID()
}
