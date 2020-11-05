// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.09-release
// source: /usr/share/vpp/api/plugins/nat66.api.json

// Package nat66 contains generated bindings for API file nat66.api.
//
// Contents:
//   8 messages
//
package nat66

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/edwarnicke/govpp/pkg/v2009/binapi/interface_types"
	ip_types "github.com/edwarnicke/govpp/pkg/v2009/binapi/ip_types"
	nat_types "github.com/edwarnicke/govpp/pkg/v2009/binapi/nat_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "nat66"
	APIVersion = "1.0.0"
	VersionCrc = 0x35ab7571
)

// Nat66AddDelInterface defines message 'nat66_add_del_interface'.
type Nat66AddDelInterface struct {
	IsAdd     bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Flags     nat_types.NatConfigFlags       `binapi:"nat_config_flags,name=flags" json:"flags,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *Nat66AddDelInterface) Reset()               { *m = Nat66AddDelInterface{} }
func (*Nat66AddDelInterface) GetMessageName() string { return "nat66_add_del_interface" }
func (*Nat66AddDelInterface) GetCrcString() string   { return "f3699b83" }
func (*Nat66AddDelInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *Nat66AddDelInterface) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 1 // m.Flags
	size += 4 // m.SwIfIndex
	return size
}
func (m *Nat66AddDelInterface) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint8(uint8(m.Flags))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *Nat66AddDelInterface) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Flags = nat_types.NatConfigFlags(buf.DecodeUint8())
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// Nat66AddDelInterfaceReply defines message 'nat66_add_del_interface_reply'.
type Nat66AddDelInterfaceReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *Nat66AddDelInterfaceReply) Reset()               { *m = Nat66AddDelInterfaceReply{} }
func (*Nat66AddDelInterfaceReply) GetMessageName() string { return "nat66_add_del_interface_reply" }
func (*Nat66AddDelInterfaceReply) GetCrcString() string   { return "e8d4e804" }
func (*Nat66AddDelInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *Nat66AddDelInterfaceReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *Nat66AddDelInterfaceReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *Nat66AddDelInterfaceReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Nat66AddDelStaticMapping defines message 'nat66_add_del_static_mapping'.
type Nat66AddDelStaticMapping struct {
	IsAdd             bool                `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	LocalIPAddress    ip_types.IP6Address `binapi:"ip6_address,name=local_ip_address" json:"local_ip_address,omitempty"`
	ExternalIPAddress ip_types.IP6Address `binapi:"ip6_address,name=external_ip_address" json:"external_ip_address,omitempty"`
	VrfID             uint32              `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
}

func (m *Nat66AddDelStaticMapping) Reset()               { *m = Nat66AddDelStaticMapping{} }
func (*Nat66AddDelStaticMapping) GetMessageName() string { return "nat66_add_del_static_mapping" }
func (*Nat66AddDelStaticMapping) GetCrcString() string   { return "fb64e50b" }
func (*Nat66AddDelStaticMapping) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *Nat66AddDelStaticMapping) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsAdd
	size += 1 * 16 // m.LocalIPAddress
	size += 1 * 16 // m.ExternalIPAddress
	size += 4      // m.VrfID
	return size
}
func (m *Nat66AddDelStaticMapping) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeBytes(m.LocalIPAddress[:], 16)
	buf.EncodeBytes(m.ExternalIPAddress[:], 16)
	buf.EncodeUint32(m.VrfID)
	return buf.Bytes(), nil
}
func (m *Nat66AddDelStaticMapping) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	copy(m.LocalIPAddress[:], buf.DecodeBytes(16))
	copy(m.ExternalIPAddress[:], buf.DecodeBytes(16))
	m.VrfID = buf.DecodeUint32()
	return nil
}

// Nat66AddDelStaticMappingReply defines message 'nat66_add_del_static_mapping_reply'.
type Nat66AddDelStaticMappingReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *Nat66AddDelStaticMappingReply) Reset() { *m = Nat66AddDelStaticMappingReply{} }
func (*Nat66AddDelStaticMappingReply) GetMessageName() string {
	return "nat66_add_del_static_mapping_reply"
}
func (*Nat66AddDelStaticMappingReply) GetCrcString() string { return "e8d4e804" }
func (*Nat66AddDelStaticMappingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *Nat66AddDelStaticMappingReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *Nat66AddDelStaticMappingReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *Nat66AddDelStaticMappingReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Nat66InterfaceDetails defines message 'nat66_interface_details'.
type Nat66InterfaceDetails struct {
	Flags     nat_types.NatConfigFlags       `binapi:"nat_config_flags,name=flags" json:"flags,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *Nat66InterfaceDetails) Reset()               { *m = Nat66InterfaceDetails{} }
func (*Nat66InterfaceDetails) GetMessageName() string { return "nat66_interface_details" }
func (*Nat66InterfaceDetails) GetCrcString() string   { return "5d286289" }
func (*Nat66InterfaceDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *Nat66InterfaceDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Flags
	size += 4 // m.SwIfIndex
	return size
}
func (m *Nat66InterfaceDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Flags))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *Nat66InterfaceDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Flags = nat_types.NatConfigFlags(buf.DecodeUint8())
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// Nat66InterfaceDump defines message 'nat66_interface_dump'.
type Nat66InterfaceDump struct{}

