package vici

import (
	"fmt"
	"github.com/strongswan/govici/vici"
	"strings"
)

type Client struct {
	Network string
	Address string
	*vici.Session
}

func (c *Client) Init(opt ...vici.SessionOption) {
	session, err := vici.NewSession(opt...)
	if err != nil {
		fmt.Println(err)
		return
	}
	session.Subscribe()
	c.Session = session
}

func (c *Client) ListSA() (LoadedIKE, error) {
	msgs, err := c.StreamedCommandRequest("list-sas", "list-sa", nil)
	for _, m := range msgs.Messages() {
		if e := m.Err(); e != nil {
			continue
		}
		for _, k := range m.Keys() {
			inbound := m.Get(k).(*vici.Message)
			var ike LoadedIKE
			if e := vici.UnmarshalMessage(inbound, &ike); e != nil {
				continue
			}
			return ike, nil
		}
	}
	return LoadedIKE{}, err
}

func (c *Client) ListConn() (ListConnection, error) {
	msgs, err := c.StreamedCommandRequest("list-conns", "list-conn", nil)
	for _, m := range msgs.Messages() {

		if e := m.Err(); e != nil {
			continue
		}
		for _, k := range m.Keys() {
			inbound := m.Get(k).(*vici.Message)
			var conn ListConnection
			if e := vici.UnmarshalMessage(inbound, &conn); e != nil {
				fmt.Println(e)
				continue
			}
			for _, k := range inbound.Keys() {
				if strings.HasPrefix(k, "local-") {
					a := AuthConf{}
					if e := vici.UnmarshalMessage(inbound.Get(k).(*vici.Message), &a); e != nil {
						fmt.Println(e)
						continue
					}
					a.Name = k
					conn.LocalAuths = append(conn.LocalAuths, a)
				}
				if strings.HasPrefix(k, "remote-") {
					a := AuthConf{}
					if e := vici.UnmarshalMessage(inbound.Get(k).(*vici.Message), &a); e != nil {
						fmt.Println(e)
						continue
					}
					a.Name = k
					conn.RemoteAuths = append(conn.RemoteAuths, a)
				}
			}
			return conn, nil
		}
	}
	return ListConnection{}, err
}
