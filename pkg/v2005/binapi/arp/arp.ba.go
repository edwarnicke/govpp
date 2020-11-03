// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.05-release
// source: /usr/share/vpp/api/core/arp.api.json

// Package arp contains generated bindings for API file arp.api.
//
// Contents:
//   1 struct
//   8 messages
//
package arp

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	_ "github.com/edwarnicke/govpp/pkg/v2005/binapi/ethernet_types"
	interface_types "github.com/edwarnicke/govpp/pkg/v2005/binapi/interface_types"
	ip_types "github.com/edwarnicke/govpp/pkg/v2005/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "arp"
	APIVersion = "1.0.0"
	VersionCrc = 0x79ca86f2
)

// ProxyArp defines type 'proxy_arp'.
type ProxyArp struct {
	TableID uint32              `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	Low     ip_types.IP4Address `binapi:"ip4_address,name=low" json:"low,omitempty"`
	Hi      ip_types.IP4Address `binapi:"ip4_address,name=hi" json:"hi,omitempty"`
}

// ProxyArpAddDel defines message 'proxy_arp_add_del'.
type ProxyArpAddDel struct {
	IsAdd bool     `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Proxy ProxyArp `binapi:"proxy_arp,name=proxy" json:"proxy,omitempty"`
}

func (m *ProxyArpAddDel) Reset()               { *m = ProxyArpAddDel{} }
func (*ProxyArpAddDel) GetMessageName() string { return "proxy_arp_add_del" }
func (*ProxyArpAddDel) GetCrcString() string   { return "85486cbd" }
func (*ProxyArpAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ProxyArpAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1     // m.IsAdd
	size += 4     // m.Proxy.TableID
	size += 1 * 4 // m.Proxy.Low
	size += 1 * 4 // m.Proxy.Hi
	return size
}
func (m *ProxyArpAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.Proxy.TableID)
	buf.EncodeBytes(m.Proxy.Low[:], 4)
	buf.EncodeBytes(m.Proxy.Hi[:], 4)
	return buf.Bytes(), nil
}
func (m *ProxyArpAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Proxy.TableID = buf.DecodeUint32()
	copy(m.Proxy.Low[:], buf.DecodeBytes(4))
	copy(m.Proxy.Hi[:], buf.DecodeBytes(4))
	return nil
}

// ProxyArpAddDelReply defines message 'proxy_arp_add_del_reply'.
type ProxyArpAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *ProxyArpAddDelReply) Reset()               { *m = ProxyArpAddDelReply{} }
func (*ProxyArpAddDelReply) GetMessageName() string { return "proxy_arp_add_del_reply" }
func (*ProxyArpAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*ProxyArpAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ProxyArpAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *ProxyArpAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *ProxyArpAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// ProxyArpDetails defines message 'proxy_arp_details'.
type ProxyArpDetails struct {
	Proxy ProxyArp `binapi:"proxy_arp,name=proxy" json:"proxy,omitempty"`
}

func (m *ProxyArpDetails) Reset()               { *m = ProxyArpDetails{} }
func (*ProxyArpDetails) GetMessageName() string { return "proxy_arp_details" }
func (*ProxyArpDetails) GetCrcString() string   { return "9228c150" }
func (*ProxyArpDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ProxyArpDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.Proxy.TableID
	size += 1 * 4 // m.Proxy.Low
	size += 1 * 4 // m.Proxy.Hi
	return size
}
func (m *ProxyArpDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Proxy.TableID)
	buf.EncodeBytes(m.Proxy.Low[:], 4)
	buf.EncodeBytes(m.Proxy.Hi[:], 4)
	return buf.Bytes(), nil
}
func (m *ProxyArpDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Proxy.TableID = buf.DecodeUint32()
	copy(m.Proxy.Low[:], buf.DecodeBytes(4))
	copy(m.Proxy.Hi[:], buf.DecodeBytes(4))
	return nil
}

// ProxyArpDump defines message 'proxy_arp_dump'.
type ProxyArpDump struct{}

func (m *ProxyArpDump) Reset()               { *m = ProxyArpDump{} }
func (*ProxyArpDump) GetMessageName() string { return "proxy_arp_dump" }
func (*ProxyArpDump) GetCrcString() string   { return "51077d14" }
func (*ProxyArpDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ProxyArpDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *ProxyArpDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *ProxyArpDump) Unmarshal(b []byte) error {
	return nil
}

