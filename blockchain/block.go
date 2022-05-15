package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	Prevhash []byte
	Nonce    int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	p := NewProof(block)
	nonce, hash := p.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}
