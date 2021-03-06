// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mhubapi/mhubapi.proto

package mhubapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskStatus int32

const (
	TaskStatus_NOT_SENT  TaskStatus = 0
	TaskStatus_SENT      TaskStatus = 1
	TaskStatus_RECEIVED  TaskStatus = 2
	TaskStatus_COMPLETED TaskStatus = 3
)

var TaskStatus_name = map[int32]string{
	0: "NOT_SENT",
	1: "SENT",
	2: "RECEIVED",
	3: "COMPLETED",
}
var TaskStatus_value = map[string]int32{
	"NOT_SENT":  0,
	"SENT":      1,
	"RECEIVED":  2,
	"COMPLETED": 3,
}

func (x TaskStatus) String() string {
	return proto.EnumName(TaskStatus_name, int32(x))
}
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mhubapi_40b041ecc178f234, []int{0}
}

type AvailableDSKFStreamInput struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvailableDSKFStreamInput) Reset()         { *m = AvailableDSKFStreamInput{} }
func (m *AvailableDSKFStreamInput) String() string { return proto.CompactTextString(m) }
func (*AvailableDSKFStreamInput) ProtoMessage()    {}
func (*AvailableDSKFStreamInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_mhubapi_40b041ecc178f234, []int{0}
}
func (m *AvailableDSKFStreamInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailableDSKFStreamInput.Unmarshal(m, b)
}
func (m *AvailableDSKFStreamInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailableDSKFStreamInput.Marshal(b, m, deterministic)
}
func (dst *AvailableDSKFStreamInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableDSKFStreamInput.Merge(dst, src)
}
func (m *AvailableDSKFStreamInput) XXX_Size() int {
	return xxx_messageInfo_AvailableDSKFStreamInput.Size(m)
}
func (m *AvailableDSKFStreamInput) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableDSKFStreamInput.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableDSKFStreamInput proto.InternalMessageInfo

type AvailableDSKFStreamOutput struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	DskfFile             string   `protobuf:"bytes,2,opt,name=dskf_file,json=dskfFile,proto3" json:"dskf_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvailableDSKFStreamOutput) Reset()         { *m = AvailableDSKFStreamOutput{} }
func (m *AvailableDSKFStreamOutput) String() string { return proto.CompactTextString(m) }
func (*AvailableDSKFStreamOutput) ProtoMessage()    {}
func (*AvailableDSKFStreamOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_mhubapi_40b041ecc178f234, []int{1}
}
func (m *AvailableDSKFStreamOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailableDSKFStreamOutput.Unmarshal(m, b)
}
func (m *AvailableDSKFStreamOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailableDSKFStreamOutput.Marshal(b, m, deterministic)
}
func (dst *AvailableDSKFStreamOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableDSKFStreamOutput.Merge(dst, src)
}
func (m *AvailableDSKFStreamOutput) XXX_Size() int {
	return xxx_messageInfo_AvailableDSKFStreamOutput.Size(m)
}
func (m *AvailableDSKFStreamOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableDSKFStreamOutput.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableDSKFStreamOutput proto.InternalMessageInfo

func (m *AvailableDSKFStreamOutput) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *AvailableDSKFStreamOutput) GetDskfFile() string {
	if m != nil {
		return m.DskfFile
	}
	return ""
}

type SetTaskStatusInput struct {
	TaskId               string     `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	UserId               string     `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status               TaskStatus `protobuf:"varint,3,opt,name=status,proto3,enum=mhubapi.TaskStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SetTaskStatusInput) Reset()         { *m = SetTaskStatusInput{} }
func (m *SetTaskStatusInput) String() string { return proto.CompactTextString(m) }
func (*SetTaskStatusInput) ProtoMessage()    {}
func (*SetTaskStatusInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_mhubapi_40b041ecc178f234, []int{2}
}
func (m *SetTaskStatusInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTaskStatusInput.Unmarshal(m, b)
}
func (m *SetTaskStatusInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTaskStatusInput.Marshal(b, m, deterministic)
}
func (dst *SetTaskStatusInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTaskStatusInput.Merge(dst, src)
}
func (m *SetTaskStatusInput) XXX_Size() int {
	return xxx_messageInfo_SetTaskStatusInput.Size(m)
}
func (m *SetTaskStatusInput) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTaskStatusInput.DiscardUnknown(m)
}

var xxx_messageInfo_SetTaskStatusInput proto.InternalMessageInfo

func (m *SetTaskStatusInput) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SetTaskStatusInput) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SetTaskStatusInput) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatus_NOT_SENT
}

type DeepPingOutput struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeepPingOutput) Reset()         { *m = DeepPingOutput{} }
func (m *DeepPingOutput) String() string { return proto.CompactTextString(m) }
func (*DeepPingOutput) ProtoMessage()    {}
func (*DeepPingOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_mhubapi_40b041ecc178f234, []int{3}
}
func (m *DeepPingOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeepPingOutput.Unmarshal(m, b)
}
func (m *DeepPingOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeepPingOutput.Marshal(b, m, deterministic)
}
func (dst *DeepPingOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeepPingOutput.Merge(dst, src)
}
func (m *DeepPingOutput) XXX_Size() int {
	return xxx_messageInfo_DeepPingOutput.Size(m)
}
func (m *DeepPingOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_DeepPingOutput.DiscardUnknown(m)
}

