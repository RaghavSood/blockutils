package blockutils

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"strings"
)

// TODO: Maybe can optimize
func ReverseHex(b []byte) []byte {
	newb := make([]byte, len(b))
	copy(newb, b)
	for i := len(newb)/2 - 1; i >= 0; i-- {
		opp := len(newb) - 1 - i
		newb[i], newb[opp] = newb[opp], newb[i]
	}

	return newb
}

func DoubleSha256(data []byte) Hash256 {
	hash := sha256.New()
	hash.Write(data)
	firstSha256 := hash.Sum(nil)
	hash.Reset()
	hash.Write(firstSha256)
	return hash.Sum(nil)
}

func (script Script) String() string {
	return hex.EncodeToString(script)
}

func (hash256 Hash256) String() string {
	return hex.EncodeToString(ReverseHex(hash256))
}

func (witness WitnessScript) String() string {
	hexStrings := make([]string, len(witness))
	for i, _ := range witness {
		hexStrings[i] = hex.EncodeToString(witness[i])
	}
	// return hex.EncodeToString(witness[0])
	return "[" + strings.Join(hexStrings, " ") + "]"
}

func AllZero(s []byte) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}

func bytesToUInt64(input []byte) (ret uint64) {
	buf := bytes.NewBuffer(padByteArray(input, 8))
	binary.Read(buf, binary.LittleEndian, &ret)
	return ret
}

func copyFromIndex(input []byte, start uint64, length uint64) []byte {
	output := make([]byte, length)

	i := uint64(0)
	for i < length {
		output[i] = input[start+i]
		i += 1
	}

	return output
}

func padByteArray(input []byte, size int) []byte {
	l := len(input)
	if l == size {
		return input
	}

	tmp := make([]byte, size)
	copy(tmp, input)
	return tmp
}
