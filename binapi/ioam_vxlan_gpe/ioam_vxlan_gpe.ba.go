// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              22.06-rc0~147-ge81e78f81
// source: /usr/share/vpp/api/plugins/ioam_vxlan_gpe.api.json

// Package ioam_vxlan_gpe contains generated bindings for API file ioam_vxlan_gpe.api.
//
// Contents:
//  12 messages
//
package ioam_vxlan_gpe

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "ioam_vxlan_gpe"
	APIVersion = "1.0.0"
	VersionCrc = 0xb49eb0b9
)

// VxlanGpeIoamDisable defines message 'vxlan_gpe_ioam_disable'.
type VxlanGpeIoamDisable struct {
	ID uint16 `binapi:"u16,name=id" json:"id,omitempty"`
}

func (m *VxlanGpeIoamDisable) Reset()               { *m = VxlanGpeIoamDisable{} }
func (*VxlanGpeIoamDisable) GetMessageName() string { return "vxlan_gpe_ioam_disable" }
func (*VxlanGpeIoamDisable) GetCrcString() string   { return "6b16a45e" }
func (*VxlanGpeIoamDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGpeIoamDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 2 // m.ID
	return size
}
func (m *VxlanGpeIoamDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint16(m.ID)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint16()
	return nil
}

// VxlanGpeIoamDisableReply defines message 'vxlan_gpe_ioam_disable_reply'.
type VxlanGpeIoamDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VxlanGpeIoamDisableReply) Reset()               { *m = VxlanGpeIoamDisableReply{} }
func (*VxlanGpeIoamDisableReply) GetMessageName() string { return "vxlan_gpe_ioam_disable_reply" }
func (*VxlanGpeIoamDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*VxlanGpeIoamDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGpeIoamDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VxlanGpeIoamDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VxlanGpeIoamEnable defines message 'vxlan_gpe_ioam_enable'.
type VxlanGpeIoamEnable struct {
	ID          uint16 `binapi:"u16,name=id" json:"id,omitempty"`
	TracePpc    uint8  `binapi:"u8,name=trace_ppc" json:"trace_ppc,omitempty"`
	PowEnable   bool   `binapi:"bool,name=pow_enable" json:"pow_enable,omitempty"`
	TraceEnable bool   `binapi:"bool,name=trace_enable" json:"trace_enable,omitempty"`
}

func (m *VxlanGpeIoamEnable) Reset()               { *m = VxlanGpeIoamEnable{} }
func (*VxlanGpeIoamEnable) GetMessageName() string { return "vxlan_gpe_ioam_enable" }
func (*VxlanGpeIoamEnable) GetCrcString() string   { return "2481bef7" }
func (*VxlanGpeIoamEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGpeIoamEnable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 2 // m.ID
	size += 1 // m.TracePpc
	size += 1 // m.PowEnable
	size += 1 // m.TraceEnable
	return size
}
func (m *VxlanGpeIoamEnable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint16(m.ID)
	buf.EncodeUint8(m.TracePpc)
	buf.EncodeBool(m.PowEnable)
	buf.EncodeBool(m.TraceEnable)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamEnable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint16()
	m.TracePpc = buf.DecodeUint8()
	m.PowEnable = buf.DecodeBool()
	m.TraceEnable = buf.DecodeBool()
	return nil
}

// VxlanGpeIoamEnableReply defines message 'vxlan_gpe_ioam_enable_reply'.
type VxlanGpeIoamEnableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VxlanGpeIoamEnableReply) Reset()               { *m = VxlanGpeIoamEnableReply{} }
func (*VxlanGpeIoamEnableReply) GetMessageName() string { return "vxlan_gpe_ioam_enable_reply" }
func (*VxlanGpeIoamEnableReply) GetCrcString() string   { return "e8d4e804" }
func (*VxlanGpeIoamEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGpeIoamEnableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VxlanGpeIoamEnableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamEnableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VxlanGpeIoamTransitDisable defines message 'vxlan_gpe_ioam_transit_disable'.
type VxlanGpeIoamTransitDisable struct {
	OuterFibIndex uint32           `binapi:"u32,name=outer_fib_index" json:"outer_fib_index,omitempty"`
	DstAddr       ip_types.Address `binapi:"address,name=dst_addr" json:"dst_addr,omitempty"`
}

func (m *VxlanGpeIoamTransitDisable) Reset()               { *m = VxlanGpeIoamTransitDisable{} }
func (*VxlanGpeIoamTransitDisable) GetMessageName() string { return "vxlan_gpe_ioam_transit_disable" }
func (*VxlanGpeIoamTransitDisable) GetCrcString() string   { return "3d3ec657" }
func (*VxlanGpeIoamTransitDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGpeIoamTransitDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.OuterFibIndex
	size += 1      // m.DstAddr.Af
	size += 1 * 16 // m.DstAddr.Un
	return size
}
func (m *VxlanGpeIoamTransitDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.OuterFibIndex)
	buf.EncodeUint8(uint8(m.DstAddr.Af))
	buf.EncodeBytes(m.DstAddr.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamTransitDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.OuterFibIndex = buf.DecodeUint32()
	m.DstAddr.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.DstAddr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// VxlanGpeIoamTransitDisableReply defines message 'vxlan_gpe_ioam_transit_disable_reply'.
type VxlanGpeIoamTransitDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VxlanGpeIoamTransitDisableReply) Reset() { *m = VxlanGpeIoamTransitDisableReply{} }
func (*VxlanGpeIoamTransitDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_transit_disable_reply"
}
func (*VxlanGpeIoamTransitDisableReply) GetCrcString() string { return "e8d4e804" }
func (*VxlanGpeIoamTransitDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGpeIoamTransitDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VxlanGpeIoamTransitDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamTransitDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VxlanGpeIoamTransitEnable defines message 'vxlan_gpe_ioam_transit_enable'.
type VxlanGpeIoamTransitEnable struct {
	OuterFibIndex uint32           `binapi:"u32,name=outer_fib_index" json:"outer_fib_index,omitempty"`
	DstAddr       ip_types.Address `binapi:"address,name=dst_addr" json:"dst_addr,omitempty"`
}

func (m *VxlanGpeIoamTransitEnable) Reset()               { *m = VxlanGpeIoamTransitEnable{} }
func (*VxlanGpeIoamTransitEnable) GetMessageName() string { return "vxlan_gpe_ioam_transit_enable" }
func (*VxlanGpeIoamTransitEnable) GetCrcString() string   { return "3d3ec657" }
func (*VxlanGpeIoamTransitEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGpeIoamTransitEnable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.OuterFibIndex
	size += 1      // m.DstAddr.Af
	size += 1 * 16 // m.DstAddr.Un
	return size
}
func (m *VxlanGpeIoamTransitEnable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.OuterFibIndex)
	buf.EncodeUint8(uint8(m.DstAddr.Af))
	buf.EncodeBytes(m.DstAddr.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamTransitEnable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.OuterFibIndex = buf.DecodeUint32()
	m.DstAddr.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.DstAddr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// VxlanGpeIoamTransitEnableReply defines message 'vxlan_gpe_ioam_transit_enable_reply'.
type VxlanGpeIoamTransitEnableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VxlanGpeIoamTransitEnableReply) Reset() { *m = VxlanGpeIoamTransitEnableReply{} }
func (*VxlanGpeIoamTransitEnableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_transit_enable_reply"
}
func (*VxlanGpeIoamTransitEnableReply) GetCrcString() string { return "e8d4e804" }
func (*VxlanGpeIoamTransitEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGpeIoamTransitEnableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VxlanGpeIoamTransitEnableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamTransitEnableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VxlanGpeIoamVniDisable defines message 'vxlan_gpe_ioam_vni_disable'.
type VxlanGpeIoamVniDisable struct {
	Vni    uint32           `binapi:"u32,name=vni" json:"vni,omitempty"`
	Local  ip_types.Address `binapi:"address,name=local" json:"local,omitempty"`
	Remote ip_types.Address `binapi:"address,name=remote" json:"remote,omitempty"`
}

func (m *VxlanGpeIoamVniDisable) Reset()               { *m = VxlanGpeIoamVniDisable{} }
func (*VxlanGpeIoamVniDisable) GetMessageName() string { return "vxlan_gpe_ioam_vni_disable" }
func (*VxlanGpeIoamVniDisable) GetCrcString() string   { return "0fbb5fb1" }
func (*VxlanGpeIoamVniDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGpeIoamVniDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Vni
	size += 1      // m.Local.Af
	size += 1 * 16 // m.Local.Un
	size += 1      // m.Remote.Af
	size += 1 * 16 // m.Remote.Un
	return size
}
func (m *VxlanGpeIoamVniDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Vni)
	buf.EncodeUint8(uint8(m.Local.Af))
	buf.EncodeBytes(m.Local.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Remote.Af))
	buf.EncodeBytes(m.Remote.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamVniDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Vni = buf.DecodeUint32()
	m.Local.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Local.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Remote.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Remote.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// VxlanGpeIoamVniDisableReply defines message 'vxlan_gpe_ioam_vni_disable_reply'.
type VxlanGpeIoamVniDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VxlanGpeIoamVniDisableReply) Reset() { *m = VxlanGpeIoamVniDisableReply{} }
func (*VxlanGpeIoamVniDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_vni_disable_reply"
}
func (*VxlanGpeIoamVniDisableReply) GetCrcString() string { return "e8d4e804" }
func (*VxlanGpeIoamVniDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGpeIoamVniDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VxlanGpeIoamVniDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamVniDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VxlanGpeIoamVniEnable defines message 'vxlan_gpe_ioam_vni_enable'.
type VxlanGpeIoamVniEnable struct {
	Vni    uint32           `binapi:"u32,name=vni" json:"vni,omitempty"`
	Local  ip_types.Address `binapi:"address,name=local" json:"local,omitempty"`
	Remote ip_types.Address `binapi:"address,name=remote" json:"remote,omitempty"`
}

func (m *VxlanGpeIoamVniEnable) Reset()               { *m = VxlanGpeIoamVniEnable{} }
func (*VxlanGpeIoamVniEnable) GetMessageName() string { return "vxlan_gpe_ioam_vni_enable" }
func (*VxlanGpeIoamVniEnable) GetCrcString() string   { return "0fbb5fb1" }
func (*VxlanGpeIoamVniEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGpeIoamVniEnable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Vni
	size += 1      // m.Local.Af
	size += 1 * 16 // m.Local.Un
	size += 1      // m.Remote.Af
	size += 1 * 16 // m.Remote.Un
	return size
}
func (m *VxlanGpeIoamVniEnable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Vni)
	buf.EncodeUint8(uint8(m.Local.Af))
	buf.EncodeBytes(m.Local.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Remote.Af))
	buf.EncodeBytes(m.Remote.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamVniEnable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Vni = buf.DecodeUint32()
	m.Local.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Local.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Remote.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Remote.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// VxlanGpeIoamVniEnableReply defines message 'vxlan_gpe_ioam_vni_enable_reply'.
type VxlanGpeIoamVniEnableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VxlanGpeIoamVniEnableReply) Reset()               { *m = VxlanGpeIoamVniEnableReply{} }
func (*VxlanGpeIoamVniEnableReply) GetMessageName() string { return "vxlan_gpe_ioam_vni_enable_reply" }
func (*VxlanGpeIoamVniEnableReply) GetCrcString() string   { return "e8d4e804" }
func (*VxlanGpeIoamVniEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGpeIoamVniEnableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VxlanGpeIoamVniEnableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamVniEnableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_ioam_vxlan_gpe_binapi_init() }
func file_ioam_vxlan_gpe_binapi_init() {
	api.RegisterMessage((*VxlanGpeIoamDisable)(nil), "vxlan_gpe_ioam_disable_6b16a45e")
	api.RegisterMessage((*VxlanGpeIoamDisableReply)(nil), "vxlan_gpe_ioam_disable_reply_e8d4e804")
	api.RegisterMessage((*VxlanGpeIoamEnable)(nil), "vxlan_gpe_ioam_enable_2481bef7")
	api.RegisterMessage((*VxlanGpeIoamEnableReply)(nil), "vxlan_gpe_ioam_enable_reply_e8d4e804")
	api.RegisterMessage((*VxlanGpeIoamTransitDisable)(nil), "vxlan_gpe_ioam_transit_disable_3d3ec657")
	api.RegisterMessage((*VxlanGpeIoamTransitDisableReply)(nil), "vxlan_gpe_ioam_transit_disable_reply_e8d4e804")
	api.RegisterMessage((*VxlanGpeIoamTransitEnable)(nil), "vxlan_gpe_ioam_transit_enable_3d3ec657")
	api.RegisterMessage((*VxlanGpeIoamTransitEnableReply)(nil), "vxlan_gpe_ioam_transit_enable_reply_e8d4e804")
	api.RegisterMessage((*VxlanGpeIoamVniDisable)(nil), "vxlan_gpe_ioam_vni_disable_0fbb5fb1")
	api.RegisterMessage((*VxlanGpeIoamVniDisableReply)(nil), "vxlan_gpe_ioam_vni_disable_reply_e8d4e804")
	api.RegisterMessage((*VxlanGpeIoamVniEnable)(nil), "vxlan_gpe_ioam_vni_enable_0fbb5fb1")
	api.RegisterMessage((*VxlanGpeIoamVniEnableReply)(nil), "vxlan_gpe_ioam_vni_enable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*VxlanGpeIoamDisable)(nil),
		(*VxlanGpeIoamDisableReply)(nil),
		(*VxlanGpeIoamEnable)(nil),
		(*VxlanGpeIoamEnableReply)(nil),
		(*VxlanGpeIoamTransitDisable)(nil),
		(*VxlanGpeIoamTransitDisableReply)(nil),
		(*VxlanGpeIoamTransitEnable)(nil),
		(*VxlanGpeIoamTransitEnableReply)(nil),
		(*VxlanGpeIoamVniDisable)(nil),
		(*VxlanGpeIoamVniDisableReply)(nil),
		(*VxlanGpeIoamVniEnable)(nil),
		(*VxlanGpeIoamVniEnableReply)(nil),
	}
}
