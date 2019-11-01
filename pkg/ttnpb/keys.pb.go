// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/keys.proto

package ttnpb

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	go_thethings_network_lorawan_stack_pkg_types "go.thethings.network/lorawan-stack/pkg/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type KeyEnvelope struct {
	// The unencrypted AES key.
	Key *go_thethings_network_lorawan_stack_pkg_types.AES128Key `protobuf:"bytes,1,opt,name=key,proto3,customtype=go.thethings.network/lorawan-stack/pkg/types.AES128Key" json:"key,omitempty"`
	// The label of the RFC 3394 key-encryption-key (KEK) that was used to encrypt the key.
	KEKLabel             string   `protobuf:"bytes,2,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
	EncryptedKey         []byte   `protobuf:"bytes,3,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEnvelope) Reset()      { *m = KeyEnvelope{} }
func (*KeyEnvelope) ProtoMessage() {}
func (*KeyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee170ee4ccd55993, []int{0}
}
func (m *KeyEnvelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyEnvelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEnvelope.Merge(m, src)
}
func (m *KeyEnvelope) XXX_Size() int {
	return m.Size()
}
func (m *KeyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEnvelope proto.InternalMessageInfo

func (m *KeyEnvelope) GetKEKLabel() string {
	if m != nil {
		return m.KEKLabel
	}
	return ""
}

func (m *KeyEnvelope) GetEncryptedKey() []byte {
	if m != nil {
		return m.EncryptedKey
	}
	return nil
}

// Root keys for a LoRaWAN device.
// These are stored on the Join Server.
type RootKeys struct {
	// Join Server issued identifier for the root keys.
	RootKeyID string `protobuf:"bytes,1,opt,name=root_key_id,json=rootKeyId,proto3" json:"root_key_id,omitempty"`
	// The (encrypted) Application Key.
	AppKey *KeyEnvelope `protobuf:"bytes,2,opt,name=app_key,json=appKey,proto3" json:"app_key,omitempty"`
	// The (encrypted) Network Key.
	NwkKey               *KeyEnvelope `protobuf:"bytes,3,opt,name=nwk_key,json=nwkKey,proto3" json:"nwk_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RootKeys) Reset()      { *m = RootKeys{} }
func (*RootKeys) ProtoMessage() {}
func (*RootKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee170ee4ccd55993, []int{1}
}
func (m *RootKeys) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RootKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RootKeys.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RootKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RootKeys.Merge(m, src)
}
func (m *RootKeys) XXX_Size() int {
	return m.Size()
}
func (m *RootKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_RootKeys.DiscardUnknown(m)
}

var xxx_messageInfo_RootKeys proto.InternalMessageInfo

func (m *RootKeys) GetRootKeyID() string {
	if m != nil {
		return m.RootKeyID
	}
	return ""
}

func (m *RootKeys) GetAppKey() *KeyEnvelope {
	if m != nil {
		return m.AppKey
	}
	return nil
}

func (m *RootKeys) GetNwkKey() *KeyEnvelope {
	if m != nil {
		return m.NwkKey
	}
	return nil
}

// Session keys for a LoRaWAN session.
// Only the components for which the keys were meant, will have the key-encryption-key (KEK) to decrypt the individual keys.
type SessionKeys struct {
	// Join Server issued identifier for the session keys.
	// This ID can be used to request the keys from the Join Server in case the are lost.
	SessionKeyID []byte `protobuf:"bytes,1,opt,name=session_key_id,json=sessionKeyId,proto3" json:"session_key_id,omitempty"`
	// The (encrypted) Forwarding Network Session Integrity Key (or Network Session Key in 1.0 compatibility mode).
	// This key is stored by the (forwarding) Network Server.
	FNwkSIntKey *KeyEnvelope `protobuf:"bytes,2,opt,name=f_nwk_s_int_key,json=fNwkSIntKey,proto3" json:"f_nwk_s_int_key,omitempty"`
	// The (encrypted) Serving Network Session Integrity Key.
	// This key is stored by the (serving) Network Server.
	SNwkSIntKey *KeyEnvelope `protobuf:"bytes,3,opt,name=s_nwk_s_int_key,json=sNwkSIntKey,proto3" json:"s_nwk_s_int_key,omitempty"`
	// The (encrypted) Network Session Encryption Key.
	// This key is stored by the (serving) Network Server.
	NwkSEncKey *KeyEnvelope `protobuf:"bytes,4,opt,name=nwk_s_enc_key,json=nwkSEncKey,proto3" json:"nwk_s_enc_key,omitempty"`
	// The (encrypted) Application Session Key.
	// This key is stored by the Application Server.
	AppSKey              *KeyEnvelope `protobuf:"bytes,5,opt,name=app_s_key,json=appSKey,proto3" json:"app_s_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SessionKeys) Reset()      { *m = SessionKeys{} }
func (*SessionKeys) ProtoMessage() {}
func (*SessionKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee170ee4ccd55993, []int{2}
}
func (m *SessionKeys) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionKeys.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SessionKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionKeys.Merge(m, src)
}
func (m *SessionKeys) XXX_Size() int {
	return m.Size()
}
func (m *SessionKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionKeys.DiscardUnknown(m)
}

var xxx_messageInfo_SessionKeys proto.InternalMessageInfo

func (m *SessionKeys) GetSessionKeyID() []byte {
	if m != nil {
		return m.SessionKeyID
	}
	return nil
}

func (m *SessionKeys) GetFNwkSIntKey() *KeyEnvelope {
	if m != nil {
		return m.FNwkSIntKey
	}
	return nil
}

func (m *SessionKeys) GetSNwkSIntKey() *KeyEnvelope {
	if m != nil {
		return m.SNwkSIntKey
	}
	return nil
}

func (m *SessionKeys) GetNwkSEncKey() *KeyEnvelope {
	if m != nil {
		return m.NwkSEncKey
	}
	return nil
}

func (m *SessionKeys) GetAppSKey() *KeyEnvelope {
	if m != nil {
		return m.AppSKey
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyEnvelope)(nil), "ttn.lorawan.v3.KeyEnvelope")
	golang_proto.RegisterType((*KeyEnvelope)(nil), "ttn.lorawan.v3.KeyEnvelope")
	proto.RegisterType((*RootKeys)(nil), "ttn.lorawan.v3.RootKeys")
	golang_proto.RegisterType((*RootKeys)(nil), "ttn.lorawan.v3.RootKeys")
	proto.RegisterType((*SessionKeys)(nil), "ttn.lorawan.v3.SessionKeys")
	golang_proto.RegisterType((*SessionKeys)(nil), "ttn.lorawan.v3.SessionKeys")
}

