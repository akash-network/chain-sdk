// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1/bid.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// BidID stores owner and all other seq numbers.
// A successful bid becomes a Lease(ID).
type BidID struct {
	// Owner is the account bech32 address of the user who owns the deployment.
	// It is a string representing a valid bech32 account address.
	//
	// Example:
	//   "akash1..."
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	// Dseq (deployment sequence number) is a unique numeric identifier for the deployment.
	// It is used to differentiate deployments created by the same owner.
	DSeq uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	// Gseq (group sequence number) is a unique numeric identifier for the group.
	// It is used to differentiate groups created by the same owner in a deployment.
	GSeq uint32 `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	// Oseq (order sequence) distinguishes multiple orders associated with a single deployment.
	// Oseq is incremented when a lease associated with an existing deployment is closed, and a new order is generated.
	OSeq uint32 `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
	// Provider is the account bech32 address of the provider making the bid.
	// It is a string representing a valid account bech32 address.
	//
	// Example:
	//   "akash1..."
	Provider string `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider" yaml:"provider"`
}

func (m *BidID) Reset()      { *m = BidID{} }
func (*BidID) ProtoMessage() {}
func (*BidID) Descriptor() ([]byte, []int) {
	return fileDescriptor_3938cb3dd8faff6a, []int{0}
}
func (m *BidID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BidID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BidID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BidID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidID.Merge(m, src)
}
func (m *BidID) XXX_Size() int {
	return m.Size()
}
func (m *BidID) XXX_DiscardUnknown() {
	xxx_messageInfo_BidID.DiscardUnknown(m)
}

var xxx_messageInfo_BidID proto.InternalMessageInfo

func (m *BidID) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BidID) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *BidID) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *BidID) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

func (m *BidID) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func init() {
	proto.RegisterType((*BidID)(nil), "akash.market.v1.BidID")
}

func init() { proto.RegisterFile("akash/market/v1/bid.proto", fileDescriptor_3938cb3dd8faff6a) }

var fileDescriptor_3938cb3dd8faff6a = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd1, 0xb1, 0x6a, 0xc2, 0x40,
	0x18, 0x07, 0xf0, 0x9c, 0x8d, 0xc5, 0xa6, 0x2d, 0x42, 0x70, 0x50, 0xa1, 0x39, 0xc9, 0xe4, 0xd2,
	0x04, 0x71, 0x10, 0xdc, 0x1a, 0x04, 0xe9, 0x54, 0xd0, 0xad, 0x4b, 0x89, 0xbd, 0xe3, 0x1a, 0x52,
	0xfd, 0xf4, 0x12, 0x52, 0xfa, 0x16, 0x1d, 0x3b, 0xfa, 0x10, 0x7d, 0x88, 0x4e, 0x45, 0x3a, 0x75,
	0x3a, 0x4a, 0x5c, 0x8a, 0xa3, 0x4f, 0x50, 0xee, 0x2e, 0x55, 0x1c, 0x3a, 0x25, 0xdf, 0xff, 0x7f,
	0xbf, 0x83, 0xe3, 0xb3, 0x1a, 0x61, 0x1c, 0x26, 0x0f, 0xfe, 0x34, 0xe4, 0x31, 0x4d, 0xfd, 0xac,
	0xe3, 0x4f, 0x22, 0xe2, 0xcd, 0x39, 0xa4, 0x60, 0x57, 0x55, 0xe5, 0xe9, 0xca, 0xcb, 0x3a, 0xcd,
	0x1a, 0x03, 0x06, 0xaa, 0xf3, 0xe5, 0x9f, 0x3e, 0xd6, 0x6c, 0xdc, 0x43, 0x32, 0x85, 0xe4, 0x4e,
	0x17, 0x7a, 0xd0, 0x95, 0xfb, 0x51, 0xb2, 0xca, 0x41, 0x44, 0xae, 0x07, 0xf6, 0xd0, 0x2a, 0xc3,
	0xd3, 0x8c, 0xf2, 0x3a, 0x6a, 0xa1, 0xf6, 0x49, 0xd0, 0xd9, 0x08, 0xac, 0x83, 0xad, 0xc0, 0x67,
	0xcf, 0xe1, 0xf4, 0xb1, 0xef, 0xaa, 0xd1, 0xfd, 0x7c, 0xbb, 0xac, 0x15, 0x77, 0x5c, 0x11, 0xc2,
	0x69, 0x92, 0x8c, 0x53, 0x1e, 0xcd, 0xd8, 0x48, 0x1f, 0xb7, 0xbb, 0x96, 0x49, 0x12, 0xba, 0xa8,
	0x97, 0x5a, 0xa8, 0x6d, 0x06, 0x38, 0x17, 0xd8, 0x1c, 0x8c, 0xe9, 0x62, 0x23, 0xb0, 0xca, 0xb7,
	0x02, 0x9f, 0xea, 0xeb, 0xe4, 0xe4, 0x8e, 0x54, 0x28, 0x11, 0x93, 0xe8, 0xa8, 0x85, 0xda, 0xe7,
	0x1a, 0x0d, 0x0b, 0xc4, 0x0e, 0x10, 0xd3, 0x88, 0x15, 0x08, 0x24, 0x32, 0xf7, 0xe8, 0xa6, 0x40,
	0x70, 0x80, 0x40, 0x23, 0xf9, 0xb1, 0xc7, 0x56, 0x65, 0xce, 0x21, 0x8b, 0x08, 0xe5, 0xf5, 0xb2,
	0x7a, 0x6a, 0x6f, 0x23, 0xf0, 0x2e, 0xdb, 0x0a, 0x5c, 0xd5, 0xe8, 0x2f, 0xf9, 0xff, 0xc1, 0x3b,
	0xd4, 0xaf, 0xbc, 0x2e, 0xb1, 0xf1, 0xb3, 0xc4, 0x46, 0xd0, 0x7b, 0xcf, 0x1d, 0xb4, 0xca, 0x1d,
	0xf4, 0x9d, 0x3b, 0xe8, 0x65, 0xed, 0x18, 0xab, 0xb5, 0x63, 0x7c, 0xad, 0x1d, 0xe3, 0xf6, 0x62,
	0x1e, 0x33, 0x2f, 0x8c, 0x53, 0x8f, 0xd0, 0xcc, 0x67, 0xe0, 0xcf, 0x80, 0xd0, 0xfd, 0x56, 0x27,
	0xc7, 0x6a, 0x21, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x43, 0xac, 0x89, 0xef, 0x01,
	0x00, 0x00,
}

func (m *BidID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BidID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BidID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintBid(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x2a
	}
	if m.OSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintBid(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBid(dAtA []byte, offset int, v uint64) int {
	offset -= sovBid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BidID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovBid(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovBid(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovBid(uint64(m.OSeq))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	return n
}

func sovBid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBid(x uint64) (n int) {
	return sovBid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BidID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBid
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
			return fmt.Errorf("proto: BidID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BidID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSeq", wireType)
			}
			m.OSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBid
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
func skipBid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBid
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
					return 0, ErrIntOverflowBid
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
					return 0, ErrIntOverflowBid
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
				return 0, ErrInvalidLengthBid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBid = fmt.Errorf("proto: unexpected end of group")
)
