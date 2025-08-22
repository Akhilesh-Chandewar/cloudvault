package main

import (
	"github.com/Akhilesh-Chandewar/cloudvault/p2p"
	"log"
)

func main() {
	t := p2p.NewTCPTransport(":3000")
	if err := t.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {} // keep program running
}
