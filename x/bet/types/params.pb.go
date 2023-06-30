// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/params.proto

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

// Params defines the parameters for the module.
type Params struct {
	// batch_settlement_count is the batch settlement bet count.
	BatchSettlementCount uint32 `protobuf:"varint,1,opt,name=batch_settlement_count,json=batchSettlementCount,proto3" json:"batch_settlement_count,omitempty"`
	// max_bet_by_uid_query_count is the maximum bet by uid query items count.
	MaxBetByUidQueryCount uint32 `protobuf:"varint,2,opt,name=max_bet_by_uid_query_count,json=maxBetByUidQueryCount,proto3" json:"max_bet_by_uid_query_count,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4216d2638a14c9d3, []int{0}
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

func (m *Params) GetBatchSettlementCount() uint32 {
	if m != nil {
		return m.BatchSettlementCount
	}
	return 0
}

func (m *Params) GetMaxBetByUidQueryCount() uint32 {
	if m != nil {
		return m.MaxBetByUidQueryCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "sgenetwork.sge.bet.Params")
}

func init() { proto.RegisterFile("sge/bet/params.proto", fileDescriptor_4216d2638a14c9d3) }

var fileDescriptor_4216d2638a14c9d3 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xbf, 0x4a, 0xc5, 0x30,
	0x18, 0x47, 0x1b, 0x91, 0x3b, 0x14, 0x5c, 0x4a, 0x15, 0xb9, 0x43, 0x14, 0x07, 0x71, 0xb1, 0x19,
	0x74, 0xd1, 0x49, 0xea, 0x0b, 0xf8, 0x07, 0x17, 0x97, 0x90, 0xf4, 0x7e, 0xe4, 0x16, 0x4d, 0x53,
	0x93, 0x2f, 0xd8, 0x4c, 0xbe, 0x82, 0xa3, 0xa3, 0x8f, 0xe3, 0x78, 0x47, 0x47, 0x69, 0x5f, 0x44,
	0x12, 0x8b, 0xdb, 0x0f, 0xce, 0xef, 0x0c, 0x27, 0x2f, 0x9d, 0x02, 0x26, 0x01, 0x59, 0x2f, 0xac,
	0xd0, 0xae, 0xea, 0xad, 0x41, 0x53, 0x14, 0x4e, 0x41, 0x07, 0xf8, 0x6a, 0xec, 0x53, 0xe5, 0x14,
	0x54, 0x12, 0x70, 0x59, 0x2a, 0xa3, 0x4c, 0xc2, 0x2c, 0xae, 0xbf, 0xe7, 0xd1, 0x5b, 0xbe, 0xb8,
	0x49, 0x66, 0x71, 0x9e, 0xef, 0x49, 0x81, 0xcd, 0x9a, 0x3b, 0x40, 0x7c, 0x06, 0x0d, 0x1d, 0xf2,
	0xc6, 0xf8, 0x0e, 0xf7, 0xc9, 0x21, 0x39, 0xd9, 0xb9, 0x2b, 0x13, 0xbd, 0xff, 0x87, 0xd7, 0x91,
	0x15, 0x17, 0xf9, 0x52, 0x8b, 0x81, 0x4b, 0x40, 0x2e, 0x03, 0xf7, 0xed, 0x8a, 0xbf, 0x78, 0xb0,
	0x61, 0x36, 0xb7, 0x92, 0xb9, 0xab, 0xc5, 0x50, 0x03, 0xd6, 0xe1, 0xa1, 0x5d, 0xdd, 0x46, 0x9a,
	0xd4, 0xcb, 0xed, 0x8f, 0xcf, 0x83, 0xac, 0xbe, 0xfa, 0x1a, 0x29, 0xd9, 0x8c, 0x94, 0xfc, 0x8c,
	0x94, 0xbc, 0x4f, 0x34, 0xdb, 0x4c, 0x34, 0xfb, 0x9e, 0x68, 0xf6, 0x78, 0xac, 0x5a, 0x5c, 0x7b,
	0x59, 0x35, 0x46, 0x33, 0xa7, 0xe0, 0x74, 0x0e, 0x8a, 0x9b, 0x0d, 0xa9, 0x19, 0x43, 0x0f, 0x4e,
	0x2e, 0x52, 0xc9, 0xd9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x0f, 0xc3, 0xe9, 0x0b, 0x01,
	0x00, 0x00,
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
	if m.MaxBetByUidQueryCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxBetByUidQueryCount))
		i--
		dAtA[i] = 0x10
	}
	if m.BatchSettlementCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BatchSettlementCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BatchSettlementCount != 0 {
		n += 1 + sovParams(uint64(m.BatchSettlementCount))
	}
	if m.MaxBetByUidQueryCount != 0 {
		n += 1 + sovParams(uint64(m.MaxBetByUidQueryCount))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field BatchSettlementCount", wireType)
			}
			m.BatchSettlementCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchSettlementCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBetByUidQueryCount", wireType)
			}
			m.MaxBetByUidQueryCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBetByUidQueryCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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

func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
