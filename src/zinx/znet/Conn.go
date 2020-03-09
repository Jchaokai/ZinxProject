package znet

/*
	封装的conn模块，包含tcpConn,连接状态,绑定的处理业务 etc.
*/

import (
	"ZinxProject/src/zinx/utils"
	"ZinxProject/src/zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint64
	isClosed bool
	//由 reader 告知 writer 推出chan
	ExitChan chan bool
	//该链接处理业务的router
	MsgHandler ziface.IMsgHandler
	//读写分离之间的 channel
	MsgChan chan []byte
}

//读客户端数据
func (c *Connection) StartReader() {
	fmt.Printf("[client-handle reader] of [ %d ] - [%s]is running ...\n", c.ConnID, c.Conn.RemoteAddr().String())
	defer fmt.Printf("[client-handle reader] of [ %d ] - [%s] is exiting ...\n", c.ConnID, c.Conn.RemoteAddr().String())
	defer c.Stop()

	//模拟业务
	for {
		//拆包
		dp := NewDataPack()
		packHeadBinary := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.Conn, packHeadBinary)
		if err != nil {
			fmt.Println("解析包头 error:", err)
			break
		}
		Msg, er := dp.UnPack(packHeadBinary)
		if er != nil {
			fmt.Println("根据包头 拆包error:", er)
			break
		}
		if Msg.GetMsgLen() > 0 {
			data := make([]byte, Msg.GetMsgLen())
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				fmt.Println("获取包体数据 error:", err)
				break
			}
			Msg.SetData(data)
		}
		//print完整的msg
		fmt.Printf("Client-Handle Reader得到 Msg [ dataLen: %d | ID : %d | data : %s]\n",
			Msg.GetMsgLen(), Msg.GetMsgID(), string(Msg.GetMsgData()))
		//得到request
		request := &Request{
			conn: c,
			msg:  Msg,
		}
		//server中已经开启worker pool和TaskQueue,判断是否开启了
		if utils.GlobalObject.WorkPoolSize > 0 {
			//将request交给工作池
			c.MsgHandler.SendMsgToTaskQueue(request)
		}else{
			//如果用户不适用worker pool,直接开协程处理
			go c.MsgHandler.DoMsgHandle(request)
		}
	}
}

//要发送的数据先封装，在发送给MsgChan,供client-handle writer使用
func (c *Connection) SendMsg(msgId uint64, data []byte) error {
	if c.isClosed {
		return errors.New("conn closed when send msg")
	}
	dp := NewDataPack()
	packBinary, e := dp.Pack(NewMsg(msgId, data))
	if e != nil {
		return errors.New("数据封装 error")
	}
	c.MsgChan <- packBinary
	return nil
}

//向客户端写数据
func (c *Connection) StartWriter() {
	fmt.Printf("[client-handle writer] of [ %d ] - [%s]is running ...\n", c.ConnID, c.Conn.RemoteAddr().String())
	defer fmt.Printf("[client-handle writer] of [ %d ] - [%s] is exiting ...\n", c.ConnID, c.Conn.RemoteAddr().String())
	//不断阻塞,等待读写协程 之间channel的消息
	for {
		select {
		case data := <-c.MsgChan:
			//有数据写给client
			if _, e := c.Conn.Write(data); e != nil {
				fmt.Println("client-handle writer error :", e)
				return
			}
		case <-c.ExitChan:
			//reader告知writer可以推出
			return
		}
	}
}

func (c *Connection) Start() {
	//TODO 启动 读客户端goroutine
	go c.StartReader()
	//TODO 启动 写客户端goroutine
	go c.StartWriter()

}

func (c *Connection) Stop() {
	fmt.Println("conn stop ", c.ConnID)
	if c.isClosed {
		return
	}
	c.isClosed = true
	_ = c.Conn.Close()
	//告知 client-handle writer 关闭
	c.ExitChan <- true
	//回收资源
	close(c.ExitChan)
	close(c.MsgChan)
}

func (c *Connection) GetTcpConn() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint64 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func NewConnection(conn *net.TCPConn, connID uint64, msgHandle ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		MsgChan:    make(chan []byte),
		MsgHandler: msgHandle,
	}
}