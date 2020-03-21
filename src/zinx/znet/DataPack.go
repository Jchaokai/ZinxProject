package znet

import (
	"ZinxProject/src/zinx/utils"
	"ZinxProject/src/zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

//包头  8个字节，DataLen uint32 (4个) DataID uint32 (4个)
func (dp *DataPack) GetHeadLen() uint {
	return 8
}

func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	/*  将包头信息 写入buffer
	1. dataLen
	2. msgID
	3. data
	*/
	if e := binary.Write(buffer, binary.LittleEndian, msg.GetMsgLen()); e != nil {
		return nil, e
	}
	if e := binary.Write(buffer, binary.LittleEndian, msg.GetMsgID()); e != nil {
		return nil, e
	}
	if e := binary.Write(buffer, binary.LittleEndian, msg.GetMsgData()); e != nil {
		return nil, e
	}

	return buffer.Bytes(), nil
}

//UnPack返回的msg  [ dataLen | id  |  nil ]  （之后用户自己根据dataLen 继续读取conn的数据）
func (dp *DataPack) UnPack(binaryBytes []byte) (ziface.IMessage, error) {
	reader := bytes.NewReader(binaryBytes)
	msg := &Message{}
	//读 dataLen
	if e := binary.Read(reader, binary.LittleEndian, &msg.DataLen); e != nil {
		return nil, e
	}
	//读 msgID
	if e := binary.Read(reader, binary.LittleEndian, &msg.Id); e != nil {
		return nil, e
	}
	//判断 dataLen 是否超过了 utils.GlobalObject.MaxPackageSize
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("[too large msg data . . .]")
	}
	//只含包头信息的msg [ dataLen | id | nil]
	return msg, nil
}
