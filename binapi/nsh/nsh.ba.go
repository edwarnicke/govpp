// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              22.06-rc0~147-ge81e78f81
// source: /usr/share/vpp/api/plugins/nsh.api.json

// Package nsh contains generated bindings for API file nsh.api.
//
// Contents:
//   8 messages
//
package nsh

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "nsh"
	APIVersion = "1.0.0"
	VersionCrc = 0xc2f3127d
)

// NshAddDelEntry defines message 'nsh_add_del_entry'.
type NshAddDelEntry struct {
	IsAdd        bool   `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	NspNsi       uint32 `binapi:"u32,name=nsp_nsi" json:"nsp_nsi,omitempty"`
	MdType       uint8  `binapi:"u8,name=md_type" json:"md_type,omitempty"`
	VerOC        uint8  `binapi:"u8,name=ver_o_c" json:"ver_o_c,omitempty"`
	TTL          uint8  `binapi:"u8,name=ttl" json:"ttl,omitempty"`
	Length       uint8  `binapi:"u8,name=length" json:"length,omitempty"`
	NextProtocol uint8  `binapi:"u8,name=next_protocol" json:"next_protocol,omitempty"`
	C1           uint32 `binapi:"u32,name=c1" json:"c1,omitempty"`
	C2           uint32 `binapi:"u32,name=c2" json:"c2,omitempty"`
	C3           uint32 `binapi:"u32,name=c3" json:"c3,omitempty"`
	C4           uint32 `binapi:"u32,name=c4" json:"c4,omitempty"`
	TlvLength    uint8  `binapi:"u8,name=tlv_length" json:"tlv_length,omitempty"`
	Tlv          []byte `binapi:"u8[248],name=tlv" json:"tlv,omitempty"`
}

func (m *NshAddDelEntry) Reset()               { *m = NshAddDelEntry{} }
func (*NshAddDelEntry) GetMessageName() string { return "nsh_add_del_entry" }
func (*NshAddDelEntry) GetCrcString() string   { return "7dea480b" }
func (*NshAddDelEntry) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *NshAddDelEntry) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1       // m.IsAdd
	size += 4       // m.NspNsi
	size += 1       // m.MdType
	size += 1       // m.VerOC
	size += 1       // m.TTL
	size += 1       // m.Length
	size += 1       // m.NextProtocol
	size += 4       // m.C1
	size += 4       // m.C2
	size += 4       // m.C3
	size += 4       // m.C4
	size += 1       // m.TlvLength
	size += 1 * 248 // m.Tlv
	return size
}
func (m *NshAddDelEntry) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.NspNsi)
	buf.EncodeUint8(m.MdType)
	buf.EncodeUint8(m.VerOC)
	buf.EncodeUint8(m.TTL)
	buf.EncodeUint8(m.Length)
	buf.EncodeUint8(m.NextProtocol)
	buf.EncodeUint32(m.C1)
	buf.EncodeUint32(m.C2)
	buf.EncodeUint32(m.C3)
	buf.EncodeUint32(m.C4)
	buf.EncodeUint8(m.TlvLength)
	buf.EncodeBytes(m.Tlv, 248)
	return buf.Bytes(), nil
}
func (m *NshAddDelEntry) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.NspNsi = buf.DecodeUint32()
	m.MdType = buf.DecodeUint8()
	m.VerOC = buf.DecodeUint8()
	m.TTL = buf.DecodeUint8()
	m.Length = buf.DecodeUint8()
	m.NextProtocol = buf.DecodeUint8()
	m.C1 = buf.DecodeUint32()
	m.C2 = buf.DecodeUint32()
	m.C3 = buf.DecodeUint32()
	m.C4 = buf.DecodeUint32()
	m.TlvLength = buf.DecodeUint8()
	m.Tlv = make([]byte, 248)
	copy(m.Tlv, buf.DecodeBytes(len(m.Tlv)))
	return nil
}

// NshAddDelEntryReply defines message 'nsh_add_del_entry_reply'.
type NshAddDelEntryReply struct {
	Retval     int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	EntryIndex uint32 `binapi:"u32,name=entry_index" json:"entry_index,omitempty"`
}

func (m *NshAddDelEntryReply) Reset()               { *m = NshAddDelEntryReply{} }
func (*NshAddDelEntryReply) GetMessageName() string { return "nsh_add_del_entry_reply" }
func (*NshAddDelEntryReply) GetCrcString() string   { return "6296a9eb" }
func (*NshAddDelEntryReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *NshAddDelEntryReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.EntryIndex
	return size
}
func (m *NshAddDelEntryReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.EntryIndex)
	return buf.Bytes(), nil
}
func (m *NshAddDelEntryReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.EntryIndex = buf.DecodeUint32()
	return nil
}

// NshAddDelMap defines message 'nsh_add_del_map'.
type NshAddDelMap struct {
	IsAdd        bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	NspNsi       uint32                         `binapi:"u32,name=nsp_nsi" json:"nsp_nsi,omitempty"`
	MappedNspNsi uint32                         `binapi:"u32,name=mapped_nsp_nsi" json:"mapped_nsp_nsi,omitempty"`
	NshAction    uint32                         `binapi:"u32,name=nsh_action" json:"nsh_action,omitempty"`
	SwIfIndex    interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	RxSwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=rx_sw_if_index" json:"rx_sw_if_index,omitempty"`
	NextNode     uint32                         `binapi:"u32,name=next_node" json:"next_node,omitempty"`
}

