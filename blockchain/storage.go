package blockchain

import "context"

type Storage interface {
	GetLastBlock(ctx context.Context) (*Block, error)
	GetBlock(ctx context.Context, key []byte) (*Block, error)
	AddBlock(ctx context.Context, b *Block) error
}

type memoryStorage struct {
	blocks        []*Block
	mapKeyToBlock map[string]*Block
}

func newMemoryStorage() Storage {
	return &memoryStorage{
		blocks:        make([]*Block, 0),
		mapKeyToBlock: make(map[string]*Block, 0),
	}
}

func (s *memoryStorage) GetLastBlock(ctx context.Context) (*Block, error) {
	if len(s.blocks) == 0 {
		return nil, nil
	}

	return s.blocks[len(s.blocks)-1], nil
}

func (s *memoryStorage) AddBlock(ctx context.Context, b *Block) error {
	s.blocks = append(s.blocks, b)
	s.mapKeyToBlock[string(b.Hash)] = b

	return nil
}

func (s *memoryStorage) GetBlock(ctx context.Context, key []byte) (*Block, error) {
	if block, ok := s.mapKeyToBlock[string(key)]; ok {
		return block, nil
	}

	return nil, nil
}
