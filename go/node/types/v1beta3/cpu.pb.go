// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/base/v1beta3/cpu.proto

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

// CPU stores resource units and cpu config attributes
type CPU struct {
	Units      ResourceValue `protobuf:"bytes,1,opt,name=units,proto3" json:"units"`
	Attributes Attributes    `protobuf:"bytes,2,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

func (m *CPU) Reset()         { *m = CPU{} }
func (m *CPU) String() string { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()    {}
func (*CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ab47749737c9d3, []int{0}
}
func (m *CPU) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CPU.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPU.Merge(m, src)
}
func (m *CPU) XXX_Size() int {
	return m.Size()
}
func (m *CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CPU proto.InternalMessageInfo

func (m *CPU) GetUnits() ResourceValue {
	if m != nil {
		return m.Units
	}
	return ResourceValue{}
}

func (m *CPU) GetAttributes() Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*CPU)(nil), "akash.base.v1beta3.CPU")
}

func init() { proto.RegisterFile("akash/base/v1beta3/cpu.proto", fileDescriptor_80ab47749737c9d3) }

var fileDescriptor_80ab47749737c9d3 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xcc, 0x4e, 0x2c,
	0xce, 0xd0, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd6, 0x4f,
	0x2e, 0x28, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0xcb, 0xea, 0x81, 0x64, 0xf5,
	0xa0, 0xb2, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x69, 0x7d, 0x10, 0x0b, 0xa2, 0x52, 0x4a,
	0x09, 0x8b, 0x39, 0x89, 0x25, 0x25, 0x45, 0x99, 0x49, 0xa5, 0x25, 0xa9, 0x50, 0x35, 0x6a, 0x58,
	0xd4, 0x14, 0xa5, 0x16, 0xe7, 0x97, 0x16, 0x25, 0xa7, 0x96, 0x25, 0xe6, 0x94, 0x42, 0xd5, 0x29,
	0x5d, 0x65, 0xe4, 0x62, 0x76, 0x0e, 0x08, 0x15, 0xb2, 0xe5, 0x62, 0x2d, 0xcd, 0xcb, 0x2c, 0x29,
	0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd4, 0xc3, 0x74, 0x8d, 0x5e, 0x10, 0x54, 0x7f,
	0x18, 0x48, 0xbf, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x10, 0x5d, 0x42, 0x1d, 0x8c, 0x5c,
	0x5c, 0x70, 0x27, 0x14, 0x4b, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x1b, 0xc9, 0x62, 0x33, 0xc4, 0x11,
	0xa6, 0xca, 0xc9, 0x13, 0x64, 0xc0, 0xab, 0x7b, 0xf2, 0x22, 0x08, 0x8d, 0x3a, 0xf9, 0xb9, 0x99,
	0x25, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x9f, 0xee, 0xc9, 0x4b, 0x57, 0x26, 0xe6, 0xe6, 0x58, 0x29,
	0x61, 0x93, 0x55, 0x5a, 0x75, 0x5f, 0x9e, 0x0b, 0x6e, 0x52, 0x71, 0x10, 0x92, 0xdd, 0x56, 0x2c,
	0x2f, 0x16, 0xc8, 0x33, 0x3a, 0x59, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83,
	0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43,
	0x94, 0x62, 0x41, 0x76, 0xba, 0x5e, 0x62, 0x76, 0x89, 0x5e, 0x4a, 0x6a, 0x99, 0x7e, 0x7a, 0xbe,
	0x7e, 0x5e, 0x7e, 0x4a, 0xaa, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0x31, 0x2c, 0xa0, 0x92, 0xd8, 0xc0,
	0x61, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x96, 0xe9, 0x27, 0xb1, 0x01, 0x00, 0x00,
}

func (this *CPU) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CPU)
	if !ok {
		that2, ok := that.(CPU)
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
	if !this.Units.Equal(&that1.Units) {
		return false
	}
	if len(this.Attributes) != len(that1.Attributes) {
		return false
	}
	for i := range this.Attributes {
		if !this.Attributes[i].Equal(&that1.Attributes[i]) {
			return false
		}
	}
	return true
}
func (m *CPU) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CPU) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CPU) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCpu(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Units.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCpu(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintCpu(dAtA []byte, offset int, v uint64) int {
	offset -= sovCpu(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CPU) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Units.Size()
	n += 1 + l + sovCpu(uint64(l))
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovCpu(uint64(l))
		}
	}
	return n
}

func sovCpu(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCpu(x uint64) (n int) {
	return sovCpu(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CPU) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCpu
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
			return fmt.Errorf("proto: CPU: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CPU: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCpu
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
				return ErrInvalidLengthCpu
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCpu
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Units.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCpu
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
				return ErrInvalidLengthCpu
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCpu
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCpu(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCpu
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
func skipCpu(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCpu
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
					return 0, ErrIntOverflowCpu
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
					return 0, ErrIntOverflowCpu
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
				return 0, ErrInvalidLengthCpu
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCpu
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCpu
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCpu        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCpu          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCpu = fmt.Errorf("proto: unexpected end of group")
)
