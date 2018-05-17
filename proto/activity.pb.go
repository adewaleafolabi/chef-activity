// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/activity.proto

package chefhero

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

type ActivityRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Environment          string   `protobuf:"bytes,3,opt,name=environment" json:"environment,omitempty"`
	Component            string   `protobuf:"bytes,4,opt,name=component" json:"component,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
	Data                 string   `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=createdAt" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivityRequest) Reset()         { *m = ActivityRequest{} }
func (m *ActivityRequest) String() string { return proto.CompactTextString(m) }
func (*ActivityRequest) ProtoMessage()    {}
func (*ActivityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_activity_9c629b1709ef02c5, []int{0}
}
func (m *ActivityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityRequest.Unmarshal(m, b)
}
func (m *ActivityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityRequest.Marshal(b, m, deterministic)
}
func (dst *ActivityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityRequest.Merge(dst, src)
}
func (m *ActivityRequest) XXX_Size() int {
	return xxx_messageInfo_ActivityRequest.Size(m)
}
func (m *ActivityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityRequest proto.InternalMessageInfo

func (m *ActivityRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ActivityRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ActivityRequest) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *ActivityRequest) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *ActivityRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ActivityRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ActivityRequest) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type ActivityResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivityResponse) Reset()         { *m = ActivityResponse{} }
func (m *ActivityResponse) String() string { return proto.CompactTextString(m) }
func (*ActivityResponse) ProtoMessage()    {}
func (*ActivityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_activity_9c629b1709ef02c5, []int{1}
}
func (m *ActivityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityResponse.Unmarshal(m, b)
}
func (m *ActivityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityResponse.Marshal(b, m, deterministic)
}
func (dst *ActivityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityResponse.Merge(dst, src)
}
func (m *ActivityResponse) XXX_Size() int {
	return xxx_messageInfo_ActivityResponse.Size(m)
}
func (m *ActivityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityResponse proto.InternalMessageInfo

func (m *ActivityResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ActivityResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*ActivityRequest)(nil), "chefhero.ActivityRequest")
	proto.RegisterType((*ActivityResponse)(nil), "chefhero.ActivityResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ActivityServiceClient is the client API for ActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivityServiceClient interface {
	CreateActivity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	FindActivity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (ActivityService_FindActivityClient, error)
}

type activityServiceClient struct {
	cc *grpc.ClientConn
}

func NewActivityServiceClient(cc *grpc.ClientConn) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) CreateActivity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/chefhero.ActivityService/CreateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) FindActivity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (ActivityService_FindActivityClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ActivityService_serviceDesc.Streams[0], "/chefhero.ActivityService/FindActivity", opts...)
	if err != nil {
		return nil, err
	}
	x := &activityServiceFindActivityClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ActivityService_FindActivityClient interface {
	Recv() (*ActivityRequest, error)
	grpc.ClientStream
}

type activityServiceFindActivityClient struct {
	grpc.ClientStream
}

func (x *activityServiceFindActivityClient) Recv() (*ActivityRequest, error) {
	m := new(ActivityRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ActivityService service

type ActivityServiceServer interface {
	CreateActivity(context.Context, *ActivityRequest) (*ActivityResponse, error)
	FindActivity(*ActivityRequest, ActivityService_FindActivityServer) error
}

func RegisterActivityServiceServer(s *grpc.Server, srv ActivityServiceServer) {
	s.RegisterService(&_ActivityService_serviceDesc, srv)
}

func _ActivityService_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chefhero.ActivityService/CreateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).CreateActivity(ctx, req.(*ActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_FindActivity_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ActivityRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActivityServiceServer).FindActivity(m, &activityServiceFindActivityServer{stream})
}

type ActivityService_FindActivityServer interface {
	Send(*ActivityRequest) error
	grpc.ServerStream
}

type activityServiceFindActivityServer struct {
	grpc.ServerStream
}

func (x *activityServiceFindActivityServer) Send(m *ActivityRequest) error {
	return x.ServerStream.SendMsg(m)
}

var _ActivityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chefhero.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateActivity",
			Handler:    _ActivityService_CreateActivity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindActivity",
			Handler:       _ActivityService_FindActivity_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/activity.proto",
}

func init() { proto.RegisterFile("proto/activity.proto", fileDescriptor_activity_9c629b1709ef02c5) }

var fileDescriptor_activity_9c629b1709ef02c5 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x3b, 0x63, 0xdb, 0x69, 0xaf, 0x52, 0xe5, 0xd2, 0x45, 0x2c, 0x2e, 0xca, 0xac, 0x5c,
	0x8d, 0xa2, 0x5b, 0x37, 0x45, 0x10, 0xdd, 0x8e, 0x4f, 0x10, 0x93, 0xab, 0x0d, 0x38, 0xc9, 0x98,
	0xa4, 0x03, 0xbe, 0x90, 0x2f, 0xe2, 0x8b, 0x89, 0xb7, 0x0d, 0x2d, 0xfe, 0x40, 0x77, 0x39, 0xe7,
	0x83, 0x43, 0xce, 0xb9, 0x30, 0x6d, 0xbd, 0x8b, 0xee, 0x42, 0xaa, 0x68, 0x3a, 0x13, 0xdf, 0x2b,
	0x96, 0x38, 0x52, 0x4b, 0x7a, 0x5e, 0x92, 0x77, 0xe5, 0x67, 0x06, 0xc7, 0x8b, 0x0d, 0xac, 0xe9,
	0x6d, 0x45, 0x21, 0xe2, 0x04, 0x72, 0xa3, 0x45, 0x36, 0xcf, 0xce, 0xc7, 0x75, 0x6e, 0x34, 0x4e,
	0x61, 0x40, 0x8d, 0x34, 0xaf, 0x22, 0x67, 0x6b, 0x2d, 0x70, 0x0e, 0x87, 0x64, 0x3b, 0xe3, 0x9d,
	0x6d, 0xc8, 0x46, 0x71, 0xc0, 0x6c, 0xd7, 0xc2, 0x33, 0x18, 0x2b, 0xd7, 0xb4, 0xce, 0x7e, 0xf3,
	0x3e, 0xf3, 0xad, 0x81, 0x02, 0x8a, 0x86, 0x42, 0x90, 0x2f, 0x24, 0x06, 0xcc, 0x92, 0x44, 0x84,
	0xbe, 0x96, 0x51, 0x8a, 0x21, 0xdb, 0xfc, 0xe6, 0x2c, 0x4f, 0x32, 0x92, 0x5e, 0x44, 0x51, 0x6c,
	0xb2, 0x92, 0x51, 0xde, 0xc0, 0xc9, 0xb6, 0x44, 0x68, 0x9d, 0x0d, 0xf4, 0xab, 0x85, 0x80, 0x22,
	0xac, 0x94, 0xa2, 0x10, 0xb8, 0xc7, 0xa8, 0x4e, 0xf2, 0xea, 0x63, 0x67, 0x83, 0x47, 0xf2, 0x9d,
	0x51, 0x84, 0x0f, 0x30, 0xb9, 0xe5, 0xf8, 0x04, 0xf0, 0xb4, 0x4a, 0xa3, 0x55, 0x3f, 0x06, 0x9b,
	0xcd, 0xfe, 0x42, 0xeb, 0x6f, 0x94, 0x3d, 0xbc, 0x87, 0xa3, 0x3b, 0x63, 0xf5, 0x3e, 0x41, 0xff,
	0xa3, 0xb2, 0x77, 0x99, 0x3d, 0x0d, 0xf9, 0x7a, 0xd7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda,
	0x42, 0xad, 0xe0, 0xd5, 0x01, 0x00, 0x00,
}
