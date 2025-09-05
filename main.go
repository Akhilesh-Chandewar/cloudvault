package main

import (
	"log"
	"time"

	"github.com/Akhilesh-Chandewar/cloudvault/p2p"
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

	// schedule stop after 5 seconds
	go func() {
		time.Sleep(5 * time.Second)
		fileServer.Stop()
		log.Println("File server stopped")
	}()

	// start
	if err := fileServer.Start(); err != nil {
		log.Fatal(err)
	}
}
