package blockchain

// Simple version for now.
type BlockChain struct {
	blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) error {
	// @TODO: return error for this function
	// This will fail in terms of empty blocks
	// Caller need to create genesis block first.
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock, err := NewBlock(data, prevBlock.Hash)
	if err != nil {
		return err
	}

	chain.blocks = append(chain.blocks, newBlock)
	return nil
}

func Genesis() (*Block, error) {
	tracer.Trace("Creating genesis block")
	// @TODO: put this in either from env, or some config.
	return NewBlock("Geneisis", []byte{})
}

func InitBlockChain() (*BlockChain, error) {
	genesis, err := Genesis()
	if err != nil {
		return nil, err
	}

	return &BlockChain{[]*Block{genesis}}, nil
}

func (chain *BlockChain) GetBlocks() []*Block {
	return chain.blocks
}