var xxx_messageInfo_DeepPingOutput proto.InternalMessageInfo

func (m *DeepPingOutput) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_mhubapi_40b041ecc178f234, []int{4}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (dst *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(dst, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AvailableDSKFStreamInput)(nil), "mhubapi.AvailableDSKFStreamInput")
	proto.RegisterType((*AvailableDSKFStreamOutput)(nil), "mhubapi.AvailableDSKFStreamOutput")
	proto.RegisterType((*SetTaskStatusInput)(nil), "mhubapi.SetTaskStatusInput")
	proto.RegisterType((*DeepPingOutput)(nil), "mhubapi.DeepPingOutput")
	proto.RegisterType((*Void)(nil), "mhubapi.Void")
	proto.RegisterEnum("mhubapi.TaskStatus", TaskStatus_name, TaskStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MicrologProxyHubClient is the client API for MicrologProxyHub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MicrologProxyHubClient interface {
	DeepPing(ctx context.Context, in *Void, opts ...grpc.CallOption) (*DeepPingOutput, error)
	SetTaskStatus(ctx context.Context, in *SetTaskStatusInput, opts ...grpc.CallOption) (*Void, error)
	AvailableDSKFStream(ctx context.Context, in *AvailableDSKFStreamInput, opts ...grpc.CallOption) (MicrologProxyHub_AvailableDSKFStreamClient, error)
}

type micrologProxyHubClient struct {
	cc *grpc.ClientConn
}

func NewMicrologProxyHubClient(cc *grpc.ClientConn) MicrologProxyHubClient {
	return &micrologProxyHubClient{cc}
}

func (c *micrologProxyHubClient) DeepPing(ctx context.Context, in *Void, opts ...grpc.CallOption) (*DeepPingOutput, error) {
	out := new(DeepPingOutput)
	err := c.cc.Invoke(ctx, "/mhubapi.MicrologProxyHub/DeepPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *micrologProxyHubClient) SetTaskStatus(ctx context.Context, in *SetTaskStatusInput, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/mhubapi.MicrologProxyHub/SetTaskStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *micrologProxyHubClient) AvailableDSKFStream(ctx context.Context, in *AvailableDSKFStreamInput, opts ...grpc.CallOption) (MicrologProxyHub_AvailableDSKFStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MicrologProxyHub_serviceDesc.Streams[0], "/mhubapi.MicrologProxyHub/AvailableDSKFStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &micrologProxyHubAvailableDSKFStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MicrologProxyHub_AvailableDSKFStreamClient interface {
	Recv() (*AvailableDSKFStreamOutput, error)
	grpc.ClientStream
}

type micrologProxyHubAvailableDSKFStreamClient struct {
	grpc.ClientStream
}

func (x *micrologProxyHubAvailableDSKFStreamClient) Recv() (*AvailableDSKFStreamOutput, error) {
	m := new(AvailableDSKFStreamOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MicrologProxyHubServer is the server API for MicrologProxyHub service.
type MicrologProxyHubServer interface {
	DeepPing(context.Context, *Void) (*DeepPingOutput, error)
	SetTaskStatus(context.Context, *SetTaskStatusInput) (*Void, error)
	AvailableDSKFStream(*AvailableDSKFStreamInput, MicrologProxyHub_AvailableDSKFStreamServer) error
}

func RegisterMicrologProxyHubServer(s *grpc.Server, srv MicrologProxyHubServer) {
	s.RegisterService(&_MicrologProxyHub_serviceDesc, srv)
}

func _MicrologProxyHub_DeepPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicrologProxyHubServer).DeepPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mhubapi.MicrologProxyHub/DeepPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicrologProxyHubServer).DeepPing(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicrologProxyHub_SetTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTaskStatusInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicrologProxyHubServer).SetTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mhubapi.MicrologProxyHub/SetTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicrologProxyHubServer).SetTaskStatus(ctx, req.(*SetTaskStatusInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicrologProxyHub_AvailableDSKFStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AvailableDSKFStreamInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MicrologProxyHubServer).AvailableDSKFStream(m, &micrologProxyHubAvailableDSKFStreamServer{stream})
}

type MicrologProxyHub_AvailableDSKFStreamServer interface {
	Send(*AvailableDSKFStreamOutput) error
	grpc.ServerStream
}

type micrologProxyHubAvailableDSKFStreamServer struct {
	grpc.ServerStream
}

func (x *micrologProxyHubAvailableDSKFStreamServer) Send(m *AvailableDSKFStreamOutput) error {
	return x.ServerStream.SendMsg(m)
}

var _MicrologProxyHub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mhubapi.MicrologProxyHub",
	HandlerType: (*MicrologProxyHubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeepPing",
			Handler:    _MicrologProxyHub_DeepPing_Handler,
		},
		{
			MethodName: "SetTaskStatus",
			Handler:    _MicrologProxyHub_SetTaskStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AvailableDSKFStream",
			Handler:       _MicrologProxyHub_AvailableDSKFStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mhubapi/mhubapi.proto",
}

func init() { proto.RegisterFile("mhubapi/mhubapi.proto", fileDescriptor_mhubapi_40b041ecc178f234) }

var fileDescriptor_mhubapi_40b041ecc178f234 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x5d, 0x6f, 0xa2, 0x40,
	0x14, 0x15, 0x75, 0x11, 0x6f, 0x56, 0x43, 0xc6, 0xdd, 0xc8, 0xe2, 0x8b, 0x3b, 0x0f, 0x1b, 0xb3,
	0x9b, 0xb8, 0x8d, 0xed, 0x6b, 0x1f, 0x8c, 0x60, 0x4a, 0x5a, 0x3f, 0x2a, 0xc4, 0x57, 0x3b, 0x94,
	0xd1, 0x4e, 0x44, 0x21, 0x30, 0x63, 0xda, 0x9f, 0xdc, 0x7f, 0xd1, 0x80, 0x88, 0x31, 0xd5, 0x3e,
	0xcd, 0x9c, 0x7b, 0x4f, 0xce, 0x39, 0x73, 0xef, 0xc0, 0xcf, 0xcd, 0x8b, 0x70, 0x49, 0xc8, 0xfe,
	0x67, 0x67, 0x37, 0x8c, 0x02, 0x1e, 0xa0, 0x4a, 0x06, 0xb1, 0x0e, 0x5a, 0x7f, 0x47, 0x98, 0x4f,
	0x5c, 0x9f, 0x1a, 0xf6, 0xfd, 0xd0, 0xe6, 0x11, 0x25, 0x1b, 0x6b, 0x1b, 0x0a, 0x8e, 0x1f, 0xe1,
	0xd7, 0x99, 0xde, 0x44, 0xf0, 0x50, 0x70, 0xd4, 0x84, 0x0a, 0x27, 0xf1, 0x7a, 0xc1, 0x3c, 0x4d,
	0x6a, 0x4b, 0x9d, 0xea, 0x4c, 0x4e, 0xa0, 0xe5, 0xa1, 0x16, 0x54, 0xbd, 0x78, 0xbd, 0x5c, 0x2c,
	0x99, 0x4f, 0xb5, 0x62, 0xda, 0x52, 0x92, 0xc2, 0x90, 0xf9, 0x14, 0xc7, 0x80, 0x6c, 0xca, 0x1d,
	0x12, 0xaf, 0x6d, 0x4e, 0xb8, 0x88, 0x53, 0xa3, 0xcb, 0x5a, 0x4d, 0xa8, 0x88, 0x98, 0x46, 0x49,
	0x63, 0xaf, 0x24, 0x27, 0xd0, 0xf2, 0xd0, 0x3f, 0x90, 0xe3, 0x54, 0x40, 0x2b, 0xb5, 0xa5, 0x4e,
	0xbd, 0xd7, 0xe8, 0x1e, 0xde, 0x77, 0xd4, 0x9e, 0x65, 0x14, 0xfc, 0x07, 0xea, 0x06, 0xa5, 0xe1,
	0x94, 0x6d, 0x57, 0x59, 0xf8, 0x1f, 0xf0, 0x6d, 0x47, 0x7c, 0x41, 0x33, 0xbb, 0x3d, 0xc0, 0x32,
	0x94, 0xe7, 0x01, 0xf3, 0xfe, 0xf6, 0x01, 0x8e, 0x2a, 0xe8, 0x3b, 0x28, 0xe3, 0x89, 0xb3, 0xb0,
	0xcd, 0xb1, 0xa3, 0x16, 0x90, 0x02, 0xe5, 0xf4, 0x26, 0x25, 0xf5, 0x99, 0x39, 0x30, 0xad, 0xb9,
	0x69, 0xa8, 0x45, 0x54, 0x83, 0xea, 0x60, 0x32, 0x9a, 0x3e, 0x98, 0x8e, 0x69, 0xa8, 0xa5, 0xde,
	0xbb, 0x04, 0xea, 0x88, 0x3d, 0x47, 0x81, 0x1f, 0xac, 0xa6, 0x51, 0xf0, 0xfa, 0x76, 0x27, 0x5c,
	0x74, 0x03, 0xca, 0x21, 0x07, 0xaa, 0xe5, 0x81, 0x13, 0x4b, 0xbd, 0x99, 0xc3, 0xd3, 0xa4, 0xb8,
	0x80, 0x6e, 0xa1, 0x76, 0x32, 0x32, 0xd4, 0xca, 0xb9, 0x9f, 0x47, 0xa9, 0x9f, 0xea, 0xe2, 0x02,
	0x7a, 0x82, 0xc6, 0x99, 0x25, 0xa2, 0xdf, 0x39, 0xef, 0xd2, 0xfa, 0x75, 0xfc, 0x15, 0xe5, 0x10,
	0xef, 0x4a, 0x72, 0xe5, 0xf4, 0x4b, 0x5d, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x16, 0x8d, 0x8c,
	0x65, 0x6b, 0x02, 0x00, 0x00,
}
