// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sudo/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

type QueryAdminRequest struct {
}

func (m *QueryAdminRequest) Reset()         { *m = QueryAdminRequest{} }
func (m *QueryAdminRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAdminRequest) ProtoMessage()    {}
func (*QueryAdminRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e89a5ddd8e4c069, []int{0}
}
func (m *QueryAdminRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminRequest.Merge(m, src)
}
func (m *QueryAdminRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminRequest proto.InternalMessageInfo

type QueryAdminResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryAdminResponse) Reset()         { *m = QueryAdminResponse{} }
func (m *QueryAdminResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAdminResponse) ProtoMessage()    {}
func (*QueryAdminResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e89a5ddd8e4c069, []int{1}
}
func (m *QueryAdminResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminResponse.Merge(m, src)
}
func (m *QueryAdminResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminResponse proto.InternalMessageInfo

func (m *QueryAdminResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryAdminRequest)(nil), "sojahub.sojahub.sudo.QueryAdminRequest")
	proto.RegisterType((*QueryAdminResponse)(nil), "sojahub.sojahub.sudo.QueryAdminResponse")
}

func init() { proto.RegisterFile("sudo/query.proto", fileDescriptor_8e89a5ddd8e4c069) }

var fileDescriptor_8e89a5ddd8e4c069 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2e, 0x4d, 0xc9,
	0xd7, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2b, 0x2e,
	0x49, 0x4c, 0xcb, 0xcc, 0x28, 0x4d, 0xd2, 0x43, 0x30, 0x4a, 0x53, 0xf2, 0xa5, 0x64, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b,
	0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xba, 0x94, 0x84, 0xb9, 0x04, 0x03, 0x41, 0x86, 0x38, 0xa6, 0xe4,
	0x66, 0xe6, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0xe9, 0x71, 0x09, 0x21, 0x0b, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17,
	0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x46, 0xfd, 0x8c, 0x5c, 0xac, 0x60, 0x0d,
	0x42, 0xad, 0x8c, 0x5c, 0xac, 0x60, 0x5d, 0x42, 0x9a, 0x7a, 0xd8, 0xdd, 0xa3, 0x87, 0x61, 0x9d,
	0x94, 0x16, 0x31, 0x4a, 0x21, 0x8e, 0x50, 0x52, 0x6d, 0xba, 0xfc, 0x64, 0x32, 0x93, 0xbc, 0x90,
	0xac, 0x3e, 0x4c, 0x29, 0x12, 0x03, 0x14, 0x24, 0x89, 0x20, 0xe5, 0x4e, 0xce, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x99, 0x9e, 0x59, 0x02, 0xb2, 0x20, 0x39, 0x3f,
	0x17, 0x8b, 0x11, 0x15, 0x10, 0x43, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x41, 0x64,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x19, 0xcc, 0x2b, 0x6c, 0x01, 0x00, 0x00,
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
	// Queries a list of admin items.
	Admin(ctx context.Context, in *QueryAdminRequest, opts ...grpc.CallOption) (*QueryAdminResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Admin(ctx context.Context, in *QueryAdminRequest, opts ...grpc.CallOption) (*QueryAdminResponse, error) {
	out := new(QueryAdminResponse)
	err := c.cc.Invoke(ctx, "/sojahub.sojahub.sudo.Query/Admin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a list of admin items.
	Admin(context.Context, *QueryAdminRequest) (*QueryAdminResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Admin(ctx context.Context, req *QueryAdminRequest) (*QueryAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Admin not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Admin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Admin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sojahub.sojahub.sudo.Query/Admin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Admin(ctx, req.(*QueryAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sojahub.sojahub.sudo.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Admin",
			Handler:    _Query_Admin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sudo/query.proto",
}

func (m *QueryAdminRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAdminResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
func (m *QueryAdminRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAdminResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAdminRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAdminRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryAdminResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAdminResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
