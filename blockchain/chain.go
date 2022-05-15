package blockchain

// Simple version for now.
type BlockChain struct {
	blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	// @TODO: return error for this function
	// This will fail in terms of empty blocks
	// Caller need to create genesis block first.
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
	// @TODO: put this in either from env, or some config.
	return NewBlock("Geneisis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) GetBlocks() []*Block {
	return chain.blocks
}
