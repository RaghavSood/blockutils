package blockutils

import (
	"bytes"
	"encoding/binary"
	"log"
)

const length_UINT16 = 2
const length_UINT32 = 4
const length_UINT64 = 8

type ByteReader struct {
	Bytes  []byte
	Cursor uint64
}

func (r *ByteReader) readUint16() uint16 {
	val := uint16(0)
	buf := bytes.NewBuffer(r.Bytes[r.Cursor : r.Cursor+length_UINT16])
	err := binary.Read(buf, binary.LittleEndian, &val)
	if err != nil {
		log.Fatalf("Decode failed: %s", err)
	}
	r.Cursor += length_UINT16
	return val
}

func (r *ByteReader) readUint32() uint32 {
	val := uint32(0)
	buf := bytes.NewBuffer(r.Bytes[r.Cursor : r.Cursor+length_UINT32])
	err := binary.Read(buf, binary.LittleEndian, &val)
	if err != nil {
		log.Fatalf("Decode failed: %s", err)
	}
	r.Cursor += length_UINT32
	return val
}

func (r *ByteReader) readUint64() uint64 {
	val := uint64(0)
	buf := bytes.NewBuffer(r.Bytes[r.Cursor : r.Cursor+length_UINT64])
	err := binary.Read(buf, binary.LittleEndian, &val)
	if err != nil {
		log.Fatalf("Decode failed: %s", err)
	}
	r.Cursor += length_UINT64
	return val
}

func (r *ByteReader) readByte() byte {
	byteVal := r.Bytes[r.Cursor]
	r.Cursor += 1
	return byteVal
}

func (r *ByteReader) readBytes(length uint64) []byte {
	byteVals := r.Bytes[r.Cursor : r.Cursor+length]
	r.Cursor += length
	return byteVals
}

// Allows you to view data without moving the cursor and from a start point.
// Useful for cases to lookahead on data, such as checking if a tx is a
// segwit tx

func (r *ByteReader) peekBytesFrom(start uint64, length uint64) []byte {
	byteVals := r.Bytes[start : start+length]
	return byteVals
}

// Allows you to view data without moving the cursor. Useful for cases to
// lookahead on data, such as checking if a tx is a segwit tx

func (r *ByteReader) peekBytes(length uint64) []byte {
	byteVals := r.Bytes[r.Cursor : r.Cursor+length]
	return byteVals
}

// A compact size uint is defined as follows in the original satoshi code
// (it has since been somewhat replaced by CVarInt in present day bitcoin)
//
// if the first byte is >= 0x00 and <= 0xFC, it is interpreted as an uint8
// total length 1 byte
//
// if the first byte is 0xFD, the next 2 bytes interpreted as an uint16
// total length 3 bytes
//
// if the first byte is 0xFE, the next 4 bytes interpreted as an uint32
// total length 5 bytes
//
// if the first byte is 0xFF, the next 8 bytes interpreted as an uint64
// total length 9 bytes
//
// Always returns a uint64
func (r *ByteReader) readCompactSizeUint() uint64 {
	intType := r.readByte()
	switch intType {
	case 0xFF:
		return r.readUint64()
	case 0xFE:
		return uint64(r.readUint32())
	case 0xFD:
		return uint64(r.readUint16())
	default:
		return uint64(intType)
	}
}

// For segwit transactions, the canonical sha256(sha256(txhex)) returns
// an incorrect hash. The valid txid needs to be calculated from the tx
// as encodied in the original tx format. This requires us to strip the segwit
// data from the tx, which amounts to the two flag and mask bytes after the
// tx version and the segwit data between the end of the last output and the
// locktime.
func (r *ByteReader) stripSegwit(txstartpos uint64, outputendpos uint64, nlocktimepos uint64) []byte {
	txlength := nlocktimepos - txstartpos + 4
	dup := make([]byte, txlength)
	dup = copyFromIndex(r.Bytes, txstartpos, txlength)
	outputendpos = outputendpos - txstartpos
	txstartpos = 0
	noLocktime := append(dup[txstartpos:txstartpos+4], dup[txstartpos+6:outputendpos]...)
	withLocktime := append(noLocktime, dup[txlength-4:txlength]...)
	return withLocktime
}
