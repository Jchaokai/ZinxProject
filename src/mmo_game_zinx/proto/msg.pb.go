// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//同步客户端玩家ID
type SyncPid struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPid) Reset()         { *m = SyncPid{} }
func (m *SyncPid) String() string { return proto.CompactTextString(m) }
func (*SyncPid) ProtoMessage()    {}
func (*SyncPid) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *SyncPid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPid.Unmarshal(m, b)
}
func (m *SyncPid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPid.Marshal(b, m, deterministic)
}
func (m *SyncPid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPid.Merge(m, src)
}
func (m *SyncPid) XXX_Size() int {
	return xxx_messageInfo_SyncPid.Size(m)
}
func (m *SyncPid) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPid.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPid proto.InternalMessageInfo

func (m *SyncPid) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

//玩家位置
type Position struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
	V                    float32  `protobuf:"fixed32,4,opt,name=V,proto3" json:"V,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Position) GetV() float32 {
	if m != nil {
		return m.V
	}
	return 0
}

//玩家广播数据
type BroadCast struct {
	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Tp  int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*BroadCast_Content
	//	*BroadCast_P
	//	*BroadCast_ActionData
	Data                 isBroadCast_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BroadCast) Reset()         { *m = BroadCast{} }
func (m *BroadCast) String() string { return proto.CompactTextString(m) }
func (*BroadCast) ProtoMessage()    {}
func (*BroadCast) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *BroadCast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadCast.Unmarshal(m, b)
}
func (m *BroadCast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadCast.Marshal(b, m, deterministic)
}
func (m *BroadCast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadCast.Merge(m, src)
}
func (m *BroadCast) XXX_Size() int {
	return xxx_messageInfo_BroadCast.Size(m)
}
func (m *BroadCast) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadCast.DiscardUnknown(m)
}

var xxx_messageInfo_BroadCast proto.InternalMessageInfo

func (m *BroadCast) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *BroadCast) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

type isBroadCast_Data interface {
	isBroadCast_Data()
}

type BroadCast_Content struct {
	Content string `protobuf:"bytes,3,opt,name=Content,proto3,oneof"`
}

type BroadCast_P struct {
	P *Position `protobuf:"bytes,4,opt,name=P,proto3,oneof"`
}

type BroadCast_ActionData struct {
	ActionData int32 `protobuf:"varint,5,opt,name=ActionData,proto3,oneof"`
}

func (*BroadCast_Content) isBroadCast_Data() {}

func (*BroadCast_P) isBroadCast_Data() {}

func (*BroadCast_ActionData) isBroadCast_Data() {}

func (m *BroadCast) GetData() isBroadCast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BroadCast) GetContent() string {
	if x, ok := m.GetData().(*BroadCast_Content); ok {
		return x.Content
	}
	return ""
}

func (m *BroadCast) GetP() *Position {
	if x, ok := m.GetData().(*BroadCast_P); ok {
		return x.P
	}
	return nil
}

func (m *BroadCast) GetActionData() int32 {
	if x, ok := m.GetData().(*BroadCast_ActionData); ok {
		return x.ActionData
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BroadCast) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BroadCast_Content)(nil),
		(*BroadCast_P)(nil),
		(*BroadCast_ActionData)(nil),
	}
}

//玩家聊天数据
type Talk struct {
	Content              string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Talk) Reset()         { *m = Talk{} }
func (m *Talk) String() string { return proto.CompactTextString(m) }
func (*Talk) ProtoMessage()    {}
func (*Talk) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *Talk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Talk.Unmarshal(m, b)
}
func (m *Talk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Talk.Marshal(b, m, deterministic)
}
func (m *Talk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Talk.Merge(m, src)
}
func (m *Talk) XXX_Size() int {
	return xxx_messageInfo_Talk.Size(m)
}
func (m *Talk) XXX_DiscardUnknown() {
	xxx_messageInfo_Talk.DiscardUnknown(m)
}

var xxx_messageInfo_Talk proto.InternalMessageInfo

func (m *Talk) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

//玩家信息
type Player struct {
	Pid                  int32     `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	P                    *Position `protobuf:"bytes,2,opt,name=P,proto3" json:"P,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Player) GetP() *Position {
	if m != nil {
		return m.P
	}
	return nil
}

//同步玩家显示数据
type SyncPlayers struct {
	Ps                   []*Player `protobuf:"bytes,1,rep,name=ps,proto3" json:"ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SyncPlayers) Reset()         { *m = SyncPlayers{} }
func (m *SyncPlayers) String() string { return proto.CompactTextString(m) }
func (*SyncPlayers) ProtoMessage()    {}
func (*SyncPlayers) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}

