package znet

type Message struct {
	Id      uint64
	DataLen uint64
	Data    []byte
}

func NewMsg(id uint64,data []byte) *Message{
	return &Message{
		Id:      id,
		DataLen: uint64(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMsgID() uint64 {
	return m.Id
}

func (m *Message) GetMsgLen() uint64 {
	return m.DataLen
}

func (m *Message) GetMsgData() []byte {
	return m.Data
}

func (m *Message) SetMsgID(id uint64) {
	m.Id = id
}

func (m *Message) SetMsgLen(len uint64) {
	m.DataLen = len
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}


