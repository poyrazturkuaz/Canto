// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canto/fees/v1/genesis.proto

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

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// active registered contracts for fee distribution
	Fees []Fee `protobuf:"bytes,2,rep,name=fees,proto3" json:"fees"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4078d542b007a71e, []int{0}
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

func (m *GenesisState) GetFees() []Fee {
	if m != nil {
		return m.Fees
	}
	return nil
}

// Params defines the fees module params
type Params struct {
	// parameter to enable fees
	EnableFees bool `protobuf:"varint,1,opt,name=enable_fees,json=enableFees,proto3" json:"enable_fees,omitempty"`
	// developer_shares defines the proportion of the transaction fees to be
	// distributed to the registered contract owner
	DeveloperShares github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=developer_shares,json=developerShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"developer_shares"`
	// addr_derivation_cost_create defines the cost of address derivation for
	// verifying the contract deployer at fee registration
	AddrDerivationCostCreate uint64 `protobuf:"varint,3,opt,name=addr_derivation_cost_create,json=addrDerivationCostCreate,proto3" json:"addr_derivation_cost_create,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4078d542b007a71e, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnableFees() bool {
	if m != nil {
		return m.EnableFees
	}
	return false
}

func (m *Params) GetAddrDerivationCostCreate() uint64 {
	if m != nil {
		return m.AddrDerivationCostCreate
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "canto.fees.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "canto.fees.v1.Params")
}

func init() { proto.RegisterFile("canto/fees/v1/genesis.proto", fileDescriptor_4078d542b007a71e) }

var fileDescriptor_4078d542b007a71e = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x3b, 0x17, 0x42, 0xee, 0x1d, 0xae, 0xd1, 0x4c, 0x34, 0x69, 0x20, 0x29, 0x84, 0x85,
	0xe9, 0x42, 0xa6, 0x01, 0xd6, 0x6e, 0x80, 0xa0, 0x2b, 0x63, 0xca, 0x4a, 0x37, 0xcd, 0xd0, 0x1e,
	0x4b, 0x03, 0x74, 0xea, 0xcc, 0x58, 0xf5, 0x2d, 0x7c, 0x27, 0x37, 0x2c, 0x59, 0x1a, 0x17, 0xc4,
	0xc0, 0x8b, 0x98, 0x99, 0x12, 0x22, 0xae, 0x7a, 0x7a, 0xfe, 0xef, 0xfc, 0xe7, 0x64, 0x7e, 0x5c,
	0x0f, 0x59, 0xaa, 0xb8, 0xf7, 0x00, 0x20, 0xbd, 0xbc, 0xe3, 0xc5, 0x90, 0x82, 0x4c, 0x24, 0xcd,
	0x04, 0x57, 0x9c, 0x1c, 0x19, 0x91, 0x6a, 0x91, 0xe6, 0x9d, 0x9a, 0x7d, 0xc8, 0x9a, 0xb6, 0x01,
	0x6b, 0xa7, 0x31, 0x8f, 0xb9, 0x29, 0x3d, 0x5d, 0x15, 0xdd, 0xd6, 0x23, 0xfe, 0x7f, 0x55, 0xf8,
	0x8d, 0x15, 0x53, 0x40, 0x7a, 0xb8, 0x92, 0x31, 0xc1, 0x16, 0xd2, 0x46, 0x4d, 0xe4, 0x56, 0xbb,
	0x67, 0xf4, 0xc0, 0x9f, 0xde, 0x1a, 0xb1, 0x5f, 0x5e, 0xae, 0x1b, 0x96, 0xbf, 0x43, 0xc9, 0x05,
	0x2e, 0x6b, 0xdd, 0xfe, 0xd3, 0x2c, 0xb9, 0xd5, 0x2e, 0xf9, 0x35, 0x32, 0x02, 0xd8, 0xf1, 0x86,
	0x6a, 0xbd, 0x23, 0x5c, 0x29, 0x6c, 0x48, 0x03, 0x57, 0x21, 0x65, 0x93, 0x39, 0x04, 0x66, 0x5e,
	0xaf, 0xfc, 0xeb, 0xe3, 0xa2, 0x35, 0x02, 0x90, 0xe4, 0x0e, 0x9f, 0x44, 0x90, 0xc3, 0x9c, 0x67,
	0x20, 0x02, 0x39, 0x65, 0xc2, 0x6c, 0x41, 0xee, 0xbf, 0x3e, 0xd5, 0x8e, 0x9f, 0xeb, 0xc6, 0x79,
	0x9c, 0xa8, 0xe9, 0xd3, 0x84, 0x86, 0x7c, 0xe1, 0x85, 0x5c, 0x2e, 0xb8, 0xdc, 0x7d, 0xda, 0x32,
	0x9a, 0x79, 0xea, 0x35, 0x03, 0x49, 0x87, 0x10, 0xfa, 0xc7, 0x7b, 0x9f, 0xb1, 0xb1, 0x21, 0x97,
	0xb8, 0xce, 0xa2, 0x48, 0x04, 0x11, 0x88, 0x24, 0x67, 0x2a, 0xe1, 0x69, 0x10, 0x72, 0xa9, 0x82,
	0x50, 0x00, 0x53, 0x60, 0x97, 0x9a, 0xc8, 0x2d, 0xfb, 0xb6, 0x46, 0x86, 0x7b, 0x62, 0xc0, 0xa5,
	0x1a, 0x18, 0xbd, 0x7f, 0xbd, 0xdc, 0x38, 0x68, 0xb5, 0x71, 0xd0, 0xd7, 0xc6, 0x41, 0x6f, 0x5b,
	0xc7, 0x5a, 0x6d, 0x1d, 0xeb, 0x63, 0xeb, 0x58, 0xf7, 0xf4, 0xc7, 0x45, 0x03, 0xfd, 0x12, 0xed,
	0x1b, 0x50, 0xcf, 0x5c, 0xcc, 0x8a, 0x3f, 0x1d, 0xcb, 0x4b, 0x11, 0x90, 0xb9, 0x6e, 0x52, 0x31,
	0x49, 0xf4, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x37, 0x60, 0x5e, 0xe7, 0x01, 0x00, 0x00,
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
	if len(m.Fees) > 0 {
		for iNdEx := len(m.Fees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AddrDerivationCostCreate != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AddrDerivationCostCreate))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.DeveloperShares.Size()
		i -= size
		if _, err := m.DeveloperShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.EnableFees {
		i--
		if m.EnableFees {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
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
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableFees {
		n += 2
	}
	l = m.DeveloperShares.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.AddrDerivationCostCreate != 0 {
		n += 1 + sovGenesis(uint64(m.AddrDerivationCostCreate))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
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
			m.Fees = append(m.Fees, Fee{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableFees", wireType)
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
			m.EnableFees = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperShares", wireType)
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
			if err := m.DeveloperShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddrDerivationCostCreate", wireType)
			}
			m.AddrDerivationCostCreate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AddrDerivationCostCreate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
