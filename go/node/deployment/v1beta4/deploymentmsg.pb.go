// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta4/deploymentmsg.proto

package v1beta4

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

// MsgCreateDeployment defines an SDK message for creating deployment.
type MsgCreateDeployment struct {
	// ID is the unique identifier of the deployment.
	ID v1.DeploymentID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	// GroupSpec is a list of group specifications for the deployment.
	// This field is required and must be a list of GroupSpec.
	Groups GroupSpecs `protobuf:"bytes,2,rep,name=groups,proto3,castrepeated=GroupSpecs" json:"groups" yaml:"groups"`
	// Hash of the deployment.
	Hash []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash" yaml:"hash"`
	// Deposit specifies the amount of coins to include in the deployment's first deposit.
	Deposit types.Coin `protobuf:"bytes,4,opt,name=deposit,proto3" json:"deposit" yaml:"deposit"`
	// Depositor is the account address of the user who will deposit funds to the deployment.
	// This value can be different than the owner of the deployment if there is authorized spend grants applied.
	// It is a string representing a valid account address.
	//
	// Example:
	//   "akash1..."
	Depositor string `protobuf:"bytes,5,opt,name=depositor,proto3" json:"depositor" yaml:"depositor"`
}

func (m *MsgCreateDeployment) Reset()         { *m = MsgCreateDeployment{} }
func (m *MsgCreateDeployment) String() string { return proto.CompactTextString(m) }
func (*MsgCreateDeployment) ProtoMessage()    {}
func (*MsgCreateDeployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b10e8e78e405ddf, []int{0}
}
func (m *MsgCreateDeployment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateDeployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateDeployment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateDeployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateDeployment.Merge(m, src)
}
func (m *MsgCreateDeployment) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateDeployment) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateDeployment.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateDeployment proto.InternalMessageInfo

func (m *MsgCreateDeployment) GetID() v1.DeploymentID {
	if m != nil {
		return m.ID
	}
	return v1.DeploymentID{}
}

func (m *MsgCreateDeployment) GetGroups() GroupSpecs {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *MsgCreateDeployment) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *MsgCreateDeployment) GetDeposit() types.Coin {
	if m != nil {
		return m.Deposit
	}
	return types.Coin{}
}

func (m *MsgCreateDeployment) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

// MsgCreateDeploymentResponse defines the Msg/CreateDeployment response type.
type MsgCreateDeploymentResponse struct {
}

func (m *MsgCreateDeploymentResponse) Reset()         { *m = MsgCreateDeploymentResponse{} }
func (m *MsgCreateDeploymentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateDeploymentResponse) ProtoMessage()    {}
func (*MsgCreateDeploymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b10e8e78e405ddf, []int{1}
}
func (m *MsgCreateDeploymentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateDeploymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateDeploymentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateDeploymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateDeploymentResponse.Merge(m, src)
}
func (m *MsgCreateDeploymentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateDeploymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateDeploymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateDeploymentResponse proto.InternalMessageInfo

// MsgUpdateDeployment defines an SDK message for updating deployment.
type MsgUpdateDeployment struct {
	// ID is the unique identifier of the deployment.
	ID v1.DeploymentID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	// Hash of the deployment.
	Hash []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash" yaml:"hash"`
}

func (m *MsgUpdateDeployment) Reset()         { *m = MsgUpdateDeployment{} }
func (m *MsgUpdateDeployment) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDeployment) ProtoMessage()    {}
func (*MsgUpdateDeployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b10e8e78e405ddf, []int{2}
}
func (m *MsgUpdateDeployment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDeployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDeployment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDeployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDeployment.Merge(m, src)
}
func (m *MsgUpdateDeployment) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDeployment) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDeployment.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDeployment proto.InternalMessageInfo

func (m *MsgUpdateDeployment) GetID() v1.DeploymentID {
	if m != nil {
		return m.ID
	}
	return v1.DeploymentID{}
}

func (m *MsgUpdateDeployment) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// MsgUpdateDeploymentResponse defines the Msg/UpdateDeployment response type.
type MsgUpdateDeploymentResponse struct {
}

