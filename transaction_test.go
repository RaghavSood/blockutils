package blockutils

import (
	"fmt"
	"testing"
)

var digibytetxcoinbase = "010000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff1803527a5f04ea66ab5a08540000fd29000000052f6d70682f00000000020000000000000000266a24aa21a9ed735a4c6d92c7bc860c0558bf0b49feb40e553dffe846613bd6d6bac983473d2cf934eb8b120000001976a914510fffca0668d410aea742e95a2fefa7952f695e88ac0120000000000000000000000000000000000000000000000000000000000000000000000000"
var digibytetx = "0100000002be92100bc9f1b6e6e11637d3bbc841bea9cffcc0a5d710ef83e36c438d5dcd78020000006b483045022100eb4671f9bbcbcc937855ef8aad774ff81cd4aedc65f79fedf2a9c88c9cd566c6022034039dd992ab0be0db95a1d7b615bb2c39e7b16515c74c9f021dd39ac0ffe213012102f24f8135e2f62f81d6c4ff172fd2681a3e03cf7485510a2871ca2c41b5aa9733ffffffff89491ae9534c2c5b7f000352588ff7778999b5ee0d19cad1bc0396e3fdf48c9c000000006a47304402200d5fff4b02e1b89e7a5067c6f8383d08b56c14ea54b6cb6257601dd150b11a07022025c75524788cc1de76146d2849f4cb821739f47459ed8fc8057cf9b703f399f2012102629fe53bdbf029c7d3be5dd64758229f0f754529981d70788d916e48c9e9af6cffffffff025a4ccf805e0000001976a914b788297cf734149f6225228c50ff905917aa8f4088ac51e68b050a0000001976a914d00455c4000530f93bf53e32615a7dee6da2a03b88ac4d7a5f00"

// https://digiexplorer.info/tx/d0e075c1e5c52854a5b5386e89bd6436c767a2570901d38537703baef3a313ef
func TestDigiByteTx(t *testing.T) {
	tx, err := NewTransactionFromHexString(digibytetx)
	if err != nil {
		t.Errorf("Could not parse tx hex; %s", err)
	}

	if tx.Version != 1 {
		t.Errorf("Invalid tx version. Expect %d, got %d", 1, tx.Version)
	}

	if tx.TxId.String() != "d0e075c1e5c52854a5b5386e89bd6436c767a2570901d38537703baef3a313ef" {
		t.Errorf("TX ID did not match for digibyte tx, expected %s, got %s", "d0e075c1e5c52854a5b5386e89bd6436c767a2570901d38537703baef3a313ef", tx.TxId)
	}

	if len(tx.Vin) != 2 {
		t.Errorf("Invalid input count. Expected %d, found %d", 2, len(tx.Vin))
	}

	if len(tx.Vout) != 2 {
		t.Errorf("Invalid output count. Expected %d, found %d", 2, len(tx.Vin))
	}

	if tx.Locktime != 6257229 {
		t.Errorf("Incorrect lock time. Expected %d, found %d", 6257229, tx.Locktime)
	}

	if tx.Size != 373 {
		t.Errorf("Incorrect tx size. Expected %d, found %d", 373, tx.Size)
	}

	if tx.Vin[0].Hash.String() != "78cd5d8d436ce383ef10d7a5c0fccfa9be41c8bbd33716e1e6b6f1c90b1092be" {
		t.Errorf("TX ID did not match for input 0, expected %s, got %s", "78cd5d8d436ce383ef10d7a5c0fccfa9be41c8bbd33716e1e6b6f1c90b1092be", tx.Vin[0].Hash)
	}

	if tx.Vin[1].Hash.String() != "9c8cf4fde39603bcd1ca190deeb5998977f78f585203007f5b2c4c53e91a4989" {
		t.Errorf("TX ID did not match for input 1, expected %s, got %s", "9c8cf4fde39603bcd1ca190deeb5998977f78f585203007f5b2c4c53e91a4989", tx.Vin[1].Hash)
	}

	if tx.Vin[0].Index != 2 {
		t.Errorf("TX index did not match for input 0, expected %d, got %d", 2, tx.Vin[0].Index)
	}

	if tx.Vin[1].Index != 0 {
		t.Errorf("TX index did not match for input 1, expected %d, got %d", 0, tx.Vin[1].Index)
	}

	if tx.Vin[0].Script.String() != "483045022100eb4671f9bbcbcc937855ef8aad774ff81cd4aedc65f79fedf2a9c88c9cd566c6022034039dd992ab0be0db95a1d7b615bb2c39e7b16515c74c9f021dd39ac0ffe213012102f24f8135e2f62f81d6c4ff172fd2681a3e03cf7485510a2871ca2c41b5aa9733" {
		t.Errorf("TX script did not match for input 0, expected %s, got %s", "483045022100eb4671f9bbcbcc937855ef8aad774ff81cd4aedc65f79fedf2a9c88c9cd566c6022034039dd992ab0be0db95a1d7b615bb2c39e7b16515c74c9f021dd39ac0ffe213012102f24f8135e2f62f81d6c4ff172fd2681a3e03cf7485510a2871ca2c41b5aa9733", tx.Vin[0].Script)
	}

	if tx.Vin[1].Script.String() != "47304402200d5fff4b02e1b89e7a5067c6f8383d08b56c14ea54b6cb6257601dd150b11a07022025c75524788cc1de76146d2849f4cb821739f47459ed8fc8057cf9b703f399f2012102629fe53bdbf029c7d3be5dd64758229f0f754529981d70788d916e48c9e9af6c" {
		t.Errorf("TX script did not match for input 1, expected %s, got %s", "47304402200d5fff4b02e1b89e7a5067c6f8383d08b56c14ea54b6cb6257601dd150b11a07022025c75524788cc1de76146d2849f4cb821739f47459ed8fc8057cf9b703f399f2012102629fe53bdbf029c7d3be5dd64758229f0f754529981d70788d916e48c9e9af6c", tx.Vin[1].Script)
	}

	if tx.Vin[0].Sequence != 4294967295 {
		t.Errorf("TX sequence did not match for input 0, expected %d, got %d", 4294967295, tx.Vin[0].Sequence)
	}

	if tx.Vin[1].Sequence != 4294967295 {
		t.Errorf("TX sequence did not match for input 1, expected %d, got %d", 4294967295, tx.Vin[1].Sequence)
	}

	if tx.Vout[0].Value != 405887994970 {
		t.Errorf("TX value did not match for output 0, expected %d, got %d", 405887994970, tx.Vout[0].Value)
	}

	if tx.Vout[1].Value != 43042727505 {
		t.Errorf("TX value did not match for output 1, expected %d, got %d", 43042727505, tx.Vout[1].Value)
	}

	if tx.Vout[0].Script.String() != "76a914b788297cf734149f6225228c50ff905917aa8f4088ac" {
		t.Errorf("TX script did not match for output 0, expected %s, got %s", "76a914b788297cf734149f6225228c50ff905917aa8f4088ac", tx.Vout[0].Script)
	}

	if tx.Vout[1].Script.String() != "76a914d00455c4000530f93bf53e32615a7dee6da2a03b88ac" {
		t.Errorf("TX script did not match for output 1, expected %s, got %s", "76a914d00455c4000530f93bf53e32615a7dee6da2a03b88ac", tx.Vout[1].Script)
	}

	if tx.IsCoinbase() {
		t.Error("TX should not be a coinbase tx")
	}
}

