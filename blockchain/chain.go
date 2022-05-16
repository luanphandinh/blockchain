package blockchain

type BlockChain struct {
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

	// @TODO: probs need to start transaction wrapper here
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
		storage: storage,
	}, nil
}

func (chain *BlockChain) GetBlocks() []*Block {
	blocks, _ := chain.storage.GetBlocks()
	return blocks
}

func (chain *BlockChain) AddBlock(data string) error {
	prevBlock, err := chain.storage.GetLastBlock()
	if err != nil {
		return err
	}

	newBlock, err := NewBlock(data, prevBlock.Hash)
	if err != nil {
		return err
	}

	chain.storage.AddBlock(newBlock)
	return nil
}
