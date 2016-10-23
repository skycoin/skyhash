package gui

import (
	//"encoding/json"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	//"strconv"

	wh "github.com/skycoin/skycoin/src/util/http"
	"github.com/skycoin/skyhash/skyhash"
	"github.com/skycoin/skyhash/skyhashmanager"
)

//RegisterNodeManagerHandlers - create routes for NodeManager
func RegisterNodeManagerHandlers(mux *http.ServeMux, shm *skyhashmanager.SkyhashManager) {
	// enclose shm into SkyhashManager to be able to add methods
	lshm := SkyhashManager{SkyhashManager: shm}

	//  Test  Will be assigned name if present.
	mux.HandleFunc("/test", GET(lshm.testHandler))

	mux.HandleFunc("/subscriptions", GET(lshm.handlerListSubscriptions))

	//Route for starting Node
	mux.HandleFunc("/nodemanager/nodes/start", POST(lshm.handlerStartNode))

	//Route for stopping Node
	mux.HandleFunc("/nodemanager/nodes/stop", GET(lshm.handlerStopNode))

	mux.HandleFunc("/nodemanager/nodes", MethodsToHandlers(
		//Route for listing Nodes
		MethodToHandler(http.MethodGet, GET(lshm.handlerListNodes)),
	))

	mux.HandleFunc("/nodemanager/transports", MethodsToHandlers(
		//Route for listing transports from Node
		MethodToHandler(http.MethodGet, lshm.handlerListTransports),

		//Route for adding transport to Node
		MethodToHandler(http.MethodPost, lshm.handlerAddTransport),

		//Route for removing transport from Node
		MethodToHandler(http.MethodDelete, lshm.handlerRemoveTransport),
	))

}

func (self *SkyhashManager) testHandler(w http.ResponseWriter, r *http.Request) {

	wh.Error400(w, fmt.Sprint("Works!"))

	if addr := r.FormValue("addr"); addr == "" {
		wh.Error404(w)
	} else {
		//wh.SendOr404(w, nm.GetConnection(addr))
		wh.Error404(w)
	}
}

func (self *SkyhashManager) handlerListSubscriptions(w http.ResponseWriter, r *http.Request) {
	logger.Info("Get subscriptions list")

	keysSubscribedTo := make([]string, 0, len(self.Subscriptions))
	for key := range self.Subscriptions {
		keysSubscribedTo = append(keysSubscribedTo, key.Hex())
	}

	wh.SendJSON(w, keysSubscribedTo)
}

//Handler for /nodemanager/nodes/start
//method: POST
//url: /nodemanager/nodes/start
func (self *SkyhashManager) handlerStartNode(w http.ResponseWriter, r *http.Request) {
	// QUESTION: what parameters does this handler need?
	// QUESTION: should the node IDs be serial ints?
	// QUESTION: should node IDs be immutable? (i.e. a speicifc node has always the same ID)

	logger.Info("Starting Node")

	newNodeID := len(self.Nodes) + 1

	// TODO: manage node ids without conflicts or duplication

	if _, ok := self.Nodes[newNodeID]; ok {
		// node with this id already exists
		http.Error(w, "node already exists", http.StatusInternalServerError)
	}

	newNode := skyhash.NewPublicBroadcastChannelNode()
	newNode.InitConnectionPool(6060 + newNodeID)

	self.Nodes[newNodeID] = newNode

	w.WriteHeader(http.StatusOK)
	// return id of new node
	w.Write([]byte(strconv.Itoa(newNodeID)))
}

//Handler for /nodemanager/nodes/stop - stop Node
//method: GET
//url: /nodemanager/nodes/stop?id=value
func (self *SkyhashManager) handlerStopNode(w http.ResponseWriter, r *http.Request) {
	/*	logger.Info("Stoping Node")
		nodeID := r.FormValue("id")
		if nodeID == "" {
			wh.Error400(w, "Missing Node id")
			return
		}
		i, err := strconv.Atoi(nodeID)
		if err != nil {
			wh.Error400(w, "Node id must be integer")
			return
		}

		if len(shm.PubKeyList) < i {
			wh.Error400(w, "Invalid Node id")
			return
		}

		shm.NodesList[nm.PubKeyList[i]].Close()
		delete(shm.Subscriptions, shm.PubKeyList[i])
		nm.PubKeyList = append(nm.PubKeyList[:i], nm.PubKeyList[i+1:]...)*/
}