// https://digiexplorer.info/tx/b982c9ccdd9898456bf7d35daeb2bac2fa00d490cf4e2db2d1bd8c76ca5a9ffc
func TestDigiByteCoinbaseTx(t *testing.T) {
	tx, err := NewTransactionFromHexString(digibytetxcoinbase)
	if err != nil {
		t.Errorf("Could not parse tx hex; %s", err)
	}
	if tx.Version != 1 {
		t.Errorf("Invalid tx version. Expect %d, got %d", 1, tx.Version)
	}

	if tx.TxId.String() != "b982c9ccdd9898456bf7d35daeb2bac2fa00d490cf4e2db2d1bd8c76ca5a9ffc" {
		t.Errorf("TX ID did not match for digibyte tx, expected %s, got %s", "b982c9ccdd9898456bf7d35daeb2bac2fa00d490cf4e2db2d1bd8c76ca5a9ffc", tx.TxId)
	}

	if len(tx.Vin) != 1 {
		t.Errorf("Invalid input count. Expected %d, found %d", 1, len(tx.Vin))
	}

	if len(tx.Vout) != 2 {
		t.Errorf("Invalid output count. Expected %d, found %d", 2, len(tx.Vin))
	}

	if tx.Locktime != 0 {
		t.Errorf("Incorrect lock time. Expected %d, found %d", 0, tx.Locktime)
	}

	if tx.Size != 192 {
		t.Errorf("Incorrect tx size. Expected %d, found %d", 192, tx.Size)
	}

	if tx.Vin[0].Hash.String() != "0000000000000000000000000000000000000000000000000000000000000000" {
		t.Errorf("TX ID did not match for input 0, expected %s, got %s", "0000000000000000000000000000000000000000000000000000000000000000", tx.Vin[0].Hash)
	}

	if tx.Vin[0].Index != 4294967295 {
		t.Errorf("TX index did not match for input 0, expected %d, got %d", 4294967295, tx.Vin[0].Index)
	}

	if tx.Vin[0].Script.String() != "03527a5f04ea66ab5a08540000fd29000000052f6d70682f" {
		t.Errorf("TX script did not match for input 0, expected %s, got %s", "03527a5f04ea66ab5a08540000fd29000000052f6d70682f", tx.Vin[0].Script)
	}

	if tx.Vin[0].Sequence != 0 {
		t.Errorf("TX sequence did not match for input 0, expected %d, got %d", 0, tx.Vin[0].Sequence)
	}

	if tx.Vout[0].Value != 0 {
		t.Errorf("TX value did not match for output 0, expected %d, got %d", 0, tx.Vout[0].Value)
	}

	if tx.Vout[1].Value != 79656858873 {
		t.Errorf("TX value did not match for output 1, expected %d, got %d", 79656858873, tx.Vout[1].Value)
	}

	if tx.Vout[0].Script.String() != "6a24aa21a9ed735a4c6d92c7bc860c0558bf0b49feb40e553dffe846613bd6d6bac983473d2c" {
		t.Errorf("TX script did not match for output 0, expected %s, got %s", "6a24aa21a9ed735a4c6d92c7bc860c0558bf0b49feb40e553dffe846613bd6d6bac983473d2c", tx.Vout[0].Script)
	}

	if tx.Vout[1].Script.String() != "76a914510fffca0668d410aea742e95a2fefa7952f695e88ac" {
		t.Errorf("TX script did not match for output 1, expected %s, got %s", "76a914510fffca0668d410aea742e95a2fefa7952f695e88ac", tx.Vout[1].Script)
	}

	if !tx.IsCoinbase() {
		t.Error("TX should be a coinbase tx")
	}
}