// ProxyArpIntfcDetails defines message 'proxy_arp_intfc_details'.
type ProxyArpIntfcDetails struct {
	SwIfIndex uint32 `binapi:"u32,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *ProxyArpIntfcDetails) Reset()               { *m = ProxyArpIntfcDetails{} }
func (*ProxyArpIntfcDetails) GetMessageName() string { return "proxy_arp_intfc_details" }
func (*ProxyArpIntfcDetails) GetCrcString() string   { return "f6458e5f" }
func (*ProxyArpIntfcDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ProxyArpIntfcDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *ProxyArpIntfcDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.SwIfIndex)
	return buf.Bytes(), nil
}
func (m *ProxyArpIntfcDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = buf.DecodeUint32()
	return nil
}

// ProxyArpIntfcDump defines message 'proxy_arp_intfc_dump'.
type ProxyArpIntfcDump struct{}

func (m *ProxyArpIntfcDump) Reset()               { *m = ProxyArpIntfcDump{} }
func (*ProxyArpIntfcDump) GetMessageName() string { return "proxy_arp_intfc_dump" }
func (*ProxyArpIntfcDump) GetCrcString() string   { return "51077d14" }
func (*ProxyArpIntfcDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ProxyArpIntfcDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *ProxyArpIntfcDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *ProxyArpIntfcDump) Unmarshal(b []byte) error {
	return nil
}

// ProxyArpIntfcEnableDisable defines message 'proxy_arp_intfc_enable_disable'.
type ProxyArpIntfcEnableDisable struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Enable    bool                           `binapi:"bool,name=enable" json:"enable,omitempty"`
}

func (m *ProxyArpIntfcEnableDisable) Reset()               { *m = ProxyArpIntfcEnableDisable{} }
func (*ProxyArpIntfcEnableDisable) GetMessageName() string { return "proxy_arp_intfc_enable_disable" }
func (*ProxyArpIntfcEnableDisable) GetCrcString() string   { return "ae6cfcfb" }
func (*ProxyArpIntfcEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ProxyArpIntfcEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.Enable
	return size
}
func (m *ProxyArpIntfcEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.Enable)
	return buf.Bytes(), nil
}
func (m *ProxyArpIntfcEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Enable = buf.DecodeBool()
	return nil
}

// ProxyArpIntfcEnableDisableReply defines message 'proxy_arp_intfc_enable_disable_reply'.
type ProxyArpIntfcEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *ProxyArpIntfcEnableDisableReply) Reset() { *m = ProxyArpIntfcEnableDisableReply{} }
func (*ProxyArpIntfcEnableDisableReply) GetMessageName() string {
	return "proxy_arp_intfc_enable_disable_reply"
}
func (*ProxyArpIntfcEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*ProxyArpIntfcEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ProxyArpIntfcEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *ProxyArpIntfcEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *ProxyArpIntfcEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_arp_binapi_init() }
func file_arp_binapi_init() {
	api.RegisterMessage((*ProxyArpAddDel)(nil), "proxy_arp_add_del_85486cbd")
	api.RegisterMessage((*ProxyArpAddDelReply)(nil), "proxy_arp_add_del_reply_e8d4e804")
	api.RegisterMessage((*ProxyArpDetails)(nil), "proxy_arp_details_9228c150")
	api.RegisterMessage((*ProxyArpDump)(nil), "proxy_arp_dump_51077d14")
	api.RegisterMessage((*ProxyArpIntfcDetails)(nil), "proxy_arp_intfc_details_f6458e5f")
	api.RegisterMessage((*ProxyArpIntfcDump)(nil), "proxy_arp_intfc_dump_51077d14")
	api.RegisterMessage((*ProxyArpIntfcEnableDisable)(nil), "proxy_arp_intfc_enable_disable_ae6cfcfb")
	api.RegisterMessage((*ProxyArpIntfcEnableDisableReply)(nil), "proxy_arp_intfc_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*ProxyArpAddDel)(nil),
		(*ProxyArpAddDelReply)(nil),
		(*ProxyArpDetails)(nil),
		(*ProxyArpDump)(nil),
		(*ProxyArpIntfcDetails)(nil),
		(*ProxyArpIntfcDump)(nil),
		(*ProxyArpIntfcEnableDisable)(nil),
		(*ProxyArpIntfcEnableDisableReply)(nil),
	}
}