func (m *NshAddDelMap) Reset()               { *m = NshAddDelMap{} }
func (*NshAddDelMap) GetMessageName() string { return "nsh_add_del_map" }
func (*NshAddDelMap) GetCrcString() string   { return "0a0f42b0" }
func (*NshAddDelMap) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *NshAddDelMap) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 4 // m.NspNsi
	size += 4 // m.MappedNspNsi
	size += 4 // m.NshAction
	size += 4 // m.SwIfIndex
	size += 4 // m.RxSwIfIndex
	size += 4 // m.NextNode
	return size
}
func (m *NshAddDelMap) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.NspNsi)
	buf.EncodeUint32(m.MappedNspNsi)
	buf.EncodeUint32(m.NshAction)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(uint32(m.RxSwIfIndex))
	buf.EncodeUint32(m.NextNode)
	return buf.Bytes(), nil
}
func (m *NshAddDelMap) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.NspNsi = buf.DecodeUint32()
	m.MappedNspNsi = buf.DecodeUint32()
	m.NshAction = buf.DecodeUint32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.RxSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.NextNode = buf.DecodeUint32()
	return nil
}

// NshAddDelMapReply defines message 'nsh_add_del_map_reply'.
type NshAddDelMapReply struct {
	Retval   int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	MapIndex uint32 `binapi:"u32,name=map_index" json:"map_index,omitempty"`
}

func (m *NshAddDelMapReply) Reset()               { *m = NshAddDelMapReply{} }
func (*NshAddDelMapReply) GetMessageName() string { return "nsh_add_del_map_reply" }
func (*NshAddDelMapReply) GetCrcString() string   { return "b2b127ef" }
func (*NshAddDelMapReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *NshAddDelMapReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.MapIndex
	return size
}
func (m *NshAddDelMapReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.MapIndex)
	return buf.Bytes(), nil
}
func (m *NshAddDelMapReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.MapIndex = buf.DecodeUint32()
	return nil
}

