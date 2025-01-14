// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mining/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the mining module's genesis state.
type GenesisState struct {
	Params               Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	StakePoolList        []*StakePool       `protobuf:"bytes,2,rep,name=stakePoolList,proto3" json:"stakePoolList,omitempty"`
	StakeItemList        []*StakeItem       `protobuf:"bytes,3,rep,name=stakeItemList,proto3" json:"stakeItemList,omitempty"`
	UserStakeRecordList  []*UserStakeRecord `protobuf:"bytes,4,rep,name=userStakeRecordList,proto3" json:"userStakeRecordList,omitempty"`
	MiningProviderList   []string           `protobuf:"bytes,5,rep,name=miningProviderList,proto3" json:"miningProviderList,omitempty"`
	RewardTokenList      []*RewardToken     `protobuf:"bytes,6,rep,name=rewardTokenList,proto3" json:"rewardTokenList,omitempty"`
	MaxRewardPoolNumber  uint32             `protobuf:"varint,7,opt,name=maxRewardPoolNumber,proto3" json:"maxRewardPoolNumber,omitempty"`
	MiningProviderSwitch bool               `protobuf:"varint,8,opt,name=miningProviderSwitch,proto3" json:"miningProviderSwitch,omitempty"`
	MaxStakeItemNumber   uint32             `protobuf:"varint,9,opt,name=maxStakeItemNumber,proto3" json:"maxStakeItemNumber,omitempty"`
	StakeTokenList       []string           `protobuf:"bytes,10,rep,name=stakeTokenList,proto3" json:"stakeTokenList,omitempty"`
	StakeItemLimit       *StakeItemLimit    `protobuf:"bytes,11,opt,name=stakeItemLimit,proto3" json:"stakeItemLimit,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2628b9fc0ab431ed, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetStakePoolList() []*StakePool {
	if m != nil {
		return m.StakePoolList
	}
	return nil
}

func (m *GenesisState) GetStakeItemList() []*StakeItem {
	if m != nil {
		return m.StakeItemList
	}
	return nil
}

func (m *GenesisState) GetUserStakeRecordList() []*UserStakeRecord {
	if m != nil {
		return m.UserStakeRecordList
	}
	return nil
}

func (m *GenesisState) GetMiningProviderList() []string {
	if m != nil {
		return m.MiningProviderList
	}
	return nil
}

func (m *GenesisState) GetRewardTokenList() []*RewardToken {
	if m != nil {
		return m.RewardTokenList
	}
	return nil
}

func (m *GenesisState) GetMaxRewardPoolNumber() uint32 {
	if m != nil {
		return m.MaxRewardPoolNumber
	}
	return 0
}

func (m *GenesisState) GetMiningProviderSwitch() bool {
	if m != nil {
		return m.MiningProviderSwitch
	}
	return false
}

func (m *GenesisState) GetMaxStakeItemNumber() uint32 {
	if m != nil {
		return m.MaxStakeItemNumber
	}
	return 0
}

func (m *GenesisState) GetStakeTokenList() []string {
	if m != nil {
		return m.StakeTokenList
	}
	return nil
}

func (m *GenesisState) GetStakeItemLimit() *StakeItemLimit {
	if m != nil {
		return m.StakeItemLimit
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sojahub.sojahub.mining.GenesisState")
}

func init() { proto.RegisterFile("mining/genesis.proto", fileDescriptor_2628b9fc0ab431ed) }

var fileDescriptor_2628b9fc0ab431ed = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0xab, 0xd3, 0x40,
	0x10, 0xc7, 0x13, 0xfb, 0x5e, 0x7d, 0x6f, 0xeb, 0x53, 0xd8, 0xf6, 0x10, 0x7a, 0x88, 0x41, 0x51,
	0x22, 0x42, 0x22, 0xf5, 0xee, 0xa1, 0x20, 0x52, 0x10, 0x0d, 0x5b, 0xbd, 0xe8, 0x69, 0xdb, 0xac,
	0xe9, 0xd2, 0x6e, 0xb6, 0xec, 0x6e, 0x6d, 0xfd, 0x16, 0x7e, 0xac, 0x1e, 0x8b, 0x27, 0x4f, 0x22,
	0xed, 0x17, 0x79, 0x64, 0x92, 0xb4, 0x34, 0x24, 0xd0, 0xdb, 0x32, 0xf3, 0xff, 0xff, 0x66, 0x76,
	0x66, 0x17, 0xf5, 0x04, 0x4f, 0x79, 0x9a, 0x84, 0x09, 0x4b, 0x99, 0xe6, 0x3a, 0x58, 0x2a, 0x69,
	0x24, 0x76, 0xb4, 0xa1, 0x3f, 0xf8, 0x6c, 0x35, 0x09, 0x8e, 0x87, 0x5c, 0xd7, 0xef, 0x25, 0x32,
	0x91, 0x20, 0x0a, 0xb3, 0x53, 0xae, 0xef, 0x77, 0x0b, 0xca, 0x92, 0x2a, 0x2a, 0x74, 0x25, 0x28,
	0x64, 0xcc, 0x16, 0x45, 0xf0, 0xd9, 0x9f, 0x6b, 0xf4, 0xe8, 0x43, 0x5e, 0x6b, 0x6c, 0xa8, 0x61,
	0xf8, 0x1d, 0x6a, 0xe7, 0x2e, 0xc7, 0xf6, 0x6c, 0xbf, 0x33, 0xf0, 0x82, 0xa6, 0xda, 0x41, 0x04,
	0xba, 0xe1, 0xd5, 0xf6, 0xdf, 0x53, 0x8b, 0x14, 0x2e, 0x3c, 0x42, 0x77, 0xda, 0xd0, 0x39, 0x8b,
	0xa4, 0x5c, 0x7c, 0xe4, 0xda, 0x38, 0x0f, 0xbc, 0x96, 0xdf, 0x19, 0x3c, 0x6f, 0xc6, 0x8c, 0x4b,
	0x39, 0x39, 0x77, 0x1e, 0x51, 0x23, 0xc3, 0x04, 0xa0, 0x5a, 0x17, 0xa1, 0x32, 0x39, 0x39, 0x77,
	0xe2, 0xef, 0xa8, 0xbb, 0xd2, 0x4c, 0x41, 0x9e, 0xb0, 0xa9, 0x54, 0x31, 0x00, 0xaf, 0x00, 0xf8,
	0xaa, 0x19, 0xf8, 0xf5, 0xdc, 0x44, 0xea, 0x28, 0x38, 0x40, 0x38, 0x97, 0x47, 0x4a, 0xfe, 0xe4,
	0x31, 0x53, 0xc0, 0xbe, 0xf6, 0x5a, 0xfe, 0x2d, 0xa9, 0xc9, 0xe0, 0xcf, 0xe8, 0x89, 0x62, 0x6b,
	0xaa, 0xe2, 0x2f, 0x72, 0xce, 0x52, 0x10, 0xb7, 0xa1, 0x91, 0x17, 0xcd, 0x8d, 0x90, 0x93, 0x81,
	0x54, 0xdd, 0xf8, 0x0d, 0xea, 0x0a, 0xba, 0xc9, 0x25, 0xd9, 0xf4, 0x3e, 0xad, 0xc4, 0x84, 0x29,
	0xe7, 0xa1, 0x67, 0xfb, 0x77, 0xa4, 0x2e, 0x85, 0x07, 0xe5, 0x43, 0x2b, 0x1b, 0x1b, 0xaf, 0xb9,
	0x99, 0xce, 0x9c, 0x1b, 0xcf, 0xf6, 0x6f, 0x48, 0x6d, 0x0e, 0xae, 0x49, 0x37, 0xc7, 0x11, 0x17,
	0x45, 0x6e, 0xa1, 0x48, 0x4d, 0x06, 0xbf, 0x44, 0x8f, 0x61, 0x09, 0xa7, 0x5b, 0x22, 0x18, 0x49,
	0x25, 0x8a, 0xa3, 0x42, 0x97, 0x2f, 0x4b, 0x70, 0xe3, 0x74, 0xe0, 0xe5, 0xf9, 0x17, 0xec, 0x19,
	0xf4, 0xa4, 0xe2, 0x1f, 0xbe, 0xdf, 0xee, 0x5d, 0x7b, 0xb7, 0x77, 0xed, 0xff, 0x7b, 0xd7, 0xfe,
	0x7d, 0x70, 0xad, 0xdd, 0xc1, 0xb5, 0xfe, 0x1e, 0x5c, 0xeb, 0xdb, 0xeb, 0x84, 0x9b, 0x8c, 0x33,
	0x95, 0x22, 0x2c, 0xa1, 0xa7, 0xc3, 0x26, 0x2c, 0xbe, 0x88, 0xf9, 0xb5, 0x64, 0x7a, 0xd2, 0x86,
	0x2f, 0xf2, 0xf6, 0x3e, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x32, 0xdc, 0xec, 0x94, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StakeItemLimit != nil {
		{
			size, err := m.StakeItemLimit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if len(m.StakeTokenList) > 0 {
		for iNdEx := len(m.StakeTokenList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.StakeTokenList[iNdEx])
			copy(dAtA[i:], m.StakeTokenList[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.StakeTokenList[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.MaxStakeItemNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxStakeItemNumber))
		i--
		dAtA[i] = 0x48
	}
	if m.MiningProviderSwitch {
		i--
		if m.MiningProviderSwitch {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.MaxRewardPoolNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxRewardPoolNumber))
		i--
		dAtA[i] = 0x38
	}
	if len(m.RewardTokenList) > 0 {
		for iNdEx := len(m.RewardTokenList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RewardTokenList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.MiningProviderList) > 0 {
		for iNdEx := len(m.MiningProviderList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.MiningProviderList[iNdEx])
			copy(dAtA[i:], m.MiningProviderList[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.MiningProviderList[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.UserStakeRecordList) > 0 {
		for iNdEx := len(m.UserStakeRecordList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserStakeRecordList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.StakeItemList) > 0 {
		for iNdEx := len(m.StakeItemList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StakeItemList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StakePoolList) > 0 {
		for iNdEx := len(m.StakePoolList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StakePoolList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.StakePoolList) > 0 {
		for _, e := range m.StakePoolList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StakeItemList) > 0 {
		for _, e := range m.StakeItemList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserStakeRecordList) > 0 {
		for _, e := range m.UserStakeRecordList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MiningProviderList) > 0 {
		for _, s := range m.MiningProviderList {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RewardTokenList) > 0 {
		for _, e := range m.RewardTokenList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.MaxRewardPoolNumber != 0 {
		n += 1 + sovGenesis(uint64(m.MaxRewardPoolNumber))
	}
	if m.MiningProviderSwitch {
		n += 2
	}
	if m.MaxStakeItemNumber != 0 {
		n += 1 + sovGenesis(uint64(m.MaxStakeItemNumber))
	}
	if len(m.StakeTokenList) > 0 {
		for _, s := range m.StakeTokenList {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.StakeItemLimit != nil {
		l = m.StakeItemLimit.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakePoolList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakePoolList = append(m.StakePoolList, &StakePool{})
			if err := m.StakePoolList[len(m.StakePoolList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeItemList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakeItemList = append(m.StakeItemList, &StakeItem{})
			if err := m.StakeItemList[len(m.StakeItemList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserStakeRecordList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserStakeRecordList = append(m.UserStakeRecordList, &UserStakeRecord{})
			if err := m.UserStakeRecordList[len(m.UserStakeRecordList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MiningProviderList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MiningProviderList = append(m.MiningProviderList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardTokenList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardTokenList = append(m.RewardTokenList, &RewardToken{})
			if err := m.RewardTokenList[len(m.RewardTokenList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRewardPoolNumber", wireType)
			}
			m.MaxRewardPoolNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRewardPoolNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MiningProviderSwitch", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MiningProviderSwitch = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStakeItemNumber", wireType)
			}
			m.MaxStakeItemNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxStakeItemNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeTokenList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakeTokenList = append(m.StakeTokenList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeItemLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StakeItemLimit == nil {
				m.StakeItemLimit = &StakeItemLimit{}
			}
			if err := m.StakeItemLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
