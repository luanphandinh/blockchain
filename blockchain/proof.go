package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

type ProofOfWork interface {
	Run(*Block) error
	Validate(*Block) (validated bool, err error)
}

var Difficulty uint = 12

func SetDificulty(difficulty uint) {
	Difficulty = difficulty
}

// Take the data from the block
// create a counter (nonce) which starts at 0
// create a hash of data plus counter
// check the hash to see if it meets a set of requirements (quite vauge)

// Requirements:
// The First few bits must contains 0s
type SimpleProofOfWork struct {
	Target *big.Int
}

func (p *SimpleProofOfWork) Run(block *Block) error {
	tracer.Trace("Starting to run proof of work...")
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data, err := p.InitData(block, nonce)
		if err != nil {
			return err
		}
		hash = sha256.Sum256(data)

		tracer.TraceCarriagef("\r%x", hash)

		intHash.SetBytes(hash[:])
		if intHash.Cmp(p.Target) == -1 {
			break
		}
		nonce++
	}

	tracer.Trace("")
	tracer.Trace("Finish proof of work")

	block.Nonce = nonce
	block.Hash = hash[:]

	return nil
}

func (p *SimpleProofOfWork) Validate(block *Block) (bool, error) {
	var intHash big.Int

	data, err := p.InitData(block, block.Nonce)
	if err != nil {
		return false, err
	}
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(p.Target) == -1, nil
}

// @TODO: Make this flexible
func NewProof() *SimpleProofOfWork {
	// Underlying byte look like this
	target := big.NewInt(1)
	// 256 bit of sha256 - difficulty
	// Then left shift
	// This will make sure out requirements to be met
	// Shift all the way to the left, this will leave Difficulty bits as 0 from the beginning
	target.Lsh(target, 256-Difficulty)
	return &SimpleProofOfWork{
		Target: target,
	}
}

func (p *SimpleProofOfWork) InitData(block *Block, nonce int) ([]byte, error) {
	nonceBytes, err := toHex(int64(nonce))
	if err != nil {
		return nil, err
	}

	diffBytes, err := toHex(int64(Difficulty))
	if err != nil {
		return nil, err
	}

	data := bytes.Join([][]byte{
		block.Prevhash,
		block.Data,
		nonceBytes,
		diffBytes,
	}, []byte{})

	return data, nil
}
