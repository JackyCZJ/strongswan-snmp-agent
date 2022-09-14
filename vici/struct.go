package vici

type LoadedIKE struct {
	Name         string
	UniqueId     string                 `vici:"uniqueid"`
	Version      int                    `vici:"version"`
	State        string                 `vici:"state"`
	LocalHost    string                 `vici:"local-host"`
	LocalPort    int                    `vici:"local-port"`
	LocalId      string                 `vici:"local-id"`
	RemoteHost   string                 `vici:"remote-host"`
	RemotePort   int                    `vici:"remote-port"`
	RemoteId     string                 `vici:"remote-id"`
	Initiator    string                 `vici:"initiator"`
	NatRemote    string                 `vici:"nat-remote"`
	NatFake      string                 `vici:"nat-fake"`
	NatAny       string                 `vici:"nat-any"`
	EncAlg       string                 `vici:"encr-alg"`
	EncKey       int                    `vici:"encr-keysize"`
	IntegAlg     string                 `vici:"integ-alg"`
	IntegKey     int                    `vici:"integ-keysize"`
	DHGroup      string                 `vici:"dh-group"`
	EstablishSec int64                  `vici:"established"`
	RekeySec     int64                  `vici:"rekey-time"`
	ReauthSec    int64                  `vici:"reauth-time"`
	LocalVips    []string               `vici:"local-vips"`
	RemoteVips   []string               `vici:"remote-vips"`
	Children     map[string]LoadedChild `vici:"child-sas"`
}
type LoadedChild struct {
	Name         string   `vici:"name"`
	UniqueId     string   `vici:"uniqueid"`
	State        string   `vici:"state"`
	Mode         string   `vici:"mode"`
	Protocol     string   `vici:"protocol"`
	Encap        string   `vici:"encap"`
	EncAlg       string   `vici:"encr-alg"`
	EncKey       int      `vici:"encr-keysize"`
	IntegAlg     string   `vici:"integ-alg"`
	IntegKey     int      `vici:"integ-keysize"`
	DHGroup      string   `vici:"dh-group"`
	BytesIn      int64    `vici:"bytes-in"`
	PacketsIn    int64    `vici:"bytes-out"`
	LastInSec    int64    `vici:"use-in"`
	BytesOut     int64    `vici:"bytes-out"`
	PacketsOut   int64    `vici:"bytes-out"`
	LastOutSec   int64    `vici:"use-out"`
	EstablishSec int64    `vici:"install-time"`
	RekeySec     int64    `vici:"rekey-time"`
	LifetimeSec  int64    `vici:"life-time"`
	LocalTS      []string `vici:"local-ts"`
	RemoteTS     []string `vici:"remote-ts"`
}

type ListConnection struct {
	Name        string
	LocalAddrs  []string             `vici:"local_addrs"`
	RemoteAddrs []string             `vici:"remote_addrs"`
	Version     string               `vici:"version"`
	ReauthTime  int64                `vici:"reauth_time"`
	RekeyTime   int64                `vici:"rekey_time"`
	LocalAuths  map[string]ConnAuth  `vici:"local_auth"`
	RemoteAuths map[string]ConnAuth  `vici:"remote_auth"`
	Children    map[string]ConnChild `vici:"children"`
}

type ConnAuth struct {
	Class      string   `vici:"class"`
	EapType    string   `vici:"eap_type"`
	EapVendor  string   `vici:"eap_vendor"`
	Xauth      string   `vici:"xauth"`
	Revocation string   `vici:"revocation"`
	Id         string   `vici:"id"`
	AaaId      string   `vici:"aaa_id"`
	EapId      string   `vici:"eap_id"`
	XauthId    string   `vici:"xauth_id"`
	Groups     []string `vici:"groups"`
	Certs      []string `vici:"certs"`
	Cacerts    []string `vici:"cacerts"`
}

type ConnChild struct {
	Mode       string   `vici:"mode"`
	Label      string   `vici:"label"`
	RekeyTime  int64    `vici:"rekey_time"`
	RekeyBytes int64    `vici:"rekey_bytes"`
	RekeyPkt   int64    `vici:"rekey_packets"`
	LocalTS    []string `vici:"local_ts"`
	RemoteTS   []string `vici:"remote_ts"`
}
