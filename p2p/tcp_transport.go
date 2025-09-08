package p2p

import (
	"errors"
	"log"
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
	OnPeerConnect func(Peer) error
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
}

func NewTCPTransport(ops TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		tcpTransportOpts: ops,
		rpcchan:          make(chan RPC),
	}
}

// consumes implements the transport interface which will return the readonly RPC channel
// For reading the incomming messages from another peer in the network
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcchan
}

func (t *TCPTransport) Close() error {
	return t.listener.Close()
}

func (t *TCPTransport) Dial(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	go t.handleConn(conn, true)
	return nil
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.tcpTransportOpts.ListenAddr)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()

	log.Printf("Listening for incoming connections on %s\n", t.tcpTransportOpts.ListenAddr)

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			log.Println("Listener closed, stopping accept loop")
			return
		}
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		log.Println("new connection from:", conn.RemoteAddr())
		go t.handleConn(conn, false)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn, outbound bool) {
	var err error
	defer func() {
		log.Println("Dropping peer connection", err)
		conn.Close()
	}()

	peer := NewTcpPeer(conn, outbound)
	if err := t.tcpTransportOpts.HandshakeFunc(peer); err != nil {
		return
	}

	if t.tcpTransportOpts.OnPeerConnect != nil {
		if err := t.tcpTransportOpts.OnPeerConnect(peer); err != nil {
			return
		}
	}

	rpc := RPC{}
	for {
		if err := t.tcpTransportOpts.Decoder.Decode(peer.conn, &rpc); err != nil {
			log.Printf("decode error: %v\n", err)
			continue
		}
		rpc.From = conn.RemoteAddr()
		t.rpcchan <- rpc
	}
}
