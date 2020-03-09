package znet

import (
	"ZinxProject/src/zinx/utils"
	"ZinxProject/src/zinx/ziface"
	"fmt"
	"net"
	"time"
)

type Server struct {
	Name      string
	Version   string
	IPVersion string
	IP        string
	Port      int
	//添加一个router map，不同msgID对应不同处理业务
	MsgHandler ziface.IMsgHandler
}

func (s *Server) Start() {
	fmt.Printf("%s %s server Listenning  at IP:%s , port:%d  "+
		"[starting . . .]\n   [maxConn] : %d\n   [maxPackageSize] : %d \n",
		s.Name, s.Version, s.IP, s.Port, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	time.Sleep(1 * time.Second)
	go func() {
		//server 开启工作池 TaskQueue
		s.MsgHandler.StartWorkerPool()

		//获取tcp的addr
		tcpAddr, e := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if e != nil {
			fmt.Println("获取tcp的addr", e.Error())
			return
		}
		//监听服务器的addr
		tcpListener, e := net.ListenTCP(s.IPVersion, tcpAddr)
		if e != nil {
			fmt.Println("监听服务器的addr", e.Error())
			return
		}
		fmt.Println(s.Name, " [start success !!!]")

		var connID uint64
		connID = 0
		//阻塞等待 客户端连接 处理一些业务
		for {
			tcpConn, e := tcpListener.AcceptTCP()
			if e != nil {
				fmt.Println("tcp conn ", e.Error())
				continue
			}
			connection := NewConnection(tcpConn, connID, s.MsgHandler)
			connID++
			go connection.Start()
		}
	}()
}
func (s *Server) Stop() {
	//ToDo 将服务器资源，信息....关闭
}

//不暴露 start stop 给用户
func (s *Server) Server() {
	//启动server基本功能
	s.Start()
	//ToDo 做一些额外的功能
	//一直阻塞
	select {}
}

func (s *Server) AddRouter(msgID uint64, r ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID,r)
}

//初始化 server模块的func
func NewServer() ziface.IServer {
	return &Server{
		Name:       utils.GlobalObject.Name,
		Version:    utils.GlobalObject.Version,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.Port,
		MsgHandler: NewMsgHandler(),
	}
}