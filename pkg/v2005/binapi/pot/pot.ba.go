// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.05-release
// source: /usr/share/vpp/api/plugins/pot.api.json

// Package pot contains generated bindings for API file pot.api.
//
// Contents:
//   8 messages
//
package pot

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "pot"
	APIVersion = "1.0.0"
	VersionCrc = 0xa9d8e55c
)

// PotProfileActivate defines message 'pot_profile_activate'.
type PotProfileActivate struct {
	ID       uint8  `binapi:"u8,name=id" json:"id,omitempty"`
	ListName string `binapi:"string[],name=list_name" json:"list_name,omitempty"`
}

func (m *PotProfileActivate) Reset()               { *m = PotProfileActivate{} }
func (*PotProfileActivate) GetMessageName() string { return "pot_profile_activate" }
func (*PotProfileActivate) GetCrcString() string   { return "0770af98" }
func (*PotProfileActivate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PotProfileActivate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1                   // m.ID
	size += 4 + len(m.ListName) // m.ListName
	return size
}
func (m *PotProfileActivate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.ID)
	buf.EncodeString(m.ListName, 0)
	return buf.Bytes(), nil
}
func (m *PotProfileActivate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint8()
	m.ListName = buf.DecodeString(0)
	return nil
}

// PotProfileActivateReply defines message 'pot_profile_activate_reply'.
type PotProfileActivateReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PotProfileActivateReply) Reset()               { *m = PotProfileActivateReply{} }
func (*PotProfileActivateReply) GetMessageName() string { return "pot_profile_activate_reply" }
func (*PotProfileActivateReply) GetCrcString() string   { return "e8d4e804" }
func (*PotProfileActivateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PotProfileActivateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PotProfileActivateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PotProfileActivateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PotProfileAdd defines message 'pot_profile_add'.
type PotProfileAdd struct {
	ID               uint8  `binapi:"u8,name=id" json:"id,omitempty"`
	Validator        uint8  `binapi:"u8,name=validator" json:"validator,omitempty"`
	SecretKey        uint64 `binapi:"u64,name=secret_key" json:"secret_key,omitempty"`
	SecretShare      uint64 `binapi:"u64,name=secret_share" json:"secret_share,omitempty"`
	Prime            uint64 `binapi:"u64,name=prime" json:"prime,omitempty"`
	MaxBits          uint8  `binapi:"u8,name=max_bits" json:"max_bits,omitempty"`
	Lpc              uint64 `binapi:"u64,name=lpc" json:"lpc,omitempty"`
	PolynomialPublic uint64 `binapi:"u64,name=polynomial_public" json:"polynomial_public,omitempty"`
	ListName         string `binapi:"string[],name=list_name" json:"list_name,omitempty"`
}

func (m *PotProfileAdd) Reset()               { *m = PotProfileAdd{} }
func (*PotProfileAdd) GetMessageName() string { return "pot_profile_add" }
func (*PotProfileAdd) GetCrcString() string   { return "ad5da3a3" }
func (*PotProfileAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PotProfileAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1                   // m.ID
	size += 1                   // m.Validator
	size += 8                   // m.SecretKey
	size += 8                   // m.SecretShare
	size += 8                   // m.Prime
	size += 1                   // m.MaxBits
	size += 8                   // m.Lpc
	size += 8                   // m.PolynomialPublic
	size += 4 + len(m.ListName) // m.ListName
	return size
}
func (m *PotProfileAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.ID)
	buf.EncodeUint8(m.Validator)
	buf.EncodeUint64(m.SecretKey)
	buf.EncodeUint64(m.SecretShare)
	buf.EncodeUint64(m.Prime)
	buf.EncodeUint8(m.MaxBits)
	buf.EncodeUint64(m.Lpc)
	buf.EncodeUint64(m.PolynomialPublic)
	buf.EncodeString(m.ListName, 0)
	return buf.Bytes(), nil
}
func (m *PotProfileAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint8()
	m.Validator = buf.DecodeUint8()
	m.SecretKey = buf.DecodeUint64()
	m.SecretShare = buf.DecodeUint64()
	m.Prime = buf.DecodeUint64()
	m.MaxBits = buf.DecodeUint8()
	m.Lpc = buf.DecodeUint64()
	m.PolynomialPublic = buf.DecodeUint64()
	m.ListName = buf.DecodeString(0)
	return nil
}

// PotProfileAddReply defines message 'pot_profile_add_reply'.
type PotProfileAddReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PotProfileAddReply) Reset()               { *m = PotProfileAddReply{} }
func (*PotProfileAddReply) GetMessageName() string { return "pot_profile_add_reply" }
func (*PotProfileAddReply) GetCrcString() string   { return "e8d4e804" }
func (*PotProfileAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PotProfileAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PotProfileAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PotProfileAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PotProfileDel defines message 'pot_profile_del'.
type PotProfileDel struct {
	ListName string `binapi:"string[],name=list_name" json:"list_name,omitempty"`
}

func (m *PotProfileDel) Reset()               { *m = PotProfileDel{} }
func (*PotProfileDel) GetMessageName() string { return "pot_profile_del" }
func (*PotProfileDel) GetCrcString() string   { return "cd63f53b" }
func (*PotProfileDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PotProfileDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 + len(m.ListName) // m.ListName
	return size
}
func (m *PotProfileDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.ListName, 0)
	return buf.Bytes(), nil
}
func (m *PotProfileDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ListName = buf.DecodeString(0)
	return nil
}

