// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/strategicreserve/genesis.proto

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

// GenesisState defines the strategicreserve module's genesis state.
type GenesisState struct {
	// params defines all the parameters related to orderbook.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// book_list defines the books active at genesis.
	BookList []OrderBook `protobuf:"bytes,2,rep,name=book_list,json=bookList,proto3" json:"book_list"`
	// book_participation_list defines the book participations active at genesis.
	BookParticipationList []BookParticipation `protobuf:"bytes,3,rep,name=book_participation_list,json=bookParticipationList,proto3" json:"book_participation_list"`
	// book_exposure_list defines the book exposures active at genesis.
	BookExposureList []BookOddsExposure `protobuf:"bytes,4,rep,name=book_exposure_list,json=bookExposureList,proto3" json:"book_exposure_list"`
	// participation_exposure_list defines the participation exposures active at
	// genesis.
	ParticipationExposureList []ParticipationExposure `protobuf:"bytes,5,rep,name=participation_exposure_list,json=participationExposureList,proto3" json:"participation_exposure_list"`
	// participation_exposure_by_index_list defines the participation exposures by
	// index active at genesis.
	ParticipationExposureByIndexList []ParticipationExposure `protobuf:"bytes,6,rep,name=participation_exposure_by_index_list,json=participationExposureByIndexList,proto3" json:"participation_exposure_by_index_list"`
	// historical_participation_exposure_list defines the historical participation
	// exposures active at
	// genesis.
	HistoricalParticipationExposureList []ParticipationExposure `protobuf:"bytes,7,rep,name=historical_participation_exposure_list,json=historicalParticipationExposureList,proto3" json:"historical_participation_exposure_list"`
	// historical_participation_exposure_list defines the historical participation
	// exposures active at
	// genesis.
	ParticipationBetPairExposureList []ParticipationBetPair `protobuf:"bytes,8,rep,name=participation_bet_pair_exposure_list,json=participationBetPairExposureList,proto3" json:"participation_bet_pair_exposure_list"`
	PayoutLock                       [][]byte               `protobuf:"bytes,9,rep,name=payout_lock,json=payoutLock,proto3" json:"payout_lock,omitempty"`
	// stats is the statistics of the order-book
	Stats OrderBookStats `protobuf:"bytes,10,opt,name=stats,proto3" json:"stats"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_708845e0eb254a47, []int{0}
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

func (m *GenesisState) GetBookList() []OrderBook {
	if m != nil {
		return m.BookList
	}
	return nil
}

func (m *GenesisState) GetBookParticipationList() []BookParticipation {
	if m != nil {
		return m.BookParticipationList
	}
	return nil
}

func (m *GenesisState) GetBookExposureList() []BookOddsExposure {
	if m != nil {
		return m.BookExposureList
	}
	return nil
}

func (m *GenesisState) GetParticipationExposureList() []ParticipationExposure {
	if m != nil {
		return m.ParticipationExposureList
	}
	return nil
}

func (m *GenesisState) GetParticipationExposureByIndexList() []ParticipationExposure {
	if m != nil {
		return m.ParticipationExposureByIndexList
	}
	return nil
}

func (m *GenesisState) GetHistoricalParticipationExposureList() []ParticipationExposure {
	if m != nil {
		return m.HistoricalParticipationExposureList
	}
	return nil
}

func (m *GenesisState) GetParticipationBetPairExposureList() []ParticipationBetPair {
	if m != nil {
		return m.ParticipationBetPairExposureList
	}
	return nil
}

func (m *GenesisState) GetPayoutLock() [][]byte {
	if m != nil {
		return m.PayoutLock
	}
	return nil
}

func (m *GenesisState) GetStats() OrderBookStats {
	if m != nil {
		return m.Stats
	}
	return OrderBookStats{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sgenetwork.sge.strategicreserve.GenesisState")
}

func init() {
	proto.RegisterFile("sge/strategicreserve/genesis.proto", fileDescriptor_708845e0eb254a47)
}

var fileDescriptor_708845e0eb254a47 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x63, 0xd2, 0x86, 0x74, 0xd2, 0x05, 0xb2, 0x40, 0x84, 0x20, 0xb9, 0xa6, 0x2d, 0x10,
	0x21, 0x61, 0x8b, 0x54, 0xf0, 0x00, 0x96, 0x2a, 0x84, 0x68, 0xd5, 0x08, 0x76, 0x6c, 0xac, 0xb1,
	0x73, 0xe5, 0x8e, 0x9c, 0x66, 0xac, 0x99, 0x1b, 0x48, 0x24, 0xde, 0x00, 0x16, 0xf0, 0x06, 0x3c,
	0x4e, 0x97, 0x5d, 0xb2, 0x42, 0x28, 0x79, 0x11, 0x34, 0x3f, 0x51, 0x49, 0xea, 0x90, 0x20, 0x65,
	0xe7, 0x64, 0xce, 0x3d, 0xdf, 0xf1, 0xb1, 0xe6, 0x92, 0x7d, 0x99, 0x41, 0x28, 0x51, 0x50, 0x84,
	0x8c, 0xa5, 0x02, 0x24, 0x88, 0x8f, 0x10, 0x66, 0x30, 0x00, 0xc9, 0x64, 0x50, 0x08, 0x8e, 0xdc,
	0xdd, 0x93, 0xea, 0x37, 0x7e, 0xe2, 0x22, 0x0f, 0x64, 0x06, 0xc1, 0xa2, 0xbc, 0x75, 0x37, 0xe3,
	0x19, 0xd7, 0xda, 0x50, 0x3d, 0x99, 0xb1, 0xd6, 0xa3, 0x52, 0xeb, 0x82, 0x0a, 0x7a, 0x61, 0x9d,
	0x5b, 0x87, 0xa5, 0x12, 0x2e, 0x7a, 0x20, 0x12, 0xce, 0x73, 0xab, 0x6a, 0x2f, 0x33, 0x42, 0x96,
	0xb2, 0x82, 0x22, 0xe3, 0x03, 0xab, 0xf4, 0x4b, 0x95, 0x12, 0x29, 0xce, 0x88, 0x07, 0xa5, 0x0a,
	0x18, 0x15, 0x5c, 0x0e, 0x05, 0x18, 0xd1, 0xfe, 0x8f, 0x3a, 0xd9, 0x7d, 0x6d, 0x2a, 0x78, 0x8f,
	0x14, 0xc1, 0x3d, 0x26, 0x35, 0x93, 0xbb, 0xe9, 0xf8, 0x4e, 0xbb, 0xd1, 0x79, 0x1a, 0xac, 0xa8,
	0x24, 0xe8, 0x6a, 0x79, 0xb4, 0x75, 0xf9, 0x6b, 0xaf, 0xf2, 0xce, 0x0e, 0xbb, 0xa7, 0x64, 0x47,
	0xbd, 0x56, 0xdc, 0x67, 0x12, 0x9b, 0xb7, 0xfc, 0x6a, 0xbb, 0xd1, 0x79, 0xb6, 0xd2, 0xe9, 0x4c,
	0xb5, 0x11, 0x71, 0x9e, 0x5b, 0xb3, 0xba, 0xb2, 0x38, 0x61, 0x12, 0xdd, 0x82, 0xdc, 0xd7, 0x76,
	0x73, 0x4d, 0x18, 0xf3, 0xaa, 0x36, 0xef, 0xac, 0x34, 0x57, 0xbe, 0xdd, 0xbf, 0xc7, 0x2d, 0xe4,
	0x5e, 0xb2, 0x78, 0xa0, 0x89, 0x40, 0x5c, 0x4d, 0x9c, 0xf5, 0x65, 0x60, 0x5b, 0x1a, 0xf6, 0x62,
	0x2d, 0xd8, 0x59, 0xaf, 0x27, 0x8f, 0xed, 0xb4, 0x65, 0xdd, 0x51, 0x96, 0xb3, 0xff, 0x34, 0xe6,
	0x33, 0x79, 0x38, 0xff, 0x4e, 0xf3, 0xbc, 0x6d, 0xcd, 0x7b, 0xb5, 0xce, 0x37, 0xb8, 0xf6, 0x58,
	0x80, 0x3e, 0x28, 0xca, 0x0e, 0x35, 0xfd, 0xab, 0x43, 0x0e, 0x97, 0xe0, 0x93, 0x71, 0xcc, 0x06,
	0x3d, 0x18, 0x99, 0x1c, 0xb5, 0x0d, 0xe4, 0xf0, 0x4b, 0x73, 0x44, 0xe3, 0x37, 0x0a, 0xa3, 0xe3,
	0x7c, 0x77, 0xc8, 0x93, 0x73, 0x26, 0x91, 0x0b, 0x96, 0xd2, 0x7e, 0xfc, 0xaf, 0x62, 0x6e, 0x6f,
	0x20, 0xd0, 0xc1, 0x35, 0xab, 0xbb, 0xb4, 0xa2, 0x2f, 0x37, 0x2a, 0x4a, 0x00, 0xe3, 0x82, 0x32,
	0xb1, 0x90, 0xa8, 0xae, 0x13, 0xbd, 0xfc, 0xbf, 0x44, 0x11, 0x60, 0x97, 0x32, 0x51, 0xda, 0x90,
	0x3d, 0x9b, 0x4b, 0xf3, 0x98, 0x34, 0x0a, 0x3a, 0xe6, 0x43, 0x8c, 0xfb, 0x3c, 0xcd, 0x9b, 0x3b,
	0x7e, 0xb5, 0xbd, 0x6b, 0x87, 0x89, 0x39, 0x38, 0xe1, 0x69, 0xee, 0xbe, 0x25, 0xdb, 0x7a, 0x13,
	0x34, 0x89, 0xbe, 0xc3, 0xe1, 0xfa, 0x37, 0x4f, 0x2d, 0x81, 0xd9, 0x5d, 0x36, 0x1e, 0xd1, 0xe9,
	0xe5, 0xc4, 0x73, 0xae, 0x26, 0x9e, 0xf3, 0x7b, 0xe2, 0x39, 0xdf, 0xa6, 0x5e, 0xe5, 0x6a, 0xea,
	0x55, 0x7e, 0x4e, 0xbd, 0xca, 0x87, 0xa3, 0x8c, 0xe1, 0xf9, 0x30, 0x09, 0x52, 0x7e, 0x11, 0xca,
	0x0c, 0x9e, 0x5b, 0x84, 0x7a, 0x0e, 0x47, 0x37, 0x57, 0x0f, 0x8e, 0x0b, 0x90, 0x49, 0x4d, 0x2f,
	0x9e, 0xa3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x2f, 0x88, 0x26, 0x8f, 0x05, 0x00, 0x00,
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
		size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if len(m.PayoutLock) > 0 {
		for iNdEx := len(m.PayoutLock) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PayoutLock[iNdEx])
			copy(dAtA[i:], m.PayoutLock[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.PayoutLock[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ParticipationBetPairExposureList) > 0 {
		for iNdEx := len(m.ParticipationBetPairExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipationBetPairExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.HistoricalParticipationExposureList) > 0 {
		for iNdEx := len(m.HistoricalParticipationExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HistoricalParticipationExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ParticipationExposureByIndexList) > 0 {
		for iNdEx := len(m.ParticipationExposureByIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipationExposureByIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ParticipationExposureList) > 0 {
		for iNdEx := len(m.ParticipationExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipationExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BookExposureList) > 0 {
		for iNdEx := len(m.BookExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BookExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.BookParticipationList) > 0 {
		for iNdEx := len(m.BookParticipationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BookParticipationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.BookList) > 0 {
		for iNdEx := len(m.BookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.BookList) > 0 {
		for _, e := range m.BookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BookParticipationList) > 0 {
		for _, e := range m.BookParticipationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BookExposureList) > 0 {
		for _, e := range m.BookExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ParticipationExposureList) > 0 {
		for _, e := range m.ParticipationExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ParticipationExposureByIndexList) > 0 {
		for _, e := range m.ParticipationExposureByIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.HistoricalParticipationExposureList) > 0 {
		for _, e := range m.HistoricalParticipationExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ParticipationBetPairExposureList) > 0 {
		for _, e := range m.ParticipationBetPairExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PayoutLock) > 0 {
		for _, b := range m.PayoutLock {
			l = len(b)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Stats.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field BookList", wireType)
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
			m.BookList = append(m.BookList, OrderBook{})
			if err := m.BookList[len(m.BookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookParticipationList", wireType)
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
			m.BookParticipationList = append(m.BookParticipationList, BookParticipation{})
			if err := m.BookParticipationList[len(m.BookParticipationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookExposureList", wireType)
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
			m.BookExposureList = append(m.BookExposureList, BookOddsExposure{})
			if err := m.BookExposureList[len(m.BookExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationExposureList", wireType)
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
			m.ParticipationExposureList = append(m.ParticipationExposureList, ParticipationExposure{})
			if err := m.ParticipationExposureList[len(m.ParticipationExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationExposureByIndexList", wireType)
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
			m.ParticipationExposureByIndexList = append(m.ParticipationExposureByIndexList, ParticipationExposure{})
			if err := m.ParticipationExposureByIndexList[len(m.ParticipationExposureByIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoricalParticipationExposureList", wireType)
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
			m.HistoricalParticipationExposureList = append(m.HistoricalParticipationExposureList, ParticipationExposure{})
			if err := m.HistoricalParticipationExposureList[len(m.HistoricalParticipationExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationBetPairExposureList", wireType)
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
			m.ParticipationBetPairExposureList = append(m.ParticipationBetPairExposureList, ParticipationBetPair{})
			if err := m.ParticipationBetPairExposureList[len(m.ParticipationBetPairExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayoutLock", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayoutLock = append(m.PayoutLock, make([]byte, postIndex-iNdEx))
			copy(m.PayoutLock[len(m.PayoutLock)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
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
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