func (m *MsgUpdateDeploymentResponse) Reset()         { *m = MsgUpdateDeploymentResponse{} }
func (m *MsgUpdateDeploymentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDeploymentResponse) ProtoMessage()    {}
func (*MsgUpdateDeploymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b10e8e78e405ddf, []int{3}
}
func (m *MsgUpdateDeploymentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDeploymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDeploymentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDeploymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDeploymentResponse.Merge(m, src)
}
func (m *MsgUpdateDeploymentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDeploymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDeploymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDeploymentResponse proto.InternalMessageInfo

// MsgCloseDeployment defines an SDK message for closing deployment
type MsgCloseDeployment struct {
	// ID is the unique identifier of the deployment.
	ID v1.DeploymentID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *MsgCloseDeployment) Reset()         { *m = MsgCloseDeployment{} }
func (m *MsgCloseDeployment) String() string { return proto.CompactTextString(m) }
func (*MsgCloseDeployment) ProtoMessage()    {}
func (*MsgCloseDeployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b10e8e78e405ddf, []int{4}
}
func (m *MsgCloseDeployment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseDeployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseDeployment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseDeployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseDeployment.Merge(m, src)
}
func (m *MsgCloseDeployment) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseDeployment) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseDeployment.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseDeployment proto.InternalMessageInfo

func (m *MsgCloseDeployment) GetID() v1.DeploymentID {
	if m != nil {
		return m.ID
	}
	return v1.DeploymentID{}
}

// MsgCloseDeploymentResponse defines the Msg/CloseDeployment response type.
type MsgCloseDeploymentResponse struct {
}

func (m *MsgCloseDeploymentResponse) Reset()         { *m = MsgCloseDeploymentResponse{} }
func (m *MsgCloseDeploymentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCloseDeploymentResponse) ProtoMessage()    {}
func (*MsgCloseDeploymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b10e8e78e405ddf, []int{5}
}
func (m *MsgCloseDeploymentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseDeploymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseDeploymentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseDeploymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseDeploymentResponse.Merge(m, src)
}
func (m *MsgCloseDeploymentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseDeploymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseDeploymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseDeploymentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateDeployment)(nil), "akash.deployment.v1beta4.MsgCreateDeployment")
	proto.RegisterType((*MsgCreateDeploymentResponse)(nil), "akash.deployment.v1beta4.MsgCreateDeploymentResponse")
	proto.RegisterType((*MsgUpdateDeployment)(nil), "akash.deployment.v1beta4.MsgUpdateDeployment")
	proto.RegisterType((*MsgUpdateDeploymentResponse)(nil), "akash.deployment.v1beta4.MsgUpdateDeploymentResponse")
	proto.RegisterType((*MsgCloseDeployment)(nil), "akash.deployment.v1beta4.MsgCloseDeployment")
	proto.RegisterType((*MsgCloseDeploymentResponse)(nil), "akash.deployment.v1beta4.MsgCloseDeploymentResponse")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta4/deploymentmsg.proto", fileDescriptor_9b10e8e78e405ddf)
}

