package gui

import (
	//"encoding/json"

	"fmt"
	"net/http"
	//"strconv"

	wh "github.com/skycoin/skycoin/src/util/http"
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
	mux.HandleFunc("/nodemanager/start", GET(lshm.handlerStartNode))

	//Route for stopping Node
	mux.HandleFunc("/nodemanager/stop", GET(lshm.handlerStopNode))

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

func (shm *SkyhashManager) testHandler(w http.ResponseWriter, r *http.Request) {

	wh.Error400(w, fmt.Sprint("Works!"))

	if addr := r.FormValue("addr"); addr == "" {
		wh.Error404(w)
	} else {
		//wh.SendOr404(w, nm.GetConnection(addr))
		wh.Error404(w)
	}
}

func (shm *SkyhashManager) handlerListSubscriptions(w http.ResponseWriter, r *http.Request) {
	logger.Info("Get subscriptions list")
	wh.SendJSON(w, shm.Subscriptions)
}

//Handler for /nodemanager/start
//method: GET
//url: /nodemanager/start
func (shm *SkyhashManager) handlerStartNode(w http.ResponseWriter, r *http.Request) {
	logger.Info("Starting Node")
}

//Handler for /nodemanager/stop - stop Node
//method: GET
//url: /nodemanager/stop?id=value
func (shm *SkyhashManager) handlerStopNode(w http.ResponseWriter, r *http.Request) {
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
func (shm *SkyhashManager) handlerListNodes(w http.ResponseWriter, r *http.Request) {
	/*	logger.Info("Get list nodes")
		var list ListNodes
		for _, PubKey := range nm.PubKeyList {
			list.PubKey = append(list.PubKey, string(PubKey[:]))
		}
		wh.SendJSON(w, nm.PubKeyList)*/
}

//Handler for /nodemanager/transports
//method: GET
//url: /nodemanager/transports?id=value
func (shm *SkyhashManager) handlerListTransports(w http.ResponseWriter, r *http.Request) {
	/*	logger.Info("Get transport from Node")
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

		if len(nm.PubKeyList) < i {
			wh.Error400(w, "Invalid Node id")
			return
		}

		wh.SendJSON(w, nm.GetTransportsFromNode(i))*/
}

//Handler for /nodemanager/transports
//method: POST
//url: /nodemanager/transports
func (shm *SkyhashManager) handlerAddTransport(w http.ResponseWriter, r *http.Request) {
	/*	logger.Info("Add transport to Node")

		var c ConfigWithID
		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			wh.Error400(w, "Error decoding config for transport")
		}
		if len(nm.PubKeyList) < c.NodeID {
			wh.Error400(w, "Invalid Node id")
			return
		}

		node := nm.GetNodeByIndex(c.NodeID)
		nodemanager.AddTransportToNode(node, c.Config)*/
}

//Handler for /nodemanager/transports
//method: DELETE
//url: /nodemanager/transports
func (shm *SkyhashManager) handlerRemoveTransport(w http.ResponseWriter, r *http.Request) {
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
