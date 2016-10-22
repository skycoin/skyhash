package skyhashmanager

import (
	"github.com/skycoin/skycoin/src/cipher"
)

type Peer struct {
	Pubkey cipher.PubKey
	IpAddr string
	Port   string
}

var test_publicKey_1, _ = cipher.PubKeyFromHex("1")
var test_publicKey_2, _ = cipher.PubKeyFromHex("2")

var PeerData = []Peer{
	Peer{test_publicKey_1, "1.2.3.4", "6061"},
	Peer{test_publicKey_2, "1.2.3.5", "6061"},
}
