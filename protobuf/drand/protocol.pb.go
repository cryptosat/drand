// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drand/protocol.proto

package drand

import (
	fmt "fmt"
	dkg "github.com/drand/drand/protobuf/crypto/dkg"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IdentityRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentityRequest) Reset()         { *m = IdentityRequest{} }
func (m *IdentityRequest) String() string { return proto.CompactTextString(m) }
func (*IdentityRequest) ProtoMessage()    {}
func (*IdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{0}
}

func (m *IdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityRequest.Unmarshal(m, b)
}
func (m *IdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityRequest.Marshal(b, m, deterministic)
}
func (m *IdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityRequest.Merge(m, src)
}
func (m *IdentityRequest) XXX_Size() int {
	return xxx_messageInfo_IdentityRequest.Size(m)
}
func (m *IdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityRequest proto.InternalMessageInfo

// SignalDKGPacket is the packet nodes send to a coordinator that collects all
// keys and setups the group and sends them back to the nodes such that they can
// start the DKG automatically.
type SignalDKGPacket struct {
	Node        *Identity `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	SecretProof []byte    `protobuf:"bytes,2,opt,name=secret_proof,json=secretProof,proto3" json:"secret_proof,omitempty"`
	// In resharing cases, previous_group_hash is the hash of the previous group.
	// It is to make sure the nodes build on top of the correct previous group.
	PreviousGroupHash    []byte   `protobuf:"bytes,3,opt,name=previous_group_hash,json=previousGroupHash,proto3" json:"previous_group_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalDKGPacket) Reset()         { *m = SignalDKGPacket{} }
func (m *SignalDKGPacket) String() string { return proto.CompactTextString(m) }
func (*SignalDKGPacket) ProtoMessage()    {}
func (*SignalDKGPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{1}
}

func (m *SignalDKGPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalDKGPacket.Unmarshal(m, b)
}
func (m *SignalDKGPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalDKGPacket.Marshal(b, m, deterministic)
}
func (m *SignalDKGPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalDKGPacket.Merge(m, src)
}
func (m *SignalDKGPacket) XXX_Size() int {
	return xxx_messageInfo_SignalDKGPacket.Size(m)
}
func (m *SignalDKGPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalDKGPacket.DiscardUnknown(m)
}

var xxx_messageInfo_SignalDKGPacket proto.InternalMessageInfo

func (m *SignalDKGPacket) GetNode() *Identity {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *SignalDKGPacket) GetSecretProof() []byte {
	if m != nil {
		return m.SecretProof
	}
	return nil
}

func (m *SignalDKGPacket) GetPreviousGroupHash() []byte {
	if m != nil {
		return m.PreviousGroupHash
	}
	return nil
}

// PushDKGInfor is the packet the coordinator sends that contains the group over
// which to run the DKG on, the secret proof (to prove it's he's part of the
// expected group, and it's not a random packet) and as well the time at which
// every node should start the DKG.
type DKGInfoPacket struct {
	NewGroup    *GroupPacket `protobuf:"bytes,1,opt,name=new_group,json=newGroup,proto3" json:"new_group,omitempty"`
	SecretProof []byte       `protobuf:"bytes,2,opt,name=secret_proof,json=secretProof,proto3" json:"secret_proof,omitempty"`
	// timeout in seconds
	DkgTimeout uint32 `protobuf:"varint,3,opt,name=dkg_timeout,json=dkgTimeout,proto3" json:"dkg_timeout,omitempty"`
	// signature from the coordinator to prove he is the one sending that group
	// file.
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DKGInfoPacket) Reset()         { *m = DKGInfoPacket{} }
func (m *DKGInfoPacket) String() string { return proto.CompactTextString(m) }
func (*DKGInfoPacket) ProtoMessage()    {}
func (*DKGInfoPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{2}
}

func (m *DKGInfoPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DKGInfoPacket.Unmarshal(m, b)
}
func (m *DKGInfoPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DKGInfoPacket.Marshal(b, m, deterministic)
}
func (m *DKGInfoPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DKGInfoPacket.Merge(m, src)
}
func (m *DKGInfoPacket) XXX_Size() int {
	return xxx_messageInfo_DKGInfoPacket.Size(m)
}
func (m *DKGInfoPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_DKGInfoPacket.DiscardUnknown(m)
}

var xxx_messageInfo_DKGInfoPacket proto.InternalMessageInfo

func (m *DKGInfoPacket) GetNewGroup() *GroupPacket {
	if m != nil {
		return m.NewGroup
	}
	return nil
}

func (m *DKGInfoPacket) GetSecretProof() []byte {
	if m != nil {
		return m.SecretProof
	}
	return nil
}

func (m *DKGInfoPacket) GetDkgTimeout() uint32 {
	if m != nil {
		return m.DkgTimeout
	}
	return 0
}

func (m *DKGInfoPacket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type PartialBeaconPacket struct {
	// Round is the round for which the beacon will be created from the partial
	// signatures
	Round uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	// signature of the previous round - could be removed at some point but now
	// is used to verify the signature even before accessing the store
	PreviousSig []byte `protobuf:"bytes,2,opt,name=previous_sig,json=previousSig,proto3" json:"previous_sig,omitempty"`
	// partial signature - a threshold of them needs to be aggregated to produce
	// the final beacon at the given round.Beautiful
	PartialSig           []byte   `protobuf:"bytes,3,opt,name=partial_sig,json=partialSig,proto3" json:"partial_sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartialBeaconPacket) Reset()         { *m = PartialBeaconPacket{} }
func (m *PartialBeaconPacket) String() string { return proto.CompactTextString(m) }
func (*PartialBeaconPacket) ProtoMessage()    {}
func (*PartialBeaconPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{3}
}

func (m *PartialBeaconPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartialBeaconPacket.Unmarshal(m, b)
}
func (m *PartialBeaconPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartialBeaconPacket.Marshal(b, m, deterministic)
}
func (m *PartialBeaconPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialBeaconPacket.Merge(m, src)
}
func (m *PartialBeaconPacket) XXX_Size() int {
	return xxx_messageInfo_PartialBeaconPacket.Size(m)
}
func (m *PartialBeaconPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialBeaconPacket.DiscardUnknown(m)
}

var xxx_messageInfo_PartialBeaconPacket proto.InternalMessageInfo

func (m *PartialBeaconPacket) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *PartialBeaconPacket) GetPreviousSig() []byte {
	if m != nil {
		return m.PreviousSig
	}
	return nil
}

func (m *PartialBeaconPacket) GetPartialSig() []byte {
	if m != nil {
		return m.PartialSig
	}
	return nil
}

// BroadcastPacket is the packet that nodes send to others nodes as part of the
// broadcasting protocol.
type BroadcastPacket struct {
	Dkg *dkg.Packet `protobuf:"bytes,1,opt,name=dkg,proto3" json:"dkg,omitempty"`
	// transmitter is the address of the node participating to the broadcasting
	// channel - this is to find the identity of the signer of this packet.
	Transmitter string `protobuf:"bytes,2,opt,name=transmitter,proto3" json:"transmitter,omitempty"`
	// dealer indicates if the issuer is a dealer or not during the broadcast
	Dealer bool `protobuf:"varint,3,opt,name=dealer,proto3" json:"dealer,omitempty"`
	// signature over hash of the DKG packet
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastPacket) Reset()         { *m = BroadcastPacket{} }
func (m *BroadcastPacket) String() string { return proto.CompactTextString(m) }
func (*BroadcastPacket) ProtoMessage()    {}
func (*BroadcastPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{4}
}

func (m *BroadcastPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastPacket.Unmarshal(m, b)
}
func (m *BroadcastPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastPacket.Marshal(b, m, deterministic)
}
func (m *BroadcastPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastPacket.Merge(m, src)
}
func (m *BroadcastPacket) XXX_Size() int {
	return xxx_messageInfo_BroadcastPacket.Size(m)
}
func (m *BroadcastPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastPacket.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastPacket proto.InternalMessageInfo

func (m *BroadcastPacket) GetDkg() *dkg.Packet {
	if m != nil {
		return m.Dkg
	}
	return nil
}

func (m *BroadcastPacket) GetTransmitter() string {
	if m != nil {
		return m.Transmitter
	}
	return ""
}

func (m *BroadcastPacket) GetDealer() bool {
	if m != nil {
		return m.Dealer
	}
	return false
}

func (m *BroadcastPacket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// SyncRequest is from a node that needs to sync up with the current head of the
// chain
type SyncRequest struct {
	FromRound            uint64   `protobuf:"varint,1,opt,name=from_round,json=fromRound,proto3" json:"from_round,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRequest) Reset()         { *m = SyncRequest{} }
func (m *SyncRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()    {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{5}
}

func (m *SyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRequest.Unmarshal(m, b)
}
func (m *SyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRequest.Marshal(b, m, deterministic)
}
func (m *SyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRequest.Merge(m, src)
}
func (m *SyncRequest) XXX_Size() int {
	return xxx_messageInfo_SyncRequest.Size(m)
}
func (m *SyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRequest proto.InternalMessageInfo

func (m *SyncRequest) GetFromRound() uint64 {
	if m != nil {
		return m.FromRound
	}
	return 0
}

type BeaconPacket struct {
	PreviousSig          []byte   `protobuf:"bytes,1,opt,name=previous_sig,json=previousSig,proto3" json:"previous_sig,omitempty"`
	Round                uint64   `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconPacket) Reset()         { *m = BeaconPacket{} }
func (m *BeaconPacket) String() string { return proto.CompactTextString(m) }
func (*BeaconPacket) ProtoMessage()    {}
func (*BeaconPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{6}
}

func (m *BeaconPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconPacket.Unmarshal(m, b)
}
func (m *BeaconPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconPacket.Marshal(b, m, deterministic)
}
func (m *BeaconPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconPacket.Merge(m, src)
}
func (m *BeaconPacket) XXX_Size() int {
	return xxx_messageInfo_BeaconPacket.Size(m)
}
func (m *BeaconPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconPacket.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconPacket proto.InternalMessageInfo

func (m *BeaconPacket) GetPreviousSig() []byte {
	if m != nil {
		return m.PreviousSig
	}
	return nil
}

func (m *BeaconPacket) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *BeaconPacket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*IdentityRequest)(nil), "drand.IdentityRequest")
	proto.RegisterType((*SignalDKGPacket)(nil), "drand.SignalDKGPacket")
	proto.RegisterType((*DKGInfoPacket)(nil), "drand.DKGInfoPacket")
	proto.RegisterType((*PartialBeaconPacket)(nil), "drand.PartialBeaconPacket")
	proto.RegisterType((*BroadcastPacket)(nil), "drand.BroadcastPacket")
	proto.RegisterType((*SyncRequest)(nil), "drand.SyncRequest")
	proto.RegisterType((*BeaconPacket)(nil), "drand.BeaconPacket")
}

func init() {
	proto.RegisterFile("drand/protocol.proto", fileDescriptor_e344a98fea1e2f3a)
}

var fileDescriptor_e344a98fea1e2f3a = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x25, 0xdb, 0xee, 0xd2, 0xdc, 0xb4, 0x94, 0x9d, 0x96, 0xa5, 0x04, 0x17, 0x6b, 0x44, 0xdc,
	0x07, 0x49, 0x75, 0x85, 0x05, 0xc1, 0xa7, 0xba, 0x52, 0x97, 0x7d, 0x29, 0xa9, 0x4f, 0xbe, 0x94,
	0x69, 0x32, 0x4d, 0x86, 0x36, 0x33, 0x71, 0x32, 0x71, 0xe9, 0x1f, 0xe8, 0x6f, 0xf8, 0x57, 0xfe,
	0x8d, 0x64, 0x66, 0x12, 0xd3, 0xac, 0xa0, 0x0f, 0x85, 0xde, 0x73, 0xe7, 0xde, 0x39, 0x39, 0xf7,
	0xdc, 0x81, 0x71, 0x24, 0x30, 0x8b, 0x66, 0x99, 0xe0, 0x92, 0x87, 0x7c, 0xef, 0xab, 0x3f, 0xe8,
	0x54, 0xa1, 0xee, 0x38, 0x14, 0x87, 0x4c, 0xf2, 0x59, 0xb4, 0x8b, 0xcb, 0x9f, 0x4e, 0xba, 0x48,
	0x97, 0x84, 0x3c, 0x4d, 0x39, 0xd3, 0x98, 0x77, 0x0e, 0xc3, 0xbb, 0x88, 0x30, 0x49, 0xe5, 0x21,
	0x20, 0x5f, 0x0b, 0x92, 0x4b, 0xef, 0x87, 0x05, 0xc3, 0x15, 0x8d, 0x19, 0xde, 0xdf, 0xde, 0x2f,
	0x96, 0x38, 0xdc, 0x11, 0x89, 0x9e, 0x43, 0x97, 0xf1, 0x88, 0x4c, 0xac, 0xa9, 0x75, 0xe5, 0x5c,
	0x0f, 0x7d, 0xd5, 0xc9, 0xaf, 0x2b, 0x55, 0x12, 0x3d, 0x83, 0x7e, 0x4e, 0x42, 0x41, 0xe4, 0x3a,
	0x13, 0x9c, 0x6f, 0x27, 0x27, 0x53, 0xeb, 0xaa, 0x1f, 0x38, 0x1a, 0x5b, 0x96, 0x10, 0xf2, 0x61,
	0x94, 0x09, 0xf2, 0x8d, 0xf2, 0x22, 0x5f, 0xc7, 0x82, 0x17, 0xd9, 0x3a, 0xc1, 0x79, 0x32, 0xe9,
	0xa8, 0x93, 0xe7, 0x55, 0x6a, 0x51, 0x66, 0x3e, 0xe1, 0x3c, 0xf1, 0x7e, 0x5a, 0x30, 0xb8, 0xbd,
	0x5f, 0xdc, 0xb1, 0x2d, 0x37, 0x4c, 0x66, 0x60, 0x33, 0xf2, 0xa0, 0x8b, 0x0d, 0x1d, 0x64, 0xe8,
	0xa8, 0x32, 0x7d, 0x2c, 0xe8, 0x31, 0xf2, 0xa0, 0xe2, 0xff, 0x61, 0xf5, 0x14, 0x9c, 0x68, 0x17,
	0xaf, 0x25, 0x4d, 0x09, 0x2f, 0xa4, 0x62, 0x33, 0x08, 0x20, 0xda, 0xc5, 0x9f, 0x35, 0x82, 0x9e,
	0x80, 0x9d, 0x97, 0x8a, 0xc8, 0x42, 0x90, 0x49, 0x57, 0x35, 0xf8, 0x03, 0x78, 0x1c, 0x46, 0x4b,
	0x2c, 0x24, 0xc5, 0xfb, 0x39, 0xc1, 0x21, 0x67, 0x86, 0xe9, 0x18, 0x4e, 0x05, 0x2f, 0x58, 0xa4,
	0x58, 0x76, 0x03, 0x1d, 0x94, 0x74, 0x6a, 0x05, 0x72, 0x1a, 0x57, 0x74, 0x2a, 0x6c, 0x45, 0xe3,
	0x92, 0x4e, 0xa6, 0xfb, 0xa9, 0x13, 0x5a, 0x1c, 0x30, 0xd0, 0x8a, 0xc6, 0xde, 0x77, 0x0b, 0x86,
	0x73, 0xc1, 0x71, 0x14, 0xe2, 0x5c, 0x9a, 0xdb, 0x2e, 0xa1, 0x13, 0xed, 0x62, 0xa3, 0x88, 0xe3,
	0x97, 0x53, 0x37, 0x52, 0x94, 0x38, 0x9a, 0x82, 0x23, 0x05, 0x66, 0x79, 0x4a, 0xa5, 0x24, 0x42,
	0xdd, 0x6a, 0x07, 0x4d, 0x08, 0x5d, 0xc0, 0x59, 0x44, 0xf0, 0x9e, 0x08, 0x75, 0x61, 0x2f, 0x30,
	0xd1, 0x3f, 0xbe, 0xfd, 0x15, 0x38, 0xab, 0x03, 0x0b, 0x8d, 0x77, 0xd0, 0x25, 0xc0, 0x56, 0xf0,
	0x74, 0xdd, 0xfc, 0x70, 0xbb, 0x44, 0x82, 0x12, 0xf0, 0x08, 0xf4, 0x8f, 0x24, 0x6a, 0x8b, 0x61,
	0x3d, 0x16, 0xa3, 0x56, 0xf1, 0xa4, 0xa9, 0xe2, 0x11, 0xa9, 0x4e, 0x8b, 0xd4, 0xf5, 0xaf, 0x13,
	0xe8, 0x2d, 0xcd, 0x62, 0xa0, 0x1b, 0x70, 0x16, 0x44, 0x56, 0x56, 0x45, 0x17, 0x6d, 0xef, 0x6a,
	0xe6, 0x6e, 0xdb, 0xd3, 0xe8, 0x3d, 0x8c, 0x1b, 0x5b, 0x20, 0x24, 0x0d, 0x69, 0x86, 0x99, 0xac,
	0x1b, 0xb4, 0x56, 0xc4, 0xed, 0x1b, 0xfc, 0x63, 0x9a, 0xc9, 0x03, 0x7a, 0x03, 0xce, 0xb2, 0xc8,
	0x13, 0xe3, 0x5d, 0x34, 0x36, 0xc9, 0x23, 0x2f, 0x3f, 0x2a, 0xb1, 0xeb, 0xa1, 0xd6, 0xb7, 0xb4,
	0xc6, 0xdc, 0x2a, 0x79, 0x07, 0x83, 0x23, 0xe7, 0x21, 0xd7, 0xa4, 0xff, 0xe2, 0xc7, 0x56, 0xe9,
	0x0d, 0xd8, 0xe5, 0xe0, 0x3e, 0x24, 0x98, 0x32, 0x54, 0x6d, 0x50, 0x63, 0x94, 0xee, 0xa8, 0x62,
	0xd0, 0xe8, 0xf1, 0xda, 0x9a, 0xbf, 0xfc, 0xf2, 0x22, 0xa6, 0x32, 0x29, 0x36, 0x7e, 0xc8, 0xd3,
	0x99, 0x7e, 0x51, 0x1a, 0x4f, 0xd1, 0xa6, 0xd8, 0xea, 0x70, 0x73, 0xa6, 0xe2, 0xb7, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x9a, 0x67, 0x74, 0x0d, 0xa9, 0x04, 0x00, 0x00,
}