func init() { proto.RegisterFile("lorawan-stack/api/keys.proto", fileDescriptor_ee170ee4ccd55993) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/keys.proto", fileDescriptor_ee170ee4ccd55993)
}

var fileDescriptor_ee170ee4ccd55993 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x3f, 0x4c, 0xdb, 0x4e,
	0x14, 0xc7, 0xef, 0xf8, 0xf7, 0x23, 0x97, 0xc0, 0xaf, 0x8a, 0x5a, 0x29, 0xa2, 0xd5, 0x03, 0xd1,
	0x85, 0xa1, 0xb1, 0x05, 0xf4, 0x9f, 0x18, 0x90, 0xb0, 0xc8, 0x80, 0x8c, 0x3a, 0x38, 0x5b, 0x97,
	0xc8, 0x49, 0x0e, 0x63, 0x5d, 0x7a, 0x67, 0xf9, 0x8e, 0xa4, 0xee, 0xc4, 0xc8, 0xd8, 0xb1, 0x23,
	0xea, 0xc4, 0x52, 0x09, 0xa9, 0x0b, 0x23, 0x23, 0x23, 0xea, 0x84, 0x3a, 0x20, 0x7c, 0x5e, 0x18,
	0x19, 0x11, 0x53, 0x75, 0x4e, 0x04, 0xa1, 0x1d, 0x60, 0x7b, 0xcf, 0xfe, 0x7e, 0xbf, 0xf7, 0x79,
	0xf7, 0x74, 0xe4, 0x45, 0x47, 0xc4, 0x7e, 0xcf, 0xe7, 0x55, 0xa9, 0xfc, 0x16, 0xb3, 0xfd, 0x28,
	0xb4, 0x19, 0x4d, 0xa4, 0x15, 0xc5, 0x42, 0x89, 0xf2, 0xb4, 0x52, 0xdc, 0x1a, 0x28, 0xac, 0xee,
	0xf2, 0xcc, 0x5a, 0x10, 0xaa, 0xed, 0x9d, 0xa6, 0xd5, 0x12, 0x9f, 0x6c, 0xca, 0xbb, 0x22, 0x89,
	0x62, 0xf1, 0x39, 0xb1, 0x73, 0x71, 0xab, 0x1a, 0x50, 0x5e, 0xed, 0xfa, 0x9d, 0xb0, 0xed, 0x2b,
	0x6a, 0xff, 0x53, 0xf4, 0x23, 0x67, 0xaa, 0x43, 0x11, 0x81, 0x08, 0x44, 0xdf, 0xdc, 0xdc, 0xd9,
	0xca, 0xbb, 0xbc, 0xc9, 0xab, 0xbe, 0x7c, 0xfe, 0x27, 0x26, 0x45, 0x97, 0x26, 0x35, 0xde, 0xa5,
	0x1d, 0x11, 0xd1, 0xf2, 0x26, 0x19, 0x65, 0x34, 0xa9, 0xe0, 0x39, 0xbc, 0x50, 0x72, 0x56, 0x7e,
	0x9f, 0xcf, 0xbe, 0x0d, 0x84, 0xa5, 0xb6, 0xa9, 0xda, 0x0e, 0x79, 0x20, 0x2d, 0x4e, 0x55, 0x4f,
	0xc4, 0xcc, 0xbe, 0x3f, 0x55, 0xc4, 0x02, 0x5b, 0x25, 0x11, 0x95, 0xd6, 0x5a, 0xad, 0xbe, 0xb8,
	0xf4, 0xde, 0xa5, 0x89, 0x67, 0x62, 0xca, 0x8b, 0xa4, 0xc0, 0x28, 0x6b, 0x74, 0xfc, 0x26, 0xed,
	0x54, 0x46, 0xe6, 0xf0, 0x42, 0xc1, 0x79, 0x7a, 0xe3, 0x8c, 0xc7, 0xa3, 0x95, 0xdd, 0x27, 0xfa,
	0x7c, 0x76, 0xd2, 0xad, 0xb9, 0x9b, 0xe6, 0x9f, 0x37, 0xc9, 0x28, 0xcb, 0xab, 0xf2, 0x4b, 0x32,
	0x45, 0x79, 0x2b, 0x4e, 0x22, 0x45, 0xdb, 0x0d, 0x83, 0x32, 0x6a, 0x50, 0xbc, 0xd2, 0xed, 0x47,
	0x97, 0x26, 0xf3, 0x3f, 0x30, 0x99, 0xf4, 0x84, 0x50, 0x2e, 0x4d, 0x64, 0xf9, 0x0d, 0x29, 0xc6,
	0x42, 0x28, 0x23, 0x6e, 0x84, 0xed, 0x1c, 0xbd, 0xe0, 0x3c, 0x1b, 0x3a, 0xa6, 0x30, 0x90, 0x6e,
	0xac, 0x7b, 0x85, 0x78, 0x50, 0xb6, 0xcb, 0xaf, 0xc9, 0x7f, 0x7e, 0x14, 0xe5, 0x47, 0x18, 0xb2,
	0xe2, 0xd2, 0x73, 0xeb, 0xfe, 0x36, 0xac, 0xa1, 0x7b, 0xf1, 0x26, 0xfc, 0x28, 0x72, 0x69, 0x62,
	0x5c, 0xbc, 0xc7, 0x6e, 0xc1, 0x1e, 0x72, 0xf1, 0x1e, 0x33, 0xbc, 0xbf, 0x46, 0x48, 0xb1, 0x4e,
	0xa5, 0x0c, 0x05, 0xcf, 0x91, 0x57, 0xc9, 0xb4, 0xec, 0xb7, 0xc3, 0xd4, 0x25, 0xa7, 0x72, 0xe3,
	0x8c, 0x7f, 0x19, 0x50, 0x97, 0xee, 0x0c, 0x1b, 0xeb, 0x5e, 0x49, 0xde, 0x75, 0xed, 0xf2, 0x1a,
	0xf9, 0x7f, 0xab, 0x61, 0x38, 0x64, 0x23, 0xe4, 0xea, 0xb1, 0x33, 0x14, 0xb7, 0x3e, 0xf4, 0x58,
	0x7d, 0x83, 0x9b, 0x0b, 0x30, 0x11, 0xf2, 0xaf, 0x88, 0x47, 0x0c, 0x54, 0x94, 0x43, 0x11, 0xab,
	0x64, 0xaa, 0x1f, 0x40, 0x79, 0x2b, 0x0f, 0x18, 0x7b, 0x38, 0x80, 0xf0, 0x1e, 0xab, 0xd7, 0x78,
	0xcb, 0xf8, 0xdf, 0x91, 0x82, 0xd9, 0x80, 0xcc, 0xbd, 0xe3, 0x0f, 0x7b, 0xcd, 0xbe, 0xea, 0x2e,
	0x4d, 0x56, 0xc6, 0x8e, 0xf6, 0x67, 0x91, 0xf3, 0x1d, 0x9f, 0xa4, 0x80, 0x4f, 0x53, 0xc0, 0x67,
	0x29, 0xa0, 0x8b, 0x14, 0xd0, 0x65, 0x0a, 0xe8, 0x2a, 0x05, 0x74, 0x9d, 0x02, 0xde, 0xd5, 0x80,
	0xf7, 0x34, 0xa0, 0x03, 0x0d, 0xf8, 0x50, 0x03, 0x3a, 0xd2, 0x80, 0x8e, 0x35, 0xa0, 0x13, 0x0d,
	0xf8, 0x54, 0x03, 0x3e, 0xd3, 0x80, 0x2e, 0x34, 0xe0, 0x4b, 0x0d, 0xe8, 0x4a, 0x03, 0xbe, 0xd6,
	0x80, 0x76, 0x33, 0x40, 0x7b, 0x19, 0xe0, 0xaf, 0x19, 0xa0, 0x6f, 0x19, 0xe0, 0xfd, 0x0c, 0xd0,
	0x41, 0x06, 0xe8, 0x30, 0x03, 0x7c, 0x94, 0x01, 0x3e, 0xce, 0x00, 0x7f, 0x7c, 0xf5, 0xd8, 0x07,
	0xa1, 0x78, 0xd4, 0x6c, 0x4e, 0xe4, 0xcf, 0x6c, 0xf9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f,
	0x3e, 0xb8, 0xae, 0x08, 0x04, 0x00, 0x00,
}

