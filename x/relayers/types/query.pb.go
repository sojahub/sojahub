// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayers/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type QueryRelayersRequest struct {
	Arena string `protobuf:"bytes,1,opt,name=arena,proto3" json:"arena,omitempty"`
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryRelayersRequest) Reset()         { *m = QueryRelayersRequest{} }
func (m *QueryRelayersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRelayersRequest) ProtoMessage()    {}
func (*QueryRelayersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e796b53400be0772, []int{0}
}
func (m *QueryRelayersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRelayersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRelayersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRelayersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRelayersRequest.Merge(m, src)
}
func (m *QueryRelayersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRelayersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRelayersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRelayersRequest proto.InternalMessageInfo

func (m *QueryRelayersRequest) GetArena() string {
	if m != nil {
		return m.Arena
	}
	return ""
}

func (m *QueryRelayersRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type QueryRelayersResponse struct {
	Relayers []string `protobuf:"bytes,1,rep,name=relayers,proto3" json:"relayers,omitempty"`
}

func (m *QueryRelayersResponse) Reset()         { *m = QueryRelayersResponse{} }
func (m *QueryRelayersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRelayersResponse) ProtoMessage()    {}
func (*QueryRelayersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e796b53400be0772, []int{1}
}
func (m *QueryRelayersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRelayersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRelayersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRelayersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRelayersResponse.Merge(m, src)
}
func (m *QueryRelayersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRelayersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRelayersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRelayersResponse proto.InternalMessageInfo

func (m *QueryRelayersResponse) GetRelayers() []string {
	if m != nil {
		return m.Relayers
	}
	return nil
}

type QueryThresholdRequest struct {
	Arena string `protobuf:"bytes,1,opt,name=arena,proto3" json:"arena,omitempty"`
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryThresholdRequest) Reset()         { *m = QueryThresholdRequest{} }
func (m *QueryThresholdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryThresholdRequest) ProtoMessage()    {}
func (*QueryThresholdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e796b53400be0772, []int{2}
}
func (m *QueryThresholdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryThresholdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryThresholdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryThresholdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryThresholdRequest.Merge(m, src)
}
func (m *QueryThresholdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryThresholdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryThresholdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryThresholdRequest proto.InternalMessageInfo

func (m *QueryThresholdRequest) GetArena() string {
	if m != nil {
		return m.Arena
	}
	return ""
}

func (m *QueryThresholdRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type QueryThresholdResponse struct {
	Threshold uint32 `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
}

func (m *QueryThresholdResponse) Reset()         { *m = QueryThresholdResponse{} }
func (m *QueryThresholdResponse) String() string { return proto.CompactTextString(m) }
func (*QueryThresholdResponse) ProtoMessage()    {}
func (*QueryThresholdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e796b53400be0772, []int{3}
}
func (m *QueryThresholdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryThresholdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryThresholdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryThresholdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryThresholdResponse.Merge(m, src)
}
func (m *QueryThresholdResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryThresholdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryThresholdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryThresholdResponse proto.InternalMessageInfo

func (m *QueryThresholdResponse) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryRelayersRequest)(nil), "sojahub.sojahub.relayers.QueryRelayersRequest")
	proto.RegisterType((*QueryRelayersResponse)(nil), "sojahub.sojahub.relayers.QueryRelayersResponse")
	proto.RegisterType((*QueryThresholdRequest)(nil), "sojahub.sojahub.relayers.QueryThresholdRequest")
	proto.RegisterType((*QueryThresholdResponse)(nil), "sojahub.sojahub.relayers.QueryThresholdResponse")
}

func init() { proto.RegisterFile("relayers/query.proto", fileDescriptor_e796b53400be0772) }

var fileDescriptor_e796b53400be0772 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xce, 0xf4, 0xde, 0x5e, 0x9a, 0x81, 0xbb, 0x19, 0x72, 0x2f, 0x25, 0x94, 0x58, 0xb2, 0x2a,
	0x4a, 0x33, 0xb6, 0x45, 0xc5, 0x85, 0x9b, 0xba, 0x70, 0x6d, 0x70, 0xe5, 0x6e, 0xd2, 0x8e, 0x69,
	0xa0, 0xcd, 0xa4, 0x99, 0xa9, 0x58, 0x4a, 0x37, 0x3e, 0x81, 0xe0, 0x2b, 0xf8, 0x02, 0xbe, 0x45,
	0x97, 0x05, 0x37, 0xae, 0x44, 0x52, 0x1f, 0x44, 0x32, 0xf9, 0xa9, 0x84, 0xaa, 0xd4, 0xdd, 0x99,
	0x8f, 0xf3, 0x7d, 0xe7, 0xfb, 0xce, 0x1c, 0xa8, 0x85, 0x74, 0x48, 0xa6, 0x34, 0xe4, 0x78, 0x3c,
	0xa1, 0xe1, 0xd4, 0x0a, 0x42, 0x26, 0x18, 0xd2, 0xb9, 0x20, 0x57, 0xde, 0x60, 0xe2, 0x58, 0x79,
	0x91, 0xf5, 0xe9, 0x35, 0x97, 0x31, 0x77, 0x48, 0x31, 0x09, 0x3c, 0x4c, 0x7c, 0x9f, 0x09, 0x22,
	0x3c, 0xe6, 0xf3, 0x84, 0xa9, 0xef, 0xf6, 0x18, 0x1f, 0x31, 0x8e, 0x1d, 0xc2, 0x69, 0x22, 0x89,
	0xaf, 0x5b, 0x0e, 0x15, 0xa4, 0x85, 0x03, 0xe2, 0x7a, 0xbe, 0x6c, 0x4e, 0x7b, 0x35, 0x97, 0xb9,
	0x4c, 0x96, 0x38, 0xae, 0x12, 0xd4, 0xec, 0x42, 0xed, 0x3c, 0xe6, 0xd9, 0xe9, 0x40, 0x9b, 0x8e,
	0x27, 0x94, 0x0b, 0xa4, 0xc1, 0x32, 0x09, 0xa9, 0x4f, 0xaa, 0xa0, 0x0e, 0x1a, 0xaa, 0x9d, 0x3c,
	0x62, 0xb4, 0x4f, 0x7d, 0x36, 0xaa, 0x96, 0x12, 0x54, 0x3e, 0xcc, 0x63, 0xf8, 0xaf, 0xa0, 0xc1,
	0x03, 0xe6, 0x73, 0x8a, 0xea, 0xb0, 0x92, 0x05, 0xa9, 0x82, 0xfa, 0xaf, 0x86, 0xda, 0xfd, 0xbd,
	0x78, 0xd9, 0x51, 0xec, 0x1c, 0x35, 0x4f, 0x53, 0xea, 0xc5, 0x20, 0xa4, 0x7c, 0xc0, 0x86, 0xfd,
	0x9f, 0xcc, 0x3f, 0x84, 0xff, 0x8b, 0x22, 0xa9, 0x81, 0x1a, 0x54, 0x45, 0x06, 0x4a, 0xa5, 0xbf,
	0xf6, 0x1a, 0x68, 0x47, 0x25, 0x58, 0x96, 0x44, 0xf4, 0x00, 0x60, 0x25, 0x73, 0x8f, 0xf6, 0xad,
	0xcf, 0xff, 0xc3, 0xda, 0xb4, 0x2c, 0xbd, 0xb5, 0x05, 0x23, 0x71, 0x66, 0x76, 0x6e, 0x9f, 0xde,
	0xee, 0x4b, 0x4d, 0xb4, 0x87, 0x33, 0xc6, 0xba, 0xc8, 0x8f, 0x64, 0x26, 0x53, 0xcf, 0xf1, 0x4c,
	0xe6, 0x9c, 0xa3, 0x47, 0x00, 0xd5, 0x3c, 0x24, 0xfa, 0x7e, 0x6a, 0x71, 0xab, 0x7a, 0x7b, 0x1b,
	0x4a, 0xea, 0xf4, 0x44, 0x3a, 0x3d, 0x42, 0x07, 0x5f, 0x39, 0xcd, 0x97, 0x5a, 0xf4, 0xdc, 0x3d,
	0x5b, 0x44, 0x06, 0x58, 0x46, 0x06, 0x78, 0x8d, 0x0c, 0x70, 0xb7, 0x32, 0x94, 0xe5, 0xca, 0x50,
	0x9e, 0x57, 0x86, 0x72, 0xd9, 0x74, 0x3d, 0x11, 0x1b, 0xe8, 0xb1, 0xd1, 0x06, 0xe9, 0x9b, 0x0f,
	0xe2, 0xd3, 0x80, 0x72, 0xe7, 0x8f, 0x3c, 0xd8, 0xce, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99,
	0xec, 0xcc, 0x51, 0x44, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a list of relayersByDenom items.
	Relayers(ctx context.Context, in *QueryRelayersRequest, opts ...grpc.CallOption) (*QueryRelayersResponse, error)
	// Queries a threshold by denom.
	Threshold(ctx context.Context, in *QueryThresholdRequest, opts ...grpc.CallOption) (*QueryThresholdResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Relayers(ctx context.Context, in *QueryRelayersRequest, opts ...grpc.CallOption) (*QueryRelayersResponse, error) {
	out := new(QueryRelayersResponse)
	err := c.cc.Invoke(ctx, "/sojahub.sojahub.relayers.Query/Relayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Threshold(ctx context.Context, in *QueryThresholdRequest, opts ...grpc.CallOption) (*QueryThresholdResponse, error) {
	out := new(QueryThresholdResponse)
	err := c.cc.Invoke(ctx, "/sojahub.sojahub.relayers.Query/Threshold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a list of relayersByDenom items.
	Relayers(context.Context, *QueryRelayersRequest) (*QueryRelayersResponse, error)
	// Queries a threshold by denom.
	Threshold(context.Context, *QueryThresholdRequest) (*QueryThresholdResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Relayers(ctx context.Context, req *QueryRelayersRequest) (*QueryRelayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Relayers not implemented")
}
func (*UnimplementedQueryServer) Threshold(ctx context.Context, req *QueryThresholdRequest) (*QueryThresholdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Threshold not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Relayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRelayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Relayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sojahub.sojahub.relayers.Query/Relayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Relayers(ctx, req.(*QueryRelayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Threshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryThresholdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Threshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sojahub.sojahub.relayers.Query/Threshold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Threshold(ctx, req.(*QueryThresholdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sojahub.sojahub.relayers.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Relayers",
			Handler:    _Query_Relayers_Handler,
		},
		{
			MethodName: "Threshold",
			Handler:    _Query_Threshold_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relayers/query.proto",
}

func (m *QueryRelayersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRelayersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRelayersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Arena) > 0 {
		i -= len(m.Arena)
		copy(dAtA[i:], m.Arena)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Arena)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRelayersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRelayersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRelayersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relayers) > 0 {
		for iNdEx := len(m.Relayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Relayers[iNdEx])
			copy(dAtA[i:], m.Relayers[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Relayers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryThresholdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryThresholdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryThresholdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Arena) > 0 {
		i -= len(m.Arena)
		copy(dAtA[i:], m.Arena)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Arena)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryThresholdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryThresholdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryThresholdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Threshold != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRelayersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Arena)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRelayersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Relayers) > 0 {
		for _, s := range m.Relayers {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryThresholdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Arena)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryThresholdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Threshold != 0 {
		n += 1 + sovQuery(uint64(m.Threshold))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRelayersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRelayersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRelayersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arena", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Arena = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRelayersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRelayersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRelayersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayers = append(m.Relayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryThresholdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryThresholdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryThresholdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arena", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Arena = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryThresholdResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryThresholdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryThresholdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