var fileDescriptor_9b10e8e78e405ddf = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xb4, 0x0c, 0xd5, 0x05, 0x84, 0xc2, 0xc4, 0xb2, 0xc2, 0xe2, 0x12, 0x10, 0x8a,
	0x04, 0x73, 0xd4, 0xc1, 0x85, 0x9d, 0x20, 0xab, 0x84, 0x26, 0x84, 0x84, 0x32, 0x21, 0x24, 0x2e,
	0x28, 0xad, 0x2d, 0x37, 0xb4, 0x89, 0xad, 0x38, 0x54, 0xda, 0xb7, 0xe0, 0xcc, 0x89, 0x33, 0xe2,
	0xc8, 0x87, 0xd8, 0x71, 0xe2, 0xc4, 0xc9, 0xa0, 0xf6, 0x82, 0x7a, 0xcc, 0x27, 0x40, 0x89, 0xbd,
	0x76, 0x83, 0xee, 0xc0, 0x61, 0x37, 0xff, 0xfd, 0x7e, 0xef, 0xbd, 0xff, 0xf3, 0x93, 0xc1, 0xc3,
	0x68, 0x14, 0x89, 0xa1, 0x8f, 0x09, 0x1f, 0xb3, 0xc3, 0x84, 0xa4, 0xb9, 0x3f, 0xe9, 0xf6, 0x49,
	0x1e, 0x3d, 0x3e, 0x75, 0x95, 0x08, 0x8a, 0x78, 0xc6, 0x72, 0x66, 0xd9, 0x15, 0x8d, 0x96, 0x21,
	0xa4, 0xe9, 0xf6, 0x3a, 0x65, 0x94, 0x55, 0x90, 0x5f, 0x9e, 0x14, 0xdf, 0xde, 0x1c, 0x30, 0x91,
	0x30, 0xf1, 0x4e, 0x05, 0x94, 0xd0, 0xa1, 0x0d, 0xa5, 0xfc, 0x44, 0x50, 0x7f, 0xd2, 0xf5, 0x17,
	0x3d, 0xda, 0x8e, 0x0e, 0xf4, 0x23, 0x41, 0xb4, 0x99, 0xae, 0x3f, 0x60, 0x71, 0xaa, 0xe3, 0xf7,
	0x56, 0x38, 0x3e, 0xa5, 0x34, 0xe5, 0x9d, 0x3b, 0x17, 0xcd, 0xd8, 0x07, 0x2e, 0x38, 0x19, 0x28,
	0xd2, 0xfd, 0x5a, 0x07, 0x37, 0x5e, 0x0a, 0xba, 0x97, 0x91, 0x28, 0x27, 0xbd, 0x05, 0x6f, 0xbd,
	0x02, 0x66, 0x8c, 0x6d, 0xa3, 0x63, 0x78, 0xad, 0x9d, 0x3b, 0x68, 0xc5, 0xe0, 0x68, 0x09, 0xef,
	0xf7, 0x82, 0xad, 0x23, 0x09, 0x6b, 0x53, 0x09, 0xcd, 0xfd, 0xde, 0x5c, 0x42, 0x33, 0xc6, 0x85,
	0x84, 0xcd, 0xc3, 0x28, 0x19, 0xef, 0xba, 0x31, 0x76, 0x43, 0x33, 0xc6, 0xd6, 0x7b, 0xb0, 0xa6,
	0x9a, 0xdb, 0x66, 0xa7, 0xee, 0xb5, 0x76, 0xee, 0xa2, 0xf3, 0x9e, 0x13, 0x3d, 0x2f, 0xb9, 0x03,
	0x4e, 0x06, 0xc1, 0x76, 0x59, 0x77, 0x2e, 0xa1, 0x4e, 0x2d, 0x24, 0xbc, 0xaa, 0xaa, 0x2a, 0xed,
	0x7e, 0xf9, 0x09, 0xc1, 0x82, 0x16, 0xa1, 0xc6, 0xac, 0x07, 0xa0, 0x31, 0x8c, 0xc4, 0xd0, 0xae,
	0x77, 0x0c, 0xef, 0x4a, 0xb0, 0x31, 0x97, 0xb0, 0xd2, 0x85, 0x84, 0x2d, 0x95, 0x5e, 0x2a, 0x37,
	0xac, 0x2e, 0xad, 0x17, 0xe0, 0x32, 0x26, 0x9c, 0x89, 0x38, 0xb7, 0x1b, 0xd5, 0xbc, 0x9b, 0x48,
	0xef, 0xaa, 0x5c, 0x82, 0x36, 0xd5, 0x45, 0x7b, 0x2c, 0x4e, 0x83, 0x9b, 0xa5, 0x9f, 0x42, 0xc2,
	0x6b, 0xaa, 0x8c, 0xce, 0x73, 0xc3, 0x93, 0x0a, 0xd6, 0x1b, 0xd0, 0xd4, 0x47, 0x96, 0xd9, 0x97,
	0x3a, 0x86, 0xd7, 0x0c, 0x9e, 0xcc, 0x25, 0x5c, 0x5e, 0x16, 0x12, 0x5e, 0x3f, 0x93, 0xcc, 0x32,
	0xf7, 0xfb, 0xb7, 0xed, 0x75, 0xdd, 0xf2, 0x19, 0xc6, 0x19, 0x11, 0xe2, 0x20, 0xcf, 0xe2, 0x94,
	0x86, 0xcb, 0xb4, 0xdd, 0xc6, 0xef, 0xcf, 0xb0, 0xe6, 0x6e, 0x81, 0x5b, 0x2b, 0xb6, 0x15, 0x12,
	0xc1, 0x59, 0x2a, 0x88, 0xfb, 0xc9, 0xa8, 0xb6, 0xf9, 0x9a, 0xe3, 0x8b, 0xde, 0xe6, 0xff, 0xbc,
	0xf0, 0x19, 0xef, 0x7f, 0x7b, 0x5b, 0x78, 0x1f, 0x03, 0xab, 0x1c, 0x6d, 0xcc, 0xc4, 0x85, 0x3a,
	0xd7, 0x66, 0x6e, 0x83, 0xf6, 0xbf, 0xdd, 0x4e, 0xbc, 0x04, 0x4f, 0x8f, 0xa6, 0x8e, 0x71, 0x3c,
	0x75, 0x8c, 0x5f, 0x53, 0xc7, 0xf8, 0x38, 0x73, 0x6a, 0xc7, 0x33, 0xa7, 0xf6, 0x63, 0xe6, 0xd4,
	0xde, 0xde, 0xe7, 0x23, 0x8a, 0xa2, 0x51, 0x8e, 0x30, 0x99, 0xf8, 0x94, 0xf9, 0x29, 0xc3, 0x64,
	0xc5, 0x3f, 0xeb, 0xaf, 0x55, 0xdf, 0xeb, 0xd1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xd6,
	0x1f, 0x9c, 0x62, 0x04, 0x00, 0x00,
}

