// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/orderbook/orderbook.proto

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

// OrderBookStatus is the enum type for the status of the order book.
type OrderBookStatus int32

const (
	// invalid
	OrderBookStatus_ORDER_BOOK_STATUS_UNSPECIFIED OrderBookStatus = 0
	// active
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_ACTIVE OrderBookStatus = 1
	// resolved not settled
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_RESOLVED OrderBookStatus = 2
	// resolved and settled
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_SETTLED OrderBookStatus = 3
)

var OrderBookStatus_name = map[int32]string{
	0: "ORDER_BOOK_STATUS_UNSPECIFIED",
	1: "ORDER_BOOK_STATUS_STATUS_ACTIVE",
	2: "ORDER_BOOK_STATUS_STATUS_RESOLVED",
	3: "ORDER_BOOK_STATUS_STATUS_SETTLED",
}

var OrderBookStatus_value = map[string]int32{
	"ORDER_BOOK_STATUS_UNSPECIFIED":     0,
	"ORDER_BOOK_STATUS_STATUS_ACTIVE":   1,
	"ORDER_BOOK_STATUS_STATUS_RESOLVED": 2,
	"ORDER_BOOK_STATUS_STATUS_SETTLED":  3,
}

func (x OrderBookStatus) String() string {
	return proto.EnumName(OrderBookStatus_name, int32(x))
}

func (OrderBookStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7247ccc164993ca5, []int{0}
}

// OrderBook represents the order book maintained against a market.
type OrderBook struct {
	// uid is the universal unique identifier of the order book.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// participation_count is the count of participations in the order book.
	ParticipationCount uint64 `protobuf:"varint,2,opt,name=participation_count,json=participationCount,proto3" json:"participation_count,omitempty" yaml:"participation_count"`
	// odds_count is the count of the odds in the order book.
	OddsCount uint64 `protobuf:"varint,3,opt,name=odds_count,json=oddsCount,proto3" json:"odds_count,omitempty" yaml:"odds_count"`
	// status represents the status of the order book.
	Status OrderBookStatus `protobuf:"varint,4,opt,name=status,proto3,enum=sgenetwork.sge.orderbook.OrderBookStatus" json:"status,omitempty"`
}

func (m *OrderBook) Reset()      { *m = OrderBook{} }
func (*OrderBook) ProtoMessage() {}
func (*OrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_7247ccc164993ca5, []int{0}
}

func (m *OrderBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *OrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *OrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBook.Merge(m, src)
}

func (m *OrderBook) XXX_Size() int {
	return m.Size()
}

func (m *OrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBook proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("sgenetwork.sge.orderbook.OrderBookStatus", OrderBookStatus_name, OrderBookStatus_value)
	proto.RegisterType((*OrderBook)(nil), "sgenetwork.sge.orderbook.OrderBook")
}

func init() { proto.RegisterFile("sge/orderbook/orderbook.proto", fileDescriptor_7247ccc164993ca5) }

