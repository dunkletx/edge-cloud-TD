// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dynamic-location-group.proto

package distributed_match_engine

import (
	context "context"
	"encoding/json"
	"errors"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	"github.com/mobiledgex/edge-cloud/util"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	"strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Need acknowledgement
type DlgMessage_DlgAck int32

const (
	DlgMessage_DLG_ACK_EACH_MESSAGE    DlgMessage_DlgAck = 0
	DlgMessage_DLG_ASY_EVERY_N_MESSAGE DlgMessage_DlgAck = 1
	DlgMessage_DLG_NO_ACK              DlgMessage_DlgAck = 2
)

var DlgMessage_DlgAck_name = map[int32]string{
	0: "DLG_ACK_EACH_MESSAGE",
	1: "DLG_ASY_EVERY_N_MESSAGE",
	2: "DLG_NO_ACK",
}

var DlgMessage_DlgAck_value = map[string]int32{
	"DLG_ACK_EACH_MESSAGE":    0,
	"DLG_ASY_EVERY_N_MESSAGE": 1,
	"DLG_NO_ACK":              2,
}

func (x DlgMessage_DlgAck) String() string {
	return proto.EnumName(DlgMessage_DlgAck_name, int32(x))
}

func (DlgMessage_DlgAck) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75937b7725a23625, []int{0, 0}
}

type DlgMessage struct {
	Ver uint32 `protobuf:"varint,1,opt,name=ver,proto3" json:"ver,omitempty"`
	// Dynamic Location Group Id
	LgId uint64 `protobuf:"varint,2,opt,name=lg_id,json=lgId,proto3" json:"lg_id,omitempty"`
	// Group Cookie if secure
	GroupCookie string `protobuf:"bytes,3,opt,name=group_cookie,json=groupCookie,proto3" json:"group_cookie,omitempty"`
	// Message ID
	MessageId uint64            `protobuf:"varint,4,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	AckType   DlgMessage_DlgAck `protobuf:"varint,5,opt,name=ack_type,json=ackType,proto3,enum=distributed_match_engine.DlgMessage_DlgAck" json:"ack_type,omitempty"`
	// Message
	Message              string   `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DlgMessage) Reset()         { *m = DlgMessage{} }
