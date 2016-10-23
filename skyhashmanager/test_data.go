package skyhashmanager

import (
	"github.com/skycoin/skycoin/src/cipher"
)

type Peer struct {
	Pubkey cipher.PubKey
	IpAddr string
	Port   string
}
