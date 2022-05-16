package blockchain

type BlockChain struct {
	storage Storage
	pow     ProofOfWork
}

type BlockChainIterator struct {
	currentBlock *Block
	chain        *BlockChain
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

func (chain *BlockChain) ValidateBlock(block *Block) (bool, error) {
	return chain.pow.Validate(block)
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

func (chain *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{nil, chain}
}

func (iterator *BlockChainIterator) Next() (*Block, error) {
	var block *Block
	var err error
	if iterator.currentBlock == nil {
		block, err = iterator.chain.storage.GetLastBlock()
	} else {
		block, err = iterator.chain.storage.GetBlock(iterator.currentBlock.Prevhash)
	}

	iterator.currentBlock = block
	if err != nil {
		return nil, err
	}

	return block, nil
}
