package p2p

import "net"

// Peer represents remote node
type Peer interface {
    Conn() net.Conn
    Outbound() bool
}

// Transport that handles communication between peers beween nodes in the network.
// This can in form of (TCP , UDP , WebSockets)
type Transport interface {
	ListenAndAccept() error
}
