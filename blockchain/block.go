package blockchain

import (
	"bytes"
	"context"
	"encoding/gob"
)

type Block interface {
	GetHash() []byte
	GetPrevHash() []byte
	GetData() []byte
	GetNonce() int

	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type SimpleBlock struct {
	Hash     []byte
	Data     []byte
	Prevhash []byte
	Nonce    int
}

func (b *SimpleBlock) GetHash() []byte {
	return b.Hash
}

func (b *SimpleBlock) GetPrevHash() []byte {
	return b.Prevhash
}

func (b *SimpleBlock) GetData() []byte {
	return b.Data
}

func (b *SimpleBlock) GetNonce() int {
	return b.Nonce
}

func (b *SimpleBlock) Marshal() ([]byte, error) {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	if err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}

func (b *SimpleBlock) Unmarshal(data []byte) error {
	var block SimpleBlock

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		return err
	}

	b.Hash = block.Hash
	b.Prevhash = block.Prevhash
	b.Data = block.Data
	b.Nonce = block.Nonce

	return nil
}

func defaultGenBlock(ctx context.Context, data []byte, prevHash []byte) Block {
	tracer.Tracef(ctx, "Create new block with data: %x, prevHash: %x", data, prevHash)
	block := &SimpleBlock{[]byte{}, []byte(data), prevHash, 0}

	return block
}
