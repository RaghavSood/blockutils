package blockutils

import (
	// "encoding/binary"
	"encoding/hex"
	// "errors"
	"fmt"
)

type Script []byte
type WitnessScript [][]byte

type Transaction struct {
	Hash     Hash256    `json:"hash"` // not actually in blockchain data; calculated
	TxId     Hash256    `json:"txid"`
	Version  uint32     `json:"version"`
	Locktime uint32     `json:"locktime"`
	Vin      []TxInput  `json:"vin"`
	Vout     []TxOutput `json:"vout"`
}

type TxInput struct {
	Hash          Hash256       `json:"hash"`
	Index         uint32        `json:"index"`
	Script        Script        `json:"scriptSig"`
	Sequence      uint32        `json:"sequence"`
	ScriptWitness WitnessScript `json:"witness"`
}

type TxOutput struct {
	Value  uint64 `json:"value"`
	Script Script `json:"script"`
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

func readWitnessData(txreader *ByteReader, vinsize uint64) (witnessData [][][]byte, err error) {
	i := uint64(0)
	witnessData = make([][][]byte, vinsize)
	for i < vinsize { //There is one witness stack for each input
		stackSize := txreader.ReadCompactSizeUint() // Each stack has a length defined by a compact int
		witnessData[i] = make([][]byte, stackSize)
		j := uint64(0)
		for j < stackSize {
			stackItemLength := txreader.ReadCompactSizeUint() //Each stack item's length is also defined by a compact int
			witnessData[i][j] = make([]byte, stackItemLength)
			stackItem := txreader.ReadBytes(stackItemLength) // Read the actual stack item
			witnessData[i][j] = stackItem
			j += 1
		}
		i += 1
	}
	return witnessData, nil
}

func NewTransactionFromBytes(txbytes []byte) (*Transaction, error) {
	var err error
	hash := DoubleSha256(txbytes)
	txid := hash
	isSegwit := false

	txreader := ByteReader{
		Bytes:  txbytes,
		Cursor: 0,
	}

	// First 4 bytes of a tx are the tx version; most chains only have version 1
	version := txreader.ReadUint32()

	// If this is a segwit tx, the first two bytes following the version will be 0x00 0x01
	// We can peek these bytes to see if it is a purely segwit tx
	// This works because the immediate next byte after the version will never be 0x00 (except)
	// for coinbase transactions, where the following byte will then never be 0x01, as the input
	// tx is a null hash in coinbase transactions
	potentialSegwitFlag := txreader.PeekBytes(2)
	if potentialSegwitFlag[0] == 0x00 && potentialSegwitFlag[1] == 0x01 {
		isSegwit = true
		txreader.ReadBytes(2)
	}

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

	if isSegwit {
		witnessData, err := readWitnessData(&txreader, vinsize)
		if err != nil {
			return nil, err
		}

		for i, _ := range txins {
			txins[i].ScriptWitness = witnessData[i]
		}
	}

	locktime := txreader.ReadUint32()

	tx := &Transaction{
		Version:  version,
		Hash:     hash,
		TxId:     txid,
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