func ExampleNewTransactionFromHexString() {
	ltcsegwittx := "02000000000101b539b9e41717be24d14c06cd72aed10a1d9593a860067850116e458d96b56d660000000017160014336d166ab51b21b3ef2f0c885b7004bd3ad38b3dfeffffff0200c2eb0b000000001976a914f6a3510afba93284b4a1969bcf411a225423acd188ac4924fe020000000017a9148a4275e9d10794c5d54d0b2ef9d33cb028258c5a870247304402202a91f2110e7a06b926bb8166fbffac12552326c6099ff1f077f2f8e9a5ac74be02202d19aad053f65d30d89b99205696c8c18bebaca1a188c4f0886a0542b01d3dcc01210271f262fee7b7aba93564d0ed468018f3ccca489ef9c87032a8c9db2dc820f7a0ba671400"
	tx, err := NewTransactionFromHexString(ltcsegwittx)
	if err != nil {
		fmt.Errorf("Could not parse tx hex; %s", err)
	}
	fmt.Printf("TxId: %s\n", tx.TxId)
	fmt.Printf("Size: %d\n", tx.Size)
	fmt.Printf("Version: %d\n", tx.Version)
	fmt.Printf("Locktime: %d\n", tx.Locktime)
	for _, txin := range tx.Vin {
		fmt.Println("\n\tTransaction Inputs:")
		fmt.Printf("\tHash: %s\n", txin.Hash)
		fmt.Printf("\tIndex: %d\n", txin.Index)
		fmt.Printf("\tScript: %s\n", txin.Script)
		fmt.Printf("\tSequence: %d\n", txin.Sequence)
		fmt.Printf("\tScriptWitness: %s\n", txin.ScriptWitness)
	}
	for _, txout := range tx.Vout {
		fmt.Println("\n\tTransaction Outputs:")
		fmt.Printf("\tValue: %d\n", txout.Value)
		fmt.Printf("\tScript: %s\n", txout.Script)
	}

	// Output: TxId: 17b78667eb3a2b93de08d8b02c8171843f8bd84fd5797e8a4c3e455dc1d54903
	// Size: 249
	// Version: 2
	// Locktime: 1337274
	//
	// 	Transaction Inputs:
	// 	Hash: 666db5968d456e1150780660a893951d0ad1ae72cd064cd124be1717e4b939b5
	// 	Index: 0
	// 	Script: 160014336d166ab51b21b3ef2f0c885b7004bd3ad38b3d
	// 	Sequence: 4294967294
	// 	ScriptWitness: [304402202a91f2110e7a06b926bb8166fbffac12552326c6099ff1f077f2f8e9a5ac74be02202d19aad053f65d30d89b99205696c8c18bebaca1a188c4f0886a0542b01d3dcc01 0271f262fee7b7aba93564d0ed468018f3ccca489ef9c87032a8c9db2dc820f7a0]
	//
	// 	Transaction Outputs:
	// 	Value: 200000000
	// 	Script: 76a914f6a3510afba93284b4a1969bcf411a225423acd188ac
	//
	// 	Transaction Outputs:
	// 	Value: 50209865
	// 	Script: a9148a4275e9d10794c5d54d0b2ef9d33cb028258c5a87
}