// PotProfileDelReply defines message 'pot_profile_del_reply'.
type PotProfileDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PotProfileDelReply) Reset()               { *m = PotProfileDelReply{} }
func (*PotProfileDelReply) GetMessageName() string { return "pot_profile_del_reply" }
func (*PotProfileDelReply) GetCrcString() string   { return "e8d4e804" }
func (*PotProfileDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PotProfileDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PotProfileDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PotProfileDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PotProfileShowConfigDetails defines message 'pot_profile_show_config_details'.
type PotProfileShowConfigDetails struct {
	Retval           int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ID               uint8  `binapi:"u8,name=id" json:"id,omitempty"`
	Validator        uint8  `binapi:"u8,name=validator" json:"validator,omitempty"`
	SecretKey        uint64 `binapi:"u64,name=secret_key" json:"secret_key,omitempty"`
	SecretShare      uint64 `binapi:"u64,name=secret_share" json:"secret_share,omitempty"`
	Prime            uint64 `binapi:"u64,name=prime" json:"prime,omitempty"`
	BitMask          uint64 `binapi:"u64,name=bit_mask" json:"bit_mask,omitempty"`
	Lpc              uint64 `binapi:"u64,name=lpc" json:"lpc,omitempty"`
	PolynomialPublic uint64 `binapi:"u64,name=polynomial_public" json:"polynomial_public,omitempty"`
}

func (m *PotProfileShowConfigDetails) Reset()               { *m = PotProfileShowConfigDetails{} }
func (*PotProfileShowConfigDetails) GetMessageName() string { return "pot_profile_show_config_details" }
func (*PotProfileShowConfigDetails) GetCrcString() string   { return "b7ce0618" }
func (*PotProfileShowConfigDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PotProfileShowConfigDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 1 // m.ID
	size += 1 // m.Validator
	size += 8 // m.SecretKey
	size += 8 // m.SecretShare
	size += 8 // m.Prime
	size += 8 // m.BitMask
	size += 8 // m.Lpc
	size += 8 // m.PolynomialPublic
	return size
}
func (m *PotProfileShowConfigDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint8(m.ID)
	buf.EncodeUint8(m.Validator)
	buf.EncodeUint64(m.SecretKey)
	buf.EncodeUint64(m.SecretShare)
	buf.EncodeUint64(m.Prime)
	buf.EncodeUint64(m.BitMask)
	buf.EncodeUint64(m.Lpc)
	buf.EncodeUint64(m.PolynomialPublic)
	return buf.Bytes(), nil
}
func (m *PotProfileShowConfigDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ID = buf.DecodeUint8()
	m.Validator = buf.DecodeUint8()
	m.SecretKey = buf.DecodeUint64()
	m.SecretShare = buf.DecodeUint64()
	m.Prime = buf.DecodeUint64()
	m.BitMask = buf.DecodeUint64()
	m.Lpc = buf.DecodeUint64()
	m.PolynomialPublic = buf.DecodeUint64()
	return nil
}

// PotProfileShowConfigDump defines message 'pot_profile_show_config_dump'.
type PotProfileShowConfigDump struct {
	ID uint8 `binapi:"u8,name=id" json:"id,omitempty"`
}

func (m *PotProfileShowConfigDump) Reset()               { *m = PotProfileShowConfigDump{} }
func (*PotProfileShowConfigDump) GetMessageName() string { return "pot_profile_show_config_dump" }
func (*PotProfileShowConfigDump) GetCrcString() string   { return "005b7d59" }
func (*PotProfileShowConfigDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PotProfileShowConfigDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.ID
	return size
}
func (m *PotProfileShowConfigDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.ID)
	return buf.Bytes(), nil
}
func (m *PotProfileShowConfigDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint8()
	return nil
}

func init() { file_pot_binapi_init() }
func file_pot_binapi_init() {
	api.RegisterMessage((*PotProfileActivate)(nil), "pot_profile_activate_0770af98")
	api.RegisterMessage((*PotProfileActivateReply)(nil), "pot_profile_activate_reply_e8d4e804")
	api.RegisterMessage((*PotProfileAdd)(nil), "pot_profile_add_ad5da3a3")
	api.RegisterMessage((*PotProfileAddReply)(nil), "pot_profile_add_reply_e8d4e804")
	api.RegisterMessage((*PotProfileDel)(nil), "pot_profile_del_cd63f53b")
	api.RegisterMessage((*PotProfileDelReply)(nil), "pot_profile_del_reply_e8d4e804")
	api.RegisterMessage((*PotProfileShowConfigDetails)(nil), "pot_profile_show_config_details_b7ce0618")
	api.RegisterMessage((*PotProfileShowConfigDump)(nil), "pot_profile_show_config_dump_005b7d59")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PotProfileActivate)(nil),
		(*PotProfileActivateReply)(nil),
		(*PotProfileAdd)(nil),
		(*PotProfileAddReply)(nil),
		(*PotProfileDel)(nil),
		(*PotProfileDelReply)(nil),
		(*PotProfileShowConfigDetails)(nil),
		(*PotProfileShowConfigDump)(nil),
	}
}
