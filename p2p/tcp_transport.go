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

// --- TCPTransport ---
type TCPTransport struct {
	tcpTransportOpts TCPTransportOpts
	listener         net.Listener
	// mu               sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(ops TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		tcpTransportOpts: ops,
		peers:            make(map[net.Addr]Peer),
	}
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

	msg := &Message{}
	for {
		if err := t.tcpTransportOpts.Decoder.Decode(peer.conn, msg); err != nil {
			fmt.Printf("decode error: %v\n", err)
			continue
		}
		msg.from = conn.RemoteAddr()
		fmt.Printf("message: %v from %v\n", string(msg.payload) , msg.from)
	}

	// buf := make([]byte , 20000)
	// for {
	// 	n, err := peer.conn.Read(buf)
	// 	if err != nil {
	// 		fmt.Println("read error:", err)
	// 		continue
	// 	}
	// 	fmt.Println(string(buf[:n]))
	// }
}
