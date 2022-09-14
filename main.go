package main

import (
	"github.com/jackyczj/strongswan-snmp-agent/mib"
	"github.com/jackyczj/strongswan-snmp-agent/snmp"
	"github.com/jackyczj/strongswan-snmp-agent/vici"
	"github.com/posteo/go-agentx"
	"github.com/posteo/go-agentx/value"
	"time"
)

func main() {
	client := vici.Client{}
	client.Init()
	agentxSession, err := snmp.NewSession("unix", "/var/run/agentx.sock")
	if err != nil {
		panic(err)
	}
	defer agentxSession.Close()
	if err := agentxSession.Register(127, value.MustParseOID("1.3.6.1.4.1.9.9.171.1.2.3.1")); err != nil {
		panic(err)
	}

	listHandler := &agentx.ListHandler{}
	for {
		ike, _ := client.ListSA()
		CikeTunnelInject(&ike, listHandler)
		agentxSession.Handler = listHandler

		time.Sleep(3 * time.Second)
	}
}

func CikeTunnelInject(ike *vici.LoadedIKE, listHandler *agentx.ListHandler) {
	//mib.SortOid(mib.CiKeTunnel)
	for _, v := range mib.CiKeTunnel {
		var item = &agentx.ListItem{}
		switch v.Name {
		case "cikeTunLocalValue":
			item = listHandler.Add(v.Oid.String())
			item.Type = v.Type
			item.Value = ike.LocalId
		case "cikeTunRemoteValue":
			item = listHandler.Add(v.Oid.String())
			item.Type = v.Type
			item.Value = ike.RemoteId
		case "cikeTunActiveTime":
			item = listHandler.Add(v.Oid.String())
			item.Type = v.Type
			item.Value = uint32(ike.EstablishSec)
		case "cikeTunAuthMethod":
			item = listHandler.Add(v.Oid.String())
			item.Type = v.Type
			item.Value = ike.EncAlg
		case "cikeTunLifeTime":
			item = listHandler.Add(v.Oid.String())
			item.Type = v.Type
			item.Value = uint32(ike.ReauthSec)
		case "cikeTunInPkts":
			item = listHandler.Add(v.Oid.String())
			item.Type = v.Type
			var val uint32
			for _, v := range ike.Children {
				val += uint32(v.PacketsIn)
			}
			item.Value = val
		case "cikeTunOutPkts":
			item = listHandler.Add(v.Oid.String())
			item.Type = v.Type
			var val uint32
			for _, v := range ike.Children {
				val += uint32(v.PacketsOut)
			}
			item.Value = val
		}
	}
}
