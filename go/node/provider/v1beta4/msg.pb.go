// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/provider/v1beta4/msg.proto

package v1beta4

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	pkg_akt_dev_go_node_types_attributes_v1 "pkg.akt.dev/go/node/types/attributes/v1"
	v1 "pkg.akt.dev/go/node/types/attributes/v1"
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

// MsgCreateProvider defines an SDK message for creating a provider.
type MsgCreateProvider struct {
	// Owner is the bech32 address of the account of the provider.
	// It is a string representing a valid account address.
	//
	// Example:
	//   "akash1..."
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	// HostURI is the Uniform Resource Identifier for provider connection.
	// This URI is used to directly connect to the provider to perform tasks such as sending the manifest.
	HostURI string `protobuf:"bytes,2,opt,name=host_uri,json=hostUri,proto3" json:"host_uri" yaml:"host_uri"`
	// Attributes is a list of arbitrary attribute key-value pairs.
	Attributes pkg_akt_dev_go_node_types_attributes_v1.Attributes `protobuf:"bytes,3,rep,name=attributes,proto3,castrepeated=pkg.akt.dev/go/node/types/attributes/v1.Attributes" json:"attributes" yaml:"attributes"`
	// Info contains additional provider information.
	Info Info `protobuf:"bytes,4,opt,name=info,proto3" json:"info" yaml:"info"`
}

func (m *MsgCreateProvider) Reset()         { *m = MsgCreateProvider{} }
func (m *MsgCreateProvider) String() string { return proto.CompactTextString(m) }
func (*MsgCreateProvider) ProtoMessage()    {}
func (*MsgCreateProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c874c91147ead42, []int{0}
}
func (m *MsgCreateProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateProvider.Merge(m, src)
}
func (m *MsgCreateProvider) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateProvider.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateProvider proto.InternalMessageInfo

func (m *MsgCreateProvider) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgCreateProvider) GetHostURI() string {
	if m != nil {
		return m.HostURI
	}
	return ""
}

func (m *MsgCreateProvider) GetAttributes() pkg_akt_dev_go_node_types_attributes_v1.Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *MsgCreateProvider) GetInfo() Info {
	if m != nil {
		return m.Info
	}
	return Info{}
}

// MsgCreateProviderResponse defines the Msg/CreateProvider response type.
type MsgCreateProviderResponse struct {
}

