package ziface
/*
	封包，拆包 模块
	处理TCP粘包问题
*/

type IDataPack interface {
	//获取包头 长度
	GetHeadLen() uint
	//封包
	Pack(IMessage) ([]byte,error)
	//拆包
	UnPack([]byte) (IMessage,error)
}