var fileDescriptor_7247ccc164993ca5 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x4e, 0x4f, 0xd5,
	0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x4a, 0xca, 0xcf, 0xcf, 0x46, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x24, 0x8a, 0xd3, 0x53, 0xf3, 0x52, 0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0xf5, 0x8a, 0xd3,
	0x53, 0xf5, 0xe0, 0xf2, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x45, 0xfa, 0x20, 0x16, 0x44,
	0xbd, 0x52, 0x2b, 0x13, 0x17, 0xa7, 0x3f, 0x48, 0x8d, 0x53, 0x7e, 0x7e, 0xb6, 0x90, 0x02, 0x17,
	0x73, 0x69, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x13, 0xdf, 0xa3, 0x7b, 0xf2, 0xcc,
	0xa1, 0x9e, 0x2e, 0xaf, 0xee, 0xc9, 0x83, 0x44, 0x83, 0x40, 0x84, 0x90, 0x3f, 0x97, 0x70, 0x41,
	0x62, 0x51, 0x49, 0x66, 0x72, 0x66, 0x41, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x7c, 0x72, 0x7e, 0x69,
	0x5e, 0x89, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x8b, 0x93, 0xdc, 0xa7, 0x7b, 0xf2, 0x52, 0x95, 0x89,
	0xb9, 0x39, 0x56, 0x4a, 0x58, 0x14, 0x29, 0x05, 0x09, 0xa1, 0x88, 0x3a, 0x83, 0x04, 0x85, 0x4c,
	0xb8, 0xb8, 0xf2, 0x53, 0x52, 0x8a, 0xa1, 0xe6, 0x30, 0x83, 0xcd, 0x11, 0xfd, 0x74, 0x4f, 0x5e,
	0x10, 0x62, 0x0e, 0x42, 0x4e, 0x29, 0x88, 0x13, 0xc4, 0x81, 0xe8, 0x72, 0xe4, 0x62, 0x2b, 0x2e,
	0x49, 0x2c, 0x29, 0x2d, 0x96, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x33, 0xd2, 0xd4, 0xc3, 0xe5, 0x6f,
	0x3d, 0xb8, 0xef, 0x82, 0xc1, 0x1a, 0x82, 0xa0, 0x1a, 0xad, 0x78, 0x3a, 0x16, 0xc8, 0x33, 0xcc,
	0x58, 0x20, 0xcf, 0xf0, 0x62, 0x81, 0x3c, 0x83, 0xd6, 0x32, 0x46, 0x2e, 0x7e, 0x34, 0x95, 0x42,
	0x8a, 0x5c, 0xb2, 0xfe, 0x41, 0x2e, 0xae, 0x41, 0xf1, 0x4e, 0xfe, 0xfe, 0xde, 0xf1, 0xc1, 0x21,
	0x8e, 0x21, 0xa1, 0xc1, 0xf1, 0xa1, 0x7e, 0xc1, 0x01, 0xae, 0xce, 0x9e, 0x6e, 0x9e, 0xae, 0x2e,
	0x02, 0x0c, 0x42, 0xca, 0x5c, 0xf2, 0x98, 0x4a, 0xa0, 0x94, 0xa3, 0x73, 0x88, 0x67, 0x98, 0xab,
	0x00, 0xa3, 0x90, 0x2a, 0x97, 0x22, 0x4e, 0x45, 0x41, 0xae, 0xc1, 0xfe, 0x3e, 0x61, 0xae, 0x2e,
	0x02, 0x4c, 0x42, 0x2a, 0x5c, 0x0a, 0x38, 0x95, 0x05, 0xbb, 0x86, 0x84, 0xf8, 0xb8, 0xba, 0x08,
	0x30, 0x3b, 0xb9, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x4e, 0x7a,
	0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x71, 0x7a, 0xaa, 0x2e, 0x34, 0x38,
	0x40, 0x6c, 0xfd, 0x0a, 0xa4, 0x24, 0x53, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x8e, 0x7f,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x78, 0x55, 0x09, 0x50, 0x02, 0x00, 0x00,
}

func (m *OrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.OddsCount != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.OddsCount))
		i--
		dAtA[i] = 0x18
	}
	if m.ParticipationCount != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.ParticipationCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintOrderbook(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrderbook(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderbook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *OrderBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovOrderbook(uint64(l))
	}
	if m.ParticipationCount != 0 {
		n += 1 + sovOrderbook(uint64(m.ParticipationCount))
	}
	if m.OddsCount != 0 {
		n += 1 + sovOrderbook(uint64(m.OddsCount))
	}
	if m.Status != 0 {
		n += 1 + sovOrderbook(uint64(m.Status))
	}
	return n
}

func sovOrderbook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozOrderbook(x uint64) (n int) {
	return sovOrderbook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *OrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderbook
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
			return fmt.Errorf("proto: OrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
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
				return ErrInvalidLengthOrderbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationCount", wireType)
			}
			m.ParticipationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipationCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsCount", wireType)
			}
			m.OddsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OddsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OrderBookStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrderbook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderbook
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

func skipOrderbook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderbook
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
					return 0, ErrIntOverflowOrderbook
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
					return 0, ErrIntOverflowOrderbook
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
				return 0, ErrInvalidLengthOrderbook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderbook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderbook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderbook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderbook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderbook = fmt.Errorf("proto: unexpected end of group")
)
