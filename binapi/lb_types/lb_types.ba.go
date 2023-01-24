// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              23.02-rc0~189-g57127b32a
// source: /usr/share/vpp/api/plugins/lb_types.api.json

// Package lb_types contains generated bindings for API file lb_types.api.
//
// Contents:
//   5 enums
//   1 struct
//
package lb_types

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

// LbEncapType defines enum 'lb_encap_type'.
type LbEncapType uint32

const (
	LB_API_ENCAP_TYPE_GRE4  LbEncapType = 0
	LB_API_ENCAP_TYPE_GRE6  LbEncapType = 1
	LB_API_ENCAP_TYPE_L3DSR LbEncapType = 2
	LB_API_ENCAP_TYPE_NAT4  LbEncapType = 3
	LB_API_ENCAP_TYPE_NAT6  LbEncapType = 4
	LB_API_ENCAP_N_TYPES    LbEncapType = 5
)

var (
	LbEncapType_name = map[uint32]string{
		0: "LB_API_ENCAP_TYPE_GRE4",
		1: "LB_API_ENCAP_TYPE_GRE6",
		2: "LB_API_ENCAP_TYPE_L3DSR",
		3: "LB_API_ENCAP_TYPE_NAT4",
		4: "LB_API_ENCAP_TYPE_NAT6",
		5: "LB_API_ENCAP_N_TYPES",
	}
	LbEncapType_value = map[string]uint32{
		"LB_API_ENCAP_TYPE_GRE4":  0,
		"LB_API_ENCAP_TYPE_GRE6":  1,
		"LB_API_ENCAP_TYPE_L3DSR": 2,
		"LB_API_ENCAP_TYPE_NAT4":  3,
		"LB_API_ENCAP_TYPE_NAT6":  4,
		"LB_API_ENCAP_N_TYPES":    5,
	}
)

func (x LbEncapType) String() string {
	s, ok := LbEncapType_name[uint32(x)]
	if ok {
		return s
	}
	return "LbEncapType(" + strconv.Itoa(int(x)) + ")"
}

// LbLkpTypeT defines enum 'lb_lkp_type_t'.
type LbLkpTypeT uint32

const (
	LB_API_LKP_SAME_IP_PORT LbLkpTypeT = 0
	LB_API_LKP_DIFF_IP_PORT LbLkpTypeT = 1
	LB_API_LKP_ALL_PORT_IP  LbLkpTypeT = 2
	LB_API_LKP_N_TYPES      LbLkpTypeT = 3
)

var (
	LbLkpTypeT_name = map[uint32]string{
		0: "LB_API_LKP_SAME_IP_PORT",
		1: "LB_API_LKP_DIFF_IP_PORT",
		2: "LB_API_LKP_ALL_PORT_IP",
		3: "LB_API_LKP_N_TYPES",
	}
	LbLkpTypeT_value = map[string]uint32{
		"LB_API_LKP_SAME_IP_PORT": 0,
		"LB_API_LKP_DIFF_IP_PORT": 1,
		"LB_API_LKP_ALL_PORT_IP":  2,
		"LB_API_LKP_N_TYPES":      3,
	}
)

func (x LbLkpTypeT) String() string {
	s, ok := LbLkpTypeT_name[uint32(x)]
	if ok {
		return s
	}
	return "LbLkpTypeT(" + strconv.Itoa(int(x)) + ")"
}

// LbNatProtocol defines enum 'lb_nat_protocol'.
type LbNatProtocol uint32

const (
	LB_API_NAT_PROTOCOL_UDP LbNatProtocol = 6
	LB_API_NAT_PROTOCOL_TCP LbNatProtocol = 23
	LB_API_NAT_PROTOCOL_ANY LbNatProtocol = 4294967295
)

var (
	LbNatProtocol_name = map[uint32]string{
		6:          "LB_API_NAT_PROTOCOL_UDP",
		23:         "LB_API_NAT_PROTOCOL_TCP",
		4294967295: "LB_API_NAT_PROTOCOL_ANY",
	}
	LbNatProtocol_value = map[string]uint32{
		"LB_API_NAT_PROTOCOL_UDP": 6,
		"LB_API_NAT_PROTOCOL_TCP": 23,
		"LB_API_NAT_PROTOCOL_ANY": 4294967295,
	}
)

