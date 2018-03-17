package blockutils

import (
	"fmt"
	"testing"
)

var dgb6257234 = "020400208cf785c17e9dfc2ff64aa063c1adf182d1a97dedc5740871d1a05b378565bf620b9bdd81ed697d7a0fec140499d0427875b9a4bdb6211ca14bbef7fa8ce30a49e766ab5a5cae461aa79914f303010000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff1803527a5f04ea66ab5a08540000fd29000000052f6d70682f00000000020000000000000000266a24aa21a9ed735a4c6d92c7bc860c0558bf0b49feb40e553dffe846613bd6d6bac983473d2cf934eb8b120000001976a914510fffca0668d410aea742e95a2fefa7952f695e88ac01200000000000000000000000000000000000000000000000000000000000000000000000000100000002be92100bc9f1b6e6e11637d3bbc841bea9cffcc0a5d710ef83e36c438d5dcd78020000006b483045022100eb4671f9bbcbcc937855ef8aad774ff81cd4aedc65f79fedf2a9c88c9cd566c6022034039dd992ab0be0db95a1d7b615bb2c39e7b16515c74c9f021dd39ac0ffe213012102f24f8135e2f62f81d6c4ff172fd2681a3e03cf7485510a2871ca2c41b5aa9733ffffffff89491ae9534c2c5b7f000352588ff7778999b5ee0d19cad1bc0396e3fdf48c9c000000006a47304402200d5fff4b02e1b89e7a5067c6f8383d08b56c14ea54b6cb6257601dd150b11a07022025c75524788cc1de76146d2849f4cb821739f47459ed8fc8057cf9b703f399f2012102629fe53bdbf029c7d3be5dd64758229f0f754529981d70788d916e48c9e9af6cffffffff025a4ccf805e0000001976a914b788297cf734149f6225228c50ff905917aa8f4088ac51e68b050a0000001976a914d00455c4000530f93bf53e32615a7dee6da2a03b88ac4d7a5f0001000000018406eac46f6f3b15a5e571810af69bd2f9259bbae075642ae59da6b000e418b5010000006b483045022100f8d8ee39f2b85c8ce78858d0842de3cb0b86d183aaf36e0c04a49e7cfb6e39fd02207dc9653c46258c6256ad6a4221115569cde836fbb16ee83c7a751aaf526e384a012102098e6d1444c81f8996daad02c03cfb97cff923440e4771e6c1195c977173c868ffffffff02b67e6b2d010000001976a914d9614692f408a27dd13b2d0f0492583131f591a888ac8f7ed5f3050000001976a91413bb88fcb733994225713acea00aa1fc102bbea388ac4f7a5f00"
var dgb6257234TxHashes = []string{"b982c9ccdd9898456bf7d35daeb2bac2fa00d490cf4e2db2d1bd8c76ca5a9ffc", "d0e075c1e5c52854a5b5386e89bd6436c767a2570901d38537703baef3a313ef", "34814eb7cb7f90b275cbc08c7c50507879f9eed1a23db2420e44b0abe2cfdcc3"}

func TestDGB6257234(t *testing.T) {
	block, err := NewBlockFromHexString(dgb6257234)
	if err != nil {
		t.Errorf("Could not parse block hex; %s", err)
	}

	if block.Height != 6257234 {
		t.Errorf("Incorrect block height. Expected %d, got %d", 6257234, block.Height)
	}

	if block.Version != 536871938 {
		t.Errorf("Incorrect block version. Expected %d, got %d", 536871938, block.Version)
	}

	if block.PrevBlockHash.String() != "62bf6585375ba0d1710874c5ed7da9d182f1adc163a04af62ffc9d7ec185f78c" {
		t.Errorf("Incorrect PrevBlockHash. Expected %s, got %s", "62bf6585375ba0d1710874c5ed7da9d182f1adc163a04af62ffc9d7ec185f78c", block.PrevBlockHash)
	}

	if block.MerkleRoot.String() != "490ae38cfaf7be4ba11c21b6bda4b9757842d0990414ec0f7a7d69ed81dd9b0b" {
		t.Errorf("Incorrect MerkleRoot. Expected %s, got %s", "490ae38cfaf7be4ba11c21b6bda4b9757842d0990414ec0f7a7d69ed81dd9b0b", block.MerkleRoot)
	}

	if block.Hash.String() != "7443ce7b891fbfb09a180320709d99e794974a1df2a87972cd3dd2c08e788c11" {
		t.Errorf("Incorrect block hash. Expected %s, got %s", "7443ce7b891fbfb09a180320709d99e794974a1df2a87972cd3dd2c08e788c11", block.Hash)
	}

	if block.Time != 1521182439 {
		t.Errorf("Incorrect block time. Expected %d, got %d", 1521182439, block.Time)
	}

	if block.NBits != 440839772 {
		t.Errorf("Incorrect block time. Expected %d, got %d", 440839772, block.NBits)
	}

	if block.Nonce != 4078213543 {
		t.Errorf("Incorrect block time. Expected %d, got %d", 4078213543, block.Nonce)
	}

	if block.TxCount != 3 {
		t.Errorf("Incorrect block time. Expected %d, got %d", 3, block.TxCount)
	}

	for i, transaction := range block.Transactions {
		if transaction.TxId.String() != dgb6257234TxHashes[i] {
			t.Errorf("Incorrect txid for tx %d. Expected %s, got %s", i, dgb6257234TxHashes[i], transaction.TxId)
		}
	}
}

