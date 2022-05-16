package blockchain

import "context"

type BlockChainConfigs struct {
	Storage     Storage
	ProofOfWork ProofOfWork
	GenBlock    func(ctx context.Context, data []byte, prevHash []byte) Block
}

type BlockChain struct {
	storage  Storage
	pow      ProofOfWork
	genBlock func(ctx context.Context, data []byte, prevHash []byte) Block
}

type BlockChainIterator struct {
	currentBlock Block
	chain        *BlockChain
}

func InitBlockChain(
	ctx context.Context,
	config *BlockChainConfigs,
) (*BlockChain, error) {
	chain := &BlockChain{}
	chain.applyConfigs(ctx, config)

	lastBlock, err := chain.storage.GetLastBlock(ctx)
	if err != nil {
		return nil, err
	}

	if lastBlock == nil {
		genesis, err := chain.Genesis(ctx)
		if err != nil {
			return nil, err
		}

		err = chain.storage.AddBlock(ctx, genesis)
		if err != nil {
			return nil, err
		}
	}

	return chain, nil
}

func (chain *BlockChain) applyConfigs(ctx context.Context, opts *BlockChainConfigs) {
	if opts.Storage == nil {
		tracer.Trace(ctx, "InitBlockChain with default memoryStorage")
		opts.Storage = newMemoryStorage()
	}

	if opts.ProofOfWork == nil {
		tracer.Trace(ctx, "InitBlockChain with default simpleProofOfWork")
		opts.ProofOfWork = NewProof()
	}

	if opts.GenBlock == nil {
		tracer.Trace(ctx, "InitBlockChain with default genSimpleBlock")
		opts.GenBlock = defaultGenBlock
	}

	chain.storage = opts.Storage
	chain.pow = opts.ProofOfWork
	chain.genBlock = opts.GenBlock
}

func (chain *BlockChain) Genesis(ctx context.Context) (Block, error) {
	tracer.Trace(ctx, "Creating genesis block")
	// @TODO: put this in either from env, or some config.
	return chain.NewBlock(ctx, []byte("Geneisis"), []byte{})
}

func (chain *BlockChain) NewBlock(
	ctx context.Context,
	data []byte,
	prevHash []byte,
) (Block, error) {
	newBlock := chain.genBlock(ctx, data, prevHash)

	err := chain.pow.Run(ctx, newBlock)
	if err != nil {
		return nil, err
	}

	return newBlock, nil
}

func (chain *BlockChain) ValidateBlock(ctx context.Context, block Block) (bool, error) {
	return chain.pow.Validate(ctx, block)
}

func (chain *BlockChain) AddBlock(ctx context.Context, data []byte) error {
	prevBlock, err := chain.storage.GetLastBlock(ctx)
	if err != nil {
		return err
	}

	newBlock, err := chain.NewBlock(ctx, data, prevBlock.GetHash())
	if err != nil {
		return err
	}

	return chain.storage.AddBlock(ctx, newBlock)
}

func (chain *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{nil, chain}
}

func (iterator *BlockChainIterator) Next(ctx context.Context) (Block, error) {
	var block Block
	var err error
	if iterator.currentBlock == nil {
		block, err = iterator.chain.storage.GetLastBlock(ctx)
	} else {
		block, err = iterator.chain.storage.GetBlock(ctx, iterator.currentBlock.GetPrevHash())
	}

	iterator.currentBlock = block
	if err != nil {
		return nil, err
	}

	return block, nil
}
