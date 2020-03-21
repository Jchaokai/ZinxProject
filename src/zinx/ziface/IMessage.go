package ziface

/*
	将 客户端请求的数据封装到Message
*/

type IMessage interface {
	GetMsgID() uint32
	GetMsgLen() uint32
	GetMsgData() []byte
	SetMsgID(uint32)
	SetMsgLen(uint32)
	SetData([]byte)
}