func (m *SyncPlayers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPlayers.Unmarshal(m, b)
}
func (m *SyncPlayers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPlayers.Marshal(b, m, deterministic)
}
func (m *SyncPlayers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPlayers.Merge(m, src)
}
func (m *SyncPlayers) XXX_Size() int {
	return xxx_messageInfo_SyncPlayers.Size(m)
}
func (m *SyncPlayers) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPlayers.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPlayers proto.InternalMessageInfo

func (m *SyncPlayers) GetPs() []*Player {
	if m != nil {
		return m.Ps
	}
	return nil
}

func init() {
	proto.RegisterType((*SyncPid)(nil), "proto.SyncPid")
	proto.RegisterType((*Position)(nil), "proto.Position")
	proto.RegisterType((*BroadCast)(nil), "proto.BroadCast")
	proto.RegisterType((*Talk)(nil), "proto.Talk")
	proto.RegisterType((*Player)(nil), "proto.Player")
	proto.RegisterType((*SyncPlayers)(nil), "proto.SyncPlayers")
}

func init() {
	proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899)
}

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x63, 0x27, 0x4e, 0xbf, 0xdc, 0x7c, 0xfc, 0x91, 0x27, 0x0b, 0x54, 0x11, 0x79, 0xea,
	0x80, 0x3a, 0x94, 0x89, 0x91, 0x94, 0x21, 0xa3, 0x65, 0xa2, 0xaa, 0xed, 0xe6, 0x36, 0x15, 0x8a,
	0x28, 0xb1, 0x15, 0x7b, 0xe9, 0x7b, 0xf0, 0x14, 0x3c, 0x25, 0xb2, 0xa3, 0x00, 0x12, 0x9d, 0x6e,
	0x7e, 0xf7, 0xc4, 0xe7, 0x1c, 0x1b, 0xb2, 0x77, 0xfb, 0x3a, 0x37, 0xbd, 0x76, 0x9a, 0x92, 0x30,
	0xf8, 0x2d, 0x4c, 0x5e, 0x4e, 0xdd, 0x5e, 0xb4, 0x0d, 0xbd, 0x86, 0x58, 0xb4, 0x0d, 0x43, 0x05,
	0x9a, 0x11, 0xe9, 0x3f, 0x79, 0x09, 0xff, 0x84, 0xb6, 0xad, 0x6b, 0x75, 0x47, 0xff, 0x03, 0x5a,
	0x07, 0x0d, 0x4b, 0xb4, 0xf6, 0xb4, 0x61, 0x78, 0xa0, 0x8d, 0xa7, 0x2d, 0x8b, 0x07, 0xda, 0x7a,
	0x5a, 0xb1, 0x64, 0xa0, 0x15, 0xff, 0x40, 0x90, 0x95, 0xbd, 0x56, 0xcd, 0x52, 0x59, 0xf7, 0x37,
	0x83, 0x5e, 0x02, 0xae, 0x4d, 0xb0, 0x22, 0x12, 0xd7, 0x86, 0xde, 0xc0, 0x64, 0xa9, 0x3b, 0x77,
	0xe8, 0x5c, 0x70, 0xcc, 0xaa, 0x48, 0x8e, 0x0b, 0x7a, 0x07, 0x48, 0x04, 0xe7, 0x7c, 0x71, 0x35,
	0x5c, 0x63, 0x3e, 0xf6, 0xab, 0x22, 0x89, 0x04, 0x2d, 0x00, 0x9e, 0xf6, 0x1e, 0x9f, 0x95, 0x53,
	0x8c, 0x78, 0xd3, 0x2a, 0x92, 0xbf, 0x76, 0x65, 0x0a, 0x89, 0x9f, 0xbc, 0x80, 0xa4, 0x56, 0xc7,
	0x37, 0xca, 0x7e, 0xe2, 0x7c, 0xa9, 0xec, 0x3b, 0x8c, 0x3f, 0x42, 0x2a, 0x8e, 0xea, 0x74, 0xe8,
	0xcf, 0x94, 0x9e, 0xfa, 0x22, 0xf8, 0x6c, 0x11, 0x89, 0x04, 0xbf, 0x87, 0x3c, 0x3c, 0x6a, 0x38,
	0x6e, 0xe9, 0x14, 0xb0, 0xb1, 0x0c, 0x15, 0xf1, 0x2c, 0x5f, 0x5c, 0x8c, 0xbf, 0x07, 0x4d, 0x62,
	0x63, 0x4b, 0xf2, 0x89, 0xb1, 0xd8, 0xed, 0xd2, 0x20, 0x3c, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xbe, 0x20, 0xb2, 0xad, 0xa4, 0x01, 0x00, 0x00,
}
