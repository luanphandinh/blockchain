package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// Take the data from the block
// create a counter (nounce) which starts at 0
// create a hash of data plus counter
// check the hash to see if it meets a set of requirements (quite vauge)

// Requirements:
// The First few bits must contains 0s

const Difficulty = 12

type ProofOfWork struct {
	// @TODO: make these private
	Block  *Block
	Target *big.Int
}

func (p *ProofOfWork) Run() (int, []byte, error) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data, err := p.InitData(nonce)
		if err != nil {
			return -1, nil, err
		}
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)

		intHash.SetBytes(hash[:])
		if intHash.Cmp(p.Target) == -1 {
			break
		}
		nonce++
	}

	fmt.Println()

	return nonce, hash[:], nil
}

func (p *ProofOfWork) Validate() (bool, error) {
	var intHash big.Int

	data, err := p.InitData(p.Block.Nonce)
	if err != nil {
		return false, err
	}
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(p.Target) == -1, nil
}

func NewProof(block *Block) *ProofOfWork {
	// Underlying byte look like this
	target := big.NewInt(1)
	// 256 bit of sha256 - difficulty
	// Then left shift
	// This will make sure out requirements to be met
	// Shift all the way to the left, this will leave Difficulty bits as 0 from the beginning
	target.Lsh(target, 256-Difficulty)
	return &ProofOfWork{
		Block:  block,
		Target: target,
	}
}

func (p *ProofOfWork) InitData(nonce int) ([]byte, error) {
	nonceBytes, err := toHex(int64(nonce))
	if err != nil {
		return nil, err
	}

	diffBytes, err := toHex(Difficulty)
	if err != nil {
		return nil, err
	}

	data := bytes.Join([][]byte{
		p.Block.Prevhash,
		p.Block.Data,
		nonceBytes,
		diffBytes,
	}, []byte{})

	return data, nil
}
