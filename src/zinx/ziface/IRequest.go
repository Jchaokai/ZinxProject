package ziface

/*
	将 Conn与数据绑定
*/
type IRequest interface {
	//get当前连接
	GetConn() IConn
	//get连接的msg
	GetData() []byte
	GetMsgID() uint32
}