func (m *Nat66InterfaceDump) Reset()               { *m = Nat66InterfaceDump{} }
func (*Nat66InterfaceDump) GetMessageName() string { return "nat66_interface_dump" }
func (*Nat66InterfaceDump) GetCrcString() string   { return "51077d14" }
func (*Nat66InterfaceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *Nat66InterfaceDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *Nat66InterfaceDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *Nat66InterfaceDump) Unmarshal(b []byte) error {
	return nil
}

// Nat66StaticMappingDetails defines message 'nat66_static_mapping_details'.
type Nat66StaticMappingDetails struct {
	LocalIPAddress    ip_types.IP6Address `binapi:"ip6_address,name=local_ip_address" json:"local_ip_address,omitempty"`
	ExternalIPAddress ip_types.IP6Address `binapi:"ip6_address,name=external_ip_address" json:"external_ip_address,omitempty"`
	VrfID             uint32              `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	TotalBytes        uint64              `binapi:"u64,name=total_bytes" json:"total_bytes,omitempty"`
	TotalPkts         uint64              `binapi:"u64,name=total_pkts" json:"total_pkts,omitempty"`
}

func (m *Nat66StaticMappingDetails) Reset()               { *m = Nat66StaticMappingDetails{} }
func (*Nat66StaticMappingDetails) GetMessageName() string { return "nat66_static_mapping_details" }
func (*Nat66StaticMappingDetails) GetCrcString() string   { return "5c568448" }
func (*Nat66StaticMappingDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *Nat66StaticMappingDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.LocalIPAddress
	size += 1 * 16 // m.ExternalIPAddress
	size += 4      // m.VrfID
	size += 8      // m.TotalBytes
	size += 8      // m.TotalPkts
	return size
}
func (m *Nat66StaticMappingDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.LocalIPAddress[:], 16)
	buf.EncodeBytes(m.ExternalIPAddress[:], 16)
	buf.EncodeUint32(m.VrfID)
	buf.EncodeUint64(m.TotalBytes)
	buf.EncodeUint64(m.TotalPkts)
	return buf.Bytes(), nil
}
func (m *Nat66StaticMappingDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.LocalIPAddress[:], buf.DecodeBytes(16))
	copy(m.ExternalIPAddress[:], buf.DecodeBytes(16))
	m.VrfID = buf.DecodeUint32()
	m.TotalBytes = buf.DecodeUint64()
	m.TotalPkts = buf.DecodeUint64()
	return nil
}

// Nat66StaticMappingDump defines message 'nat66_static_mapping_dump'.
type Nat66StaticMappingDump struct{}

func (m *Nat66StaticMappingDump) Reset()               { *m = Nat66StaticMappingDump{} }
func (*Nat66StaticMappingDump) GetMessageName() string { return "nat66_static_mapping_dump" }
func (*Nat66StaticMappingDump) GetCrcString() string   { return "51077d14" }
func (*Nat66StaticMappingDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *Nat66StaticMappingDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *Nat66StaticMappingDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *Nat66StaticMappingDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_nat66_binapi_init() }
func file_nat66_binapi_init() {
	api.RegisterMessage((*Nat66AddDelInterface)(nil), "nat66_add_del_interface_f3699b83")
	api.RegisterMessage((*Nat66AddDelInterfaceReply)(nil), "nat66_add_del_interface_reply_e8d4e804")
	api.RegisterMessage((*Nat66AddDelStaticMapping)(nil), "nat66_add_del_static_mapping_fb64e50b")
	api.RegisterMessage((*Nat66AddDelStaticMappingReply)(nil), "nat66_add_del_static_mapping_reply_e8d4e804")
	api.RegisterMessage((*Nat66InterfaceDetails)(nil), "nat66_interface_details_5d286289")
	api.RegisterMessage((*Nat66InterfaceDump)(nil), "nat66_interface_dump_51077d14")
	api.RegisterMessage((*Nat66StaticMappingDetails)(nil), "nat66_static_mapping_details_5c568448")
	api.RegisterMessage((*Nat66StaticMappingDump)(nil), "nat66_static_mapping_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*Nat66AddDelInterface)(nil),
		(*Nat66AddDelInterfaceReply)(nil),
		(*Nat66AddDelStaticMapping)(nil),
		(*Nat66AddDelStaticMappingReply)(nil),
		(*Nat66InterfaceDetails)(nil),
		(*Nat66InterfaceDump)(nil),
		(*Nat66StaticMappingDetails)(nil),
		(*Nat66StaticMappingDump)(nil),
	}
}
