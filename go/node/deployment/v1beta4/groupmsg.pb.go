// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta4/groupmsg.proto

package v1beta4

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	v1 "pkg.akt.dev/go/node/deployment/v1"
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

// MsgCloseGroup defines SDK message to close a single Group within a Deployment.
type MsgCloseGroup struct {
	// Id is the unique identifier of the Group.
	ID v1.GroupID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *MsgCloseGroup) Reset()         { *m = MsgCloseGroup{} }
func (m *MsgCloseGroup) String() string { return proto.CompactTextString(m) }
func (*MsgCloseGroup) ProtoMessage()    {}
func (*MsgCloseGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2e655b03e10552, []int{0}
}
func (m *MsgCloseGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseGroup.Merge(m, src)
}
func (m *MsgCloseGroup) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseGroup proto.InternalMessageInfo

func (m *MsgCloseGroup) GetID() v1.GroupID {
	if m != nil {
		return m.ID
	}
	return v1.GroupID{}
}

// MsgCloseGroupResponse defines the Msg/CloseGroup response type.
type MsgCloseGroupResponse struct {
}

func (m *MsgCloseGroupResponse) Reset()         { *m = MsgCloseGroupResponse{} }
func (m *MsgCloseGroupResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCloseGroupResponse) ProtoMessage()    {}
func (*MsgCloseGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2e655b03e10552, []int{1}
}
func (m *MsgCloseGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseGroupResponse.Merge(m, src)
}
func (m *MsgCloseGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseGroupResponse proto.InternalMessageInfo

// MsgPauseGroup defines SDK message to pause a single Group within a Deployment.
type MsgPauseGroup struct {
	// Id is the unique identifier of the Group.
	ID v1.GroupID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *MsgPauseGroup) Reset()         { *m = MsgPauseGroup{} }
func (m *MsgPauseGroup) String() string { return proto.CompactTextString(m) }
func (*MsgPauseGroup) ProtoMessage()    {}
func (*MsgPauseGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2e655b03e10552, []int{2}
}
func (m *MsgPauseGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPauseGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPauseGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPauseGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPauseGroup.Merge(m, src)
}
func (m *MsgPauseGroup) XXX_Size() int {
	return m.Size()
}
func (m *MsgPauseGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPauseGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPauseGroup proto.InternalMessageInfo

func (m *MsgPauseGroup) GetID() v1.GroupID {
	if m != nil {
		return m.ID
	}
	return v1.GroupID{}
}

// MsgPauseGroupResponse defines the Msg/PauseGroup response type.
type MsgPauseGroupResponse struct {
}

func (m *MsgPauseGroupResponse) Reset()         { *m = MsgPauseGroupResponse{} }
func (m *MsgPauseGroupResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPauseGroupResponse) ProtoMessage()    {}
func (*MsgPauseGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2e655b03e10552, []int{3}
}
func (m *MsgPauseGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPauseGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPauseGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPauseGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPauseGroupResponse.Merge(m, src)
}
func (m *MsgPauseGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPauseGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPauseGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPauseGroupResponse proto.InternalMessageInfo

// MsgStartGroup defines SDK message to start a single Group within a Deployment.
type MsgStartGroup struct {
	// Id is the unique identifier of the Group.
	ID v1.GroupID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *MsgStartGroup) Reset()         { *m = MsgStartGroup{} }
func (m *MsgStartGroup) String() string { return proto.CompactTextString(m) }
func (*MsgStartGroup) ProtoMessage()    {}
func (*MsgStartGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2e655b03e10552, []int{4}
}
func (m *MsgStartGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStartGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStartGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStartGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStartGroup.Merge(m, src)
}
func (m *MsgStartGroup) XXX_Size() int {
	return m.Size()
}
func (m *MsgStartGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStartGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStartGroup proto.InternalMessageInfo

func (m *MsgStartGroup) GetID() v1.GroupID {
	if m != nil {
		return m.ID
	}
	return v1.GroupID{}
}

// MsgStartGroupResponse defines the Msg/StartGroup response type.
type MsgStartGroupResponse struct {
}

func (m *MsgStartGroupResponse) Reset()         { *m = MsgStartGroupResponse{} }
func (m *MsgStartGroupResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStartGroupResponse) ProtoMessage()    {}
func (*MsgStartGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2e655b03e10552, []int{5}
}
func (m *MsgStartGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStartGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStartGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStartGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStartGroupResponse.Merge(m, src)
}
func (m *MsgStartGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStartGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStartGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStartGroupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCloseGroup)(nil), "akash.deployment.v1beta4.MsgCloseGroup")
	proto.RegisterType((*MsgCloseGroupResponse)(nil), "akash.deployment.v1beta4.MsgCloseGroupResponse")
	proto.RegisterType((*MsgPauseGroup)(nil), "akash.deployment.v1beta4.MsgPauseGroup")
	proto.RegisterType((*MsgPauseGroupResponse)(nil), "akash.deployment.v1beta4.MsgPauseGroupResponse")
	proto.RegisterType((*MsgStartGroup)(nil), "akash.deployment.v1beta4.MsgStartGroup")
	proto.RegisterType((*MsgStartGroupResponse)(nil), "akash.deployment.v1beta4.MsgStartGroupResponse")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta4/groupmsg.proto", fileDescriptor_ec2e655b03e10552)
}

var fileDescriptor_ec2e655b03e10552 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcc, 0x4e, 0x2c,
	0xce, 0xd0, 0x4f, 0x49, 0x2d, 0xc8, 0xc9, 0xaf, 0xcc, 0x4d, 0xcd, 0x2b, 0xd1, 0x2f, 0x33, 0x4c,
	0x4a, 0x2d, 0x49, 0x34, 0xd1, 0x4f, 0x2f, 0xca, 0x2f, 0x2d, 0xc8, 0x2d, 0x4e, 0xd7, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x00, 0x2b, 0xd4, 0x43, 0x28, 0xd4, 0x83, 0x2a, 0x94, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd2, 0x07, 0xb1, 0x20, 0xea, 0xa5, 0xc4, 0x93, 0xf3, 0x8b, 0x73,
	0xf3, 0x8b, 0xf5, 0x73, 0x8b, 0xd3, 0xf5, 0xcb, 0x0c, 0xf5, 0xe1, 0x06, 0x49, 0xc9, 0x63, 0xb1,
	0x11, 0x62, 0x19, 0x44, 0x81, 0x52, 0x12, 0x17, 0xaf, 0x6f, 0x71, 0xba, 0x73, 0x4e, 0x7e, 0x71,
	0xaa, 0x3b, 0x48, 0x58, 0xc8, 0x9b, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb,
	0x48, 0x46, 0x0f, 0x8b, 0x3b, 0xf4, 0xc0, 0xea, 0x3c, 0x5d, 0x9c, 0x64, 0x4f, 0xdc, 0x93, 0x67,
	0x78, 0x74, 0x4f, 0x9e, 0xc9, 0xd3, 0xe5, 0xd5, 0x3d, 0x79, 0xa6, 0xcc, 0x94, 0x4f, 0xf7, 0xe4,
	0x39, 0x2b, 0x13, 0x73, 0x73, 0xac, 0x94, 0x32, 0x53, 0x94, 0x82, 0x98, 0x32, 0x53, 0xac, 0x58,
	0x5e, 0x2c, 0x90, 0x67, 0x50, 0x12, 0xe7, 0x12, 0x45, 0xb1, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x15, 0x6a, 0x79, 0x40, 0x62, 0x29, 0xcd, 0x2d, 0x47, 0xd8, 0x81, 0x66, 0x79, 0x70,
	0x49, 0x62, 0x51, 0x09, 0x8d, 0x2d, 0x47, 0xd8, 0x01, 0xb3, 0xdc, 0xc9, 0xe1, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xd4, 0x0a, 0xb2, 0xd3, 0xf5, 0x12, 0xb3, 0x4b, 0xf4,
	0x52, 0x52, 0xcb, 0xf4, 0xd3, 0xf3, 0xf5, 0xf3, 0xf2, 0x53, 0x52, 0xb1, 0xa4, 0x98, 0x24, 0x36,
	0x70, 0xfc, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x25, 0x59, 0x4f, 0x17, 0x54, 0x02, 0x00,
	0x00,
}

func (m *MsgCloseGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroupmsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgCloseGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgPauseGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPauseGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPauseGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroupmsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgPauseGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPauseGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPauseGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgStartGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStartGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStartGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroupmsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgStartGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStartGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStartGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintGroupmsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroupmsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCloseGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovGroupmsg(uint64(l))
	return n
}

func (m *MsgCloseGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgPauseGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovGroupmsg(uint64(l))
	return n
}

func (m *MsgPauseGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgStartGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovGroupmsg(uint64(l))
	return n
}

func (m *MsgStartGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovGroupmsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroupmsg(x uint64) (n int) {
	return sovGroupmsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCloseGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupmsg
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
			return fmt.Errorf("proto: MsgCloseGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupmsg
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
				return ErrInvalidLengthGroupmsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroupmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroupmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroupmsg
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
func (m *MsgCloseGroupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupmsg
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
			return fmt.Errorf("proto: MsgCloseGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGroupmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroupmsg
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
func (m *MsgPauseGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupmsg
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
			return fmt.Errorf("proto: MsgPauseGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPauseGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupmsg
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
				return ErrInvalidLengthGroupmsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroupmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroupmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroupmsg
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
func (m *MsgPauseGroupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupmsg
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
			return fmt.Errorf("proto: MsgPauseGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPauseGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGroupmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroupmsg
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
func (m *MsgStartGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupmsg
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
			return fmt.Errorf("proto: MsgStartGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStartGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupmsg
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
				return ErrInvalidLengthGroupmsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroupmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroupmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroupmsg
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
func (m *MsgStartGroupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupmsg
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
			return fmt.Errorf("proto: MsgStartGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStartGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGroupmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroupmsg
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
func skipGroupmsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroupmsg
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
					return 0, ErrIntOverflowGroupmsg
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
					return 0, ErrIntOverflowGroupmsg
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
				return 0, ErrInvalidLengthGroupmsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroupmsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroupmsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroupmsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroupmsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroupmsg = fmt.Errorf("proto: unexpected end of group")
)
