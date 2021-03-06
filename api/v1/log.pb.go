// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/log.proto

package log_v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ProduceRequest struct {
	RecordBatch          *RecordBatch `protobuf:"bytes,1,opt,name=record_batch,json=recordBatch,proto3" json:"record_batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProduceRequest) Reset()         { *m = ProduceRequest{} }
func (m *ProduceRequest) String() string { return proto.CompactTextString(m) }
func (*ProduceRequest) ProtoMessage()    {}
func (*ProduceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{0}
}
func (m *ProduceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProduceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProduceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProduceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceRequest.Merge(m, src)
}
func (m *ProduceRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProduceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceRequest proto.InternalMessageInfo

func (m *ProduceRequest) GetRecordBatch() *RecordBatch {
	if m != nil {
		return m.RecordBatch
	}
	return nil
}

type ProduceResponse struct {
	FirstOffset          uint64   `protobuf:"varint,1,opt,name=first_offset,json=firstOffset,proto3" json:"first_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceResponse) Reset()         { *m = ProduceResponse{} }
func (m *ProduceResponse) String() string { return proto.CompactTextString(m) }
func (*ProduceResponse) ProtoMessage()    {}
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{1}
}
func (m *ProduceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProduceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProduceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProduceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceResponse.Merge(m, src)
}
func (m *ProduceResponse) XXX_Size() int {
	return m.Size()
}
func (m *ProduceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceResponse proto.InternalMessageInfo

func (m *ProduceResponse) GetFirstOffset() uint64 {
	if m != nil {
		return m.FirstOffset
	}
	return 0
}

type ConsumeRequest struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeRequest) Reset()         { *m = ConsumeRequest{} }
func (m *ConsumeRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumeRequest) ProtoMessage()    {}
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{2}
}
func (m *ConsumeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsumeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeRequest.Merge(m, src)
}
func (m *ConsumeRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConsumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeRequest proto.InternalMessageInfo

func (m *ConsumeRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ConsumeResponse struct {
	RecordBatch          *RecordBatch `protobuf:"bytes,2,opt,name=record_batch,json=recordBatch,proto3" json:"record_batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConsumeResponse) Reset()         { *m = ConsumeResponse{} }
func (m *ConsumeResponse) String() string { return proto.CompactTextString(m) }
func (*ConsumeResponse) ProtoMessage()    {}
func (*ConsumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{3}
}
func (m *ConsumeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsumeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeResponse.Merge(m, src)
}
func (m *ConsumeResponse) XXX_Size() int {
	return m.Size()
}
func (m *ConsumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeResponse proto.InternalMessageInfo

func (m *ConsumeResponse) GetRecordBatch() *RecordBatch {
	if m != nil {
		return m.RecordBatch
	}
	return nil
}

type RecordBatch struct {
	FirstOffset          uint64    `protobuf:"varint,1,opt,name=first_offset,json=firstOffset,proto3" json:"first_offset,omitempty"`
	Records              []*Record `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RecordBatch) Reset()         { *m = RecordBatch{} }
func (m *RecordBatch) String() string { return proto.CompactTextString(m) }
func (*RecordBatch) ProtoMessage()    {}
func (*RecordBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{4}
}
func (m *RecordBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecordBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecordBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecordBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordBatch.Merge(m, src)
}
func (m *RecordBatch) XXX_Size() int {
	return m.Size()
}
func (m *RecordBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordBatch.DiscardUnknown(m)
}

var xxx_messageInfo_RecordBatch proto.InternalMessageInfo

func (m *RecordBatch) GetFirstOffset() uint64 {
	if m != nil {
		return m.FirstOffset
	}
	return 0
}

func (m *RecordBatch) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type Record struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	OffsetDelta          uint32   `protobuf:"varint,2,opt,name=offset_delta,json=offsetDelta,proto3" json:"offset_delta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{5}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetOffsetDelta() uint32 {
	if m != nil {
		return m.OffsetDelta
	}
	return 0
}

func init() {
	proto.RegisterType((*ProduceRequest)(nil), "log.v1.ProduceRequest")
	proto.RegisterType((*ProduceResponse)(nil), "log.v1.ProduceResponse")
	proto.RegisterType((*ConsumeRequest)(nil), "log.v1.ConsumeRequest")
	proto.RegisterType((*ConsumeResponse)(nil), "log.v1.ConsumeResponse")
	proto.RegisterType((*RecordBatch)(nil), "log.v1.RecordBatch")
	proto.RegisterType((*Record)(nil), "log.v1.Record")
}

func init() { proto.RegisterFile("api/v1/log.proto", fileDescriptor_19a5c3fde3f7ae80) }

var fileDescriptor_19a5c3fde3f7ae80 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0x6e, 0xe2, 0x30,
	0x14, 0xc4, 0xb0, 0x1b, 0xa4, 0x97, 0x00, 0x2b, 0xef, 0x8a, 0x45, 0x1c, 0x10, 0xf5, 0x29, 0x97,
	0x26, 0x40, 0xab, 0x9e, 0x7a, 0x29, 0x45, 0x55, 0x2b, 0x55, 0x6a, 0xe5, 0xde, 0x7a, 0x41, 0x49,
	0x70, 0x02, 0x12, 0x60, 0x6a, 0x3b, 0xfc, 0x46, 0x7f, 0xab, 0xc7, 0x7e, 0x42, 0xc5, 0x97, 0x54,
	0xb1, 0x03, 0x34, 0xf4, 0x82, 0x7a, 0xf3, 0x9b, 0xf1, 0x9b, 0x99, 0x4c, 0x0c, 0x7f, 0x82, 0xd5,
	0xcc, 0x5f, 0xf7, 0xfd, 0x39, 0x4f, 0xbc, 0x95, 0xe0, 0x8a, 0x63, 0x2b, 0x3b, 0xae, 0xfb, 0xed,
	0xd3, 0x64, 0xa6, 0xa6, 0x69, 0xe8, 0x45, 0x7c, 0xe1, 0x27, 0x3c, 0xe1, 0xbe, 0xa6, 0xc3, 0x34,
	0xd6, 0x93, 0x1e, 0xf4, 0xc9, 0xac, 0x91, 0x5b, 0xa8, 0x3f, 0x0a, 0x3e, 0x49, 0x23, 0x46, 0xd9,
	0x4b, 0xca, 0xa4, 0xc2, 0x17, 0xe0, 0x08, 0x16, 0x71, 0x31, 0x19, 0x87, 0x81, 0x8a, 0xa6, 0x2d,
	0xd4, 0x45, 0xae, 0x3d, 0xf8, 0xeb, 0x19, 0x7d, 0x8f, 0x6a, 0x6e, 0x98, 0x51, 0xd4, 0x16, 0xfb,
	0x81, 0x9c, 0x43, 0x63, 0xa7, 0x24, 0x57, 0x7c, 0x29, 0x19, 0x3e, 0x01, 0x27, 0x9e, 0x09, 0xa9,
	0xc6, 0x3c, 0x8e, 0x25, 0x53, 0x5a, 0xea, 0x17, 0xb5, 0x35, 0xf6, 0xa0, 0x21, 0xe2, 0x42, 0xfd,
	0x9a, 0x2f, 0x65, 0xba, 0xd8, 0xf9, 0x37, 0xc1, 0x2a, 0x5c, 0xcf, 0x27, 0x72, 0x07, 0x8d, 0xdd,
	0xcd, 0x5c, 0xff, 0x30, 0x6a, 0xf9, 0xc8, 0xa8, 0xcf, 0x60, 0x7f, 0xe1, 0x8e, 0x88, 0x89, 0x5d,
	0xa8, 0x1a, 0x01, 0xd9, 0xaa, 0x74, 0x2b, 0xae, 0x3d, 0xa8, 0x17, 0x4d, 0xe8, 0x96, 0x26, 0x57,
	0x60, 0x19, 0x08, 0xff, 0x83, 0xdf, 0xeb, 0x60, 0x9e, 0x32, 0xad, 0xe7, 0x50, 0x33, 0x64, 0x66,
	0xc6, 0x66, 0x3c, 0x61, 0x73, 0x15, 0xe8, 0xcc, 0x35, 0x6a, 0x1b, 0x6c, 0x94, 0x41, 0x83, 0xd7,
	0x32, 0x54, 0xee, 0x79, 0x82, 0x2f, 0xa1, 0x9a, 0x37, 0x8a, 0x9b, 0x5b, 0xbb, 0xe2, 0xcf, 0x6a,
	0xff, 0xff, 0x86, 0x9b, 0x6a, 0x48, 0x29, 0xdb, 0xce, 0xfb, 0xda, 0x6f, 0x17, 0xab, 0xde, 0x6f,
	0x1f, 0x14, 0x4b, 0x4a, 0x78, 0x04, 0xb5, 0x1c, 0x7c, 0x52, 0x82, 0x05, 0x8b, 0x1f, 0x68, 0xf4,
	0x10, 0xbe, 0x81, 0x5a, 0x1e, 0xec, 0x50, 0xe5, 0xe8, 0xef, 0x70, 0x51, 0x0f, 0x0d, 0x9d, 0xb7,
	0x4d, 0x07, 0xbd, 0x6f, 0x3a, 0xe8, 0x63, 0xd3, 0x41, 0xa1, 0xa5, 0x9f, 0xee, 0xd9, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7c, 0x67, 0x7b, 0xa5, 0x05, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogClient interface {
	Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error)
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error)
	ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Log_ConsumeStreamClient, error)
	ProduceStream(ctx context.Context, opts ...grpc.CallOption) (Log_ProduceStreamClient, error)
}

type logClient struct {
	cc *grpc.ClientConn
}

func NewLogClient(cc *grpc.ClientConn) LogClient {
	return &logClient{cc}
}

func (c *logClient) Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error) {
	out := new(ProduceResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Log/Produce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error) {
	out := new(ConsumeResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Log/Consume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Log_ConsumeStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Log_serviceDesc.Streams[0], "/log.v1.Log/ConsumeStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &logConsumeStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Log_ConsumeStreamClient interface {
	Recv() (*ConsumeResponse, error)
	grpc.ClientStream
}

type logConsumeStreamClient struct {
	grpc.ClientStream
}

func (x *logConsumeStreamClient) Recv() (*ConsumeResponse, error) {
	m := new(ConsumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logClient) ProduceStream(ctx context.Context, opts ...grpc.CallOption) (Log_ProduceStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Log_serviceDesc.Streams[1], "/log.v1.Log/ProduceStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &logProduceStreamClient{stream}
	return x, nil
}

type Log_ProduceStreamClient interface {
	Send(*ProduceRequest) error
	Recv() (*ProduceResponse, error)
	grpc.ClientStream
}

type logProduceStreamClient struct {
	grpc.ClientStream
}

func (x *logProduceStreamClient) Send(m *ProduceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logProduceStreamClient) Recv() (*ProduceResponse, error) {
	m := new(ProduceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServer is the server API for Log service.
type LogServer interface {
	Produce(context.Context, *ProduceRequest) (*ProduceResponse, error)
	Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error)
	ConsumeStream(*ConsumeRequest, Log_ConsumeStreamServer) error
	ProduceStream(Log_ProduceStreamServer) error
}

func RegisterLogServer(s *grpc.Server, srv LogServer) {
	s.RegisterService(&_Log_serviceDesc, srv)
}

func _Log_Produce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.v1.Log/Produce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Produce(ctx, req.(*ProduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_Consume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Consume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.v1.Log/Consume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Consume(ctx, req.(*ConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_ConsumeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServer).ConsumeStream(m, &logConsumeStreamServer{stream})
}

type Log_ConsumeStreamServer interface {
	Send(*ConsumeResponse) error
	grpc.ServerStream
}

type logConsumeStreamServer struct {
	grpc.ServerStream
}

func (x *logConsumeStreamServer) Send(m *ConsumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Log_ProduceStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogServer).ProduceStream(&logProduceStreamServer{stream})
}

type Log_ProduceStreamServer interface {
	Send(*ProduceResponse) error
	Recv() (*ProduceRequest, error)
	grpc.ServerStream
}

type logProduceStreamServer struct {
	grpc.ServerStream
}

func (x *logProduceStreamServer) Send(m *ProduceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logProduceStreamServer) Recv() (*ProduceRequest, error) {
	m := new(ProduceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Log_serviceDesc = grpc.ServiceDesc{
	ServiceName: "log.v1.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _Log_Produce_Handler,
		},
		{
			MethodName: "Consume",
			Handler:    _Log_Consume_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsumeStream",
			Handler:       _Log_ConsumeStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ProduceStream",
			Handler:       _Log_ProduceStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/v1/log.proto",
}

func (m *ProduceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProduceRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RecordBatch != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.RecordBatch.Size()))
		n1, err := m.RecordBatch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ProduceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProduceResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FirstOffset != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.FirstOffset))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConsumeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsumeRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Offset != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.Offset))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConsumeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsumeResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RecordBatch != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.RecordBatch.Size()))
		n2, err := m.RecordBatch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RecordBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecordBatch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FirstOffset != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.FirstOffset))
	}
	if len(m.Records) > 0 {
		for _, msg := range m.Records {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintLog(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.OffsetDelta != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.OffsetDelta))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintLog(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ProduceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RecordBatch != nil {
		l = m.RecordBatch.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProduceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FirstOffset != 0 {
		n += 1 + sovLog(uint64(m.FirstOffset))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConsumeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Offset != 0 {
		n += 1 + sovLog(uint64(m.Offset))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConsumeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RecordBatch != nil {
		l = m.RecordBatch.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RecordBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FirstOffset != 0 {
		n += 1 + sovLog(uint64(m.FirstOffset))
	}
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if m.OffsetDelta != 0 {
		n += 1 + sovLog(uint64(m.OffsetDelta))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLog(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLog(x uint64) (n int) {
	return sovLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProduceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: ProduceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProduceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordBatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RecordBatch == nil {
				m.RecordBatch = &RecordBatch{}
			}
			if err := m.RecordBatch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLog
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
func (m *ProduceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: ProduceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProduceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstOffset", wireType)
			}
			m.FirstOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FirstOffset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLog
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
func (m *ConsumeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: ConsumeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsumeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLog
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
func (m *ConsumeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: ConsumeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsumeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordBatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RecordBatch == nil {
				m.RecordBatch = &RecordBatch{}
			}
			if err := m.RecordBatch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLog
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
func (m *RecordBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: RecordBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecordBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstOffset", wireType)
			}
			m.FirstOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FirstOffset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, &Record{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLog
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
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OffsetDelta", wireType)
			}
			m.OffsetDelta = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OffsetDelta |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLog
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
func skipLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLog
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
					return 0, ErrIntOverflowLog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLog
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
				return 0, ErrInvalidLengthLog
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLog
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLog
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLog(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLog
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLog = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLog   = fmt.Errorf("proto: integer overflow")
)