func (m *MsgCreateDeployment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateDeployment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateDeployment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintDeploymentmsg(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDeploymentmsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintDeploymentmsg(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Groups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeploymentmsg(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDeploymentmsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgCreateDeploymentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateDeploymentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateDeploymentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateDeployment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDeployment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDeployment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintDeploymentmsg(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDeploymentmsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgUpdateDeploymentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDeploymentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDeploymentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCloseDeployment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseDeployment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseDeployment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintDeploymentmsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgCloseDeploymentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseDeploymentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseDeploymentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintDeploymentmsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeploymentmsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateDeployment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovDeploymentmsg(uint64(l))
	if len(m.Groups) > 0 {
		for _, e := range m.Groups {
			l = e.Size()
			n += 1 + l + sovDeploymentmsg(uint64(l))
		}
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovDeploymentmsg(uint64(l))
	}
	l = m.Deposit.Size()
	n += 1 + l + sovDeploymentmsg(uint64(l))
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovDeploymentmsg(uint64(l))
	}
	return n
}

func (m *MsgCreateDeploymentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateDeployment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovDeploymentmsg(uint64(l))
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovDeploymentmsg(uint64(l))
	}
	return n
}

func (m *MsgUpdateDeploymentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCloseDeployment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovDeploymentmsg(uint64(l))
	return n
}

func (m *MsgCloseDeploymentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovDeploymentmsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeploymentmsg(x uint64) (n int) {
	return sovDeploymentmsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateDeployment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploymentmsg
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
			return fmt.Errorf("proto: MsgCreateDeployment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateDeployment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentmsg
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
				return ErrInvalidLengthDeploymentmsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeploymentmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentmsg
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
				return ErrInvalidLengthDeploymentmsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeploymentmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groups = append(m.Groups, GroupSpec{})
			if err := m.Groups[len(m.Groups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentmsg
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
				return ErrInvalidLengthDeploymentmsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDeploymentmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentmsg
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
				return ErrInvalidLengthDeploymentmsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeploymentmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentmsg
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
				return ErrInvalidLengthDeploymentmsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeploymentmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeploymentmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeploymentmsg
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
func (m *MsgCreateDeploymentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploymentmsg
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
			return fmt.Errorf("proto: MsgCreateDeploymentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateDeploymentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDeploymentmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeploymentmsg
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
func (m *MsgUpdateDeployment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploymentmsg
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
			return fmt.Errorf("proto: MsgUpdateDeployment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDeployment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentmsg
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
				return ErrInvalidLengthDeploymentmsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeploymentmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentmsg
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
				return ErrInvalidLengthDeploymentmsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDeploymentmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeploymentmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeploymentmsg
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
func (m *MsgUpdateDeploymentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploymentmsg
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
			return fmt.Errorf("proto: MsgUpdateDeploymentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDeploymentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDeploymentmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeploymentmsg
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
func (m *MsgCloseDeployment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploymentmsg
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
			return fmt.Errorf("proto: MsgCloseDeployment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseDeployment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentmsg
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
				return ErrInvalidLengthDeploymentmsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeploymentmsg
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
			skippy, err := skipDeploymentmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeploymentmsg
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
func (m *MsgCloseDeploymentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploymentmsg
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
			return fmt.Errorf("proto: MsgCloseDeploymentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseDeploymentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDeploymentmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeploymentmsg
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
func skipDeploymentmsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeploymentmsg
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
					return 0, ErrIntOverflowDeploymentmsg
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
					return 0, ErrIntOverflowDeploymentmsg
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
				return 0, ErrInvalidLengthDeploymentmsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeploymentmsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeploymentmsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeploymentmsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeploymentmsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeploymentmsg = fmt.Errorf("proto: unexpected end of group")
)
