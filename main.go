package main

import (
	"fmt"
	"log"

	"github.com/Akhilesh-Chandewar/cloudvault/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.NOPDecoder{},
		OnPeerConnect: OnPeerConnect,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	go func() {
		for {
			msg := <-tr.Consume()
			log.Printf("from=%s payload=%s\n", msg.From.String(), string(msg.Payload))
		}
	}()
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}

func OnPeerConnect(p2p.Peer) error {
	fmt.Println("Doing something with peer ")
	return nil
}