// NshEntryDetails defines message 'nsh_entry_details'.
type NshEntryDetails struct {
	EntryIndex   uint32 `binapi:"u32,name=entry_index" json:"entry_index,omitempty"`
	NspNsi       uint32 `binapi:"u32,name=nsp_nsi" json:"nsp_nsi,omitempty"`
	MdType       uint8  `binapi:"u8,name=md_type" json:"md_type,omitempty"`
	VerOC        uint8  `binapi:"u8,name=ver_o_c" json:"ver_o_c,omitempty"`
	TTL          uint8  `binapi:"u8,name=ttl" json:"ttl,omitempty"`
	Length       uint8  `binapi:"u8,name=length" json:"length,omitempty"`
	NextProtocol uint8  `binapi:"u8,name=next_protocol" json:"next_protocol,omitempty"`
	C1           uint32 `binapi:"u32,name=c1" json:"c1,omitempty"`
	C2           uint32 `binapi:"u32,name=c2" json:"c2,omitempty"`
	C3           uint32 `binapi:"u32,name=c3" json:"c3,omitempty"`
	C4           uint32 `binapi:"u32,name=c4" json:"c4,omitempty"`
	TlvLength    uint8  `binapi:"u8,name=tlv_length" json:"tlv_length,omitempty"`
	Tlv          []byte `binapi:"u8[248],name=tlv" json:"tlv,omitempty"`
}

func (m *NshEntryDetails) Reset()               { *m = NshEntryDetails{} }
func (*NshEntryDetails) GetMessageName() string { return "nsh_entry_details" }
func (*NshEntryDetails) GetCrcString() string   { return "046fb556" }
func (*NshEntryDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *NshEntryDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4       // m.EntryIndex
	size += 4       // m.NspNsi
	size += 1       // m.MdType
	size += 1       // m.VerOC
	size += 1       // m.TTL
	size += 1       // m.Length
	size += 1       // m.NextProtocol
	size += 4       // m.C1
	size += 4       // m.C2
	size += 4       // m.C3
	size += 4       // m.C4
	size += 1       // m.TlvLength
	size += 1 * 248 // m.Tlv
	return size
}
func (m *NshEntryDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.EntryIndex)
	buf.EncodeUint32(m.NspNsi)
	buf.EncodeUint8(m.MdType)
	buf.EncodeUint8(m.VerOC)
	buf.EncodeUint8(m.TTL)
	buf.EncodeUint8(m.Length)
	buf.EncodeUint8(m.NextProtocol)
	buf.EncodeUint32(m.C1)
	buf.EncodeUint32(m.C2)
	buf.EncodeUint32(m.C3)
	buf.EncodeUint32(m.C4)
	buf.EncodeUint8(m.TlvLength)
	buf.EncodeBytes(m.Tlv, 248)
	return buf.Bytes(), nil
}
func (m *NshEntryDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EntryIndex = buf.DecodeUint32()
	m.NspNsi = buf.DecodeUint32()
	m.MdType = buf.DecodeUint8()
	m.VerOC = buf.DecodeUint8()
	m.TTL = buf.DecodeUint8()
	m.Length = buf.DecodeUint8()
	m.NextProtocol = buf.DecodeUint8()
	m.C1 = buf.DecodeUint32()
	m.C2 = buf.DecodeUint32()
	m.C3 = buf.DecodeUint32()
	m.C4 = buf.DecodeUint32()
	m.TlvLength = buf.DecodeUint8()
	m.Tlv = make([]byte, 248)
	copy(m.Tlv, buf.DecodeBytes(len(m.Tlv)))
	return nil
}

// NshEntryDump defines message 'nsh_entry_dump'.
type NshEntryDump struct {
	EntryIndex uint32 `binapi:"u32,name=entry_index" json:"entry_index,omitempty"`
}

func (m *NshEntryDump) Reset()               { *m = NshEntryDump{} }
func (*NshEntryDump) GetMessageName() string { return "nsh_entry_dump" }
func (*NshEntryDump) GetCrcString() string   { return "cdaf8ccb" }
func (*NshEntryDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *NshEntryDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.EntryIndex
	return size
}
func (m *NshEntryDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.EntryIndex)
	return buf.Bytes(), nil
}
func (m *NshEntryDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EntryIndex = buf.DecodeUint32()
	return nil
}

