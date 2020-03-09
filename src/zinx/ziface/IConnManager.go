package ziface

/*
	连接管理模块
*/
type IConnManager interface {
	Add(conn IConn)
	Remove(conn IConn)
	Get(connID uint64) (IConn, error)
	Len() uint64
	ClearConn()
}
