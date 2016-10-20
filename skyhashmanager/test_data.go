package skyhashmanager

import (
	"github.com/skycoin/skycoin/src/cipher"
)

type Peer struct {
	Pubkey cipher.Pubkey
	IpAddr string
	Port   string
}

var PeerData = []Peer{
	Peer{"1", "1.2.3.4", "6061"},
	Peer{"2", "1.2.3.5", "6061"},
}
