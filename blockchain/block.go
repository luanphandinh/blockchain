package blockchain

import (
	"bytes"
	"encoding/gob"
)

type Block struct {
	Hash     []byte
	Data     []byte
	Prevhash []byte
	Nonce    int
}

func NewBlock(data string, prevHash []byte) (*Block, error) {
	tracer.Tracef("Create new block with data: %x, prevHash: %x", data, prevHash)
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	p := NewProof(block)

	nonce, hash, err := p.Run()
	if err != nil {
		return nil, err
	}
	block.Hash = hash
	block.Nonce = nonce

	return block, nil
}

func (b *Block) Serialize() ([]byte, error) {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	if err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}

func DeserializeBlock(data []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}
