package ziface

/*
	将 客户端请求的数据封装到Message
*/

type IMessage interface {
	GetMsgID() uint64
	GetMsgLen() uint64
	GetMsgData() []byte
	SetMsgID(uint64)
	SetMsgLen(uint64)
	SetData([]byte)
}
