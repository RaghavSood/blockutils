package blockutils

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"golang.org/x/crypto/ripemd160"
	"strings"
)

// ReverseHex reverses the order of a given byte array.
// A fair number of data in the blockchain is reversed
func ReverseHex(b []byte) []byte {
	newb := make([]byte, len(b))
	copy(newb, b)
	for i := len(newb)/2 - 1; i >= 0; i-- {
		opp := len(newb) - 1 - i
		newb[i], newb[opp] = newb[opp], newb[i]
	}

	return newb
}

// Computes a ripemd160(data) for the given data.
// Used for address and script serialization, etc.
func Ripemd160(data []byte) Script {
	hash := ripemd160.New()
	hash.Write(data)
	return hash.Sum(nil)
}

// hash160 returns the RIPEMD160 hash of the SHA-256 HASH of the given data.
func Hash160(data []byte) []byte {
	h := Sha256(data)
	return Ripemd160(h[:])
}

// Computes a sha256(data) for the given data.
// Used for bech32, etc.
func Sha256(data []byte) Hash256 {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

// Computes a sha256(sha256(data)) for the given data.
// Used for block headers, etc.
func DoubleSha256(data []byte) Hash256 {
	hash := sha256.New()
	hash.Write(data)
	firstSha256 := hash.Sum(nil)
	hash.Reset()
	hash.Write(firstSha256)
	return hash.Sum(nil)
}

func ToHexString(data []byte) string {
	return hex.EncodeToString(data)
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
	return "[" + strings.Join(hexStrings, " ") + "]"
}

// Returns true if a byte array is all 0.
// Useful for checking coinbase inputs
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

// We need to pad the block number bytes to fit uint64
func padByteArray(input []byte, size int) []byte {
	l := len(input)
	if l == size {
		return input
	}

	tmp := make([]byte, size)
	copy(tmp, input)
	return tmp
}
