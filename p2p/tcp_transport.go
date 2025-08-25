package p2p

import (
	"fmt"
	"net"
	// "sync"
)

// Tcp peer that representes remote node over tcp established connection
type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
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

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

// --- TCPTransport ---
type TCPTransport struct {
	tcpTransportOpts TCPTransportOpts
	listener         net.Listener
	rpcchan          chan RPC
	// mu               sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(ops TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		tcpTransportOpts: ops,
		rpcchan:          make(chan RPC),
		peers:            make(map[net.Addr]Peer),
	}
}

//consumes implements the transport interface
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcchan
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.tcpTransportOpts.ListenAddr)
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

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTcpPeer(conn, true)
	if err := t.tcpTransportOpts.HandshakeFunc(peer); err != nil {
		fmt.Println("handshake failed:", err)
		conn.Close()
		return
	}

	rpc := &RPC{}
	for {
		if err := t.tcpTransportOpts.Decoder.Decode(peer.conn, rpc); err != nil {
			fmt.Printf("decode error: %v\n", err)
			continue
		}
		rpc.from = conn.RemoteAddr()
		fmt.Printf("message from=%s payload=%s\n", rpc.from.String(), string(rpc.payload))
	}
}