func ExampleNewBlockFromHexString() {
	btc200 := "01000000eb68047fb29d78480b567ef6b76be556a2ec975656424508cc1c69b700000000bad58718fc3c6f5474918f06c44400c70b4c86d55a3f3ca3493b1d40c2061f2ba00f6b49ffff001d064b3a6d0101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0704ffff001d0138ffffffff0100f2052a010000004341045e071dedd1ed03721c6e9bba28fc276795421a378637fb41090192bb9f208630dcbac5862a3baeb9df3ca6e4e256b7fd2404824c20198ca1b004ee2197866433ac00000000"

	block, err := NewBlockFromHexString(btc200)
	if err != nil {
		fmt.Errorf("Could not parse block hex; %s", err)
	}

	fmt.Printf("Hash: %s\n", block.Hash)
	fmt.Printf("Version: %d\n", block.Version)
	fmt.Printf("Height: %d\n", block.Height)
	fmt.Printf("PrevBlockHash: %s\n", block.PrevBlockHash)
	fmt.Printf("MerkleRoot: %s\n", block.MerkleRoot)
	fmt.Printf("Time: %d\n", block.Time)
	fmt.Printf("NBits: %d\n", block.NBits)
	fmt.Printf("Nonce: %d\n", block.Nonce)
	fmt.Printf("TxCount: %d\n", block.TxCount)

	fmt.Println("\nTransactions:")

	for _, tx := range block.Transactions {
		fmt.Printf("TxId: %s\n", tx.TxId)
		fmt.Printf("Size: %d\n", tx.Size)
		fmt.Printf("Version: %d\n", tx.Version)
		fmt.Printf("Locktime: %d\n", tx.Locktime)
		for _, txin := range tx.Vin {
			fmt.Println("\nTransaction Inputs:")
			fmt.Printf("Hash: %s\n", txin.Hash)
			fmt.Printf("Index: %d\n", txin.Index)
			fmt.Printf("Script: %s\n", txin.Script)
			fmt.Printf("Sequence: %d\n", txin.Sequence)
			fmt.Printf("ScriptWitness: %s\n", txin.ScriptWitness)
		}
		for _, txout := range tx.Vout {
			fmt.Println("\nTransaction Outputs:")
			fmt.Printf("Value: %d\n", txout.Value)
			fmt.Printf("Script: %s\n", txout.Script)
		}
	}

	// Output: Hash: 000000008f1a7008320c16b8402b7f11e82951f44ca2663caf6860ab2eeef320
	// Version: 1
	// Height: 0
	// PrevBlockHash: 00000000b7691ccc084542565697eca256e56bb7f67e560b48789db27f0468eb
	// MerkleRoot: 2b1f06c2401d3b49a33c3f5ad5864c0bc70044c4068f9174546f3cfc1887d5ba
	// Time: 1231753120
	// NBits: 486604799
	// Nonce: 1832536838
	// TxCount: 1
	//
	// Transactions:
	// TxId: 2b1f06c2401d3b49a33c3f5ad5864c0bc70044c4068f9174546f3cfc1887d5ba
	// Size: 134
	// Version: 1
	// Locktime: 0
	//
	// Transaction Inputs:
	// Hash: 0000000000000000000000000000000000000000000000000000000000000000
	// Index: 4294967295
	// Script: 04ffff001d0138
	// Sequence: 4294967295
	// ScriptWitness: []
	//
	// Transaction Outputs:
	// Value: 5000000000
	// Script: 41045e071dedd1ed03721c6e9bba28fc276795421a378637fb41090192bb9f208630dcbac5862a3baeb9df3ca6e4e256b7fd2404824c20198ca1b004ee2197866433ac
}
