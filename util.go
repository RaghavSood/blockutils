package blockutils

import (
	"crypto/sha256"
	"encoding/hex"
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
