package main

import (
	"fmt"
	"log"

	"github.com/luanphandinh/blockchain/blockchain"
)

func main() {
	chain, err := blockchain.InitBlockChain()
	if err != nil {
		log.Fatal(err)
	}

	blocks := []string{
		"First block after Genesis",
		"Second block after Genesis",
		"Third block after Genesis",
	}

	for _, block := range blocks {
		err := chain.AddBlock(block)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, block := range chain.GetBlocks() {
		fmt.Printf("PrevHash: %x\n", block.Prevhash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		p := blockchain.NewProof(block)
		validated, err := p.Validate()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("target: %x, validated: %v\n", p.Target, validated)
	}
}