func (m *MsgCreateProviderResponse) Reset()         { *m = MsgCreateProviderResponse{} }
func (m *MsgCreateProviderResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateProviderResponse) ProtoMessage()    {}
func (*MsgCreateProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c874c91147ead42, []int{1}
}
func (m *MsgCreateProviderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateProviderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateProviderResponse.Merge(m, src)
}
func (m *MsgCreateProviderResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateProviderResponse proto.InternalMessageInfo

// MsgUpdateProvider defines an SDK message for updating a provider
type MsgUpdateProvider struct {
	Owner      string                                             `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	HostURI    string                                             `protobuf:"bytes,2,opt,name=host_uri,json=hostUri,proto3" json:"host_uri" yaml:"host_uri"`
	Attributes pkg_akt_dev_go_node_types_attributes_v1.Attributes `protobuf:"bytes,3,rep,name=attributes,proto3,castrepeated=pkg.akt.dev/go/node/types/attributes/v1.Attributes" json:"attributes" yaml:"attributes"`
	Info       Info                                               `protobuf:"bytes,4,opt,name=info,proto3" json:"info" yaml:"info"`
}

func (m *MsgUpdateProvider) Reset()         { *m = MsgUpdateProvider{} }
func (m *MsgUpdateProvider) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateProvider) ProtoMessage()    {}
func (*MsgUpdateProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c874c91147ead42, []int{2}
}
func (m *MsgUpdateProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateProvider.Merge(m, src)
}
func (m *MsgUpdateProvider) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateProvider.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateProvider proto.InternalMessageInfo

func (m *MsgUpdateProvider) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgUpdateProvider) GetHostURI() string {
	if m != nil {
		return m.HostURI
	}
	return ""
}

func (m *MsgUpdateProvider) GetAttributes() pkg_akt_dev_go_node_types_attributes_v1.Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *MsgUpdateProvider) GetInfo() Info {
	if m != nil {
		return m.Info
	}
	return Info{}
}

// MsgUpdateProviderResponse defines the Msg/UpdateProvider response type.
type MsgUpdateProviderResponse struct {
}

func (m *MsgUpdateProviderResponse) Reset()         { *m = MsgUpdateProviderResponse{} }
func (m *MsgUpdateProviderResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateProviderResponse) ProtoMessage()    {}
func (*MsgUpdateProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c874c91147ead42, []int{3}
}
func (m *MsgUpdateProviderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateProviderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateProviderResponse.Merge(m, src)
}
func (m *MsgUpdateProviderResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateProviderResponse proto.InternalMessageInfo

// MsgDeleteProvider defines an SDK message for deleting a provider
type MsgDeleteProvider struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
}

func (m *MsgDeleteProvider) Reset()         { *m = MsgDeleteProvider{} }
func (m *MsgDeleteProvider) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteProvider) ProtoMessage()    {}
func (*MsgDeleteProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c874c91147ead42, []int{4}
}
func (m *MsgDeleteProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteProvider.Merge(m, src)
}
func (m *MsgDeleteProvider) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteProvider.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteProvider proto.InternalMessageInfo

func (m *MsgDeleteProvider) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// MsgDeleteProviderResponse defines the Msg/DeleteProvider response type.
type MsgDeleteProviderResponse struct {
}

func (m *MsgDeleteProviderResponse) Reset()         { *m = MsgDeleteProviderResponse{} }
func (m *MsgDeleteProviderResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteProviderResponse) ProtoMessage()    {}
func (*MsgDeleteProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c874c91147ead42, []int{5}
}
func (m *MsgDeleteProviderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteProviderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteProviderResponse.Merge(m, src)
}
func (m *MsgDeleteProviderResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteProviderResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateProvider)(nil), "akash.provider.v1beta4.MsgCreateProvider")
	proto.RegisterType((*MsgCreateProviderResponse)(nil), "akash.provider.v1beta4.MsgCreateProviderResponse")
	proto.RegisterType((*MsgUpdateProvider)(nil), "akash.provider.v1beta4.MsgUpdateProvider")
	proto.RegisterType((*MsgUpdateProviderResponse)(nil), "akash.provider.v1beta4.MsgUpdateProviderResponse")
	proto.RegisterType((*MsgDeleteProvider)(nil), "akash.provider.v1beta4.MsgDeleteProvider")
	proto.RegisterType((*MsgDeleteProviderResponse)(nil), "akash.provider.v1beta4.MsgDeleteProviderResponse")
}

func init() { proto.RegisterFile("akash/provider/v1beta4/msg.proto", fileDescriptor_5c874c91147ead42) }

var fileDescriptor_5c874c91147ead42 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0x4f, 0x6b, 0x13, 0x41,
	0x1c, 0xdd, 0x35, 0xd5, 0xea, 0x44, 0x94, 0x2e, 0x45, 0xd3, 0x54, 0x76, 0xc2, 0xaa, 0x10, 0x0a,
	0xce, 0x90, 0xe8, 0xa9, 0x07, 0xa1, 0x51, 0xd1, 0x1e, 0x0a, 0xb2, 0x92, 0x8b, 0x20, 0x65, 0xe2,
	0x4e, 0xb7, 0x4b, 0x92, 0x9d, 0x65, 0x66, 0xba, 0xd2, 0xab, 0x9f, 0xc0, 0x8f, 0x20, 0x5e, 0x04,
	0x4f, 0x1e, 0xfc, 0x10, 0x3d, 0x16, 0x4f, 0x9e, 0x46, 0x49, 0x0e, 0x4a, 0x8e, 0xf9, 0x04, 0xb2,
	0x33, 0xfb, 0xa7, 0xd1, 0xf6, 0xe8, 0x41, 0xe8, 0x6d, 0x7e, 0xbf, 0xdf, 0xfb, 0xbd, 0x79, 0xfb,
	0xde, 0x32, 0xa0, 0x45, 0x86, 0x44, 0xec, 0xe3, 0x84, 0xb3, 0x34, 0x0a, 0x28, 0xc7, 0x69, 0x67,
	0x40, 0x25, 0x79, 0x80, 0xc7, 0x22, 0x44, 0x09, 0x67, 0x92, 0x39, 0x37, 0x34, 0x02, 0x15, 0x08,
	0x94, 0x23, 0x9a, 0xab, 0x21, 0x0b, 0x99, 0x86, 0xe0, 0xec, 0x64, 0xd0, 0xcd, 0x9b, 0xaf, 0x99,
	0x18, 0x33, 0x91, 0xed, 0xe3, 0xb4, 0x53, 0xd1, 0x34, 0xd7, 0xcc, 0x60, 0xd7, 0x6c, 0x98, 0x22,
	0x1f, 0xb5, 0x8d, 0x86, 0x01, 0x11, 0x14, 0x13, 0x29, 0x79, 0x34, 0x38, 0x90, 0x54, 0x64, 0xeb,
	0x65, 0x95, 0x23, 0xef, 0x9e, 0xa1, 0xb6, 0x14, 0xa7, 0x61, 0xde, 0xc7, 0x1a, 0x58, 0xd9, 0x11,
	0xe1, 0x23, 0x4e, 0x89, 0xa4, 0xcf, 0xf3, 0x99, 0xf3, 0x14, 0x5c, 0x64, 0x6f, 0x62, 0xca, 0x1b,
	0x76, 0xcb, 0x6e, 0x5f, 0xe9, 0x75, 0x66, 0x0a, 0x9a, 0xc6, 0x5c, 0xc1, 0xab, 0x87, 0x64, 0x3c,
	0xda, 0xf4, 0x74, 0xe9, 0x7d, 0xfd, 0x72, 0x6f, 0x35, 0x17, 0xb8, 0x15, 0x04, 0x9c, 0x0a, 0xf1,
	0x42, 0xf2, 0x28, 0x0e, 0x7d, 0x03, 0x77, 0x9e, 0x80, 0xcb, 0xfb, 0x4c, 0xc8, 0xdd, 0x03, 0x1e,
	0x35, 0x2e, 0x68, 0xae, 0x8d, 0x89, 0x82, 0xcb, 0xcf, 0x98, 0x90, 0x7d, 0x7f, 0x7b, 0xa6, 0x60,
	0x39, 0x9e, 0x2b, 0x78, 0xdd, 0x30, 0x17, 0x1d, 0xcf, 0x5f, 0xce, 0x8e, 0x7d, 0x1e, 0x39, 0x1f,
	0x6c, 0x00, 0xaa, 0xcf, 0x6d, 0xd4, 0x5a, 0xb5, 0x76, 0xbd, 0x7b, 0x1b, 0x19, 0xbb, 0x33, 0x33,
	0x50, 0x35, 0x45, 0x69, 0x07, 0x6d, 0x15, 0x55, 0xef, 0xd5, 0x91, 0x82, 0xd6, 0x4c, 0xc1, 0x13,
	0xeb, 0x73, 0x05, 0x57, 0xcc, 0x4d, 0x55, 0xcf, 0xfb, 0xf4, 0x1d, 0x76, 0x93, 0x61, 0x88, 0xc8,
	0x50, 0xa2, 0x80, 0xa6, 0x38, 0x64, 0x38, 0x66, 0x01, 0xc5, 0xf2, 0x30, 0xa1, 0x62, 0xd1, 0xea,
	0x8a, 0x5d, 0xf8, 0x27, 0x68, 0x9d, 0x1d, 0xb0, 0x14, 0xc5, 0x7b, 0xac, 0xb1, 0xd4, 0xb2, 0xdb,
	0xf5, 0xee, 0x2d, 0x74, 0xfa, 0xcf, 0x80, 0xb6, 0xe3, 0x3d, 0xd6, 0x5b, 0xcf, 0x65, 0xe9, 0x8d,
	0xb9, 0x82, 0x75, 0x23, 0x28, 0xab, 0x3c, 0x5f, 0x37, 0x37, 0xaf, 0xfd, 0x7a, 0x0f, 0xad, 0xb7,
	0x3f, 0x3f, 0x6f, 0x18, 0x2b, 0xbd, 0x75, 0xb0, 0xf6, 0x57, 0x50, 0x3e, 0x15, 0x09, 0x8b, 0x05,
	0x2d, 0x62, 0xec, 0x27, 0xc1, 0x79, 0x8c, 0xff, 0x43, 0x8c, 0x8b, 0x41, 0x95, 0x31, 0x8e, 0x74,
	0x8a, 0x8f, 0xe9, 0x88, 0xfe, 0x83, 0x14, 0xcf, 0x90, 0xb2, 0x78, 0x5b, 0x21, 0xa5, 0xf7, 0xf0,
	0x68, 0xe2, 0xda, 0xc7, 0x13, 0xd7, 0xfe, 0x31, 0x71, 0xed, 0x77, 0x53, 0xd7, 0x3a, 0x9e, 0xba,
	0xd6, 0xb7, 0xa9, 0x6b, 0xbd, 0xbc, 0x73, 0x9a, 0xc1, 0x7f, 0xbe, 0x33, 0x83, 0x4b, 0xfa, 0x7d,
	0xb9, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xb1, 0x80, 0x88, 0x36, 0x05, 0x00, 0x00,
}

func (m *MsgCreateProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsg(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.HostURI) > 0 {
		i -= len(m.HostURI)
		copy(dAtA[i:], m.HostURI)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.HostURI)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateProviderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateProviderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateProviderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsg(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.HostURI) > 0 {
		i -= len(m.HostURI)
		copy(dAtA[i:], m.HostURI)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.HostURI)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateProviderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateProviderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateProviderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteProviderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteProviderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteProviderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.HostURI)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	l = m.Info.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgCreateProviderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.HostURI)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	l = m.Info.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgUpdateProviderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgDeleteProviderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgCreateProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, v1.Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgCreateProviderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgCreateProviderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateProviderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgUpdateProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgUpdateProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, v1.Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgUpdateProviderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgUpdateProviderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateProviderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgDeleteProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgDeleteProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgDeleteProviderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgDeleteProviderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteProviderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)
