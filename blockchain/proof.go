package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

var Difficulty uint = 12

func SetDificulty(difficulty uint) {
	Difficulty = difficulty
}

// @TODO: make this interface, the Block and Chain should rely on POW interface{} only
type ProofOfWork struct {
	// @TODO: make these private
	Block  *Block
	Target *big.Int
}

// Take the data from the block
// create a counter (nonce) which starts at 0
// create a hash of data plus counter
// check the hash to see if it meets a set of requirements (quite vauge)

// Requirements:
// The First few bits must contains 0s
func (p *ProofOfWork) Run() (int, []byte, error) {
	tracer.Trace("Starting to run proof of work...")
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data, err := p.InitData(nonce)
		if err != nil {
			return -1, nil, err
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

// @TODO: Make this flexible
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

	diffBytes, err := toHex(int64(Difficulty))
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
