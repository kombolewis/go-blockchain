package blockchain

import (
	"bytes"
	"crypto/sha256"
)

/**
 *Represents blockchain
 */

type Blockchain struct {
	Blocks []*Block
}

/**
 *Block chain is composed of multiple different blocks
 *Each block contains the data we want to pass around inside of
 *our database as well as a hash associated with the block itself
 *
 */
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

/**
 *Helps create hash based on previous hash and the data
 */
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

/*
 *make actual block
 */
func CreateBlock(data string, PrevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), PrevHash}
	block.DeriveHash()
	return block
}

/**
 *Adds block to the chain
 */

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
