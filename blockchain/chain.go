package blockchain

type BlockChain struct {
	storage Storage
	pow     ProofOfWork
}

func InitBlockChain(storage Storage, pow ProofOfWork) (*BlockChain, error) {
	if storage == nil {
		tracer.Trace("InitBlockChain with default memoryStorage")
		storage = newMemoryStorage()
	}

	if pow == nil {
		tracer.Trace("InitBlockChain with default simpleProofOfWork")
		pow = NewProof()
	}

	chain := &BlockChain{
		storage: storage,
		pow:     pow,
	}

	lastBlock, err := storage.GetLastBlock()
	if err != nil {
		return nil, err
	}

	if lastBlock == nil {
		genesis, err := chain.Genesis()
		if err != nil {
			return nil, err
		}

		err = storage.AddBlock(genesis)
		if err != nil {
			return nil, err
		}
	}

	return chain, nil
}

func (chain *BlockChain) Genesis() (*Block, error) {
	tracer.Trace("Creating genesis block")
	// @TODO: put this in either from env, or some config.
	return chain.NewBlock("Geneisis", []byte{})
}

func (chain *BlockChain) NewBlock(data string, prevHash []byte) (*Block, error) {
	newBlock := NewBlock(data, prevHash)

	err := chain.pow.Run(newBlock)
	if err != nil {
		return nil, err
	}

	return newBlock, nil
}

func (chain *BlockChain) AddBlock(data string) error {
	prevBlock, err := chain.storage.GetLastBlock()
	if err != nil {
		return err
	}

	newBlock, err := chain.NewBlock(data, prevBlock.Hash)
	if err != nil {
		return err
	}

	chain.storage.AddBlock(newBlock)
	return nil
}
