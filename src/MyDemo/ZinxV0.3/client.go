package main
/*
	模拟客户端
 */
import (
	"fmt"
	"net"
	"time"
)

func main(){
	fmt.Println("client starting . . . . . ")
	conn, e := net.Dial("tcp", "127.0.0.1:8999")
	if e != nil {
		fmt.Println("net dial err",e.Error())
		return
	}
	//client每隔一段时间，发送信息, 接受服务器的回写信息
	for {
		_, e := conn.Write([]byte(time.Now().String()))
		if e != nil {
			fmt.Println("tcpConn write err",e.Error())
			return
		}
		//读取 client-handle writer内容
		buff := make([]byte, 512)
		cnt,e := conn.Read(buff)
		if e!= nil{
			fmt.Print("接受服务器的回写信息 err",e.Error())
			return
		}
		fmt.Printf("[client-handle wirter] 的内容: %s \n",buff[:cnt])
		//阻塞释放资源，避免无限循环
		time.Sleep(1*time.Second)
	}

}