func (x LbNatProtocol) String() string {
	s, ok := LbNatProtocol_name[uint32(x)]
	if ok {
		return s
	}
	return "LbNatProtocol(" + strconv.Itoa(int(x)) + ")"
}

// LbSrvType defines enum 'lb_srv_type'.
type LbSrvType uint32

const (
	LB_API_SRV_TYPE_CLUSTERIP LbSrvType = 0
	LB_API_SRV_TYPE_NODEPORT  LbSrvType = 1
	LB_API_SRV_N_TYPES        LbSrvType = 2
)

var (
	LbSrvType_name = map[uint32]string{
		0: "LB_API_SRV_TYPE_CLUSTERIP",
		1: "LB_API_SRV_TYPE_NODEPORT",
		2: "LB_API_SRV_N_TYPES",
	}
	LbSrvType_value = map[string]uint32{
		"LB_API_SRV_TYPE_CLUSTERIP": 0,
		"LB_API_SRV_TYPE_NODEPORT":  1,
		"LB_API_SRV_N_TYPES":        2,
	}
)

func (x LbSrvType) String() string {
	s, ok := LbSrvType_name[uint32(x)]
	if ok {
		return s
	}
	return "LbSrvType(" + strconv.Itoa(int(x)) + ")"
}

// LbVipType defines enum 'lb_vip_type'.
type LbVipType uint32

const (
	LB_API_VIP_TYPE_IP6_GRE6  LbVipType = 0
	LB_API_VIP_TYPE_IP6_GRE4  LbVipType = 1
	LB_API_VIP_TYPE_IP4_GRE6  LbVipType = 2
	LB_API_VIP_TYPE_IP4_GRE4  LbVipType = 3
	LB_API_VIP_TYPE_IP4_L3DSR LbVipType = 4
	LB_API_VIP_TYPE_IP4_NAT4  LbVipType = 5
	LB_API_VIP_TYPE_IP6_NAT6  LbVipType = 6
	LB_API_VIP_N_TYPES        LbVipType = 7
)

var (
	LbVipType_name = map[uint32]string{
		0: "LB_API_VIP_TYPE_IP6_GRE6",
		1: "LB_API_VIP_TYPE_IP6_GRE4",
		2: "LB_API_VIP_TYPE_IP4_GRE6",
		3: "LB_API_VIP_TYPE_IP4_GRE4",
		4: "LB_API_VIP_TYPE_IP4_L3DSR",
		5: "LB_API_VIP_TYPE_IP4_NAT4",
		6: "LB_API_VIP_TYPE_IP6_NAT6",
		7: "LB_API_VIP_N_TYPES",
	}
	LbVipType_value = map[string]uint32{
		"LB_API_VIP_TYPE_IP6_GRE6":  0,
		"LB_API_VIP_TYPE_IP6_GRE4":  1,
		"LB_API_VIP_TYPE_IP4_GRE6":  2,
		"LB_API_VIP_TYPE_IP4_GRE4":  3,
		"LB_API_VIP_TYPE_IP4_L3DSR": 4,
		"LB_API_VIP_TYPE_IP4_NAT4":  5,
		"LB_API_VIP_TYPE_IP6_NAT6":  6,
		"LB_API_VIP_N_TYPES":        7,
	}
)

func (x LbVipType) String() string {
	s, ok := LbVipType_name[uint32(x)]
	if ok {
		return s
	}
	return "LbVipType(" + strconv.Itoa(int(x)) + ")"
}

// LbVip defines type 'lb_vip'.
type LbVip struct {
	Pfx      ip_types.AddressWithPrefix `binapi:"address_with_prefix,name=pfx" json:"pfx,omitempty"`
	Protocol ip_types.IPProto           `binapi:"ip_proto,name=protocol" json:"protocol,omitempty"`
	Port     uint16                     `binapi:"u16,name=port" json:"port,omitempty"`
}
