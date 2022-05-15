package blockchain

type BlockChain struct {
	blocks  []*Block
	storage Storage
}

func Genesis() (*Block, error) {
	tracer.Trace("Creating genesis block")
	// @TODO: put this in either from env, or some config.
	return NewBlock("Geneisis", []byte{})
}

func InitBlockChain(storage Storage) (*BlockChain, error) {
	if storage == nil {
		tracer.Trace("InitBlockChain with default memoryStorage")
		storage = newMemoryStorage()
	}

	lastBlock, err := storage.GetLastBlock()
	if err != nil {
		return nil, err
	}

	blocks := make([]*Block, 0)
	if lastBlock == nil {
		genesis, err := Genesis()
		if err != nil {
			return nil, err
		}

		err = storage.AddBlock(genesis)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, genesis)
	} else {
		blocks = append(blocks, lastBlock)
	}

	return &BlockChain{
		blocks:  blocks,
		storage: storage,
	}, nil
}

func (chain *BlockChain) GetBlocks() []*Block {
	return chain.blocks
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
