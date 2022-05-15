package blockchain

type Storage interface {
	GetLastBlock() (*Block, error)
	GetBlock(key []byte) (*Block, error)
	AddBlock(b *Block) error
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

func (s *memoryStorage) GetLastBlock() (*Block, error) {
	if len(s.blocks) == 0 {
		return nil, nil
	}

	return s.blocks[len(s.blocks)-1], nil
}

func (s *memoryStorage) AddBlock(b *Block) error {
	s.blocks = append(s.blocks, b)
	s.mapKeyToBlock[string(b.Hash)] = b

	return nil
}

func (s *memoryStorage) GetBlock(key []byte) (*Block, error) {
	if block, ok := s.mapKeyToBlock[string(key)]; ok {
		return block, nil
	}

	return nil, nil
}
