package mib

import (
	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
	"sort"
)

type Section struct {
	Name string
	Oid  value.OID
	Type pdu.VariableType
}

//var CikeGlobalActiveTunnels = Section{"cikeGlobalActiveTunnels", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 2, 1, 1, 0}, pdu.VariableTypeGauge32}
//var CikeGlobalOutOctets = Section{"cikeGlobalOutOctets", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 2, 1, 11, 0}, pdu.VariableTypeGauge32}

//1.3.6.1.4.1.9.9.171.1.2.3.1.19	cikeTunInOctets	0	0	The total number of octets received by
//this IPsec Phase-1 IKE Tunnel.
//1.3.6.1.4.1.9.9.171.1.2.3.1.20	cikeTunInPkts	0	0	The total number of packets received by
//this IPsec Phase-1 IKE Tunnel.
//1.3.6.1.4.1.9.9.171.1.2.3.1.21	cikeTunInDropPkts	0	0	The total number of packets dropped
//by this IPsec Phase-1 IKE Tunnel during
//receive processing.

var CiKeTunnel = []Section{
	{Name: "cikeTunLocalValue", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.3"), Type: pdu.VariableTypeOctetString},
	{Name: "cikeTunRemoteValue", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.7"), Type: pdu.VariableTypeOctetString},
	{Name: "cikeTunAuthMethod", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.14"), Type: pdu.VariableTypeOctetString},
	{Name: "cikeTunLifeTime", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.15"), Type: pdu.VariableTypeCounter32},
	{Name: "cikeTunActiveTime", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 2, 3, 1, 16}, Type: pdu.VariableTypeCounter32},
	//{Name: "cikeTunInOctets", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.19"), Type: pdu.VariableTypeCounter32},
	{Name: "cikeTunInPkts", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.20"), Type: pdu.VariableTypeCounter32},
	//{Name: "cikeTunInDropPkts", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.21"), Type: pdu.VariableTypeCounter32},
	{Name: "cikeTunOutPkts", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.28"), Type: pdu.VariableTypeCounter32},
}

func SortOid(s []Section) {
	sort.Slice(s, func(i, j int) bool {
		return value.CompareOIDs(s[i].Oid, s[j].Oid) == -1
	})
}

//var CipSecTunnelEntry = Section{"cipSecTunnelEntry", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 2, 3, 1}}
//var CipSecEndPtEntry = Section{"cipSecEndPtEntry", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 3, 1}}
//var CipSecTunIkeTunnelIndex = Section{"cipSecTunIkeTunnelIndex", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 2}}
//var CipSecGlobalOutOctets = Section{"cipSecGlobalOutOctets", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 1, 16, 0}}
//var CipSecGlobalInOctets = Section{"cipSecGlobalInOctets", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 1, 3, 0}}

var CiPsecTunnel = map[string]Section{
	"cipSecTunInOctets":    {Name: "cipSecTunInOctets", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 26}},
	"cipSecTunHcInOctets":  {Name: "cipSecTunHcInOctets", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 27}},
	"cipSecTunInOctWraps":  {Name: "cipSecTunInOctWraps", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 28}},
	"cipSecTunInDropPkts":  {Name: "cipSecTunInDropPkts", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 33}},
	"cipSecTunOutOctets":   {Name: "cipSecTunOutOctets", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 39}},
	"cipSecTunHcOutOctets": {Name: "cipSecTunHcOutOctets", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 40}},
	"cipSecTunOutOctWraps": {Name: "cipSecTunOutOctWraps", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 41}},
	"cipSecTunOutDropPkts": {Name: "cipSecTunOutDropPkts", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 46}},
}

var CiPsecEndPt = map[string]Section{
	"cipSecEndPtLocalAddr1":  Section{Name: "cipSecEndPtLocalAddr1", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 3, 1, 4}},
	"cipSecEndPtLocalAddr2":  Section{Name: "cipSecEndPtLocalAddr2", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 3, 1, 5}},
	"cipSecEndPtRemoteAddr1": Section{Name: "cipSecEndPtRemoteAddr1", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 3, 1, 10}},
	"cipSecEndPtRemoteAddr2": Section{Name: "cipSecEndPtRemoteAddr2", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 3, 1, 11}},
}
