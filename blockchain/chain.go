package blockchain

import "context"

type BlockChain struct {
	storage Storage
	pow     ProofOfWork
}

type BlockChainIterator struct {
	currentBlock *Block
	chain        *BlockChain
}

func InitBlockChain(
	ctx context.Context,
	storage Storage,
	pow ProofOfWork,
) (*BlockChain, error) {
	if storage == nil {
		tracer.Trace(ctx, "InitBlockChain with default memoryStorage")
		storage = newMemoryStorage()
	}

	if pow == nil {
		tracer.Trace(ctx, "InitBlockChain with default simpleProofOfWork")
		pow = NewProof()
	}

	chain := &BlockChain{
		storage: storage,
		pow:     pow,
	}

	lastBlock, err := storage.GetLastBlock(ctx)
	if err != nil {
		return nil, err
	}

	if lastBlock == nil {
		genesis, err := chain.Genesis(ctx)
		if err != nil {
			return nil, err
		}

		err = storage.AddBlock(ctx, genesis)
		if err != nil {
			return nil, err
		}
	}

	return chain, nil
}

func (chain *BlockChain) Genesis(ctx context.Context) (*Block, error) {
	tracer.Trace(ctx, "Creating genesis block")
	// @TODO: put this in either from env, or some config.
	return chain.NewBlock(ctx, "Geneisis", []byte{})
}

func (chain *BlockChain) NewBlock(
	ctx context.Context,
	data string,
	prevHash []byte,
) (*Block, error) {
	newBlock := NewBlock(ctx, data, prevHash)

	err := chain.pow.Run(ctx, newBlock)
	if err != nil {
		return nil, err
	}

	return newBlock, nil
}

func (chain *BlockChain) ValidateBlock(ctx context.Context, block *Block) (bool, error) {
	return chain.pow.Validate(ctx, block)
}

func (chain *BlockChain) AddBlock(ctx context.Context, data string) error {
	prevBlock, err := chain.storage.GetLastBlock(ctx)
	if err != nil {
		return err
	}

	newBlock, err := chain.NewBlock(ctx, data, prevBlock.Hash)
	if err != nil {
		return err
	}

	return chain.storage.AddBlock(ctx, newBlock)
}

func (chain *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{nil, chain}
}

func (iterator *BlockChainIterator) Next(ctx context.Context) (*Block, error) {
	var block *Block
	var err error
	if iterator.currentBlock == nil {
		block, err = iterator.chain.storage.GetLastBlock(ctx)
	} else {
		block, err = iterator.chain.storage.GetBlock(ctx, iterator.currentBlock.Prevhash)
	}

	iterator.currentBlock = block
	if err != nil {
		return nil, err
	}

	return block, nil
}
