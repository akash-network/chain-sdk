// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/base/v1beta3/endpoint.proto

package v1beta3

import (
	fmt "fmt"
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

// This describes how the endpoint is implemented when the lease is deployed
type Endpoint_Kind int32

const (
	// Describes an endpoint that becomes a Kubernetes Ingress
	Endpoint_SHARED_HTTP Endpoint_Kind = 0
	// Describes an endpoint that becomes a Kubernetes NodePort
	Endpoint_RANDOM_PORT Endpoint_Kind = 1
	// Describes an endpoint that becomes a leased IP
	Endpoint_LEASED_IP Endpoint_Kind = 2
)

var Endpoint_Kind_name = map[int32]string{
	0: "SHARED_HTTP",
	1: "RANDOM_PORT",
	2: "LEASED_IP",
}

var Endpoint_Kind_value = map[string]int32{
	"SHARED_HTTP": 0,
	"RANDOM_PORT": 1,
	"LEASED_IP":   2,
}

func (x Endpoint_Kind) String() string {
	return proto.EnumName(Endpoint_Kind_name, int32(x))
}

func (Endpoint_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f12485ed8bcdc6e7, []int{0, 0}
}

// Endpoint describes a publicly accessible IP service
type Endpoint struct {
	Kind           Endpoint_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=akash.base.v1beta3.Endpoint_Kind" json:"kind,omitempty"`
	SequenceNumber uint32        `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number" yaml:"sequence_number"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_f12485ed8bcdc6e7, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return m.Size()
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetKind() Endpoint_Kind {
	if m != nil {
		return m.Kind
	}
	return Endpoint_SHARED_HTTP
}

func (m *Endpoint) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func init() {
	proto.RegisterEnum("akash.base.v1beta3.Endpoint_Kind", Endpoint_Kind_name, Endpoint_Kind_value)
	proto.RegisterType((*Endpoint)(nil), "akash.base.v1beta3.Endpoint")
}

func init() { proto.RegisterFile("akash/base/v1beta3/endpoint.proto", fileDescriptor_f12485ed8bcdc6e7) }

var fileDescriptor_f12485ed8bcdc6e7 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xcc, 0x4e, 0x2c,
	0xce, 0xd0, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd6, 0x4f,
	0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02,
	0x2b, 0xd1, 0x03, 0x29, 0xd1, 0x83, 0x2a, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xeb,
	0x83, 0x58, 0x10, 0x95, 0x4a, 0xaf, 0x18, 0xb9, 0x38, 0x5c, 0xa1, 0x9a, 0x85, 0x4c, 0xb9, 0x58,
	0xb2, 0x33, 0xf3, 0x52, 0x24, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x14, 0xf5, 0x30, 0x4d, 0xd1,
	0x83, 0xa9, 0xd5, 0xf3, 0xce, 0xcc, 0x4b, 0x09, 0x02, 0x2b, 0x17, 0xca, 0xe0, 0xe2, 0x2f, 0x4e,
	0x2d, 0x2c, 0x4d, 0xcd, 0x4b, 0x4e, 0x8d, 0xcf, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x92, 0x60, 0x52,
	0x60, 0xd4, 0xe0, 0x75, 0xb2, 0x7f, 0x74, 0x4f, 0x9e, 0x2f, 0x18, 0x2a, 0xe5, 0x07, 0x96, 0x79,
	0x75, 0x4f, 0x1e, 0x5d, 0xf1, 0xa7, 0x7b, 0xf2, 0x62, 0x95, 0x89, 0xb9, 0x39, 0x56, 0x4a, 0x68,
	0x12, 0x4a, 0x41, 0x7c, 0xc5, 0x28, 0x9a, 0x95, 0xcc, 0xb9, 0x58, 0x40, 0xf6, 0x0a, 0xf1, 0x73,
	0x71, 0x07, 0x7b, 0x38, 0x06, 0xb9, 0xba, 0xc4, 0x7b, 0x84, 0x84, 0x04, 0x08, 0x30, 0x80, 0x04,
	0x82, 0x1c, 0xfd, 0x5c, 0xfc, 0x7d, 0xe3, 0x03, 0xfc, 0x83, 0x42, 0x04, 0x18, 0x85, 0x78, 0xb9,
	0x38, 0x7d, 0x5c, 0x1d, 0x83, 0x5d, 0x5d, 0xe2, 0x3d, 0x03, 0x04, 0x98, 0xac, 0x58, 0x5e, 0x2c,
	0x90, 0x67, 0x74, 0xb2, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xc5,
	0x82, 0xec, 0x74, 0xbd, 0xc4, 0xec, 0x12, 0xbd, 0x94, 0xd4, 0x32, 0xfd, 0xf4, 0x7c, 0xfd, 0xbc,
	0xfc, 0x94, 0x54, 0xfd, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x58, 0x10, 0x27, 0xb1, 0x81, 0x03, 0xcc,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x29, 0x79, 0x7e, 0x7f, 0x01, 0x00, 0x00,
}

func (this *Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint)
	if !ok {
		that2, ok := that.(Endpoint)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	if this.SequenceNumber != that1.SequenceNumber {
		return false
	}
	return true
}
func (m *Endpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Endpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Endpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SequenceNumber != 0 {
		i = encodeVarintEndpoint(dAtA, i, uint64(m.SequenceNumber))
		i--
		dAtA[i] = 0x10
	}
	if m.Kind != 0 {
		i = encodeVarintEndpoint(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEndpoint(dAtA []byte, offset int, v uint64) int {
	offset -= sovEndpoint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Endpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovEndpoint(uint64(m.Kind))
	}
	if m.SequenceNumber != 0 {
		n += 1 + sovEndpoint(uint64(m.SequenceNumber))
	}
	return n
}

func sovEndpoint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEndpoint(x uint64) (n int) {
	return sovEndpoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Endpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: Endpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Endpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= Endpoint_Kind(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNumber", wireType)
			}
			m.SequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEndpoint
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
func skipEndpoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
				return 0, ErrInvalidLengthEndpoint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEndpoint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEndpoint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEndpoint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEndpoint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEndpoint = fmt.Errorf("proto: unexpected end of group")
)
