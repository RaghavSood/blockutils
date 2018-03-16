package blockutils

import (
	// "encoding/binary"
	"encoding/hex"
	// "errors"
	"fmt"
)

type Script []byte

func (script Script) String() string {
	return hex.EncodeToString(script)
}

type Transaction struct {
	Hash     Hash256 // not actually in blockchain data; calculated
	Hashstr  string  // not actually in blockchain data; calculated
	Version  uint32
	Locktime uint32
	Vin      []TxInput
	Vout     []TxOutput
}

type TxInput struct {
	Hash          Hash256
	Index         uint32
	Script        Script
	Sequence      uint32
	ScriptWitness [][]byte
}

type TxOutput struct {
	Value  int64
	Script Script
}

func NewTransactionFromBytes(txbytes []byte) (*Transaction, error) {
	txid := DoubleSha256(txbytes)
	txidstr := hex.EncodeToString(ReverseHex(txid))

	txreader := ByteReader{
		Bytes:  txbytes,
		Cursor: 0,
	}

	version := txreader.ReadUint32()

	tx := &Transaction{
		Version: version,
		Hash:    txid,
		Hashstr: txidstr,
	}

	return tx, nil
}

func NewTransactionFromHexString(hexstring string) (*Transaction, error) {
	fmt.Println("BlockUtils")
	txbytes, err := hex.DecodeString(hexstring)
	if err != nil {
		return nil, err
	}

	return NewTransactionFromBytes(txbytes)
}