func (this *KeyEnvelope) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KeyEnvelope)
	if !ok {
		that2, ok := that.(KeyEnvelope)
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
	if that1.Key == nil {
		if this.Key != nil {
			return false
		}
	} else if !this.Key.Equal(*that1.Key) {
		return false
	}
	if this.KEKLabel != that1.KEKLabel {
		return false
	}
	if !bytes.Equal(this.EncryptedKey, that1.EncryptedKey) {
		return false
	}
	return true
}
func (this *RootKeys) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RootKeys)
	if !ok {
		that2, ok := that.(RootKeys)
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
	if this.RootKeyID != that1.RootKeyID {
		return false
	}
	if !this.AppKey.Equal(that1.AppKey) {
		return false
	}
	if !this.NwkKey.Equal(that1.NwkKey) {
		return false
	}
	return true
}
func (this *SessionKeys) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SessionKeys)
	if !ok {
		that2, ok := that.(SessionKeys)
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
	if !bytes.Equal(this.SessionKeyID, that1.SessionKeyID) {
		return false
	}
	if !this.FNwkSIntKey.Equal(that1.FNwkSIntKey) {
		return false
	}
	if !this.SNwkSIntKey.Equal(that1.SNwkSIntKey) {
		return false
	}
	if !this.NwkSEncKey.Equal(that1.NwkSEncKey) {
		return false
	}
	if !this.AppSKey.Equal(that1.AppSKey) {
		return false
	}
	return true
}
func (m *KeyEnvelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyEnvelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyEnvelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedKey) > 0 {
		i -= len(m.EncryptedKey)
		copy(dAtA[i:], m.EncryptedKey)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.EncryptedKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.KEKLabel) > 0 {
		i -= len(m.KEKLabel)
		copy(dAtA[i:], m.KEKLabel)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.KEKLabel)))
		i--
		dAtA[i] = 0x12
	}
	if m.Key != nil {
		{
			size := m.Key.Size()
			i -= size
			if _, err := m.Key.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintKeys(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RootKeys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RootKeys) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RootKeys) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NwkKey != nil {
		{
			size, err := m.NwkKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeys(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.AppKey != nil {
		{
			size, err := m.AppKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeys(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RootKeyID) > 0 {
		i -= len(m.RootKeyID)
		copy(dAtA[i:], m.RootKeyID)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.RootKeyID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SessionKeys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionKeys) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SessionKeys) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AppSKey != nil {
		{
			size, err := m.AppSKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeys(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.NwkSEncKey != nil {
		{
			size, err := m.NwkSEncKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeys(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.SNwkSIntKey != nil {
		{
			size, err := m.SNwkSIntKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeys(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.FNwkSIntKey != nil {
		{
			size, err := m.FNwkSIntKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeys(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SessionKeyID) > 0 {
		i -= len(m.SessionKeyID)
		copy(dAtA[i:], m.SessionKeyID)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.SessionKeyID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeys(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeys(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedKeyEnvelope(r randyKeys, easy bool) *KeyEnvelope {
	this := &KeyEnvelope{}
	this.Key = go_thethings_network_lorawan_stack_pkg_types.NewPopulatedAES128Key(r)
	this.KEKLabel = randStringKeys(r)
	v1 := r.Intn(100)
	this.EncryptedKey = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.EncryptedKey[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedRootKeys(r randyKeys, easy bool) *RootKeys {
	this := &RootKeys{}
	this.RootKeyID = randStringKeys(r)
	if r.Intn(5) != 0 {
		this.AppKey = NewPopulatedKeyEnvelope(r, easy)
	}
	if r.Intn(5) != 0 {
		this.NwkKey = NewPopulatedKeyEnvelope(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyKeys interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneKeys(r randyKeys) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringKeys(r randyKeys) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneKeys(r)
	}
	return string(tmps)
}
func randUnrecognizedKeys(r randyKeys, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldKeys(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldKeys(dAtA []byte, r randyKeys, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateKeys(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *KeyEnvelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	l = len(m.KEKLabel)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	l = len(m.EncryptedKey)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}

func (m *RootKeys) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RootKeyID)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.AppKey != nil {
		l = m.AppKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.NwkKey != nil {
		l = m.NwkKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}

func (m *SessionKeys) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SessionKeyID)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.FNwkSIntKey != nil {
		l = m.FNwkSIntKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.SNwkSIntKey != nil {
		l = m.SNwkSIntKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.NwkSEncKey != nil {
		l = m.NwkSEncKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.AppSKey != nil {
		l = m.AppSKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}

func sovKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeys(x uint64) (n int) {
	return sovKeys((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *KeyEnvelope) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyEnvelope{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`KEKLabel:` + fmt.Sprintf("%v", this.KEKLabel) + `,`,
		`EncryptedKey:` + fmt.Sprintf("%v", this.EncryptedKey) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RootKeys) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RootKeys{`,
		`RootKeyID:` + fmt.Sprintf("%v", this.RootKeyID) + `,`,
		`AppKey:` + strings.Replace(this.AppKey.String(), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`NwkKey:` + strings.Replace(this.NwkKey.String(), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SessionKeys) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SessionKeys{`,
		`SessionKeyID:` + fmt.Sprintf("%v", this.SessionKeyID) + `,`,
		`FNwkSIntKey:` + strings.Replace(this.FNwkSIntKey.String(), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`SNwkSIntKey:` + strings.Replace(this.SNwkSIntKey.String(), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`NwkSEncKey:` + strings.Replace(this.NwkSEncKey.String(), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`AppSKey:` + strings.Replace(this.AppSKey.String(), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringKeys(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *KeyEnvelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: KeyEnvelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyEnvelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v go_thethings_network_lorawan_stack_pkg_types.AES128Key
			m.Key = &v
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KEKLabel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KEKLabel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedKey = append(m.EncryptedKey[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptedKey == nil {
				m.EncryptedKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
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
func (m *RootKeys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: RootKeys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RootKeys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootKeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AppKey == nil {
				m.AppKey = &KeyEnvelope{}
			}
			if err := m.AppKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NwkKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NwkKey == nil {
				m.NwkKey = &KeyEnvelope{}
			}
			if err := m.NwkKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
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
func (m *SessionKeys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: SessionKeys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionKeys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionKeyID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionKeyID = append(m.SessionKeyID[:0], dAtA[iNdEx:postIndex]...)
			if m.SessionKeyID == nil {
				m.SessionKeyID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FNwkSIntKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FNwkSIntKey == nil {
				m.FNwkSIntKey = &KeyEnvelope{}
			}
			if err := m.FNwkSIntKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SNwkSIntKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SNwkSIntKey == nil {
				m.SNwkSIntKey = &KeyEnvelope{}
			}
			if err := m.SNwkSIntKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NwkSEncKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NwkSEncKey == nil {
				m.NwkSEncKey = &KeyEnvelope{}
			}
			if err := m.NwkSEncKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppSKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AppSKey == nil {
				m.AppSKey = &KeyEnvelope{}
			}
			if err := m.AppSKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
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
func skipKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
				return 0, ErrInvalidLengthKeys
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeys
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeys
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeys        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeys          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeys = fmt.Errorf("proto: unexpected end of group")
)
