package skyhashmanager

import (
	"errors"

	logging "gopkg.in/op/go-logging.v1"

	"github.com/labstack/gommon/log"
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
	Config        *SkyhashManagerConfig
	Subscriptions []*skyhash.PublicBroadcastChannelNode
}

func NewSkyhashManager(SkyhashManagerConfig *SkyhashManagerConfig) *SkyhashManager {
	shm := SkyhashManager{Config: SkyhashManagerConfig}
	return &shm
}

func (self *SkyhashManager) Start() {
	logger.Info("SkyhashManager started")
}

func (self *SkyhashManager) Shutdown() {
	logger.Info("SkyhashManager stopped")
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

func (self *SkyhashManager) Subscribe(pubkey cipher.PubKey) {
	node := skyhash.NewPublicBroadcastChannelNode()
	port := self.Config.Port + len(self.Subscriptions)

	addr, err := self.LookupAddr(pubkey)
	if err != nil {
		log.Error(err.Error())
	}

	node.AddConnection(addr)
	node.InitConnectionPool(port)

	self.Subscriptions = append(self.Subscriptions, node)
}

func (self *SkyhashManager) LookupAddr(pubkey cipher.PubKey) (string, error) {
	for _, peer := range PeerData {
		if peer.Pubkey == pubkey {
			return peer.IpAddr + ":" + peer.Port, nil
		}
	}

	return "", errors.New("Address not found")
}
