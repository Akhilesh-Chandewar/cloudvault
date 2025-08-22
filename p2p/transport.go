package p2p

// Peer represents remote node
type Peer struct {
}

// Transport that handles communication between peers beween nodes in the network.
// This can in form of (TCP , UDP , WebSockets)
type Transport interface {
	ListenAndAccept() error
}
