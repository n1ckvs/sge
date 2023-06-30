// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/ovm/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the ovm module's genesis state.
type GenesisState struct {
	// params contains parameters of ovm module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// key_vault trusted keys from genesis file.
	KeyVault KeyVault `protobuf:"bytes,2,opt,name=key_vault,json=keyVault,proto3" json:"key_vault"`
	// pubkeys_change_proposals is the finished proposal list for the
	// public keys change.
	PubkeysChangeProposals []PublicKeysChangeProposal `protobuf:"bytes,3,rep,name=pubkeys_change_proposals,json=pubkeysChangeProposals,proto3" json:"pubkeys_change_proposals"`
	// proposal_stats holds the proposal statistics.
	ProposalStats ProposalStats `protobuf:"bytes,4,opt,name=proposal_stats,json=proposalStats,proto3" json:"proposal_stats"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_89d2511229f4bb0b, []int{0}
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

func (m *GenesisState) GetKeyVault() KeyVault {
	if m != nil {
		return m.KeyVault
	}
	return KeyVault{}
}

func (m *GenesisState) GetPubkeysChangeProposals() []PublicKeysChangeProposal {
	if m != nil {
		return m.PubkeysChangeProposals
	}
	return nil
}

func (m *GenesisState) GetProposalStats() ProposalStats {
	if m != nil {
		return m.ProposalStats
	}
	return ProposalStats{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sgenetwork.sge.ovm.GenesisState")
}

func init() { proto.RegisterFile("sge/ovm/genesis.proto", fileDescriptor_89d2511229f4bb0b) }

var fileDescriptor_89d2511229f4bb0b = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x4d, 0xda, 0x52, 0x34, 0x55, 0x0f, 0xb1, 0xd6, 0x10, 0x24, 0x56, 0x0f, 0xd2, 0x83, 0x26,
	0x50, 0x2f, 0xde, 0x14, 0x3d, 0x78, 0x28, 0x48, 0x51, 0xf0, 0xe0, 0x25, 0x6c, 0xca, 0xb0, 0x2d,
	0x49, 0xba, 0x21, 0xb3, 0xa9, 0xe6, 0x2f, 0xfc, 0x2b, 0x7b, 0xec, 0xd1, 0x93, 0x48, 0xfb, 0x23,
	0x92, 0xed, 0xa4, 0x50, 0xcc, 0x6d, 0xf7, 0xcd, 0x7b, 0x6f, 0xde, 0xcc, 0x18, 0x47, 0xc8, 0xc1,
	0x13, 0xb3, 0xd8, 0xe3, 0x30, 0x05, 0x9c, 0xa0, 0x9b, 0xa4, 0x42, 0x0a, 0xd3, 0xc4, 0xe2, 0x2f,
	0xdf, 0x45, 0x1a, 0xba, 0xc8, 0xc1, 0x15, 0xb3, 0xd8, 0x6e, 0x73, 0xc1, 0x85, 0x2a, 0x7b, 0xc5,
	0x6b, 0xcd, 0xb4, 0xdb, 0xa5, 0x41, 0xc2, 0x52, 0x16, 0x93, 0xde, 0x3e, 0x2e, 0xd1, 0x10, 0x72,
	0x7f, 0xc6, 0xb2, 0x48, 0x52, 0xa1, 0xb3, 0xa1, 0xa7, 0x22, 0x11, 0xc8, 0x22, 0xc2, 0x0f, 0x4b,
	0x1c, 0x25, 0x93, 0xe4, 0x72, 0xfe, 0x55, 0x33, 0xf6, 0x1e, 0xd7, 0xb9, 0x5e, 0x24, 0x93, 0x60,
	0xde, 0x18, 0xcd, 0x75, 0x1b, 0x4b, 0xef, 0xea, 0xbd, 0x56, 0xdf, 0x76, 0xff, 0xe7, 0x74, 0x87,
	0x8a, 0x71, 0xdf, 0x98, 0xff, 0x9c, 0x6a, 0xcf, 0xc4, 0x37, 0x6f, 0x8d, 0xdd, 0x4d, 0x14, 0xab,
	0xa6, 0xc4, 0x27, 0x55, 0xe2, 0x01, 0xe4, 0xaf, 0x05, 0x87, 0xe4, 0x3b, 0x21, 0xfd, 0xcd, 0xc8,
	0xb0, 0x92, 0x2c, 0x08, 0x21, 0x47, 0x7f, 0x34, 0x66, 0x53, 0x0e, 0x7e, 0x39, 0x01, 0x5a, 0xf5,
	0x6e, 0xbd, 0xd7, 0xea, 0x5f, 0x56, 0x86, 0xc9, 0x82, 0x68, 0x32, 0x1a, 0x40, 0x8e, 0x0f, 0x4a,
	0x35, 0x24, 0x11, 0xf9, 0x77, 0xc8, 0x73, 0xbb, 0x88, 0xe6, 0x93, 0x71, 0x50, 0xda, 0xfb, 0x6a,
	0x23, 0x56, 0x43, 0x65, 0x3e, 0xab, 0xec, 0x41, 0xcc, 0x62, 0x47, 0xe5, 0xdc, 0xfb, 0xc9, 0x16,
	0x78, 0x37, 0x5f, 0x3a, 0xfa, 0x62, 0xe9, 0xe8, 0xbf, 0x4b, 0x47, 0xff, 0x5c, 0x39, 0xda, 0x62,
	0xe5, 0x68, 0xdf, 0x2b, 0x47, 0x7b, 0xbb, 0xe0, 0x13, 0x39, 0xce, 0x02, 0x77, 0x24, 0x62, 0x0f,
	0x39, 0x5c, 0x91, 0x79, 0xf1, 0xf6, 0x3e, 0xd4, 0x45, 0x64, 0x9e, 0x00, 0x06, 0x4d, 0x75, 0x92,
	0xeb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x9e, 0xd5, 0xa7, 0x31, 0x02, 0x00, 0x00,
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
	{
		size, err := m.ProposalStats.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.PubkeysChangeProposals) > 0 {
		for iNdEx := len(m.PubkeysChangeProposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PubkeysChangeProposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.KeyVault.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
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
	l = m.KeyVault.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.PubkeysChangeProposals) > 0 {
		for _, e := range m.PubkeysChangeProposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.ProposalStats.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field KeyVault", wireType)
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
			if err := m.KeyVault.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubkeysChangeProposals", wireType)
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
			m.PubkeysChangeProposals = append(m.PubkeysChangeProposals, PublicKeysChangeProposal{})
			if err := m.PubkeysChangeProposals[len(m.PubkeysChangeProposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalStats", wireType)
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
			if err := m.ProposalStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
