package vici

type Connection struct {
	Name string // This field will NOT be marshaled!

	LocalAddrs  []string           `vici:"local_addrs"`
	RemoteAddrs []string           `vici:"remote_addrs"`
	Local       localOpts          `vici:"local"`
	Remote      remoteOpts         `vici:"remote"`
	Children    map[string]childSA `vici:"children"`
	Version     int                `vici:"version"`
	Proposals   []string           `vici:"proposals"`
	Vips        []string           `vici:"vips"`
	LocalPort   uint32             `vici:"local_port"`
	RemotePort  uint32             `vici:"remote_port"`
	Dscp        uint64             `vici:"dscp"`
	Encap       string             `vici:"encap"`
	Mobike      string             `vici:"mobike"`
	DpdDelay    uint32             `vici:"dpd_delay"`
	DpdTimeout  uint32             `vici:"dpd_timeout"`
	ReauthTime  uint32             `vici:"reauth_time"`
	RekeyTime   string             `vici:"rekey_time"`
	Pools       []string           `vici:"pools"`
}
type localOpts struct {
	Auth    string   `vici:"auth"`
	Certs   []string `vici:"certs"`
	Id      string   `vici:"id"`
	EapId   string   `vici:"eap_id"`
	AaaId   string   `vici:"aaa_id"`
	XauthId string   `vici:"xauth_id"`
	PubKeys []string `vici:"pubkeys"`
}

type remoteOpts struct {
	Auth       string   `vici:"auth"`
	Id         string   `vici:"id"`
	EapId      string   `vici:"eap_id"`
	Groups     []string `vici:"groups"`
	CertPolicy []string `vici:"cert_policy"`
	Certs      []string `vici:"certs"`
	CaCerts    []string `vici:"cacerts"`
	PubKeys    []string `vici:"pubkeys"`
}

type childSA struct {
	RemoteTrafficSelectors []string `vici:"remote_ts"`
	LocalTrafficSelectors  []string `vici:"local_ts"`
	Updown                 string   `vici:"updown"`
	ESPProposals           []string `vici:"esp_proposals"`
	AgProposals            []string `vici:"ag_proposals"`
	RekeyTime              string   `vici:"rekey_time"`
	LifeTime               string   `vici:"life_time"`
	RandTime               string   `vici:"rand_time"`
	Inactivity             uint32   `vici:"inactivity"`
	MarkIn                 uint32   `vici:"mark_in"`
	MarkInSa               string   `vici:"mark_in_sa"`
	MarkOut                uint32   `vici:"mark_out"`
	SetMarkIn              uint32   `vici:"set_mark_in"`
	SetMarkOut             uint32   `vici:"set_mark_out"`
	HwOffload              string   `vici:"hw_offload"`
}

type LoadedIKE struct {
	Name         string
	UniqueId     int                    `vici:"uniqueid"`
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
	IfIdIn       string                 `vici:"if-id-in"`
	IfIdOut      string                 `vici:"if-id-out"`
	LocalVips    []string               `vici:"local-vips"`
	RemoteVips   []string               `vici:"remote-vips"`
	TasksQueued  []string               `vici:"tasks-queued"`
	TasksActive  []string               `vici:"tasks-active"`
	TasksPassive []string               `vici:"tasks-passive"`
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
	DpdDelay    uint32               `vici:"dpd_delay"`
	DpdTimeout  uint32               `vici:"dpd_timeout"`
	Ppk         string               `vici:"ppk"`
	PpkRequired string               `vici:"ppk_required"`
	LocalAuths  []AuthConf           `vici:"local"`
	RemoteAuths []AuthConf           `vici:"remote"`
	Children    map[string]ConnChild `vici:"children"`
}

type AuthConf struct {
	Name       string   // This field will NOT be marshaled!
	Class      string   `vici:"class"`
	EapType    string   `vici:"eap-type"`
	EapVendor  string   `vici:"eap-vendor"`
	Xauth      string   `vici:"xauth"`
	Revocation string   `vici:"revocation"`
	Id         string   `vici:"id"`
	CaId       string   `vici:"ca_id"`
	AaaId      string   `vici:"aaa_id"`
	EapId      string   `vici:"eap_id"`
	XauthId    string   `vici:"xauth_id"`
	Groups     []string `vici:"groups"`
	CertPolicy []string `vici:"cert_policy"`
	Certs      []string `vici:"certs"`
	CaCerts    []string `vici:"cacerts"`
}

type ConnChild struct {
	Mode       string   `vici:"mode"`
	Label      string   `vici:"label"`
	RekeyTime  int64    `vici:"rekey_time"`
	RekeyBytes int64    `vici:"rekey_bytes"`
	RekeyPkt   int64    `vici:"rekey_packets"`
	LocalTS    []string `vici:"local-ts"`
	RemoteTS   []string `vici:"remote-ts"`
}