// NshMapDetails defines message 'nsh_map_details'.
type NshMapDetails struct {
	MapIndex     uint32                         `binapi:"u32,name=map_index" json:"map_index,omitempty"`
	NspNsi       uint32                         `binapi:"u32,name=nsp_nsi" json:"nsp_nsi,omitempty"`
	MappedNspNsi uint32                         `binapi:"u32,name=mapped_nsp_nsi" json:"mapped_nsp_nsi,omitempty"`
	NshAction    uint32                         `binapi:"u32,name=nsh_action" json:"nsh_action,omitempty"`
	SwIfIndex    interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	RxSwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=rx_sw_if_index" json:"rx_sw_if_index,omitempty"`
	NextNode     uint32                         `binapi:"u32,name=next_node" json:"next_node,omitempty"`
}

func (m *NshMapDetails) Reset()               { *m = NshMapDetails{} }
func (*NshMapDetails) GetMessageName() string { return "nsh_map_details" }
func (*NshMapDetails) GetCrcString() string   { return "2fefcf49" }
func (*NshMapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *NshMapDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.MapIndex
	size += 4 // m.NspNsi
	size += 4 // m.MappedNspNsi
	size += 4 // m.NshAction
	size += 4 // m.SwIfIndex
	size += 4 // m.RxSwIfIndex
	size += 4 // m.NextNode
	return size
}
func (m *NshMapDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.MapIndex)
	buf.EncodeUint32(m.NspNsi)
	buf.EncodeUint32(m.MappedNspNsi)
	buf.EncodeUint32(m.NshAction)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(uint32(m.RxSwIfIndex))
	buf.EncodeUint32(m.NextNode)
	return buf.Bytes(), nil
}
func (m *NshMapDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MapIndex = buf.DecodeUint32()
	m.NspNsi = buf.DecodeUint32()
	m.MappedNspNsi = buf.DecodeUint32()
	m.NshAction = buf.DecodeUint32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.RxSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.NextNode = buf.DecodeUint32()
	return nil
}

// NshMapDump defines message 'nsh_map_dump'.
type NshMapDump struct {
	MapIndex uint32 `binapi:"u32,name=map_index" json:"map_index,omitempty"`
}

func (m *NshMapDump) Reset()               { *m = NshMapDump{} }
func (*NshMapDump) GetMessageName() string { return "nsh_map_dump" }
func (*NshMapDump) GetCrcString() string   { return "8fc06b82" }
func (*NshMapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *NshMapDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.MapIndex
	return size
}
func (m *NshMapDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.MapIndex)
	return buf.Bytes(), nil
}
func (m *NshMapDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MapIndex = buf.DecodeUint32()
	return nil
}

func init() { file_nsh_binapi_init() }
func file_nsh_binapi_init() {
	api.RegisterMessage((*NshAddDelEntry)(nil), "nsh_add_del_entry_7dea480b")
	api.RegisterMessage((*NshAddDelEntryReply)(nil), "nsh_add_del_entry_reply_6296a9eb")
	api.RegisterMessage((*NshAddDelMap)(nil), "nsh_add_del_map_0a0f42b0")
	api.RegisterMessage((*NshAddDelMapReply)(nil), "nsh_add_del_map_reply_b2b127ef")
	api.RegisterMessage((*NshEntryDetails)(nil), "nsh_entry_details_046fb556")
	api.RegisterMessage((*NshEntryDump)(nil), "nsh_entry_dump_cdaf8ccb")
	api.RegisterMessage((*NshMapDetails)(nil), "nsh_map_details_2fefcf49")
	api.RegisterMessage((*NshMapDump)(nil), "nsh_map_dump_8fc06b82")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*NshAddDelEntry)(nil),
		(*NshAddDelEntryReply)(nil),
		(*NshAddDelMap)(nil),
		(*NshAddDelMapReply)(nil),
		(*NshEntryDetails)(nil),
		(*NshEntryDump)(nil),
		(*NshMapDetails)(nil),
		(*NshMapDump)(nil),
	}
}
