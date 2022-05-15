package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

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

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	if err != nil {
		// @TODO: return err, bad practice
		log.Panic(err)
	}

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		// @TODO: return err, bad practice
		log.Panic(err)
	}

	return &block
}
