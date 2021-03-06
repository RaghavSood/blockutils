package blockutils

import (
	"encoding/hex"
)

// Identical to bitcoin Script in transaction.go
// Just for better readability
type Hash256 []byte

// Represents a single block in a blockchain.
// blockutils does not validate transactions or blocks.
// Height is only provided for blocks with version 2 or higher
// and is 0 otherwise (be careful when dealing with that)
// Time is a unix timestamp
type Block struct {
	Version       uint32
	PrevBlockHash Hash256
	MerkleRoot    Hash256
	Hash          Hash256
	Time          uint32
	NBits         uint32
	Nonce         uint32
	TxCount       uint64
	Transactions  []*Transaction
	Height        uint64
	Coinbase      Script
}

// Returns a block parsed from the given hexstring (such as
// from `getblock` or insight-api)
func NewBlockFromHexString(hexstring string) (*Block, error) {
	txbytes, err := hex.DecodeString(hexstring)
	if err != nil {
		return nil, err
	}

	return NewBlockFromBytes(txbytes)
}

// Returns a block parsed from the given bytes (such as
// from reading a blockchain file)
func NewBlockFromBytes(blockbytes []byte) (*Block, error) {
	var err error
	blockreader := ByteReader{
		Bytes:  blockbytes,
		Cursor: 0,
	}

	// The block header is the first 80 bytes of a block
	// We peek at it since we still need to read it for its
	// information later on
	blockheader := blockreader.PeekBytes(80)

	// Since the block header contains the tx merkleroot, hashing the
	// header gives the block hash and automatically includes all the
	// transactions
	hash := DoubleSha256(blockheader)
	version := blockreader.ReadUint32()          // The first 4 bytes of a block are the version and signal bits
	prevhash := blockreader.ReadBytes(32)        // ... followed by the block hash of the previous block
	merkleroot := blockreader.ReadBytes(32)      // ... followed by the tx merkle root
	blockTime := blockreader.ReadUint32()        // ... followed by the unix mining time
	blockbits := blockreader.ReadUint32()        // ... followed by the nbits
	nonce := blockreader.ReadUint32()            // ... followed by the nonce. This terminates the block header
	txcount := blockreader.ReadCompactSizeUint() // We then have the number of transactions in the blocks

	txs := make([]*Transaction, txcount)
	i := uint64(0)
	for i < txcount {
		tx, err := ReadTransactionFromReader(&blockreader) // ... followed by the actual raw transactions
		if err != nil {
			return nil, err
		}

		txs[i] = tx
		i += 1
	}
	blockNumber := uint64(0)
	if version >= 2 { // The block number is only defined in the coinbase tx if v>=2
		coinbaseReader := ByteReader{
			Bytes:  txs[0].Vin[0].Script,
			Cursor: 0,
		}

		blockNumberLength := uint64(coinbaseReader.ReadByte())          // The first byte specifies the length of the block number
		blockHeightBytes := coinbaseReader.ReadBytes(blockNumberLength) // Read the actual block number bytes
		blockNumber = bytesToUInt64(blockHeightBytes)                   // Convert to uint64
	}

	block := &Block{
		Version:       version,
		Hash:          hash,
		PrevBlockHash: prevhash,
		MerkleRoot:    merkleroot,
		Time:          blockTime,
		NBits:         blockbits,
		Nonce:         nonce,
		TxCount:       txcount,
		Transactions:  txs,
		Height:        blockNumber,
		Coinbase:      txs[0].Vin[0].Script,
	}

	return block, err
}

func (block *Block) IsGenesisBlock() bool {
	return AllZero(block.PrevBlockHash)
}
