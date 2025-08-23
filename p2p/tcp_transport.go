package p2p

import (
	"fmt"
	"net"
)

// Tcp peer that representes remote node over tcp established connection
type TCPPeer struct {
	// conn is underlying connection between two peers
	conn net.Conn

	// if we dial and retrieve a connection -> outbound is true
	// if we accept and retrieve a connection -> outbound is false
	outbound bool
}

func NewTcpPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (p *TCPPeer) Conn() net.Conn {
	return p.conn
}

func (p *TCPPeer) Outbound() bool {
	return p.outbound
}

// --- TCPTransport ---
type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder
	// mu            sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NOPHandshakeFunc,
		listenAddress: listenAddress,
		peers:         make(map[net.Addr]Peer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("new connection from:", conn.RemoteAddr())
		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTcpPeer(conn, true)
	if err := t.shakeHands(peer); err != nil {
		fmt.Println("handshake failed:", err)
		return
	}
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(peer.conn, msg); err != nil {
			fmt.Printf("decode error: %v\n", err)
			continue
		}
	}
}
