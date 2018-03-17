package blockutils

import (
	"errors"
)

// Bitcoin script type backed by a byte array
// The string function is particularly helpful for working
// with the stack and getting it into a string representation
type Script []byte

func (script Script) IsP2PKH() bool {

	if len(script) != 25 {
		return false
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	firstOp := scriptReader.readByte()
	secondOp := scriptReader.readByte()

	if firstOp != 0x76 && secondOp != 0xa9 {
		return false
	}

	pushLength := uint64(scriptReader.readByte())
	scriptReader.readBytes(pushLength)

	thirdOp := scriptReader.readByte()
	fourthOp := scriptReader.readByte()

	if thirdOp != 0x88 && fourthOp != 0xac {
		return false
	}

	return true
}

func (script Script) P2PKHHash160() ([]byte, error) {
	if len(script) != 25 {
		return nil, errors.New("Invalid script length for P2PKH")
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	scriptReader.readBytes(2)
	pushLength := uint64(scriptReader.readByte())
	return scriptReader.readBytes(pushLength), nil
}

func (script Script) IsP2SH() bool {
	if len(script) != 23 {
		return false
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	firstOp := scriptReader.readByte()
	if firstOp != 0xa9 {
		return false
	}

	pushLength := uint64(scriptReader.readByte())
	scriptReader.readBytes(pushLength)

	secondOp := scriptReader.readByte()

	if secondOp != 0x87 {
		return false
	}

	return true

}

func (script Script) P2SHHash160() ([]byte, error) {
	if len(script) != 23 {
		return nil, errors.New("Invalid script length for P2SH")
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	scriptReader.readByte()
	pushLength := uint64(scriptReader.readByte())
	return scriptReader.readBytes(pushLength), nil
}
