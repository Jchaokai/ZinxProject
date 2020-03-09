package main

/*
	模拟客户端
*/
import (
	"ZinxProject/src/zinx/znet"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("client starting . . . . . ")
	conn, e := net.Dial("tcp", "127.0.0.1:8999")
	if e != nil {
		fmt.Println("net dial err", e.Error())
		return
	}
	//client每隔一段时间，发送信息, 接受服务器的回写信息
	for {
		//发送封装过的 数据
		dp := znet.NewDataPack()
		bytes, e := dp.Pack(znet.NewMsg(0, []byte("client发送[msgID=0]的内容")))
		if e != nil {
			fmt.Println("data-pack error", e)
			return
		}
		if _, e := conn.Write(bytes); e != nil {
			fmt.Println("conn write error", e)
			return
		}
		//将 client-handle writer 打印出来
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, e := io.ReadFull(conn, binaryHead); e != nil {
			fmt.Println("将 client-handle writer 打印出来 err :",e)
			return
		}
		message, e := dp.UnPack(binaryHead)
		if e != nil {
			fmt.Println("dataPack unpack error : ",e)
			return
		}
		var data []byte
		if message.GetMsgLen() >0 {
			data = make([]byte,message.GetMsgLen())
			if _, e := io.ReadFull(conn, data);e != nil {
				fmt.Println("read msg data error: ",e)
				return
			}
		}
		message.SetData(data)
		fmt.Printf("client-handle writer : [msgID %d] [data %s] \n",message.GetMsgID(),message.GetMsgData())
		//阻塞释放资源，避免无限循环
		time.Sleep(2 * time.Second)
	}

}
