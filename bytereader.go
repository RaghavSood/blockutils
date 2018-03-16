package blockutils

import (
	"bytes"
	"encoding/binary"
	"log"
)

const LENGTH_UINT32 = 4

type ByteReader struct {
	Bytes  []byte
	Cursor uint
}

func (r *ByteReader) ReadUint32() uint32 {
	val := uint32(0)
	buf := bytes.NewBuffer(r.Bytes[r.Cursor : r.Cursor+LENGTH_UINT32])
	err := binary.Read(buf, binary.LittleEndian, &val)
	if err != nil {
		log.Fatalf("Decode failed: %s", err)
	}
	r.Cursor += LENGTH_UINT32
	return val
}
