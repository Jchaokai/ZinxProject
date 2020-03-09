package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//测试DataPack 拆包 封包测试
func TestDataPack(t *testing.T) {
	/*
		模拟的服务器
	*/
	listener, e := net.Listen("tcp", "127.0.0.1:7777")
	if e != nil {
		fmt.Println("server listen error:", e)
		return
	}
	// go - 从client读取数据，拆包处理
	go func() {
		for {
			conn, i := listener.Accept()
			if i != nil {
				fmt.Println("server accept error:", i)
				continue
			}

			go func(conn net.Conn) {
				/*
					拆包 :
					1. 读出包头二进制(16个字节)
					2. 根据包头信息 继续读取conn数据得到包体内容
				*/
				dataPack := NewDataPack()
				for {
					//1.
					packHeadBinary := make([]byte, dataPack.GetHeadLen())
					if _, err := io.ReadFull(conn, packHeadBinary); err != nil {
						fmt.Println("读包头二进制 error", err)
						break
					}
					//2.
					/*	MsgWithNilData :
							Id      	--包体数据id
							DataLen 	--包体数据长度
							data		--包体数据，此时nil
					}*/
					MsgWithNilData, e := dataPack.UnPack(packHeadBinary)
					if e != nil {
						fmt.Println("解析包头信息 error:", e)
						return
					}
					if MsgWithNilData.GetMsgLen() > 0 {
						Msg := MsgWithNilData.(*Message)
						Msg.Data = make([]byte, Msg.GetMsgLen())
						//根据DataLen，继续读取conn 为包体内容
						if _, e := io.ReadFull(conn, Msg.Data); e != nil {
							fmt.Println("读取包体数据 err")
							return
						}
						//print完整的msg
						fmt.Printf("---> Client-Handle Reader得到 Msg [ dataLen: %d | ID : %d | data : %s]\n",
							Msg.DataLen, Msg.Id, string(Msg.Data))
					}
				}
			}(conn)
		}
	}()


	/*
		模拟客户端
	*/
	conn, e := net.Dial("tcp", "127.0.0.1:7777")
	if e != nil {
		fmt.Println("client dial error: ",e)
		return
	}
	dp := NewDataPack()

	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z','i','n','x'},
	}
	msg1Bytes, e := dp.Pack(msg1)
	if e != nil {
		fmt.Println("pack msg1 error:",e)
		return
	}
	msg2 := &Message{
		Id:      1,
		DataLen: 7,
		Data:    []byte{'q','u','i','t','j','c','k'},
	}
	msg2Bytes, e := dp.Pack(msg2)
	if e != nil {
		fmt.Println("pack msg2 error:",e)
		return
	}
	//模拟粘包过程，封装两个msg一同发送
	msg1Bytes = append(msg1Bytes,msg2Bytes...)
	_, e = conn.Write(msg1Bytes)
	if e != nil {
		fmt.Println("client write data error :",e)
		return
	}
	//阻塞，等待server处理
	select {}
}
