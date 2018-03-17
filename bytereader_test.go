package blockutils

import "testing"

type testpairuint64 struct {
	input  []byte
	output uint64
}

func TestReadCompactSizeUint(t *testing.T) {
	tests := []testpairuint64{
		{[]byte{0x45}, 69},
		{[]byte{0xfd, 0x03, 0x02}, 515},
		{[]byte{0xfe, 0x35, 0x64, 0x54, 0xe3}, 3813958709},
		{[]byte{0xff, 0x48, 0xfe, 0xad, 0x43, 0xec, 0xcc, 0x4d, 0x9a}, 11118768370167447112},
	}

	for _, pair := range tests {
		reader := ByteReader{
			Bytes:  pair.input,
			Cursor: 0,
		}

		intvalue := reader.readCompactSizeUint()

		if intvalue != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", intvalue,
			)
		}
	}
}
