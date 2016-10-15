package main

import (
	"flag"
	"log"
	"time"
)

var flagNodeNumber int
var flagPortNumber int

func main() {

	flag.IntVar(&flagNodeNumber, "spawnnetwork", 1, "specify number of nodes you want to run")
	flag.IntVar(&flagPortNumber, "port", 6060, "specify port number you want to run nodes on")

	flag.Parse()

	nodes := []*PublicBroadcastChannelNode{}

	for i := 0; i < flagNodeNumber; i++ {
		nodes = append(nodes, NewPublicBroadcastChannelNode())
		//nodes listening on adjacent ports
		nodes[i].InitConnectionPool(flagPortNumber + i)
	}

	//connect to peer
	con, err := nodes[0].AddConnection("127.0.0.1:6061")
	_ = con

	if err != nil {
		log.Panic(err)
	}

	//create a message to send
	tm := TestMessage{Text: []byte("Message test")}
	//Node1.ConnectionPool.SendMessage(con, 1, &tm)
	nodes[0].Dispatcher.SendMessage(con, 1, &tm)

	//d1.BroadcastMessage(3, &tm)

	time.Sleep(time.Second * 10)
}
