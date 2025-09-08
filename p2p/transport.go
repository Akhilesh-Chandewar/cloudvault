package p2p

import "net"

// Peer represents remote node
type Peer interface {
	Conn() net.Conn
	Outbound() bool
	Close() error
}

// Transport that handles communication between peers beween nodes in the network.
// This can in form of (TCP , UDP , WebSockets)
type Transport interface {
	Dial(address string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
