package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

/**
 *Block chain is composed of multiple different blocks
 *Each block contains the data we want to pass around inside of
 *our database as well as a hash associated with the block itself
 *
 */
type Block struct {
	Hash         []byte
	Transactions []*Transaction
	PrevHash     []byte
	Nonce        int
}

func (b *Block) HashTransactions() []byte {
	var txtHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txtHashes = append(txtHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txtHashes, []byte{}))

	return txHash[:]
}

/*
 *make actual block
 */
func CreateBlock(txs []*Transaction, PrevHash []byte) *Block {
	block := &Block{[]byte{}, txs, PrevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Handle(err)
	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
