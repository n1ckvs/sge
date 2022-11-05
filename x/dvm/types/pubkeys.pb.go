// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/dvm/pubkeys.proto

package types

import (
	fmt "fmt"
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

type PublicKeys struct {
	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (m *PublicKeys) Reset()         { *m = PublicKeys{} }
func (m *PublicKeys) String() string { return proto.CompactTextString(m) }
func (*PublicKeys) ProtoMessage()    {}
func (*PublicKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca49d8561b21a3b5, []int{0}
}
func (m *PublicKeys) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublicKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublicKeys.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKeys.Merge(m, src)
}
func (m *PublicKeys) XXX_Size() int {
	return m.Size()
}
func (m *PublicKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKeys.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKeys proto.InternalMessageInfo

func (m *PublicKeys) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*PublicKeys)(nil), "sgenetwork.sge.dvm.PublicKeys")
}

func init() { proto.RegisterFile("sge/dvm/pubkeys.proto", fileDescriptor_ca49d8561b21a3b5) }

var fileDescriptor_ca49d8561b21a3b5 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x4f, 0xd5,
	0x4f, 0x29, 0xcb, 0xd5, 0x2f, 0x28, 0x4d, 0xca, 0x4e, 0xad, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x2a, 0x4e, 0x4f, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0x4e,
	0x4f, 0xd5, 0x4b, 0x29, 0xcb, 0x55, 0x52, 0xe0, 0xe2, 0x0a, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xf6,
	0x4e, 0xad, 0x2c, 0x16, 0x12, 0xe2, 0x62, 0xc9, 0xc9, 0x2c, 0x2e, 0x91, 0x60, 0x54, 0x60, 0xd6,
	0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0x1c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1,
	0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21,
	0x4a, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x38, 0x3d, 0x55,
	0x17, 0x6a, 0x36, 0x88, 0xad, 0x5f, 0x01, 0xb6, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d,
	0x6c, 0xbd, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x13, 0x31, 0x83, 0xcd, 0x97, 0x00, 0x00, 0x00,
}

func (m *PublicKeys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicKeys) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKeys) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.List[iNdEx])
			copy(dAtA[i:], m.List[iNdEx])
			i = encodeVarintPubkeys(dAtA, i, uint64(len(m.List[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPubkeys(dAtA []byte, offset int, v uint64) int {
	offset -= sovPubkeys(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PublicKeys) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, s := range m.List {
			l = len(s)
			n += 1 + l + sovPubkeys(uint64(l))
		}
	}
	return n
}

func sovPubkeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPubkeys(x uint64) (n int) {
	return sovPubkeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PublicKeys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPubkeys
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
			return fmt.Errorf("proto: PublicKeys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicKeys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkeys
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
				return ErrInvalidLengthPubkeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubkeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPubkeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPubkeys
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
func skipPubkeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPubkeys
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
					return 0, ErrIntOverflowPubkeys
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
					return 0, ErrIntOverflowPubkeys
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
				return 0, ErrInvalidLengthPubkeys
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPubkeys
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPubkeys
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPubkeys        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPubkeys          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPubkeys = fmt.Errorf("proto: unexpected end of group")
)