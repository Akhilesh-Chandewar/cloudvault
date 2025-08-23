package p2p

import "net"

type Message struct {
	from    net.Addr
	payload []byte
}
