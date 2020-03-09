package znet

import (
	"ZinxProject/src/zinx/ziface"
	"errors"
	"fmt"
	"sync"
)

type ConnManager struct {
	Conns    map[uint64]ziface.IConn //管理的连接集合
	ConnLock sync.RWMutex            //保护连接集合的读写锁
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		Conns: make(map[uint64]ziface.IConn),
	}
}

func (cm *ConnManager) Add(conn ziface.IConn) {
	//加写锁
	cm.ConnLock.Lock()
	defer cm.ConnLock.Unlock()
	//将conn加入到map
	cm.Conns[conn.GetConnID()] = conn
	fmt.Printf("- - - - - - - - - - - - - - - - - - - - - - - - - -\n "+
		"[ connID : %d remoteAddr : %s ] add to ConnManager successfully, [current conn number : %d] \n",
		conn.GetConnID(), conn.GetRemoteAddr(), cm.Len())
}

func (cm *ConnManager) Remove(conn ziface.IConn) {
	//加 写锁
	cm.ConnLock.Lock()
	defer cm.ConnLock.Unlock()
	delete(cm.Conns, conn.GetConnID())
	fmt.Printf("- - - - - -- - - - -- - - - - -- - - - - - - - - \n"+
		"[ connID : %d remoteAddr : %s ] remove from ConnManager successfully, [current conn number : %d] \n",
		conn.GetConnID(), conn.GetRemoteAddr(), cm.Len())
}

func (cm *ConnManager) Get(connID uint64) (ziface.IConn, error) {
	//加 读锁
	cm.ConnLock.RLock()
	defer cm.ConnLock.RUnlock()

	if conn, ok := cm.Conns[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("conn NOT FOUND")
	}
}

func (cm *ConnManager) Len() uint64 {
	return uint64(len(cm.Conns))
}

func (cm *ConnManager) ClearConn() {
	//加 写锁
	cm.ConnLock.Lock()
	defer cm.ConnLock.Unlock()
	//删除conn,并停止相关工作
	for connID, conn := range cm.Conns {
		conn.Stop()
		delete(cm.Conns, connID)
	}
	fmt.Printf("clear all conn succ!! [conn number : %d ]\n", cm.Len())
}
