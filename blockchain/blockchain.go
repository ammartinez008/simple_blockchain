package blockchain

import (
	"time"
)

type Blockchain struct {
	chain               []Block
	currentTransactions []Block
}

type Block struct {
	index        int64
	timestamp    time.Time
	transactions []Block
	proof        string
	previousHash string
}

func New() Blockchain {
	bc := Blockchain{
		chain:               []Block{},
		currentTransactions: []Block{},
	}
	return bc
}

//create a new blockchain and add it to chain
func (b Blockchain) NewBlock(proof string, previousHash string) Block {
	block := Block{
		index:        int64(len(b.chain) + 1),
		timestamp:    time.Now(),
		transactions: b.currentTransactions,
		proof:        proof,
		previousHash: previousHash,
	}

	b.currentTransactions = []Block{}
	b.chain = append(b.chain, block)
	return block
}

func (b Blockchain) NewTransaction() {

}

func (b Blockchain) hash(block Block) {

}

func (b Blockchain) lastBlock() {

}