//Handler for /nodemanager/nodes
//method: GET
//url: /nodemanager/nodes
//return: array of PubKey
func (self *SkyhashManager) handlerListNodes(w http.ResponseWriter, r *http.Request) {
	// QUESTION: what info should be displayed about each node?

	logger.Info("Get list of nodes")

	nodeIDs := make([]int, 0, len(self.Nodes))
	for key := range self.Nodes {
		nodeIDs = append(nodeIDs, key)
	}

	wh.SendJSON(w, nodeIDs)
}

//Handler for /nodemanager/transports
//method: GET
//url: /nodemanager/transports?id=value
func (self *SkyhashManager) handlerListTransports(w http.ResponseWriter, r *http.Request) {
	logger.Info("List transports from Node")

	// TODO: nodeID, err := nodeIDFromURL(r.URL)
	nodeID := r.URL.Query().Get("id")
	if nodeID == "" {
		wh.Error400(w, "Missing Node id")
		return
	}
	i, err := strconv.Atoi(nodeID)
	if err != nil {
		wh.Error400(w, "Node id must be integer")
		return
	}

	// check whether node exists, and fetch it
	node, ok := self.Nodes[i]
	if !ok {
		wh.Error400(w, "Invalid Node id")
		return
	}

	addresses := make([]string, 0, len(node.ConnectionPool.Addresses))
	for key := range node.ConnectionPool.Addresses {
		addresses = append(addresses, key)
	}

	wh.SendJSON(w, addresses)
}

//Handler for /nodemanager/transports
//method: POST
//url: /nodemanager/transports
func (self *SkyhashManager) handlerAddTransport(w http.ResponseWriter, r *http.Request) {
	logger.Info("Add transport to Node")

	// TODO: nodeID, err := nodeIDFromURL(r.URL)
	nodeID := r.URL.Query().Get("id")
	if nodeID == "" {
		wh.Error400(w, "Missing Node id")
		return
	}
	i, err := strconv.Atoi(nodeID)
	if err != nil {
		wh.Error400(w, "Node id must be integer")
		return
	}

	// check whether node exists, and fetch it
	node, ok := self.Nodes[i]
	if !ok {
		wh.Error400(w, "Invalid Node id")
		return
	}

	// decode configuration of the new transport to be created
	var newTransportConfig struct {
		IP   string `json:"ip"`
		Port string `json:"port"`
	}
	err = json.NewDecoder(r.Body).Decode(&newTransportConfig)
	if err != nil {
		wh.Error400(w, "Error decoding config for transport")
		return
	}

	if newTransportConfig.IP == "" {
		wh.Error400(w, "ip is not set")
	}
	if newTransportConfig.Port == "" {
		wh.Error400(w, "port is not set")
		return
	}

	// add trasport to node
	_, err = node.ConnectionPool.Connect(fmt.Sprintf("%v:%v", newTransportConfig.IP, newTransportConfig.Port))
	if err != nil {
		wh.Error400(w, "Error while adding trasport to node")
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

//Handler for /nodemanager/transports
//method: DELETE
//url: /nodemanager/transports
func (self *SkyhashManager) handlerRemoveTransport(w http.ResponseWriter, r *http.Request) {
	/*	logger.Info("Remove transport from Node")

		var c TransportWithID
		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			wh.Error400(w, "Error decoding config for transport")
		}
		if len(nm.PubKeyList) < c.NodeID {
			wh.Error400(w, "Invalid Node id")
			return
		}
		logger.Info(strconv.Itoa(c.NodeID))

		nm.RemoveTransportsFromNode(c.NodeID, c.Transport)*/
}
