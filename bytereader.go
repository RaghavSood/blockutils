package blockutils

import (
	"bytes"
	"encoding/binary"
	"log"
)

const LENGTH_UINT16 = 2
const LENGTH_UINT32 = 4
const LENGTH_UINT64 = 8

type ByteReader struct {
	Bytes  []byte
	Cursor uint64
}

func (r *ByteReader) ReadUint16() uint16 {
	val := uint16(0)
	buf := bytes.NewBuffer(r.Bytes[r.Cursor : r.Cursor+LENGTH_UINT16])
	err := binary.Read(buf, binary.LittleEndian, &val)
	if err != nil {
		log.Fatalf("Decode failed: %s", err)
	}
	r.Cursor += LENGTH_UINT16
	return val
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

func (r *ByteReader) ReadUint64() uint64 {
	val := uint64(0)
	buf := bytes.NewBuffer(r.Bytes[r.Cursor : r.Cursor+LENGTH_UINT64])
	err := binary.Read(buf, binary.LittleEndian, &val)
	if err != nil {
		log.Fatalf("Decode failed: %s", err)
	}
	r.Cursor += LENGTH_UINT64
	return val
}

func (r *ByteReader) ReadByte() byte {
	byteVal := r.Bytes[r.Cursor]
	r.Cursor += 1
	return byteVal
}

func (r *ByteReader) ReadBytes(length uint64) []byte {
	byteVals := r.Bytes[r.Cursor : r.Cursor+length]
	r.Cursor += length
	return byteVals
}

func (r *ByteReader) PeekBytes(length uint64) []byte {
	byteVals := r.Bytes[r.Cursor : r.Cursor+length]
	return byteVals
}

/**
 * A compact size uint is defined as follows in the original satoshi code
 * (it has since been somewhat replaced by CVarInt in present day bitcoin)
 *
 * if the first byte is >= 0x00 and <= 0xFC, it is interpreted as an uint8
 * total length 1 byte
 *
 * if the first byte is 0xFD, the next 2 bytes interpreted as an uint16
 * total length 3 bytes
 *
 * if the first byte is 0xFE, the next 4 bytes interpreted as an uint32
 * total length 5 bytes
 *
 * if the first byte is 0xFF, the next 8 bytes interpreted as an uint64
 * total length 9 bytes
 *
 * We always return a uint64
 */

func (r *ByteReader) ReadCompactSizeUint() uint64 {
	intType := r.ReadByte()
	switch intType {
	case 0xFF:
		return r.ReadUint64()
	case 0xFE:
		return uint64(r.ReadUint32())
	case 0xFD:
		return uint64(r.ReadUint16())
	default:
		return uint64(intType)
	}
}

func (r *ByteReader) StripSegwit(outputendpos uint64) []byte {
	txlength := len(r.Bytes)
	dup := make([]byte, txlength)
	copy(dup, r.Bytes)
	noLocktime := append(dup[0:4], dup[6:outputendpos]...)
	withLocktime := append(noLocktime, dup[len(dup)-4:]...)
	return withLocktime
}
