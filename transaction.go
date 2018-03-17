package blockutils

import (
	// "encoding/binary"
	"encoding/hex"
	// "errors"
	"fmt"
)

type Script []byte

type Transaction struct {
	Hash     Hash256 // not actually in blockchain data; calculated
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
	Value  uint64
	Script Script
}

func readTxInput(txreader *ByteReader) (txin TxInput, err error) {
	previoushash := txreader.ReadBytes(32)         // The first 32 bytes of a tx input are the prev hash
	vout := txreader.ReadUint32()                  // ... followed by the vout index in the previous tx
	scriptlength := txreader.ReadCompactSizeUint() // ... followed up the scriptSig length
	script := txreader.ReadBytes(scriptlength)     // ... followed by the actual scriptSig
	sequence := txreader.ReadUint32()              // ... terminated by the sequence number

	txin = TxInput{
		Hash:     previoushash,
		Index:    vout,
		Script:   script,
		Sequence: sequence,
	}

	return txin, err
}

func readTxOutput(txreader *ByteReader) (txout TxOutput, err error) {
	value := txreader.ReadUint64()                 // First 8 bytes are the value of the output
	scriptlength := txreader.ReadCompactSizeUint() // ... followed up the script length
	script := txreader.ReadBytes(scriptlength)     // ... followed by the actual script

	txout = TxOutput{
		Value:  value,
		Script: script,
	}

	return txout, err
}

func NewTransactionFromBytes(txbytes []byte) (*Transaction, error) {
	var err error
	txid := DoubleSha256(txbytes)

	txreader := ByteReader{
		Bytes:  txbytes,
		Cursor: 0,
	}

	// First 4 bytes of a tx are the tx version; most chains only have version 1
	version := txreader.ReadUint32()

	// After the version is a variable int specifying how many inputs this tx has
	vinsize := txreader.ReadCompactSizeUint()

	i := uint64(0)
	txins := make([]TxInput, vinsize)
	for i < vinsize {
		txins[i], err = readTxInput(&txreader)
		if err != nil {
			return nil, err
		}
		i += 1
	}

	voutsize := txreader.ReadCompactSizeUint()
	txouts := make([]TxOutput, voutsize)
	i = uint64(0)
	for i < voutsize {
		txouts[i], err = readTxOutput(&txreader)
		if err != nil {
			return nil, err
		}
		i += 1
	}

	locktime := txreader.ReadUint32()

	tx := &Transaction{
		Version:  version,
		Hash:     txid,
		Vin:      txins,
		Vout:     txouts,
		Locktime: locktime,
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
