// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rdex/models.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params defines the parameters for the module.
type SwapPool struct {
	Denom         string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	RTokenBalance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=rTokenBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"rTokenBalance"`
	FisBalance    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=fisBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"fisBalance"`
	TotalUnit     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=totalUnit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"totalUnit"`
}

func (m *SwapPool) Reset()         { *m = SwapPool{} }
func (m *SwapPool) String() string { return proto.CompactTextString(m) }
func (*SwapPool) ProtoMessage()    {}
func (*SwapPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc88587beb583d4f, []int{0}
}
func (m *SwapPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapPool.Merge(m, src)
}
func (m *SwapPool) XXX_Size() int {
	return m.Size()
}
func (m *SwapPool) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapPool.DiscardUnknown(m)
}

var xxx_messageInfo_SwapPool proto.InternalMessageInfo

func (m *SwapPool) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*SwapPool)(nil), "stafihub.stafihub.rdex.SwapPool")
}

func init() { proto.RegisterFile("rdex/models.proto", fileDescriptor_cc88587beb583d4f) }

var fileDescriptor_cc88587beb583d4f = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4a, 0x49, 0xad,
	0xd0, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2b,
	0x2e, 0x49, 0x4c, 0xcb, 0xcc, 0x28, 0x4d, 0xd2, 0x83, 0x33, 0x40, 0x8a, 0xa4, 0x44, 0xd2, 0xf3,
	0xd3, 0xf3, 0xc1, 0x4a, 0xf4, 0x41, 0x2c, 0x88, 0x6a, 0xa5, 0x29, 0x4c, 0x5c, 0x1c, 0xc1, 0xe5,
	0x89, 0x05, 0x01, 0xf9, 0xf9, 0x39, 0x42, 0x22, 0x5c, 0xac, 0x29, 0xa9, 0x79, 0xf9, 0xb9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x50, 0x08, 0x17, 0x6f, 0x51, 0x48, 0x7e, 0x76,
	0x6a, 0x9e, 0x53, 0x62, 0x4e, 0x62, 0x5e, 0x72, 0xaa, 0x04, 0x13, 0x48, 0xd6, 0x49, 0xef, 0xc4,
	0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0xd5, 0xd2, 0x33, 0x4b, 0x40, 0x96, 0x24, 0xe7, 0xe7, 0xea,
	0x27, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0x43, 0x29, 0xdd, 0xe2, 0x94, 0x6c, 0xfd, 0x92, 0xca, 0x82,
	0xd4, 0x62, 0x3d, 0xcf, 0xbc, 0x92, 0x20, 0x54, 0x43, 0x84, 0xfc, 0xb8, 0xb8, 0xd2, 0x32, 0x8b,
	0x61, 0x46, 0x32, 0x93, 0x65, 0x24, 0x92, 0x09, 0x42, 0x3e, 0x5c, 0x9c, 0x25, 0xf9, 0x25, 0x89,
	0x39, 0xa1, 0x79, 0x99, 0x25, 0x12, 0x2c, 0x64, 0x19, 0x87, 0x30, 0xc0, 0xc9, 0xf9, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2,
	0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x34, 0x91, 0x0c, 0x83, 0x05, 0x30, 0x82, 0x51,
	0xa1, 0x0f, 0x8e, 0x10, 0xb0, 0x99, 0x49, 0x6c, 0xe0, 0x20, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xca, 0xb0, 0x72, 0x81, 0xa5, 0x01, 0x00, 0x00,
}

func (m *SwapPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalUnit.Size()
		i -= size
		if _, err := m.TotalUnit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModels(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.FisBalance.Size()
		i -= size
		if _, err := m.FisBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModels(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.RTokenBalance.Size()
		i -= size
		if _, err := m.RTokenBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModels(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModels(dAtA []byte, offset int, v uint64) int {
	offset -= sovModels(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SwapPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = m.RTokenBalance.Size()
	n += 1 + l + sovModels(uint64(l))
	l = m.FisBalance.Size()
	n += 1 + l + sovModels(uint64(l))
	l = m.TotalUnit.Size()
	n += 1 + l + sovModels(uint64(l))
	return n
}

func sovModels(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModels(x uint64) (n int) {
	return sovModels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SwapPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: SwapPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RTokenBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RTokenBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FisBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FisBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalUnit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalUnit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModels
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
func skipModels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
				return 0, ErrInvalidLengthModels
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModels
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModels
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModels        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModels          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModels = fmt.Errorf("proto: unexpected end of group")
)