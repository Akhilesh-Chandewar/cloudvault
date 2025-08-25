package p2p

import "net"

type RPC struct {
	from    net.Addr
	payload []byte
}
