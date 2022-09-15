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

var CikeGlobalActiveTunnels = Section{"cikeGlobalActiveTunnels", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 2, 1, 1, 0}, pdu.VariableTypeCounter32}

//var CikeGlobalOutOctets = Section{"cikeGlobalOutOctets", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 2, 1, 11, 0}, pdu.VariableTypeGauge32}

var CiKeTunnel = []Section{
	{Name: "cikeTunIndex", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.1"), Type: pdu.VariableTypeCounter32},
	{Name: "cikeTunLocalName", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.5"), Type: pdu.VariableTypeOctetString},
	{Name: "cikeTunLocalValue", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.3"), Type: pdu.VariableTypeOctetString},
	{Name: "cikeTunRemoteValue", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.7"), Type: pdu.VariableTypeOctetString},
	{Name: "cikeTunAuthMethod", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.14"), Type: pdu.VariableTypeOctetString},
	{Name: "cikeTunLifeTime", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.15"), Type: pdu.VariableTypeCounter32},
	{Name: "cikeTunActiveTime", Oid: value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1.16"), Type: pdu.VariableTypeCounter32},
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

var CipSecTunnelEntry = Section{Name: "cipSecTunnelEntry", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1}, Type: pdu.VariableTypeCounter32}

//var CipSecEndPtEntry = Section{"cipSecEndPtEntry", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 3, 1}}
var CipSecTunIkeTunnelIndex = Section{Name: "cipSecTunIkeTunnelIndex", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 2}, Type: pdu.VariableTypeCounter32}

//var CipSecGlobalOutOctets = Section{"cipSecGlobalOutOctets", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 1, 16, 0}}
//var CipSecGlobalInOctets = Section{"cipSecGlobalInOctets", value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 1, 3, 0}}

var CiPsecTunnel = []Section{
	{Name: "cipSecTunInOctets", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 26}, Type: pdu.VariableTypeCounter32},
	{Name: "cipSecTunHcInOctets", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 27}, Type: pdu.VariableTypeCounter32},
	{Name: "cipSecTunInOctWraps", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 28}},
	{Name: "cipSecTunInPkts", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 32}, Type: pdu.VariableTypeCounter32},
	{Name: "cipSecTunInDropPkts", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 33}},
	{Name: "cipSecTunOutOctets", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 39}},
	{Name: "cipSecTunHcOutOctets", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 40}},
	{Name: "cipSecTunOutOctWraps", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 41}},
	{Name: "cipSecTunOutDropPkts", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 46}, Type: pdu.VariableTypeCounter32},
	{Name: "cipSecTunStatus", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 2, 1, 51}, Type: pdu.VariableTypeOctetString},
}

var CiPsecEndPt = []Section{
	{Name: "cipSecEndPtLocalAddr1", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 3, 1, 4}},
	{Name: "cipSecEndPtLocalAddr2", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 3, 1, 5}},
	{Name: "cipSecEndPtRemoteAddr1", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 3, 1, 10}},
	{Name: "cipSecEndPtRemoteAddr2", Oid: value.OID{1, 3, 6, 1, 4, 1, 9, 9, 171, 1, 3, 3, 1, 11}},
}
