package p2p

import (
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

// type GObDecoder struct{}

// func (dec GObDecoder) Decode(r io.Reader, v *Message) error {
// 	return gob.NewDecoder(r).Decode(v)
// }

type NOPDecoder struct{}

func (dec NOPDecoder) Decode(r io.Reader, v *Message) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}
	v.payload = buf[:n]
	return nil
}
