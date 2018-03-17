package blockutils

import (
	"encoding/hex"
	"testing"
)

func TestScriptP2PKH(t *testing.T) {
	var script Script
	var err error
	script, err = hex.DecodeString("76a914bdb2b538e6b07e93d6bafcef4bec9dc936818a1988ac")
	if err != nil {
		t.Errorf("Error decoding hex: %s", err)
	}

	isP2PKH := script.IsP2PKH()
	if isP2PKH == false {
		t.Error("Incorrectly declared script as non-P2PKH:")
	}

	hash160, err := script.P2PKHHash160()
	if err != nil {
		t.Errorf("Error reading hash160: %s", err)
	}
	hash160String := ToHexString(hash160)
	if hash160String != "bdb2b538e6b07e93d6bafcef4bec9dc936818a19" {
		t.Errorf("Returned incorrect hash160. Expected %s, got %s", "bdb2b538e6b07e93d6bafcef4bec9dc936818a19", hash160String)
	}
}

func TestScriptP2SH(t *testing.T) {
	var script Script
	var err error
	script, err = hex.DecodeString("a9144aef67ed61d391d6f3d9903ead92386c1efc992587")
	if err != nil {
		t.Errorf("Error decoding hex: %s", err)
	}

	isP2SH := script.IsP2SH()
	if isP2SH == false {
		t.Error("Incorrectly declared script as non-P2SH:")
	}

	hash160, err := script.P2SHHash160()
	if err != nil {
		t.Errorf("Error reading hash160: %s", err)
	}
	hash160String := ToHexString(hash160)
	if hash160String != "4aef67ed61d391d6f3d9903ead92386c1efc9925" {
		t.Errorf("Returned incorrect hash160. Expected %s, got %s", "4aef67ed61d391d6f3d9903ead92386c1efc9925", hash160String)
	}
}
