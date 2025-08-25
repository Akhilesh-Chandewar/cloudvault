package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tr := NewTCPTransport(TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       NOPDecoder{},
	})

	assert.Equal(t, ":3000", tr.tcpTransportOpts.ListenAddr)
	assert.Nil(t, tr.ListenAndAccept())
}