func (m *DlgMessage) String() string { return proto.CompactTextString(m) }
func (*DlgMessage) ProtoMessage()    {}
func (*DlgMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_75937b7725a23625, []int{0}
}
func (m *DlgMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DlgMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DlgMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DlgMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DlgMessage.Merge(m, src)
}
func (m *DlgMessage) XXX_Size() int {
	return m.Size()
}
func (m *DlgMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DlgMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DlgMessage proto.InternalMessageInfo

type DlgReply struct {
	Ver uint32 `protobuf:"varint,1,opt,name=ver,proto3" json:"ver,omitempty"`
	// AckId
	AckId uint64 `protobuf:"varint,2,opt,name=ack_id,json=ackId,proto3" json:"ack_id,omitempty"`
	// Group Cookie for Secure comm
	GroupCookie          string   `protobuf:"bytes,3,opt,name=group_cookie,json=groupCookie,proto3" json:"group_cookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DlgReply) Reset()         { *m = DlgReply{} }
func (m *DlgReply) String() string { return proto.CompactTextString(m) }
func (*DlgReply) ProtoMessage()    {}
func (*DlgReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_75937b7725a23625, []int{1}
}
func (m *DlgReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DlgReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DlgReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DlgReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DlgReply.Merge(m, src)
}
func (m *DlgReply) XXX_Size() int {
	return m.Size()
}
func (m *DlgReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DlgReply.DiscardUnknown(m)
}

var xxx_messageInfo_DlgReply proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("distributed_match_engine.DlgMessage_DlgAck", DlgMessage_DlgAck_name, DlgMessage_DlgAck_value)
	proto.RegisterType((*DlgMessage)(nil), "distributed_match_engine.DlgMessage")
	proto.RegisterType((*DlgReply)(nil), "distributed_match_engine.DlgReply")
}

func init() { proto.RegisterFile("dynamic-location-group.proto", fileDescriptor_75937b7725a23625) }

var fileDescriptor_75937b7725a23625 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdf, 0xee, 0xd2, 0x30,
	0x18, 0xa5, 0xfc, 0x60, 0xc0, 0x87, 0x92, 0xa5, 0x6a, 0x6c, 0xfc, 0xb3, 0xcc, 0xc5, 0x8b, 0x25,
	0x06, 0x2e, 0xf0, 0x09, 0x26, 0x9b, 0x48, 0xe4, 0x4f, 0xb2, 0x11, 0x12, 0xae, 0x9a, 0xd1, 0x35,
	0xb3, 0xd9, 0x58, 0x97, 0x31, 0x4c, 0xf6, 0x46, 0x3e, 0x0a, 0x97, 0x3e, 0x82, 0xf2, 0x24, 0x66,
	0x05, 0xf1, 0x46, 0x0d, 0x77, 0x5f, 0xcf, 0xf9, 0xce, 0xc9, 0xe9, 0x69, 0xe1, 0x55, 0x54, 0x65,
	0xe1, 0x5e, 0xb0, 0x61, 0x2a, 0x59, 0x58, 0x0a, 0x99, 0x0d, 0xe3, 0x42, 0x1e, 0xf3, 0x51, 0x5e,
	0xc8, 0x52, 0x62, 0x12, 0x89, 0x43, 0x59, 0x88, 0xdd, 0xb1, 0xe4, 0x11, 0xdd, 0x87, 0x25, 0xfb,
	0x42, 0x79, 0x16, 0x8b, 0x8c, 0x5b, 0xdf, 0x9a, 0x00, 0x6e, 0x1a, 0x2f, 0xf8, 0xe1, 0x10, 0xc6,
	0x1c, 0xeb, 0xf0, 0xf0, 0x95, 0x17, 0x04, 0x99, 0xc8, 0x7e, 0xec, 0xd7, 0x23, 0x7e, 0x02, 0xed,
	0x34, 0xa6, 0x22, 0x22, 0x4d, 0x13, 0xd9, 0x2d, 0xbf, 0x95, 0xc6, 0xb3, 0x08, 0xbf, 0x81, 0x47,
	0xca, 0x9e, 0x32, 0x29, 0x13, 0xc1, 0xc9, 0x83, 0x89, 0xec, 0x9e, 0xdf, 0x57, 0xd8, 0x44, 0x41,
	0xf8, 0x35, 0xc0, 0xfe, 0x62, 0x5a, 0x8b, 0x5b, 0x4a, 0xdc, 0xbb, 0x22, 0xb3, 0x08, 0x7f, 0x84,
	0x6e, 0xc8, 0x12, 0x5a, 0x56, 0x39, 0x27, 0x6d, 0x13, 0xd9, 0x83, 0xf1, 0xbb, 0xd1, 0xbf, 0x42,
	0x8e, 0xfe, 0x04, 0xac, 0x47, 0x87, 0x25, 0x7e, 0x27, 0x64, 0xc9, 0xba, 0xca, 0x39, 0x26, 0xd0,
	0xb9, 0x9a, 0x12, 0x4d, 0x85, 0xf8, 0x7d, 0xb4, 0x56, 0xa0, 0x5d, 0x96, 0x31, 0x81, 0xa7, 0xee,
	0x7c, 0x4a, 0x9d, 0xc9, 0x67, 0xea, 0x39, 0x93, 0x4f, 0x74, 0xe1, 0x05, 0x81, 0x33, 0xf5, 0xf4,
	0x06, 0x7e, 0x09, 0xcf, 0x15, 0x13, 0x6c, 0xa9, 0xb7, 0xf1, 0xfc, 0x2d, 0x5d, 0xde, 0x48, 0x84,
	0x07, 0x00, 0x35, 0xb9, 0x5c, 0xd5, 0x4a, 0xbd, 0x69, 0x6d, 0xa0, 0xeb, 0xa6, 0xb1, 0xcf, 0xf3,
	0xb4, 0xfa, 0x4b, 0x4f, 0xcf, 0x40, 0xab, 0x2f, 0x74, 0x2b, 0xaa, 0x1d, 0xb2, 0xe4, 0xae, 0xa6,
	0xc6, 0x12, 0xb0, 0x7b, 0x79, 0xbc, 0xb9, 0x64, 0xd3, 0x9a, 0x70, 0x72, 0x81, 0xb7, 0xd0, 0x0f,
	0x78, 0x16, 0xad, 0xa5, 0x42, 0xf0, 0xdb, 0x7b, 0xda, 0x79, 0x61, 0xfd, 0x77, 0x4b, 0x45, 0xb7,
	0x1a, 0x1f, 0xf4, 0xd3, 0x4f, 0xa3, 0x71, 0x3a, 0x1b, 0xe8, 0xfb, 0xd9, 0x40, 0x3f, 0xce, 0x06,
	0xda, 0x69, 0xea, 0x9b, 0xbc, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x5c, 0x4f, 0x8d, 0x46,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DynamicLocGroupApiClient is the client API for DynamicLocGroupApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DynamicLocGroupApiClient interface {
	SendToGroup(ctx context.Context, in *DlgMessage, opts ...grpc.CallOption) (*DlgReply, error)
}

type dynamicLocGroupApiClient struct {
	cc *grpc.ClientConn
}

func NewDynamicLocGroupApiClient(cc *grpc.ClientConn) DynamicLocGroupApiClient {
	return &dynamicLocGroupApiClient{cc}
}

func (c *dynamicLocGroupApiClient) SendToGroup(ctx context.Context, in *DlgMessage, opts ...grpc.CallOption) (*DlgReply, error) {
	out := new(DlgReply)
	err := c.cc.Invoke(ctx, "/distributed_match_engine.DynamicLocGroupApi/SendToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DynamicLocGroupApiServer is the server API for DynamicLocGroupApi service.
type DynamicLocGroupApiServer interface {
	SendToGroup(context.Context, *DlgMessage) (*DlgReply, error)
}

// UnimplementedDynamicLocGroupApiServer can be embedded to have forward compatible implementations.
type UnimplementedDynamicLocGroupApiServer struct {
}

func (*UnimplementedDynamicLocGroupApiServer) SendToGroup(ctx context.Context, req *DlgMessage) (*DlgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToGroup not implemented")
}

func RegisterDynamicLocGroupApiServer(s *grpc.Server, srv DynamicLocGroupApiServer) {
	s.RegisterService(&_DynamicLocGroupApi_serviceDesc, srv)
}

func _DynamicLocGroupApi_SendToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DlgMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicLocGroupApiServer).SendToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributed_match_engine.DynamicLocGroupApi/SendToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicLocGroupApiServer).SendToGroup(ctx, req.(*DlgMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _DynamicLocGroupApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "distributed_match_engine.DynamicLocGroupApi",
	HandlerType: (*DynamicLocGroupApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendToGroup",
			Handler:    _DynamicLocGroupApi_SendToGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dynamic-location-group.proto",
}

func (m *DlgMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DlgMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DlgMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintDynamicLocationGroup(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x32
	}
	if m.AckType != 0 {
		i = encodeVarintDynamicLocationGroup(dAtA, i, uint64(m.AckType))
		i--
		dAtA[i] = 0x28
	}
	if m.MessageId != 0 {
		i = encodeVarintDynamicLocationGroup(dAtA, i, uint64(m.MessageId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.GroupCookie) > 0 {
		i -= len(m.GroupCookie)
		copy(dAtA[i:], m.GroupCookie)
		i = encodeVarintDynamicLocationGroup(dAtA, i, uint64(len(m.GroupCookie)))
		i--
		dAtA[i] = 0x1a
	}
	if m.LgId != 0 {
		i = encodeVarintDynamicLocationGroup(dAtA, i, uint64(m.LgId))
		i--
		dAtA[i] = 0x10
	}
	if m.Ver != 0 {
		i = encodeVarintDynamicLocationGroup(dAtA, i, uint64(m.Ver))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DlgReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DlgReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DlgReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.GroupCookie) > 0 {
		i -= len(m.GroupCookie)
		copy(dAtA[i:], m.GroupCookie)
		i = encodeVarintDynamicLocationGroup(dAtA, i, uint64(len(m.GroupCookie)))
		i--
		dAtA[i] = 0x1a
	}
	if m.AckId != 0 {
		i = encodeVarintDynamicLocationGroup(dAtA, i, uint64(m.AckId))
		i--
		dAtA[i] = 0x10
	}
	if m.Ver != 0 {
		i = encodeVarintDynamicLocationGroup(dAtA, i, uint64(m.Ver))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDynamicLocationGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovDynamicLocationGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DlgMessage) CopyInFields(src *DlgMessage) int {
	changed := 0
	if m.Ver != src.Ver {
		m.Ver = src.Ver
		changed++
	}
	if m.LgId != src.LgId {
		m.LgId = src.LgId
		changed++
	}
	if m.GroupCookie != src.GroupCookie {
		m.GroupCookie = src.GroupCookie
		changed++
	}
	if m.MessageId != src.MessageId {
		m.MessageId = src.MessageId
		changed++
	}
	if m.AckType != src.AckType {
		m.AckType = src.AckType
		changed++
	}
	if m.Message != src.Message {
		m.Message = src.Message
		changed++
	}
	return changed
}

func (m *DlgMessage) DeepCopyIn(src *DlgMessage) {
	m.Ver = src.Ver
	m.LgId = src.LgId
	m.GroupCookie = src.GroupCookie
	m.MessageId = src.MessageId
	m.AckType = src.AckType
	m.Message = src.Message
}

// Helper method to check that enums have valid values
func (m *DlgMessage) ValidateEnums() error {
	if _, ok := DlgMessage_DlgAck_name[int32(m.AckType)]; !ok {
		return errors.New("invalid AckType")
	}
	return nil
}

func (s *DlgMessage) ClearTagged(tags map[string]struct{}) {
}

func (m *DlgReply) CopyInFields(src *DlgReply) int {
	changed := 0
	if m.Ver != src.Ver {
		m.Ver = src.Ver
		changed++
	}
	if m.AckId != src.AckId {
		m.AckId = src.AckId
		changed++
	}
	if m.GroupCookie != src.GroupCookie {
		m.GroupCookie = src.GroupCookie
		changed++
	}
	return changed
}

func (m *DlgReply) DeepCopyIn(src *DlgReply) {
	m.Ver = src.Ver
	m.AckId = src.AckId
	m.GroupCookie = src.GroupCookie
}

// Helper method to check that enums have valid values
func (m *DlgReply) ValidateEnums() error {
	return nil
}

func (s *DlgReply) ClearTagged(tags map[string]struct{}) {
}

var DlgAckStrings = []string{
	"DLG_ACK_EACH_MESSAGE",
	"DLG_ASY_EVERY_N_MESSAGE",
	"DLG_NO_ACK",
}

const (
	DlgAckDLG_ACK_EACH_MESSAGE    uint64 = 1 << 0
	DlgAckDLG_ASY_EVERY_N_MESSAGE uint64 = 1 << 1
	DlgAckDLG_NO_ACK              uint64 = 1 << 2
)

var DlgMessage_DlgAck_CamelName = map[int32]string{
	// DLG_ACK_EACH_MESSAGE -> DlgAckEachMessage
	0: "DlgAckEachMessage",
	// DLG_ASY_EVERY_N_MESSAGE -> DlgAsyEveryNMessage
	1: "DlgAsyEveryNMessage",
	// DLG_NO_ACK -> DlgNoAck
	2: "DlgNoAck",
}
var DlgMessage_DlgAck_CamelValue = map[string]int32{
	"DlgAckEachMessage":   0,
	"DlgAsyEveryNMessage": 1,
	"DlgNoAck":            2,
}

func ParseDlgMessage_DlgAck(data interface{}) (DlgMessage_DlgAck, error) {
	if val, ok := data.(DlgMessage_DlgAck); ok {
		return val, nil
	} else if str, ok := data.(string); ok {
		val, ok := DlgMessage_DlgAck_CamelValue[util.CamelCase(str)]
		if !ok {
			// may have omitted common prefix
			val, ok = DlgMessage_DlgAck_CamelValue["Dlg"+util.CamelCase(str)]
		}
		if !ok {
			// may be int value instead of enum name
			ival, err := strconv.Atoi(str)
			val = int32(ival)
			if err == nil {
				_, ok = DlgMessage_DlgAck_CamelName[val]
			}
		}
		if !ok {
			return DlgMessage_DlgAck(0), fmt.Errorf("Invalid DlgMessage_DlgAck value %q", str)
		}
		return DlgMessage_DlgAck(val), nil
	} else if ival, ok := data.(int32); ok {
		if _, ok := DlgMessage_DlgAck_CamelName[ival]; ok {
			return DlgMessage_DlgAck(ival), nil
		} else {
			return DlgMessage_DlgAck(0), fmt.Errorf("Invalid DlgMessage_DlgAck value %d", ival)
		}
	}
	return DlgMessage_DlgAck(0), fmt.Errorf("Invalid DlgMessage_DlgAck value %v", data)
}

func (e *DlgMessage_DlgAck) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	err := unmarshal(&str)
	if err != nil {
		return err
	}
	val, err := ParseDlgMessage_DlgAck(str)
	if err != nil {
		return err
	}
	*e = val
	return nil
}

func (e DlgMessage_DlgAck) MarshalYAML() (interface{}, error) {
	str := proto.EnumName(DlgMessage_DlgAck_CamelName, int32(e))
	str = strings.TrimPrefix(str, "Dlg")
	return str, nil
}

// custom JSON encoding/decoding
func (e *DlgMessage_DlgAck) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err == nil {
		val, err := ParseDlgMessage_DlgAck(str)
		if err != nil {
			return &json.UnmarshalTypeError{
				Value: "string " + str,
				Type:  reflect.TypeOf(DlgMessage_DlgAck(0)),
			}
		}
		*e = DlgMessage_DlgAck(val)
		return nil
	}
	var ival int32
	err = json.Unmarshal(b, &ival)
	if err == nil {
		val, err := ParseDlgMessage_DlgAck(ival)
		if err == nil {
			*e = val
			return nil
		}
	}
	return &json.UnmarshalTypeError{
		Value: "value " + string(b),
		Type:  reflect.TypeOf(DlgMessage_DlgAck(0)),
	}
}

func (e DlgMessage_DlgAck) MarshalJSON() ([]byte, error) {
	str := proto.EnumName(DlgMessage_DlgAck_CamelName, int32(e))
	str = strings.TrimPrefix(str, "Dlg")
	return json.Marshal(str)
}

var DlgAckCommonPrefix = "Dlg"

func (m *DlgMessage) IsValidArgsForSendToGroup() error {
	return nil
}

func (m *DlgMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ver != 0 {
		n += 1 + sovDynamicLocationGroup(uint64(m.Ver))
	}
	if m.LgId != 0 {
		n += 1 + sovDynamicLocationGroup(uint64(m.LgId))
	}
	l = len(m.GroupCookie)
	if l > 0 {
		n += 1 + l + sovDynamicLocationGroup(uint64(l))
	}
	if m.MessageId != 0 {
		n += 1 + sovDynamicLocationGroup(uint64(m.MessageId))
	}
	if m.AckType != 0 {
		n += 1 + sovDynamicLocationGroup(uint64(m.AckType))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovDynamicLocationGroup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DlgReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ver != 0 {
		n += 1 + sovDynamicLocationGroup(uint64(m.Ver))
	}
	if m.AckId != 0 {
		n += 1 + sovDynamicLocationGroup(uint64(m.AckId))
	}
	l = len(m.GroupCookie)
	if l > 0 {
		n += 1 + l + sovDynamicLocationGroup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDynamicLocationGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDynamicLocationGroup(x uint64) (n int) {
	return sovDynamicLocationGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DlgMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicLocationGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DlgMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DlgMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ver", wireType)
			}
			m.Ver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ver |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LgId", wireType)
			}
			m.LgId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LgId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupCookie", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDynamicLocationGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicLocationGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupCookie = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			m.MessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckType", wireType)
			}
			m.AckType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AckType |= DlgMessage_DlgAck(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDynamicLocationGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicLocationGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicLocationGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicLocationGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicLocationGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DlgReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicLocationGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DlgReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DlgReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ver", wireType)
			}
			m.Ver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ver |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckId", wireType)
			}
			m.AckId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AckId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupCookie", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDynamicLocationGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicLocationGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupCookie = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicLocationGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicLocationGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicLocationGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDynamicLocationGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDynamicLocationGroup
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDynamicLocationGroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDynamicLocationGroup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDynamicLocationGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDynamicLocationGroup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDynamicLocationGroup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDynamicLocationGroup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDynamicLocationGroup = fmt.Errorf("proto: unexpected end of group")
)
