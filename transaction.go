package blockutils

import (
	"encoding/hex"
)

// Bitcoin script type backed by a byte array
// The string function is particularly helpful for working
// with the stack and getting it into a string representation
type Script []byte

// Bitcoin witness script type backed by a 2d byte array
// The string function is particularly helpful for working
// with the stack and getting it into a string representation
type WitnessScript [][]byte

// A Transaction represents a complete Bitcoin-like transaction
//
// TxId should be used for the transaction id

type Transaction struct {
	Hash     Hash256 // not actually in blockchain data; calculated
	TxId     Hash256 // not actually in blockchain data; calculated
	Version  uint32
	Locktime uint32
	Vin      []TxInput
	Vout     []TxOutput
	Size     uint64
}

// Represents a single transaction output
//
// Each tx input includes the previous out point (a null hash for coinbase txs)
// the previous tx vout index, the script for this input, and the input sequence.
// If the transaction is a segwit transaction, ScriptWitness will contain the
// segwit stack for this input, and script will not contain a signature
//
// If the transaction is a coinbase tx, Index is 4294967295 (0xFFFFFFFF),
// and Hash is set to a null hash
// (0000000000000000000000000000000000000000000000000000000000000000), and
// Script contains the coinbase script
type TxInput struct {
	Hash          Hash256
	Index         uint32
	Script        Script
	Sequence      uint32
	ScriptWitness WitnessScript
}

// Represents a single transaction output, composed of its value and script
type TxOutput struct {
	Value  uint64
	Script Script
}

func readTxInput(txreader *ByteReader) (txin TxInput, err error) {
	previoushash := txreader.readBytes(32)         // The first 32 bytes of a tx input are the prev hash
	vout := txreader.readUint32()                  // ... followed by the vout index in the previous tx
	scriptlength := txreader.readCompactSizeUint() // ... followed up the scriptSig length
	script := txreader.readBytes(scriptlength)     // ... followed by the actual scriptSig
	sequence := txreader.readUint32()              // ... terminated by the sequence number

	txin = TxInput{
		Hash:     previoushash,
		Index:    vout,
		Script:   script,
		Sequence: sequence,
	}
	return txin, err
}

func readTxOutput(txreader *ByteReader) (txout TxOutput, err error) {
	value := txreader.readUint64()                 // First 8 bytes are the value of the output
	scriptlength := txreader.readCompactSizeUint() // ... followed up the script length
	script := txreader.readBytes(scriptlength)     // ... followed by the actual script

	txout = TxOutput{
		Value:  value,
		Script: script,
	}

	return txout, err
}

func readWitnessData(txreader *ByteReader, vinsize uint64) (witnessData [][][]byte, err error) {
	i := uint64(0)
	witnessData = make([][][]byte, vinsize)
	for i < vinsize { // There is one witness stack for each input
		stackSize := txreader.readCompactSizeUint() // Each stack has a length defined by a compact int
		witnessData[i] = make([][]byte, stackSize)
		j := uint64(0)
		for j < stackSize {
			stackItemLength := txreader.readCompactSizeUint() // Each stack item's length is also defined by a compact int
			witnessData[i][j] = make([]byte, stackItemLength)
			stackItem := txreader.readBytes(stackItemLength) // Read the actual stack item
			witnessData[i][j] = stackItem
			j += 1
		}
		i += 1
	}
	return witnessData, nil
}

// Parses a given byte array into a workable transaction
// such as from a file
func NewTransactionFromBytes(txbytes []byte) (*Transaction, error) {
	txreader := ByteReader{
		Bytes:  txbytes,
		Cursor: 0,
	}

	return ReadTransactionFromReader(&txreader)
}

// Parses a given hex string into a workable transaction. Ideal for use against
// getrawtransaction and insight-api etc.
func NewTransactionFromHexString(hexstring string) (*Transaction, error) {
	txbytes, err := hex.DecodeString(hexstring)
	if err != nil {
		return nil, err
	}

	return NewTransactionFromBytes(txbytes)
}

// Parses a transaction from a ByteReader. This is to be used when parsing
// an entire block
func ReadTransactionFromReader(b *ByteReader) (*Transaction, error) {
	var err error
	isSegwit := false
	outputendpos := uint64(0)
	txstartpos := b.Cursor
	// First 4 bytes of a tx are the tx version; most chains only have version 1
	version := b.readUint32()

	// If this is a segwit tx, the first two bytes following the version will be 0x00 0x01
	// We can peek these bytes to see if it is a purely segwit tx
	// This works because the immediate next byte after the version will never be 0x00 (except)
	// for coinbase transactions, where the following byte will then never be 0x01, as the input
	// tx is a null hash in coinbase transactions
	potentialSegwitFlag := b.peekBytes(2)
	if potentialSegwitFlag[0] == 0x00 && potentialSegwitFlag[1] == 0x01 {
		isSegwit = true
		b.readBytes(2)
	}

	// After the version is a variable int specifying how many inputs this tx has
	vinsize := b.readCompactSizeUint()

	i := uint64(0)
	txins := make([]TxInput, vinsize)
	for i < vinsize {
		txins[i], err = readTxInput(b)
		if err != nil {
			return nil, err
		}
		i += 1
	}

	voutsize := b.readCompactSizeUint()
	txouts := make([]TxOutput, voutsize)
	i = uint64(0)
	for i < voutsize {
		txouts[i], err = readTxOutput(b)
		if err != nil {
			return nil, err
		}
		i += 1
	}
	outputendpos = b.Cursor

	if isSegwit {
		witnessData, err := readWitnessData(b, vinsize)
		if err != nil {
			return nil, err
		}

		for i, _ := range txins {
			txins[i].ScriptWitness = witnessData[i]
		}
	}

	nlocktimepos := b.Cursor
	locktime := b.readUint32() // The Lock time is always the last 4 bytes of a tx

	txlength := nlocktimepos - txstartpos + 4

	hash := DoubleSha256(b.peekBytesFrom(txstartpos, txlength))
	txid := hash // Tx ID is the same as the hash for non-segwit transactions

	if isSegwit {
		originalFormat := b.stripSegwit(txstartpos, outputendpos, nlocktimepos) // This duplicates the original transaction and does not modify the underlying array
		txid = DoubleSha256(originalFormat)
	}

	// if AllZero(txins[0].Hash) {
	// 	txins = make([]TxInput, 0, 0)
	// }

	tx := &Transaction{
		Version:  version,
		Hash:     hash,
		TxId:     txid,
		Vin:      txins,
		Vout:     txouts,
		Locktime: locktime,
		Size:     txlength,
	}

	return tx, nil
}
