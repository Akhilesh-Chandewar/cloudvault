package main

import (
	"log"

	"github.com/Akhilesh-Chandewar/cloudvault/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	// transport options
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder,
	}

	// create transport
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	// file server options
	fileServerOptions := FileServerOptions{
		StorageRoot:    listenAddr + "_network_storage",
		PathTransform:  CASPathTransform,
		Transport:      tcpTransport,
		TransportOpts:  tcpTransportOpts,
		BootstrapAddrs: nodes,
	}

	fs := NewFileServer(fileServerOptions)

	fs.Options.TransportOpts.OnPeerConnect = func(peer *p2p.TCPPeer) error {
		return fs.OnPeer(peer)
	}

	return fs

}

func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(s1.Start())
	}()
	s2.Start()

	log.Println("Servers stopped")
}
