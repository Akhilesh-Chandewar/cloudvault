package main

import (
	"github.com/Akhilesh-Chandewar/cloudvault/p2p"
	"log"
)

func main() {
	// transport options
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",                 
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder,
	}

	// create transport
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	// file server options
	fileServerOptions := FileServerOptions{
		StorageRoot:   "3000_network_storage",
		PathTransform: CASPathTransform,
		Transport:     tcpTransport,
		TransportOpts: tcpTransportOpts,
	}

	// create file server
	fileServer := NewFileServer(fileServerOptions)

	// start
	if err := fileServer.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}
