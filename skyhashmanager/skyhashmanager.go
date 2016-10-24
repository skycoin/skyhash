package skyhashmanager

import (
	"errors"
	"sync"

	logging "gopkg.in/op/go-logging.v1"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skyhash/skyhash"
)

var (
	logger = logging.MustGetLogger("skyhashmanager")
)

type SkyhashManagerConfig struct {
	Port   int
	Pubkey cipher.PubKey
}

type SkyhashManager struct {
	sync.RWMutex
	Config        *SkyhashManagerConfig
	Subscriptions map[*cipher.PubKey]*skyhash.PublicBroadcastChannelNode
	Nodes         map[int]*skyhash.PublicBroadcastChannelNode
	Peers         []Peer
}

func NewSkyhashManager(SkyhashManagerConfig *SkyhashManagerConfig) *SkyhashManager {
	shm := SkyhashManager{
		Config:        SkyhashManagerConfig,
		Subscriptions: make(map[*cipher.PubKey]*skyhash.PublicBroadcastChannelNode),
		Peers:         []Peer{},
		Nodes:         make(map[int]*skyhash.PublicBroadcastChannelNode),
	}

	shm.bootstrapPeers()

	return &shm
}

func (self *SkyhashManager) Start() {

	logger.Info("SkyhashManager started")
}

func (self *SkyhashManager) Shutdown() {
	logger.Info("SkyhashManager stopped")
}

func (self *SkyhashManager) bootstrapPeers() error {
	test_publicKey_1, _ := cipher.GenerateKeyPair()
	test_publicKey_2, _ := cipher.GenerateKeyPair()

	self.Peers = []Peer{
		Peer{test_publicKey_1, "1.2.3.4", "6061"},
		Peer{test_publicKey_2, "1.2.3.5", "6061"},
	}
	return nil
}

// func (self *SkyhashManager) Subscribe(targetPubkey string, sourcePubkey string) {
// 	config := domain.NodeConfig{PubKey: sourcePubkey}
// 	node, err := mesh.NewNode(config)
// 	if err != nil {
// 		return err
// 	}
// 	self.NodesList[target] = node
// 	return nil
// }

func (self *SkyhashManager) Subscribe(pubkey cipher.PubKey) error {
	node := skyhash.NewPublicBroadcastChannelNode()
	port := self.Config.Port + len(self.Subscriptions)

	addr, err := self.LookupAddr(pubkey)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	node.AddConnection(addr)
	node.InitConnectionPool(port)

	self.Subscriptions[&pubkey] = node
	return nil
}

func (self *SkyhashManager) LookupAddr(pubkey cipher.PubKey) (string, error) {
	for _, peer := range self.Peers {
		if peer.Pubkey == pubkey {
			return peer.IpAddr + ":" + peer.Port, nil
		}
	}

	return "", errors.New("Address not found")
}
